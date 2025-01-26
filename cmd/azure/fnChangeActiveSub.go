package azure

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// var docStyle = lipgloss.NewStyle().Margin(1, 2)
var (
	docStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type delegateKeyMap struct {
	choose key.Binding
	remove key.Binding
}

type item struct {
	title, subId, subName, tenantId, tenantName string
	isDefault                                   bool
}

// type menu struct {
// 	options list.Model
// }

func (i item) Title() string { return i.title }
func (i item) Description() string {
	return "Sub: " + i.subId + "(" + i.subName + ") - Tenant: " + i.tenantId + " (" + i.tenantName + ")"
}
func (i item) FilterValue() string { return i.title + i.subId + i.tenantName }
func (i item) IsDefault() bool     { return i.isDefault }

type model struct {
	list         list.Model
	azureProfile AzureProfile
	choice       AzureProfileSubscription
	choiceString string
	cursor       int
}

func (m model) Init() tea.Cmd {
	return nil
}

// var selectedSubscription int

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// fmt.Println(msg)
	switch msg := msg.(type) {
	// case tea.KeyMsg:
	// 	if msg.String() == "ctrl+c" {
	// 		return m, tea.Quit
	// 	}
	// 	if msg.String() == "enter" {
	// 		// selectedSubscription = m.list.Index()
	// 		return m, tea.Quit
	// 	}
	// case tea.WindowSizeMsg:
	// 	h, v := docStyle.GetFrameSize()
	// 	m.list.SetSize(msg.Width-h, msg.Height-v)
	// }
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.

			m.choice = m.azureProfile.Subscriptions[m.cursor]
			m.choiceString = m.choice.Name + " - " + m.choice.ID + " - Tenant: " + m.choice.TenantDisplayName
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.azureProfile.Subscriptions) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.azureProfile.Subscriptions) - 1
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	// fmt.Println(m.list)
	return m, cmd
}

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		if msg.String() == "ctrl+c" {
// 			return m, tea.Quit
// 		}
// 		if msg.String() == "enter" {
// 			// selectedSubscription = m.list.Index()
// 			return m, tea.Quit
// 		}
// 	case tea.WindowSizeMsg:
// 		h, v := docStyle.GetFrameSize()
// 		m.list.SetSize(msg.Width-h, msg.Height-v)
// 	}

// 	var cmd tea.Cmd
// 	m.list, cmd = m.list.Update(msg)
// 	return m, cmd
// }

func (m model) View() string {
	return docStyle.Render(m.list.View())
	// s := strings.Builder{}
	// s.WriteString("What kind of Bubble Tea would you like to order?\n\n")

	// for i := 0; i < len(m.subs); i++ {
	// 	if m.cursor == i {
	// 		s.WriteString("(•) ")
	// 	} else {
	// 		s.WriteString("( ) ")
	// 	}
	// 	// jsonStr
	// 	// s.WriteString(m.subs[i])
	// 	s.WriteString("\n")
	// }
	// s.WriteString("\n(press q to quit)\n")

	// return s.String()
}

// type menuItem struct {
// 	title    string
// 	desc     string
// 	isActive bool
// }

func ChangeActiveSub(azProfile AzureProfile) {
	subs := azProfile.Subscriptions
	var items []list.Item

	for _, sub := range subs {
		if sub.IsDefault {
			items = append(items, item{
				title:      sub.Name + " (Active)",
				subId:      sub.ID,
				tenantId:   sub.TenantID,
				tenantName: sub.TenantDisplayName,
			})
		}
	}

	for _, sub := range subs {
		if !sub.IsDefault {
			items = append(items, item{
				title:      sub.Name,
				subId:      sub.ID,
				tenantId:   sub.TenantID,
				tenantName: sub.TenantDisplayName,
			})
		}
	}

	// m := model{list: list.New(items, newItemDelegate, 0, 0), subs: subs}
	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0), azureProfile: azProfile}
	m.list.Title = "Select Active Azure Subscription"

	p := tea.NewProgram(m, tea.WithAltScreen())
	// if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
	// 	fmt.Println("Error running program:", err)
	// 	os.Exit(1)
	// }

	// curs := m.selectedActiveSub

	// fmt.Println(curs)

	// fmt.Println(selectedSubscription)
	// fmt.Println(m.list.Index())
	// selectedSub := m.list.SelectedItem()
	// fmt.Println(m)

	// jsonStr, _ := json.MarshalIndent(m.list.Title, "", "  ")
	// fmt.Println(string(jsonStr))

	// fmt.Println(m.list.SelectedItem())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	curs := m.choiceString

	fmt.Println(curs)
	var profileUpdated AzureProfile
	profileUpdated.InstallationID = azProfile.InstallationID

	for _, sub := range azProfile.Subscriptions {
		_ = sub
	}
}

// func GetSubsMapping(cldConfigOpts *lib.CldConfigOptions) *azure.AzureProfile {
// 	_, _, cachePath := lib.InitConfig(cldConfigOpts)

// 	cacheFile := cachePath + "/azsubs.json"
// 	var cacheFileData AzureProfile

// 	if _, err := os.Stat(cacheFile); err != nil {
// 		return nil
// 	}
// 	// fmt.Println(cachePath)
// 	// jsonStr, err := json.MarshalIndent(tokenData, "", "  ")
// 	// lib.CheckFatalError(err)

// 	// os.WriteFile(cachePath+"/azsubs.json", jsonStr, os.ModePerm)

// 	// func ListAllAuthenticatedSubscriptions(tokens *lib.AllTenantTokens) TenantList {
// 	// 	// allSubscriptions := make(map[string]string)
// 	// 	allTenantSubs := TenantList{}

// 	// 	for _, token := range *tokens {
// 	// 		subs, err := azure.ListSubscriptions(token)
// 	// 		lib.CheckFatalError(err)
// 	// 		allTenantSubs[token.TenantName] = make(map[string]string)

// 	// 		for _, sub := range subs {
// 	// 			allTenantSubs[token.TenantName][sub.DisplayName] = sub.SubscriptionID
// 	// 		}
// 	// 	}
// 	// 	return allTenantSubs
// 	// }
// }

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(item); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				return m.NewStatusMessage(statusMessageStyle("You chose " + title))

			case key.Matches(msg, keys.remove):
				index := m.Index()
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.remove.SetEnabled(false)
				}
				return m.NewStatusMessage(statusMessageStyle("Deleted " + title))
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose, keys.remove}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.remove,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.remove,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
		remove: key.NewBinding(
			key.WithKeys("x", "backspace"),
			key.WithHelp("x", "delete"),
		),
	}
}
