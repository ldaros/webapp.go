package db

import "log-api/models"

type LogStore interface {
	Insert(models.Log) (models.Log, error)
	Update(id int, values models.Log) error
	Delete(id int) error
	Get(id int) (models.Log, error)
	GetAll() ([]models.Log, error)
}
