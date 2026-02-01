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

func Start(ctx context.Context, w widget.Widget) error {
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

func runApplication(ctx context.Context, teaProgram *tea.Program, program *Program) error {
	ticker := time.NewTicker(time.Second / 60)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case t := <-ticker.C:
			shouldRerender, err := program.FrameStep()
			if err != nil {
				return err
			}
			if shouldRerender {
				teaProgram.Send(RerenderMsg{TickTime: t})
			}
		}
	}
}
