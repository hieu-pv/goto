package edithost

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// NewLabelInput - component which consists from input and label.
func NewLabelInput() labeledInput {
	inputModel := textinput.New()
	inputModel.Prompt = ""

	return labeledInput{
		Model:             inputModel,
		FocusedPrompt:     "│ ",
		LabelStyle:        noStyle,
		InputStyle:        noStyle,
		FocusedLabelStyle: focusedStyle,
		FocusedInputStyle: focusedStyle,
	}
}

type labeledInput struct {
	textinput.Model
	Label             string
	LabelStyle        lipgloss.Style
	InputStyle        lipgloss.Style
	FocusedLabelStyle lipgloss.Style
	FocusedInputStyle lipgloss.Style
	FocusedPrompt     string
}

func (l labeledInput) Update(msg tea.Msg) (labeledInput, tea.Cmd) {
	var cmd tea.Cmd

	l.Model, cmd = l.Model.Update(msg)

	return l, cmd
}

func (l labeledInput) prompt() string {
	if l.Focused() {
		return l.FocusedLabelStyle.Render(l.FocusedPrompt)
	}

	return strings.Repeat(" ", utf8.RuneCountInString(l.FocusedPrompt))
}

func (l labeledInput) labelView() string {
	if l.Focused() {
		return l.prompt() + l.FocusedLabelStyle.Render(l.Label)
	}

	return l.prompt() + l.LabelStyle.Render(l.Label)
}

func (l labeledInput) View() string {
	var view string
	if l.Focused() {
		view = lipgloss.NewStyle().Foreground(lipgloss.Color("#AD58B4")).Render(l.Model.View())
	} else {
		view = l.Model.View()
	}

	return fmt.Sprintf("%s\n%s%s", l.labelView(), l.prompt(), view)
}
