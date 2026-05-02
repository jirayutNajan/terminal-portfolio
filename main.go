package main

import (
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	aboutme int = iota
	skill
	experience
)

const width = 70

type styles struct {
	inactiveTab lipgloss.Border
	activeTab lipgloss.Border
	title lipgloss.Border
	content lipgloss.Style
}

type tab struct {
	title string
	content string
}

type Model struct {
	activeTab int
	loading bool
	width int
	height int
	styles *styles
	tabs *[]tab
}

func New() *Model {
	inactiveTabBorder := tabBorderWithBottom("┴", "─", "┴", true)
	activeTabBorder := tabBorderWithBottom("┘", " ", "└", true)
	title := tabBorderWithBottom("┌", "─", "└", false)
	title.Left = ""
	contentStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true, true).Padding(1)

	catArt := " /\\_/\\\n( o.o )\n > ^ <"
	infoText := "My name is Jirayut Najan.\nSecond year computer engineering\nstudy in Chulalongkorn University"

	return &Model{
		loading: true,
		styles: &styles{
			inactiveTab: inactiveTabBorder,
			activeTab: activeTabBorder,
			content: contentStyle,
			title: title,
		},
		tabs: &[]tab{
			{
				title: "About Me",
				content: lipgloss.JoinVertical(lipgloss.Left,
					lipgloss.JoinHorizontal(lipgloss.Top,
						catArt,
						"   ",
						infoText,
					),
				),
			},
			{
				title: "Skill",
				content: "eiei",
			},
			{
				title: "Experience",
				content: "11",
			},
		},
	}
}

func tabBorderWithBottom(left, middle, right string, round bool) lipgloss.Border {
	var border lipgloss.Border
	if round {
		border = lipgloss.RoundedBorder()
	} else {
		border = lipgloss.NormalBorder()
	}

	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right

	return border
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if m.loading {
			m.height = msg.Height
			m.width = msg.Width
			m.loading = false
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "h", "left":
			m.activeTab = max(0, m.activeTab-1)
			return m, nil
		case "l", "right":
			m.activeTab = min(len(*m.tabs)-1, m.activeTab+1)
			return m, nil
		}
	}

	return m, cmd
}


func (m Model) View() tea.View {

	if m.loading {
		v := tea.NewView("loading...")
		return v
	}

	var tabs []string

	// i: index
	for i, tab := range *m.tabs {
		var borderStyle lipgloss.Border
		if m.activeTab == i {
			borderStyle = m.styles.activeTab
			if i == len(*m.tabs)-1 {
				borderStyle.BottomRight = "│"
			}
		} else {
			borderStyle = m.styles.inactiveTab
			if i == len(*m.tabs)-1 {
				borderStyle.BottomRight = "┤"
			}
		}

		
		tabs = append(tabs, lipgloss.NewStyle().Border(borderStyle, true).Render(tab.title))
	}

	nav := append([]string{
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#14d8ff")).
		Border(m.styles.title, false, false, true, true).
		Width(35).
		Render("Jirayut-terminal-portfolio"),
	}, tabs...)

	navbar := lipgloss.JoinHorizontal(
		lipgloss.Bottom,
			nav...
		)

	navWidth := lipgloss.Width(navbar)
	content := m.styles.content.Width(navWidth).Render((*m.tabs)[m.activeTab].content)

	page := lipgloss.NewStyle().Padding(1, 2).Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			navbar,
			content,
			),
		)

	v := tea.NewView(page)
	return v
}

func main() {
	m := New()

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
