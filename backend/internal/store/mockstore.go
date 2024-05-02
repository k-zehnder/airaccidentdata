// MockStore simulates data store operations for unit testing.
// It enables isolated testing of controllers and other components by mocking database interactions,
// thus eliminating dependencies on the actual database during tests.
package store

import (
	"fmt"

	"github.com/computers33333/airaccidentdata/internal/models"
)

// MockStore simulates a data store for testing purposes.
type MockStore struct {
	Aircrafts  []*models.Aircraft
	Accidents  []*models.Accident
	QueryError error // Used to simulate database query errors
}

// NewMockStore creates a new instance of MockStore.
func NewMockStore(aircrafts []*models.Aircraft, accidents []*models.Accident, queryError error) *MockStore {
	return &MockStore{
		Aircrafts:  aircrafts,
		Accidents:  accidents,
		QueryError: queryError,
	}
}

// Aircraft-related methods

// SaveAircrafts saves aircraft data to the store.
func (ms *MockStore) SaveAircrafts(aircrafts []*models.Aircraft) error {
	if ms.QueryError != nil {
		return ms.QueryError
	}
	ms.Aircrafts = aircrafts
	return nil
}

// GetAircrafts retrieves aircraft data from the store.
func (ms *MockStore) GetAircrafts() ([]*models.Aircraft, error) {
	if ms.QueryError != nil {
		return nil, ms.QueryError
	}
	return ms.Aircrafts, nil
}

// Accident-related methods

// SaveAccidents saves accident data to the store.
func (ms *MockStore) SaveAccidents(accidents []*models.Accident) error {
	if ms.QueryError != nil {
		return ms.QueryError
	}
	ms.Accidents = accidents
	return nil
}

// GetAccidents retrieves accident data from the store.
func (ms *MockStore) GetAccidents() ([]*models.Accident, error) {
	if ms.QueryError != nil {
		return nil, ms.QueryError
	}
	return ms.Accidents, nil
}

// GetAccidentById retrieves a specific accident by ID from the store.
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
