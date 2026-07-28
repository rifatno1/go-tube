package handlers

import tea "github.com/charmbracelet/bubbletea"

func Listen_UFV(channel chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		msg, ok := <-channel
		if !ok {
			return nil
		}
		return msg
	}
}
func Push_UFV(channel chan tea.Msg, msg tea.Msg) {
	// Drain the channel if it already contains an unread older update
	select {
	case <-channel:
	default:
	}
	// Put the newest update into the channel
	select {
	case channel <- msg:
	default:
	}
}
