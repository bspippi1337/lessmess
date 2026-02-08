package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/bspippi1337/restless/internal/tui/views"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type tabs int

const (
	tabWizard tabs = iota
	tabRequest
	tabStream
)

type keymap struct {
	TabNext key.Binding
	TabPrev key.Binding
	Quit    key.Binding
}

func (k keymap) ShortHelp() []key.Binding { return []key.Binding{k.TabPrev, k.TabNext, k.Quit} }
func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.TabPrev, k.TabNext},
		{k.Quit},
	}
}

type model struct {
	quiet bool
	w, h  int
	tab   tabs

	help help.Model
	keys keymap

	wizard views.Wizard
	req    views.Request
	stream views.Stream

	face views.Face
}

func newModel(quiet bool) model {
	k := keymap{
		TabNext: key.NewBinding(key.WithKeys("tab", "l"), key.WithHelp("tab", "next tab")),
		TabPrev: key.NewBinding(key.WithKeys("shift+tab", "h"), key.WithHelp("shift+tab", "prev tab")),
		Quit:    key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	}
	m := model{
		quiet:  quiet,
		tab:    tabWizard,
		help:   help.New(),
		keys:   k,
		wizard: views.NewWizard(),
		req:    views.NewRequest(),
		stream: views.NewStream(),
		face:   views.NewFace(quiet),
	}
	return m
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	if m.quiet {
		return nil
	}
	return tea.Tick(120*time.Millisecond, func(t time.Time) tea.Msg { return tickMsg(t) })
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.w, m.h = msg.Width, msg.Height
		m.wizard.SetSize(m.w, max(8, m.h-6))
		m.req.SetSize(m.w, max(8, m.h-6))
		m.stream.SetSize(m.w, max(8, m.h-6))
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.TabNext):
			m.tab = (m.tab + 1) % 3
			return m, nil
		case key.Matches(msg, m.keys.TabPrev):
			m.tab = (m.tab + 2) % 3
			return m, nil
		}
		// delegate keys
		switch m.tab {
		case tabWizard:
			var cmd tea.Cmd
			m.wizard, cmd = m.wizard.Update(msg)
			return m, cmd
		case tabRequest:
			var cmd tea.Cmd
			m.req, cmd = m.req.Update(msg)
			return m, cmd
		case tabStream:
			var cmd tea.Cmd
			m.stream, cmd = m.stream.Update(msg)
			return m, cmd
		}
	case tickMsg:
		if !m.quiet {
			m.face.Tick()
			return m, tea.Tick(120*time.Millisecond, func(t time.Time) tea.Msg { return tickMsg(t) })
		}
	}

	// delegate other messages
	switch m.tab {
	case tabWizard:
		var cmd tea.Cmd
		m.wizard, cmd = m.wizard.Update(msg)
		return m, cmd
	case tabRequest:
		var cmd tea.Cmd
		m.req, cmd = m.req.Update(msg)
		return m, cmd
	case tabStream:
		var cmd tea.Cmd
		m.stream, cmd = m.stream.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	header := views.Header("restless", m.tabLabel(), m.face.View())
	body := m.activeView()
	footer := views.Footer(m.help.View(m.keys))
	return strings.Join([]string{header, body, footer}, "\n")
}

func (m model) tabLabel() string {
	switch m.tab {
	case tabWizard:
		return "Connect & Discover"
	case tabRequest:
		return "Request Builder"
	case tabStream:
		return "Stream (SSE)"
	default:
		return fmt.Sprintf("Tab %d", m.tab)
	}
}

func (m model) activeView() string {
	switch m.tab {
	case tabWizard:
		return m.wizard.View()
	case tabRequest:
		return m.req.View()
	case tabStream:
		return m.stream.View()
	default:
		return ""
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
