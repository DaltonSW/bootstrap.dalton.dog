// styles stores the various color and style information for output
package styles

import "github.com/charmbracelet/lipgloss"

// Colors used
var (
	SubtitleColor = lipgloss.AdaptiveColor{Light: "8", Dark: "15"}
	TitleColor    = lipgloss.AdaptiveColor{Light: "5", Dark: "5"}

	PassColor = lipgloss.AdaptiveColor{Light: "2", Dark: "2"}
	FailColor = lipgloss.AdaptiveColor{Light: "1", Dark: "1"}
	WarnColor = lipgloss.AdaptiveColor{Light: "3", Dark: "3"}
)

// Styles used
var (
	CommandStyle = lipgloss.NewStyle().Foreground(SubtitleColor).Italic(true)
	TaskStyle    = lipgloss.NewStyle().Foreground(TitleColor).Bold(true)
)
