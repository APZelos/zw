package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

// Item represents a selectable item in the list.
type Item struct {
	title string
	path  string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.path }
func (i Item) FilterValue() string { return i.title }

// Model is the Bubble Tea model for the worktree selector.
type Model struct {
	list     list.Model
	selected string
	quitting bool
}

// NewModel creates a new selector model with the given items.
func NewModel(items []Item) Model {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}

	l := list.New(listItems, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select a worktree"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)

	return Model{list: l}
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if item, ok := m.list.SelectedItem().(Item); ok {
				m.selected = item.path
			}
			m.quitting = true
			return m, tea.Quit

		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View implements tea.Model.
func (m Model) View() string {
	if m.quitting {
		return ""
	}
	return docStyle.Render(m.list.View())
}

// Selected returns the selected path.
func (m Model) Selected() string {
	return m.selected
}

// NewItem creates a new Item with the given title and path.
func NewItem(title, path string) Item {
	return Item{title: title, path: path}
}
