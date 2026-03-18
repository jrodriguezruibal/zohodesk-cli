package tui

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/config"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("15")).
			Background(lipgloss.Color("62")).
			Padding(0, 1)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("36")).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196"))
)

type setupModel struct {
	step          int
	profileName   string
	clientID      string
	clientSecret  string
	orgID         string
	region        string
	validating    bool
	valid         bool
	errorMsg      string
}

func (m setupModel) Init() tea.Cmd {
	return nil
}

func (m setupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			return m.handleEnter()
		case tea.KeyBackspace:
			return m.handleBackspace(msg)
		default:
			return m.handleInput(msg)
		}
	case validationMsg:
		m.validating = false
		if msg.success {
			m.valid = true
			return m, tea.Quit
		}
		m.errorMsg = msg.error
		return m, nil
	}
	return m, nil
}

func (m setupModel) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render(" zohodesk-cli Configuration Wizard "))
	b.WriteString("\n\n")

	switch m.step {
	case 0:
		b.WriteString(labelStyle.Render("Profile name: "))
		b.WriteString(m.profileName)
		b.WriteString("█\n")
		b.WriteString("\nPress Enter to continue or type a custom name\n")

	case 1:
		b.WriteString(labelStyle.Render("Zoho Client ID: "))
		b.WriteString(maskString(m.clientID))
		b.WriteString("█\n")
		if m.clientID == "" {
			b.WriteString("\nEnter your Zoho Client ID (starts with '1000.')\n")
		}

	case 2:
		b.WriteString(labelStyle.Render("Zoho Client Secret: "))
		b.WriteString(maskString(m.clientSecret))
		b.WriteString("█\n")
		if m.clientSecret == "" {
			b.WriteString("\nEnter your Zoho Client Secret\n")
		}

	case 3:
		b.WriteString(labelStyle.Render("Zoho Organization ID: "))
		b.WriteString(m.orgID)
		b.WriteString("█\n")
		if m.orgID == "" {
			b.WriteString("\nEnter your Zoho Organization ID\n")
		}

	case 4:
		b.WriteString(labelStyle.Render("Zoho Region: "))
		b.WriteString(m.region)
		b.WriteString("█\n")
		b.WriteString("\nRegions: com (default), eu, in, cn, au\n")

	case 5:
		if m.validating {
			b.WriteString("Validating credentials...\n")
		} else if m.valid {
			b.WriteString(successStyle.Render("✓ Configuration saved successfully!\n\n"))
			b.WriteString(fmt.Sprintf("Profile: %s\n", m.profileName))
			b.WriteString(fmt.Sprintf("Client ID: %s...%s\n", m.clientID[:10], m.clientID[len(m.clientID)-4:]))
			b.WriteString(fmt.Sprintf("Region: %s\n", m.region))
		} else if m.errorMsg != "" {
			b.WriteString(errorStyle.Render("✗ Error: "+m.errorMsg))
			b.WriteString("\n\nPress Enter to try again\n")
		}
	}

	b.WriteString("\nPress Ctrl+C to cancel\n")
	return b.String()
}

func (m *setupModel) handleEnter() (tea.Model, tea.Cmd) {
	switch m.step {
	case 0:
		if m.profileName == "" {
			m.profileName = "default"
		}
		m.step++
	case 1:
		if m.clientID == "" {
			return m, nil
		}
		m.step++
	case 2:
		if m.clientSecret == "" {
			return m, nil
		}
		m.step++
	case 3:
		if m.orgID == "" {
			return m, nil
		}
		m.step++
	case 4:
		if m.region == "" {
			m.region = "com"
		}
		m.step++
		m.validating = true
		return m, m.validateCredentials()
	case 5:
		if m.errorMsg != "" {
			m.step = 1
			m.errorMsg = ""
		}
	}
	return m, nil
}

func (m *setupModel) handleBackspace(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.step {
	case 0:
		if len(m.profileName) > 0 {
			m.profileName = m.profileName[:len(m.profileName)-1]
		}
	case 1:
		if len(m.clientID) > 0 {
			m.clientID = m.clientID[:len(m.clientID)-1]
		}
	case 2:
		if len(m.clientSecret) > 0 {
			m.clientSecret = m.clientSecret[:len(m.clientSecret)-1]
		}
	case 3:
		if len(m.orgID) > 0 {
			m.orgID = m.orgID[:len(m.orgID)-1]
		}
	case 4:
		if len(m.region) > 0 {
			m.region = m.region[:len(m.region)-1]
		}
	}
	return m, nil
}

func (m *setupModel) handleInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.step {
	case 0:
		m.profileName += string(msg.Runes)
	case 1:
		m.clientID += string(msg.Runes)
	case 2:
		m.clientSecret += string(msg.Runes)
	case 3:
		m.orgID += string(msg.Runes)
	case 4:
		m.region += string(msg.Runes)
	}
	return m, nil
}

func (m *setupModel) validateCredentials() tea.Cmd {
	return func() tea.Msg {
		profile := config.Profile{
			ClientID:     m.clientID,
			ClientSecret: m.clientSecret,
			OrgID:        m.orgID,
			Region:       m.region,
		}

		auth := api.NewAuth(&profile)
		if err := auth.ValidateCredentials(context.Background()); err != nil {
			return validationMsg{success: false, error: err.Error()}
		}

		cfg, _ := config.InitConfig()
		cfg.Profiles[m.profileName] = profile
		cfg.CurrentProfile = m.profileName

		if err := config.Save(cfg); err != nil {
			return validationMsg{success: false, error: err.Error()}
		}

		return validationMsg{success: true}
	}
}

type validationMsg struct {
	success bool
	error   string
}

func maskString(s string) string {
	if len(s) <= 8 {
		return strings.Repeat("*", len(s))
	}
	return s[:4] + strings.Repeat("*", len(s)-8) + s[len(s)-4:]
}

func RunSetup() error {
	p := tea.NewProgram(setupModel{
		step:    0,
		region:  "com",
	})

	model, err := p.Run()
	if err != nil {
		return err
	}

	if m, ok := model.(setupModel); ok && m.valid {
		fmt.Println("\nConfiguration complete! Try:")
		fmt.Printf("  zohodesk-cli tickets list --profile %s\n", m.profileName)
	}

	return nil
}