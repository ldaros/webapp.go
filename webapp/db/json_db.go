package db

import (
	"encoding/json"
	"errors"
	"io"
	"log-api/models"
	"os"
	"sort"
	"sync"
)

const dbFilePath = "./db.json"

var (
	jsonDBManager = NewManager()
	once          sync.Once
)

func init() {
	once.Do(func() {
		if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
			if file, err := os.Create(dbFilePath); err != nil {
				panic(err)
			} else {
				file.Close()
			}
		}
	})
}

type Schema struct {
	Logs       []models.Log `json:"logs"`
	LogsLastID int          `json:"logs_last_id"`
}

type Manager struct {
	mu sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) get() (Schema, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	file, err := os.Open(dbFilePath)
	if err != nil {
		return Schema{}, err
	}
	defer file.Close()

	var store Schema
	err = json.NewDecoder(file).Decode(&store)
	if errors.Is(err, io.EOF) {
		return Schema{}, nil
	}
	if err != nil {
		return Schema{}, err
	}

	return store, nil
}

func (m *Manager) set(store Schema) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dbFilePath, data, 0644)
}

type LogStoreJson struct {
	mu sync.RWMutex
}

func NewLogStoreJson() *LogStoreJson {
	return &LogStoreJson{}
}

func (s *LogStoreJson) Insert(log models.Log) (models.Log, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := jsonDBManager.get()
	if err != nil {
		return models.Log{}, err
	}

	store.LogsLastID++
	log.ID = store.LogsLastID
	store.Logs = append(store.Logs, log)

	err = jsonDBManager.set(store)
	if err != nil {
		return models.Log{}, err
	}

	return log, nil
}

func (s *LogStoreJson) Update(id int, values models.Log) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := jsonDBManager.get()
	if err != nil {
		return err
	}

	idx := sort.Search(len(store.Logs), func(i int) bool {
		return store.Logs[i].ID >= id
	})

	if idx < len(store.Logs) && store.Logs[idx].ID == id {
		store.Logs[idx] = values
		store.Logs[idx].ID = id
		return jsonDBManager.set(store)
	}

	return errors.New("log not found")
}

func (s *LogStoreJson) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := jsonDBManager.get()
	if err != nil {
		return err
	}

	idx := sort.Search(len(store.Logs), func(i int) bool {
		return store.Logs[i].ID >= id
	})

	if idx < len(store.Logs) && store.Logs[idx].ID == id {
		store.Logs = append(store.Logs[:idx], store.Logs[idx+1:]...)
		return jsonDBManager.set(store)
	}

	return errors.New("log not found")
}

func (s *LogStoreJson) Get(id int) (models.Log, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	store, err := jsonDBManager.get()
	if err != nil {
		return models.Log{}, err
	}

	idx := sort.Search(len(store.Logs), func(i int) bool {
		return store.Logs[i].ID >= id
	})

	if idx < len(store.Logs) && store.Logs[idx].ID == id {
		return store.Logs[idx], nil
	}

	return models.Log{}, errors.New("log not found")
}

func (s *LogStoreJson) GetAll() ([]models.Log, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	store, err := jsonDBManager.get()
	if err != nil {
		return nil, err
	}

	return store.Logs, nil
}
