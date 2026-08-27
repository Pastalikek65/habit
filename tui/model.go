package tui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/Pastalikek65/habit/store"
)

var (
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("99")).Padding(0,1)
	selStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	normStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	faintStyle = lipgloss.NewStyle().Faint(true)
)

type Model struct {
	store *store.Store
	habits []store.Habit
	cursor int
}

func New(s *store.Store) Model {
	hs, _ := s.List()
	return Model{store: s, habits: hs}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q","ctrl+c": return m, tea.Quit
		case "j","down": if m.cursor < len(m.habits)-1 { m.cursor++ }
		case "k","up": if m.cursor>0 { m.cursor-- }
		case "c","enter":
			if len(m.habits)>0 {
				h := m.habits[m.cursor]
				today := time.Now().Format("2006-01-02")
				m.store.Check(h.ID, today)
				hs, _ := m.store.List()
				m.habits = hs
			}
		case "d":
			if len(m.habits)>0 {
				// delete not implemented, just ignore
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render(" habit — Termux ") + "\n")
	if len(m.habits)==0 {
		b.WriteString("  no habits — habit add \"drink water\"\n")
	}
	for i,h := range m.habits {
		streak, _ := m.store.Streak(h.ID)
		line := fmt.Sprintf(" %-20s 🔥 %d", h.Name, streak)
		if i==m.cursor { b.WriteString(selStyle.Render("▶"+line)+"\n") } else { b.WriteString(normStyle.Render(" "+line)+"\n") }
	}
	b.WriteString("\n"+faintStyle.Render("j/k:nav  c:check today  q:quit")+"\n")
	return b.String()
}
