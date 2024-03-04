package store

import (
	"fmt"

	"github.com/computers33333/airaccidentdata/internal/models"
)

type MockStore struct {
	Aircrafts  []*models.Aircraft
	Accidents  []*models.AircraftAccident
	QueryError error // Used to simulate database query errors
}

func NewMockStore(aircrafts []*models.Aircraft, accidents []*models.AircraftAccident, queryError error) *MockStore {
	return &MockStore{
		Aircrafts:  aircrafts,
		Accidents:  accidents,
		QueryError: queryError,
	}
}

func (ms *MockStore) SaveAircrafts(aircrafts []*models.Aircraft) error {
	if ms.QueryError != nil {
		return ms.QueryError
	}
	ms.Aircrafts = aircrafts
	return nil
}

func (ms *MockStore) SaveAccidents(accidents []*models.AircraftAccident) error {
	if ms.QueryError != nil {
		return ms.QueryError
	}
	ms.Accidents = accidents
	return nil
}

func (ms *MockStore) GetAircrafts() ([]*models.Aircraft, error) {
	if ms.QueryError != nil {
		return nil, ms.QueryError
	}
	return ms.Aircrafts, nil
}

func (ms *MockStore) GetAccidents() ([]*models.AircraftAccident, error) {
	if ms.QueryError != nil {
		return nil, ms.QueryError
	}
	return ms.Accidents, nil
}

func (ms *MockStore) GetAircraftWithAccidents(registrationNumber string) (*models.Aircraft, error) {
	var aircraftWithAccidents *models.Aircraft

	for _, aircraft := range ms.Aircrafts {
		if aircraft.RegistrationNumber == registrationNumber {
			aircraftWithAccidents = aircraft
			break
		}
	}

	if aircraftWithAccidents == nil {
		return nil, fmt.Errorf("aircraft with registration number %s not found", registrationNumber)
	}

	var aircraftAccidents []*models.AircraftAccident
	for _, accident := range ms.Accidents {
		if accident.AircraftID == aircraftWithAccidents.ID {
			aircraftAccidents = append(aircraftAccidents, accident)
		}
	}

	aircraftWithAccidents.Accidents = aircraftAccidents

	return aircraftWithAccidents, nil
}
