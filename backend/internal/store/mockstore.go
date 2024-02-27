package store

import (
	"fmt"

	"github.com/computers33333/airaccidentdata/internal/models"
)

type MockStore struct {
	Accidents  []*models.AircraftAccident
	QueryError error // Used to simulate datbase query errors
}

// NewMockStore initializes a MockStore with predefined accidents and potential errors.
// It's designed for setting up tests with controlled data and error handling.
func NewMockStore(accidents []*models.AircraftAccident, queryError error) *MockStore {
	return &MockStore{
		Accidents:  accidents,
		QueryError: queryError,
	}
}

// SaveArticles simulates storing accidents, returning a predefined error if set.
// On success, it updates the internal slice of accidents.
func (ms *MockStore) SaveAccidents(accidents []*models.AircraftAccident) error {
	if ms.QueryError != nil {
		return ms.QueryError
	}
	ms.Accidents = accidents
	return nil
}

// GetAccidents simulates fetching accidents, returning a predefined error if set.
// On success, it retrns a slice of pointers to the internal accdents.
func (ms *MockStore) GetAccidents(page, limit int) ([]*models.AircraftAccident, int, error) {
	if ms.QueryError != nil {
		return nil, 0, ms.QueryError
	}

	start := (page - 1) * limit
	end := start + limit
	if start >= len(ms.Accidents) {
		return nil, len(ms.Accidents), nil
	}
	if end > len(ms.Accidents) {
		end = len(ms.Accidents)
	}
	return ms.Accidents[start:end], len(ms.Accidents), nil
}

// GetAccidentByRegistration simulates fetching a single accident by registration number.
func (ms *MockStore) GetAccidentByRegistration(registrationNumber string) (*models.AircraftAccident, error) {
	for _, accident := range ms.Accidents {
		if accident.RegistrationNumber == registrationNumber {
			return accident, nil
		}
	}
	return nil, fmt.Errorf("no accident found with registration number: %s", registrationNumber)
}
