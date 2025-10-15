package tasks

import "sync"

type Repository struct {
	mu     sync.Mutex
	tasks  []Task
	nextID int
}

func NewRepository() *Repository {
	return &Repository{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

func (r *Repository) GetAll() []Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.tasks
}

func (r *Repository) Create(task Task) Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	task.ID = r.nextID
	r.nextID++
	r.tasks = append(r.tasks, task)
	return task
}