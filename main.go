package main

import (
	"log"
	"os"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	aboutme int = iota
	skill
	experience
)

const SafeContentHeight = 10
const SafeContentWidth = 70
const Width = 12

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
	viewport viewport.Model
}

func ContentDetail() []tab {
	// STYLE
	headerStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#8800ff"))
	subHeaderStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#29f29b"))
	linkStyle := lipgloss.NewStyle().Underline(true).Foreground(lipgloss.Color("#0000FF"))

	tabs := []tab{
		{
			title: "About me",
			content: lipgloss.JoinVertical(lipgloss.Left,
				lipgloss.JoinHorizontal(lipgloss.Top,
					" /\\_/\\\n( o.o )\n > ^ <",
					"   ",
					"My name is Jirayut Najan.\nSecond year computer engineering\nstudy in Chulalongkorn University",
					lipgloss.NewStyle().Height(SafeContentHeight-1).Render(""),
				),
				lipgloss.JoinHorizontal(lipgloss.Left,
					linkStyle.Foreground(lipgloss.Color("#63e0ff")).Hyperlink("https://jirayutnajan.github.io/").Render("Website") + " ",
					linkStyle.Foreground(lipgloss.Color("#5e5e5e")).Hyperlink("https://github.com/jirayutNajan").Render("Github") + " ",
					linkStyle.Foreground(lipgloss.Color("#1877F2")).Hyperlink("https://www.facebook.com/jirayut.najan/").Render("Facebook") + " ",
					linkStyle.Foreground(lipgloss.Color("#EA4335")).Hyperlink("mailto:jirayutnajna05@gmail.com").Render("Email") + " ",
					),
			),
		},
		{
			title: "Experience",
			content: lipgloss.JoinVertical(lipgloss.Left, 
				lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Render("↑/↓ for scroll"),
				headerStyle.Render("Freelance Full-stack Developer | Academic Service Center, MSU"),
				subHeaderStyle.Render("[Mar, 2026]"),
				"Developed a full-stack online course registration platform for the Academic Service Center, Mahasarakham University.",
				linkStyle.Hyperlink("https://umsuregister.msu.ac.th/").Render("umsuregister.msu.ac.th"),
				"",

				headerStyle.Render("Fullstack Developer | Friday Activity Crew"),
				subHeaderStyle.Render("[September 2025 - March 2026]"),
				"Developed the frontend and backend for the Talent Journey project — an activity registration system for CEDT students in the “Friday Act” course.",
				"",

				headerStyle.Render("Intern RCC Automation | Kiatnakin Phatra Bank"),
				subHeaderStyle.Render("[June - July 2025]"),
				"Develop RccAssets Application for tracking assets in RCC department with Microsoft Power platform.",
				),
		},
		{
			title: "Projects",
			content: lipgloss.JoinVertical(lipgloss.Left, 
				headerStyle.Render("My personal projects") + lipgloss.NewStyle().Width(35).Render(" ") + lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Render("↑/↓ for scroll"),
				subHeaderStyle.Render("QuicknoteAI"),
				"Developed an AI-powered study assistant web application using ReactJS as a frontend, ExpressJS as a backend and MongoDB as a nosql database.\n",

				subHeaderStyle.Render("Traffy fondue data analysis"),
				"Collaborated in a group project to analyze Traffy Fondue public complaint data with the objective of predicting new ticket reopen probability and present as a dashboard.\n",

				subHeaderStyle.Render("Neko feed embedded project"),
				"Collaborated in a group project to develop an automated IoT cat feeder using a multi-node ESP32 architecture (Sensor, Gate-way, and Camera nodes) with real-time monitoring dashboard.\n",

				subHeaderStyle.Render("Drug Label Reader for the Elderly Mobile Application"),
				"Built an Android app with Kotlin to support elderly users in reading drug labels.\n",

				subHeaderStyle.Render("Pixel harvest 2d game"),
				"Collaborated in a group project to develop a 2D top-down farming game using Java and JavaFX.\n",
				),
		},
		{
			title: "Skill",
			content: lipgloss.JoinVertical(lipgloss.Left, 
				headerStyle.Render("Technicall Skills"),
				subHeaderStyle.Render("Languages: ") + "TypeScript, Javascript, Python, SQL, C, Java, Bash script",
				subHeaderStyle.Render("Framework/Library: ") + "Nodejs, ReactJS, NextJS, ExpressJS, Electron, \nPandas, Scikit-learn, JavaFx",
				subHeaderStyle.Render("Tools: ") + "PostgreSQL, Postman, Git, Github, Docker, AWS(ec2, s3), \nMongoDB, Redis, Google Colab, Github Actions, Linux, Claude Code, GeminiCLi",
				subHeaderStyle.Render("Soft Skill: ") + "Talk",
				),
		},
	}

	return tabs
}

func New() *Model {
	inactiveTabBorder := tabBorderWithBottom("┴", "─", "┴", true)
	activeTabBorder := tabBorderWithBottom("┘", " ", "└", true)
	title := tabBorderWithBottom("┌", "─", "└", false)
	title.Left = ""
	contentStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true, true).Padding(0, 1).Height(Width)

	tabs := ContentDetail()

	vp := viewport.New(viewport.WithHeight(SafeContentHeight), viewport.WithWidth(SafeContentWidth))
	vp.SetContent(tabs[0].content)

	return &Model{
		loading: true,
		styles: &styles{
			inactiveTab: inactiveTabBorder,
			activeTab: activeTabBorder,
			content: contentStyle,
			title: title,
		},
		tabs: &tabs,
		viewport: vp,
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
			m.viewport.GotoTop()
			// UPDATE STATE SHOULD BE ON UPDATE NOT VIEW
			wrappedContent := lipgloss.NewStyle().Width(m.viewport.Width()).Render((*m.tabs)[m.activeTab].content)
			m.viewport.SetContent(wrappedContent)
			return m, nil
		case "l", "right":
			m.activeTab = min(len(*m.tabs)-1, m.activeTab+1)
			m.viewport.GotoTop()
			wrappedContent := lipgloss.NewStyle().Width(m.viewport.Width()).Render((*m.tabs)[m.activeTab].content)
			m.viewport.SetContent(wrappedContent)
			return m, nil
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)

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

	bodyWidth := lipgloss.Width(navbar)
	// m.viewport.SetContent((*m.tabs)[m.activeTab].content)
	content := m.viewport.View()
	body := m.styles.content.Width(bodyWidth).Render(content)

	footer := lipgloss.Place(bodyWidth, 1, lipgloss.Bottom, lipgloss.Right, lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Render("←, ↑, →, ↓"))

	page := lipgloss.NewStyle().Padding(1, 2).Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			navbar,
			body,
			footer,
			),
		)

	v := tea.NewView(page)
	v.AltScreen = true
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
