package store

import (
	"fmt"

	"github.com/computers33333/airaccidentdata/internal/models"
)

type MockStore struct {
	Aircrafts  []*models.Aircraft
	Accidents  []*models.Accident
	QueryError error // Used to simulate database query errors
}

func NewMockStore(aircrafts []*models.Aircraft, accidents []*models.Accident, queryError error) *MockStore {
	return &MockStore{
		Aircrafts:  aircrafts,
		Accidents:  accidents,
		QueryError: queryError,
	}
}

// Aircraft-related methods

func (ms *MockStore) SaveAircrafts(aircrafts []*models.Aircraft) error {
	if ms.QueryError != nil {
		return ms.QueryError
	}
	ms.Aircrafts = aircrafts
	return nil
}

func (ms *MockStore) GetAircrafts() ([]*models.Aircraft, error) {
	if ms.QueryError != nil {
		return nil, ms.QueryError
	}
	return ms.Aircrafts, nil
}

// Accident-related methods

func (ms *MockStore) SaveAccidents(accidents []*models.Accident) error {
	if ms.QueryError != nil {
		return ms.QueryError
	}
	ms.Accidents = accidents
	return nil
}

func (ms *MockStore) GetAccidents() ([]*models.Accident, error) {
	if ms.QueryError != nil {
		return nil, ms.QueryError
	}
	return ms.Accidents, nil
}

func (ms *MockStore) GetAccidentById(id int) (*models.Accident, error) {
	if ms.QueryError != nil {
		return nil, ms.QueryError
	}
	for _, accident := range ms.Accidents {
		if accident.ID == id {
			return accident, nil
		}
	}
	return nil, fmt.Errorf("accident not found")
}
