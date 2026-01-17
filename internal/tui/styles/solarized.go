package styles

import (
	"charm.land/lipgloss/v2"
)

// Solarized Light Palette
var (
	solBase03  = lipgloss.Color("#002b36")
	solBase02  = lipgloss.Color("#073642")
	solBase01  = lipgloss.Color("#586e75")
	solBase00  = lipgloss.Color("#657b83")
	solBase0   = lipgloss.Color("#839496")
	solBase1   = lipgloss.Color("#93a1a1")
	solBase2   = lipgloss.Color("#eee8d5")
	solBase3   = lipgloss.Color("#fdf6e3")
	solYellow  = lipgloss.Color("#b58900")
	solOrange  = lipgloss.Color("#cb4b16")
	solRed     = lipgloss.Color("#dc322f")
	solMagenta = lipgloss.Color("#d33682")
	solViolet  = lipgloss.Color("#6c71c4")
	solBlue    = lipgloss.Color("#268bd2")
	solCyan    = lipgloss.Color("#2aa198")
	solGreen   = lipgloss.Color("#859900")
)

func NewSolarizedLightTheme() *Theme {
	t := &Theme{
		Name:   "solarized-light",
		IsDark: false,

		Primary:   solBlue,
		Secondary: solCyan,
		Tertiary:  solMagenta,
		Accent:    solOrange,

		// Backgrounds
		BgBase:        solBase3, // Background (Lightest)
		BgBaseLighter: solBase2, // Background (Primary)
		BgSubtle:      solBase2,
		BgOverlay:     solBase3,

		// Foregrounds
		FgBase:      solBase02, // Content (Darkest)
		FgMuted:     solBase0,  // Content (Medium)
		FgHalfMuted: solBase0,
		FgSubtle:    solBase1,  // Content (Medium-Light)
		FgSelected:  solViolet,

		Border:      solBase2,
		BorderFocus: solBlue,

		Success: solGreen,
		Error:   solRed,
		Warning: solYellow,
		Info:    solBlue,

		// Specific colors
		White: solBase3,

		Blue:      solBlue,
		BlueLight: solBlue,
		BlueDark:  solBlue,

		Yellow: solYellow,
		Citron: solYellow,

		Green:      solGreen,
		GreenLight: solGreen,
		GreenDark:  solGreen,

		Red:      solRed,
		RedLight: solRed,
		RedDark:  solRed,
		Cherry:   solRed,
	}

	// Text selection.
	t.TextSelection = lipgloss.NewStyle().Foreground(solBase3).Background(solBlue)

	// LSP and MCP status.
	t.ItemOfflineIcon = lipgloss.NewStyle().Foreground(solBase0).SetString("●")
	t.ItemBusyIcon = t.ItemOfflineIcon.Foreground(solYellow)
	t.ItemErrorIcon = t.ItemOfflineIcon.Foreground(solRed)
	t.ItemOnlineIcon = t.ItemOfflineIcon.Foreground(solGreen)

	// Editor: Yolo Mode.
	t.YoloIconFocused = lipgloss.NewStyle().Foreground(solBase3).Background(solYellow).Bold(true).SetString(" ! ")
	t.YoloIconBlurred = t.YoloIconFocused.Foreground(solBase3).Background(solBase0)
	t.YoloDotsFocused = lipgloss.NewStyle().Foreground(solYellow).SetString(":::")
	t.YoloDotsBlurred = t.YoloDotsFocused.Foreground(solBase0)

	// oAuth Chooser.
	t.AuthBorderSelected = lipgloss.NewStyle().BorderForeground(solGreen)
	t.AuthTextSelected = lipgloss.NewStyle().Foreground(solGreen)
	t.AuthBorderUnselected = lipgloss.NewStyle().BorderForeground(solBase2)
	t.AuthTextUnselected = lipgloss.NewStyle().Foreground(solBase0)

	return t
}
