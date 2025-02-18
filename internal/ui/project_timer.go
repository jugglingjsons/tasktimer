package ui

import (
	"time"

	"github.com/caarlos0/tasktimer/internal/model"
	tea "github.com/charmbracelet/bubbletea"
)

type projectTimerModel struct {
	tasks []model.Task
}

func (m projectTimerModel) Init() tea.Cmd {
	return nil
}

func (m projectTimerModel) Update(msg tea.Msg) (projectTimerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case projectTimerUpdateMsg:
		m.tasks = msg.tasks
	}
	return m, nil
}

func (m projectTimerModel) View() string {
	return secondaryForeground.Render("total: ") +
		activeForegroundBold.Render(sumTasksTimes(m.tasks, time.Time{}).Round(time.Second).String()) +
		separator +
		secondaryForeground.Render("today: ") +
		activeForegroundBold.Render(sumTasksTimes(m.tasks, todayAtMidnight()).Round(time.Second).String())
}

// msgs and cmds

type projectTimerUpdateMsg struct {
	tasks []model.Task
}

func updateProjectTimerCmd(tasks []model.Task) tea.Cmd {
	return func() tea.Msg {
		return projectTimerUpdateMsg{tasks}
	}
}

func sumTasksTimes(tasks []model.Task, since time.Time) time.Duration {
	td := time.Duration(0)
	for _, t := range tasks {
		for _, d := range t.Durations {

			if d.StartAt.Before(since) {
				continue
			}

			z := d.EndAt
			if z.IsZero() {
				z = time.Now()
			}
			td += z.Sub(d.StartAt)
		}
	}
	return td
}

func todayAtMidnight() time.Time {
	return time.Now().Truncate(time.Hour * 24)
}
