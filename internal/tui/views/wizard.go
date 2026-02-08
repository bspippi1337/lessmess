package views

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Wizard struct {
	w, h int
	step int

	domain textinput.Model
	base   textinput.Model
	auth   textinput.Model
}

func NewWizard() Wizard {
	d := textinput.New()
	d.Placeholder = "openai.com (or api.example.com)"
	d.Prompt = "Domain: "
	d.Focus()

	b := textinput.New()
	b.Placeholder = "https://api.openai.com/v1"
	b.Prompt = "Base URL: "

	a := textinput.New()
	a.Placeholder = "Bearer / API key / Basic"
	a.Prompt = "Auth: "

	return Wizard{step: 0, domain: d, base: b, auth: a}
}

func (wz *Wizard) SetSize(w, h int) { wz.w, wz.h = w, h }

func (wz Wizard) Update(msg tea.Msg) (Wizard, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			wz.step = (wz.step + 1) % 3
			wz.domain.Blur(); wz.base.Blur(); wz.auth.Blur()
			if wz.step == 0 { wz.domain.Focus() }
			if wz.step == 1 { wz.base.Focus() }
			if wz.step == 2 { wz.auth.Focus() }
			return wz, nil
		}
	}

	if wz.step == 0 {
		wz.domain, cmd = wz.domain.Update(msg)
	} else if wz.step == 1 {
		wz.base, cmd = wz.base.Update(msg)
	} else {
		wz.auth, cmd = wz.auth.Update(msg)
	}
	return wz, cmd
}

func (wz Wizard) View() string {
	card := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).BorderForeground(cLine)
	title := lipgloss.NewStyle().Bold(true).Foreground(cTitle).Render("Connect & Discover Wizard")
	sub := lipgloss.NewStyle().Foreground(cDim).Render("Enter domain, base URL, and auth. Next: autodiscovery + presets.")

	fields := strings.Join([]string{
		wz.domain.View(),
		wz.base.View(),
		wz.auth.View(),
		"",
		lipgloss.NewStyle().Foreground(cDim).Render("Enter = next field · Tab = next tab"),
	}, "\n")

	return card.Render(title + "\n" + sub + "\n\n" + fields)
}
