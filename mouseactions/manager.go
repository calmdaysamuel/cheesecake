package mouseactions

type Manager struct {
	OnHoverStateChangeFunc func(isHovering bool)
	OnClickFunc            func()
}

func (m *Manager) OnHoverStateChange(isHovering bool) {
	if m.OnHoverStateChangeFunc != nil {
		m.OnHoverStateChangeFunc(isHovering)
	}
}

func (m *Manager) OnClick() {
	if m.OnClickFunc != nil {
		m.OnClickFunc()
	}
}
