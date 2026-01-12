package mouseactions

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Manager struct {
	isHovering             bool
	mouseDowned            bool
	OnHoverStateChangeFunc func(isHovering bool)
	OnMouseDownFunc        func(button tea.MouseButton)
	OnMouseUpFunc          func(button tea.MouseButton)
	OnClick                func(button tea.MouseButton)
}

func (m *Manager) OnHoverStateChange(isHovering bool) {
	if m.isHovering == isHovering {
		return
	}
	m.isHovering = isHovering
	if m.OnHoverStateChangeFunc != nil {
		m.OnHoverStateChangeFunc(isHovering)
	}
}

func (m *Manager) OnMouseDown(button tea.MouseButton) {
	if m.mouseDowned {
		return
	}
	m.mouseDowned = true
	if m.OnMouseDownFunc != nil {
		m.OnMouseDownFunc(button)
	}
}

func (m *Manager) OnMouseUp(button tea.MouseButton) {
	if m.OnMouseUpFunc != nil {
		m.OnMouseUpFunc(button)
	}
	if m.mouseDowned && m.OnClick != nil {
		m.OnClick(button)
	}
	m.mouseDowned = false
}

func (m *Manager) Reset() {
	if m.isHovering {
		if m.OnHoverStateChangeFunc != nil {
			m.OnHoverStateChangeFunc(false)
		}
	}
	m.isHovering = false
	m.mouseDowned = false
}

func (m *Manager) IsHovering() bool {
	return m.isHovering
}

func (m *Manager) IsActive() bool {
	return m.mouseDowned
}
