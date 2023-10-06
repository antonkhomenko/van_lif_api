package storage

import (
	"fmt"
	"sort"
	"sync"
	"van-life-api/internal/models"
)

type Storage interface {
	Insert(van models.Van)
	Get(id int) (models.Van, error)
	Update(id int, van models.Van) error
	Delete(id int) error
}


type LocalStorage struct {
	data 	map[int]models.Van
	mx 		sync.Mutex
}

func NewLocalStorage() *LocalStorage {
	data := make(map[int]models.Van, len(models.Vans))
	for _, v := range models.Vans {
		data[v.Id] = v
	}
	return &LocalStorage{
		data: data,
	}
}

func(ls *LocalStorage) Insert(van models.Van) {
	ls.mx.Lock()
	ls.data[van.Id] = van
	ls.mx.Unlock()
}

func(ls *LocalStorage) Get(id int) (models.Van, error) {
	ls.mx.Lock()
	van, ok := ls.data[id]
	ls.mx.Unlock()
	if !ok {
		return models.Van{}, fmt.Errorf("van with such id: %d dosen't exist", id)
	}
	return van, nil
}

func(ls *LocalStorage) GetAll() []models.Van {
	ls.mx.Lock()
	result := make([]models.Van, len(ls.data))
	i := 0
	for _, v := range ls.data {
		result[i] = v
		i++
	}
	ls.mx.Unlock()
	sort.Slice(result, func(i, j int) bool {
		return result[i].Id < result[j].Id
	})
	return result
}


func(ls *LocalStorage) Update(id int, van models.Van) error {
	ls.mx.Lock()
	_, ok := ls.data[id]
	if !ok {
		return fmt.Errorf("van with such id: %d dosen't exist", id)
	}
	ls.data[id] = van
	ls.mx.Unlock()
	return nil
}

func(ls *LocalStorage) Delete(id int) error {
	ls.mx.Lock()
	_, ok := ls.data[id]
	if !ok {
		return fmt.Errorf("van with such id: %d dosen't exist", id)
	}
	delete(ls.data, id)
	ls.mx.Unlock()
	return nil
}

var LocalStrogaeInstance = NewLocalStorage()



