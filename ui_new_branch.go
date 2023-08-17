package main

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	newBranchLayout = lipgloss.NewStyle().
			Padding(1, 1, 1, 2)

	newBranchSuccessTextStyle = lipgloss.NewStyle().
					Bold(true).
					Foreground(lipgloss.AdaptiveColor{Light: "#25A065", Dark: "#2AD67F"})

	newBranchSuccessNameStyle = lipgloss.NewStyle().
					Bold(true).
					Foreground(lipgloss.AdaptiveColor{Light: "#1A1A1A", Dark: "#FFFDF5"}).
					Background(lipgloss.AdaptiveColor{Light: "#5B44FF", Dark: "#7653FF"})

	newBranchFailedTextStyle = lipgloss.NewStyle().
					Bold(true).
					Foreground(lipgloss.AdaptiveColor{Light: "#D63B3A", Dark: "#D63B3A"})

	newBranchFailedErrStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.AdaptiveColor{Light: "#1A1A1A", Dark: "#FFFDF5"}).
				Background(lipgloss.AdaptiveColor{Light: "#D63B3A", Dark: "#D63B3A"})
)

func makeSafeBranchName(branch string) string {

	var illegalChars []string = []string{" ", "\n", "\t"}

	for index := 0; index < len(illegalChars); index++ {
		target := illegalChars[index]
		branch = strings.Trim(branch, target)
		branch = strings.ReplaceAll(branch, target, "-")
	}

	branch = strings.ToLower(branch)

	return branch
}

type branchModel struct {
	branch string
	done   bool
	err    error
}

func newBranchModel(branch string) branchModel {
	return branchModel{branch: makeSafeBranchName(branch)}
}

func (m branchModel) Init() tea.Cmd {
	return func() tea.Msg { return m.branch }
}

func (m branchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	case string:
		time.Sleep(500 * time.Millisecond)
		_, m.err = createBranch(msg)
		m.done = true
		return m, tea.Quit
	}

	return m, nil
}

func (m branchModel) View() string {
	msg := newBranchLayout.Render(newBranchSuccessTextStyle.Render(" (⊙ˍ⊙) Creating new branch: " + m.branch + " "))
	if m.done {
		if m.err == nil {
			msg = newBranchLayout.Render(
				newBranchSuccessTextStyle.Render(" (^_~) Switched to a new branch: ") +
					newBranchSuccessNameStyle.Render(m.branch) + " ")
		} else {
			msg = newBranchLayout.Render(
				newBranchFailedTextStyle.Render(" (｡•́︿•̀｡) Create failed: ") +
					newBranchFailedErrStyle.Render(m.err.Error()) + " ")
		}
	}
	return msg
}
