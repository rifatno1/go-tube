package model

import "github.com/charmbracelet/lipgloss"

type color struct {
	Text_main      lipgloss.Color
	Text_alt       lipgloss.Color
	BG_terminal    lipgloss.Color
	BG_main        lipgloss.Color
	BG_alt         lipgloss.Color
	Button_BG_main lipgloss.Color
	Button_BG_alt  lipgloss.Color
	Danger_BG_main lipgloss.Color
	Danger_BG_alt  lipgloss.Color
}

var (
	LightTheme = color{
		Text_main:      lipgloss.Color("#1E2030"), // Deep Indigo Ink
		Text_alt:       lipgloss.Color("#43465E"), // Slate Secondary
		BG_terminal:    lipgloss.Color("#EEF2FF"), // Light Indigo Ice
		BG_main:        lipgloss.Color("#E0E7FF"), // Panel Background
		BG_alt:         lipgloss.Color("#C7D2FE"), // Active Row / Border Accent
		Button_BG_main: lipgloss.Color("#4F46E5"), // Rich Primary Indigo
		Button_BG_alt:  lipgloss.Color("#6366F1"), // Soft Indigo Accent
		Danger_BG_main: lipgloss.Color("#DC2626"), // Crisp Red
		Danger_BG_alt:  lipgloss.Color("#EF4444"), // Soft Red
	}
	DarkTheme = color{
		Text_main:      lipgloss.Color("#ECEFF4"), // Soft White
		Text_alt:       lipgloss.Color("#B48EAD"), // Muted Lavender
		BG_terminal:    lipgloss.Color("#1B1A23"), // Deep Dark Purple
		BG_main:        lipgloss.Color("#262335"), // Elevated Surface
		BG_alt:         lipgloss.Color("#342E48"), // Card / Selection Highlight
		Button_BG_main: lipgloss.Color("#9D7CD8"), // Vibrant Purple
		Button_BG_alt:  lipgloss.Color("#7AA2F7"), // Soft Lavender Blue
		Danger_BG_main: lipgloss.Color("#F7768E"), // Bright Coral Red
		Danger_BG_alt:  lipgloss.Color("#DB4B4B"), // Muted Red
	}
)
