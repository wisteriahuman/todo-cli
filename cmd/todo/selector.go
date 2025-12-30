package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wisteriahuman/todo-cli/internal/domain/entity"
)

var (
	cursorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	selectedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	unselectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	checkStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	uncheckStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	markedStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
)

// Single select model
type selectorModel struct {
	todos    []*entity.Todo
	cursor   int
	selected *entity.Todo
	quit     bool
}

func newSelectorModel(todos []*entity.Todo) selectorModel {
	return selectorModel{
		todos: todos,
	}
}

func (m selectorModel) Init() tea.Cmd {
	return nil
}

func (m selectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.todos[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m selectorModel) View() string {
	s := "タスクを選択してください (j/k: 移動, Enter: 決定, q: キャンセル)\n\n"

	for i, todo := range m.todos {
		cursor := "  "
		if m.cursor == i {
			cursor = "> "
		}

		status := uncheckStyle.Render("[ ]")
		if todo.Completed {
			status = checkStyle.Render("[x]")
		}

		line := fmt.Sprintf("%s %s %s", status, todo.ID[:8], todo.Title)
		if m.cursor == i {
			line = selectedStyle.Render(line)
		} else {
			line = unselectedStyle.Render(line)
		}

		s += cursor + line + "\n"
	}

	return s
}

func selectTodo(todos []*entity.Todo) (*entity.Todo, error) {
	if len(todos) == 0 {
		return nil, fmt.Errorf("タスクがありません")
	}

	m := newSelectorModel(todos)
	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		return nil, err
	}

	result := finalModel.(selectorModel)
	if result.quit {
		return nil, fmt.Errorf("キャンセルされました")
	}

	return result.selected, nil
}

// Multi select model
type multiSelectorModel struct {
	todos    []*entity.Todo
	cursor   int
	marked   map[int]bool
	quit     bool
}

func newMultiSelectorModel(todos []*entity.Todo) multiSelectorModel {
	return multiSelectorModel{
		todos:  todos,
		marked: make(map[int]bool),
	}
}

func (m multiSelectorModel) Init() tea.Cmd {
	return nil
}

func (m multiSelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}
		case " ":
			m.marked[m.cursor] = !m.marked[m.cursor]
		case "a":
			// Select all
			allSelected := len(m.marked) == len(m.todos)
			if allSelected {
				m.marked = make(map[int]bool)
			} else {
				for i := range m.todos {
					m.marked[i] = true
				}
			}
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m multiSelectorModel) View() string {
	s := "タスクを選択してください (j/k: 移動, Space: 選択, a: 全選択, Enter: 決定, q: キャンセル)\n\n"

	for i, todo := range m.todos {
		cursor := "  "
		if m.cursor == i {
			cursor = "> "
		}

		mark := "○"
		if m.marked[i] {
			mark = markedStyle.Render("●")
		}

		status := uncheckStyle.Render("[ ]")
		if todo.Completed {
			status = checkStyle.Render("[x]")
		}

		line := fmt.Sprintf("%s %s %s %s", mark, status, todo.ID[:8], todo.Title)
		if m.cursor == i {
			line = cursorStyle.Render(line)
		} else if m.marked[i] {
			line = markedStyle.Render(line)
		} else {
			line = unselectedStyle.Render(line)
		}

		s += cursor + line + "\n"
	}

	count := 0
	for _, v := range m.marked {
		if v {
			count++
		}
	}
	s += fmt.Sprintf("\n%d件選択中", count)

	return s
}

func selectTodos(todos []*entity.Todo) ([]*entity.Todo, error) {
	if len(todos) == 0 {
		return nil, fmt.Errorf("タスクがありません")
	}

	m := newMultiSelectorModel(todos)
	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		return nil, err
	}

	result := finalModel.(multiSelectorModel)
	if result.quit {
		return nil, fmt.Errorf("キャンセルされました")
	}

	var selected []*entity.Todo
	for i, marked := range result.marked {
		if marked {
			selected = append(selected, result.todos[i])
		}
	}

	if len(selected) == 0 {
		return nil, fmt.Errorf("タスクが選択されていません")
	}

	return selected, nil
}
