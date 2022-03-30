package ui

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type responseMsg struct{}

func listenForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)+100))
			sub <- struct{}{}
		}
	}
}

// A command that waits for the activity on a channel.
func waitForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return responseMsg(<-sub)
	}
}
