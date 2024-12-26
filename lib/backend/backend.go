package backend

import (
	"log"
	"time"

	"github.com/beadex/espresso/lib/database"
	"github.com/beadex/espresso/lib/models"
	"github.com/reactivex/rxgo/v2"
)

type Backend struct {
	eventStream chan rxgo.Item
}

// Initialize initializes the backend and sets up the reactive event stream.
// Initialize initializes the backend and sets up the reactive event stream.
func Initialize() *Backend {
	database.Init() // Ensure the database is initialized

	eventStream := make(chan rxgo.Item)
	b := &Backend{eventStream: eventStream}

	observable := rxgo.FromChannel(eventStream)

	// Process events
	go observable.ForEach(
		func(item interface{}) {
			switch event := item.(type) {
			case AddTaskEvent:
				handleAddTask(event)
			case UpdateTaskEvent:
				handleUpdateTask(event)
			case DeleteTaskEvent:
				handleDeleteTask(event)
			default:
				log.Printf("Unhandled event type: %T", event)
			}
		},
		func(err error) {
			log.Printf("Error in task stream: %v", err)
		},
		func() {
			log.Println("Task stream completed.")
		},
	)

	log.Println("Backend initialized.")
	return b
}

// GetAllTasks fetches all tasks from the database.
func (b *Backend) GetAllTasks() ([]models.Task, error) {
	return database.GetAllTasks()
}

// Event types
type AddTaskEvent struct {
	Name        string
	DueDate     time.Time
	IsRecurring bool
}

type UpdateTaskEvent struct {
	Task models.Task
}

type DeleteTaskEvent struct {
	ID int
}

// Handle Add Task Event
func handleAddTask(event AddTaskEvent) {
	task := models.Task{
		Name:        event.Name,
		DueDate:     event.DueDate,
		IsRecurring: event.IsRecurring,
	}
	id, err := database.InsertTask(task)
	if err != nil {
		log.Printf("Failed to add task: %v", err)
		return
	}
	log.Printf("Task added: %+v (ID: %d)", task, id)
}

// Handle Update Task Event
func handleUpdateTask(event UpdateTaskEvent) {
	err := database.UpdateTask(event.Task)
	if err != nil {
		log.Printf("Failed to update task ID %d: %v", event.Task.ID, err)
		return
	}
	log.Printf("Task updated: %+v", event.Task)
}

// Handle Delete Task Event
func handleDeleteTask(event DeleteTaskEvent) {
	err := database.DeleteTask(event.ID)
	if err != nil {
		log.Printf("Failed to delete task ID %d: %v", event.ID, err)
		return
	}
	log.Printf("Task deleted: ID %d", event.ID)
}

// AddTask emits an AddTaskEvent to the event stream.
func (b *Backend) AddTask(name string, dueDate time.Time, isRecurring bool) {
	b.eventStream <- rxgo.Of(AddTaskEvent{Name: name, DueDate: dueDate, IsRecurring: isRecurring})
}

// UpdateTask emits an UpdateTaskEvent to the event stream.
func (b *Backend) UpdateTask(task models.Task) {
	b.eventStream <- rxgo.Of(UpdateTaskEvent{Task: task})
}

// DeleteTask emits a DeleteTaskEvent to the event stream.
func (b *Backend) DeleteTask(id int) {
	b.eventStream <- rxgo.Of(DeleteTaskEvent{ID: id})
}
