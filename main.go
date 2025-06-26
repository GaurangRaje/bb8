package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type stageType string

const (
	stageAsk stageType = "ask"
	stageDone stageType = "done"
)

type model struct {
	name string
	input string
	stage stageType
}


func initialModel() model {
	return model{
		stage: stageAsk,
	}
}


func (m model) Init() tea.Cmd {
	return nil
}


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
        		return m, tea.Quit
		case tea.KeyEnter:
			if m.stage == stageAsk {
				m.name = m.input
				m.stage = stageDone
			}
			return m, nil
		case tea.KeyBackspace:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input) - 1]
			}
		default:
			m.input += msg.String()
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.stage == stageDone {
		return GreetingStyle.Render(fmt.Sprintf("\n Hallo %s!! \n\n", m.name))
	}
	return fmt.Sprintf("\n What's your name? %s", m.input)
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting program: %v\n", err)
		os.Exit(1)
	}
}
	
	
