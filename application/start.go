package application

import (
	"context"
	"os"
	"time"

	"github.com/calmdaysamuel/cheesecake/widget"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/palantir/witchcraft-go-logging/wlog"
	wlogzap "github.com/palantir/witchcraft-go-logging/wlog-zap"
	"github.com/palantir/witchcraft-go-logging/wlog/svclog/svc1log"
)

type StartOptionsFunc func(options *StartOptions)
type StartOptions struct {
	Context           context.Context
	LogFile           string
	LogLevel          wlog.LogLevel
	EnableMouseMotion bool
	FrameRate         int
	ApplicationName   string
}

func Start(w widget.Widget, options ...StartOptionsFunc) error {
	opts := defaultStartOptions()
	for _, option := range options {
		option(opts)
	}
	ctx := opts.Context

	wlog.SetDefaultLoggerProvider(wlogzap.LoggerProvider())

	logFile := "logs/" + time.Now().Format(time.DateOnly)
	if err := os.MkdirAll(logFile, 0750); err != nil {
		return err
	}
	file, err := os.OpenFile(logFile+"/logs.jsonl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	ctx = svc1log.WithLogger(ctx, svc1log.New(file, wlog.DebugLevel))
	program, err := NewProgram(ctx, w)
	if err != nil {
		svc1log.FromContext(ctx).Error("failed to start application", svc1log.Stacktrace(err))
		return err
	}
	t := tea.NewProgram(program, tea.WithContext(ctx), tea.WithAltScreen(), tea.WithoutCatchPanics(), tea.WithMouseAllMotion())
	go func() {
		if err := runApplication(ctx, t, program); err != nil {
			svc1log.FromContext(ctx).Info("Application stopped with unexpected error", svc1log.Stacktrace(err))
		}
	}()

	if _, err := t.Run(); err != nil {
		svc1log.FromContext(ctx).Error("application crashed or suddenly terminated.", svc1log.Stacktrace(err))
		return err
	}

	if program.LastError != nil {
		svc1log.FromContext(ctx).Error("application crashed with an unexpected error", svc1log.Stacktrace(program.LastError))
		return program.LastError
	}
	svc1log.FromContext(ctx).Info("application closed without error.")
	return nil
}

func defaultStartOptions() *StartOptions {
	return &StartOptions{
		Context:           context.Background(),
		LogFile:           "logs/" + time.Now().Format(time.DateOnly) + ".log",
		LogLevel:          wlog.InfoLevel,
		EnableMouseMotion: true,
		FrameRate:         60,
		ApplicationName:   "Cheese Application",
	}
}

func runApplication(ctx context.Context, teaProgram *tea.Program, program *Program) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			shouldRerender, err := program.FrameStep()
			if err != nil {
				return err
			}
			if shouldRerender {
				teaProgram.Send(RerenderMsg{TickTime: time.Now()})
			}
			time.Sleep(time.Second / 60)
		}
	}
}
