package gui

import (
	"fmt"
	"log"
	"time"

	"github.com/beadex/espresso/lib/backend"
	"github.com/rivo/tview"
)

type GUI struct {
	app      *tview.Application
	backend  *backend.Backend
	taskList *tview.List
	form     *tview.Form
}

// Initialize initializes the GUI.
func Initialize(b *backend.Backend) *GUI {
	app := tview.NewApplication()
	taskList := tview.NewList()
	form := tview.NewForm()

	gui := &GUI{
		app:      app,
		backend:  b,
		taskList: taskList,
		form:     form,
	}

	gui.setupUI()
	return gui
}

// setupUI sets up the UI layout and functionality.
func (g *GUI) setupUI() {
	// Task List on the left
	g.taskList.SetBorder(true).SetTitle("Tasks")
	g.updateTaskList()

	// Input Form on the right
	g.form = tview.NewForm().
		AddInputField("Task Name", "", 20, nil, nil).
		AddInputField("Due Date (YYYY-MM-DD)", "", 10, nil, nil).
		AddCheckbox("Recurring", false, nil).
		AddButton("Add Task", func() {
			taskName := g.form.GetFormItem(0).(*tview.InputField).GetText()
			dueDateStr := g.form.GetFormItem(1).(*tview.InputField).GetText()
			recurring := g.form.GetFormItem(2).(*tview.Checkbox).IsChecked()

			dueDate, err := time.Parse("2006-01-02", dueDateStr)
			if err != nil {
				log.Printf("Invalid date format: %v", err)
				return
			}

			g.backend.AddTask(taskName, dueDate, recurring)
			g.updateTaskList()
		}).
		AddButton("Quit", func() {
			g.app.Stop()
		})

	g.form.SetBorder(true).SetTitle("Task Details")
	g.form.SetFocus(0) // Set focus to the first input field

	// Layout setup
	layout := tview.NewFlex().
		AddItem(g.taskList, 0, 1, true).
		AddItem(g.form, 0, 2, false)

	// Set layout as the root
	g.app.SetRoot(layout, true)
}

// updateTaskList refreshes the task list.
func (g *GUI) updateTaskList() {
	g.taskList.Clear()
	tasks, err := g.backend.GetAllTasks()
	if err != nil {
		log.Printf("Failed to fetch tasks: %v", err)
		return
	}

	for _, task := range tasks {
		taskStr := fmt.Sprintf("%s (Due: %s)", task.Name, task.DueDate.Format("2006-01-02"))
		g.taskList.AddItem(taskStr, "", 0, nil)
	}
}

// Run starts the GUI application.
func (g *GUI) Run() error {
	// Set the root of the application and run it
	if err := g.app.Run(); err != nil {
		return err
	}
	return nil
}
