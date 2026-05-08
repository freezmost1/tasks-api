package storage

import (
	"sync"
	"tasks-api/internal/models"
)

type memoryStorage struct {
	tasks  map[int]models.Task
	nextID int
	mu     sync.RWMutex
}

func NewMemoryStorage() Storage {
	return &memoryStorage{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (s *memoryStorage) List() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var tasks []models.Task
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *memoryStorage) Create(task models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = s.nextID
	s.tasks[task.ID] = task
	s.nextID++
	return task, nil
}

func (s *memoryStorage) Get(id int) (models.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	return task, exists
}

func (s *memoryStorage) Update(id int, task models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return models.Task{}, nil
	}

	task.ID = id
	s.tasks[id] = task
	return task, nil
}

func (s *memoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tasks, id)
	return nil
}
