package main

import "github.com/charmbracelet/lipgloss"

var (
    TitleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#00bcd4"))

    InputStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#ffffff"))

    GreetingStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#F400A1"))
)
