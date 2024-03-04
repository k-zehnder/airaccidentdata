package store

import (
	"errors"
	"testing"

	"github.com/computers33333/airaccidentdata/internal/models"
)

func TestMockStore_GetAccidents(t *testing.T) {
	expectedAccidents := []*models.AircraftAccident{
		{ID: 1},
	}
	mockStore := NewMockStore(nil, expectedAccidents, nil)
	accidents, err := mockStore.GetAccidents()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(accidents) != len(expectedAccidents) {
		t.Fatalf("Expected %d accidents, got %d", len(expectedAccidents), len(accidents))
	}
	for i, accident := range accidents {
		if accident.ID != expectedAccidents[i].ID {
			t.Errorf("Expected ID %d, got %d", expectedAccidents[i].ID, accident.ID)
		}
	}
}

func TestMockStore_SaveAccidents(t *testing.T) {
	mockStore := NewMockStore(nil, nil, nil)
	accidents := []*models.AircraftAccident{
		{ID: 2},
	}
	err := mockStore.SaveAccidents(accidents)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(mockStore.Accidents) != len(accidents) {
		t.Fatalf("Expected %d accidents, got %d", len(accidents), len(mockStore.Accidents))
	}
	for i, accident := range accidents {
		if accident.ID != accidents[i].ID {
			t.Fatalf("Expected ID %d, got %d", accidents[i].ID, accident.ID)
		}
	}
}

func TestMockStore_GetAircrafts(t *testing.T) {
	expectedAircrafts := []*models.Aircraft{
		{RegistrationNumber: "ABC123"},
	}
	mockStore := NewMockStore(expectedAircrafts, nil, nil)
	aircrafts, err := mockStore.GetAircrafts()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(aircrafts) != len(expectedAircrafts) {
		t.Fatalf("Expected %d aircrafts, got %d", len(expectedAircrafts), len(aircrafts))
	}
	for i, aircraft := range aircrafts {
		if aircraft.RegistrationNumber != expectedAircrafts[i].RegistrationNumber {
			t.Errorf("Expected Registration Number %s, got %s", expectedAircrafts[i].RegistrationNumber, aircraft.RegistrationNumber)
		}
	}
}

func TestMockStore_SaveAircrafts(t *testing.T) {
	mockStore := NewMockStore(nil, nil, nil)
	aircrafts := []*models.Aircraft{
		{RegistrationNumber: "XYZ789"},
	}
	err := mockStore.SaveAircrafts(aircrafts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(mockStore.Aircrafts) != len(aircrafts) {
		t.Fatalf("Expected %d aircrafts, got %d", len(aircrafts), len(mockStore.Aircrafts))
	}
	for i, aircraft := range aircrafts {
		if aircraft.RegistrationNumber != aircrafts[i].RegistrationNumber {
			t.Fatalf("Expected registration number %s, got %s", aircrafts[i].RegistrationNumber, aircraft.RegistrationNumber)
		}
	}
}

func TestMockStore_GetAircrafts_Error(t *testing.T) {
	expectedError := errors.New("query error")
	mockStore := NewMockStore(nil, nil, expectedError)
	_, err := mockStore.GetAircrafts()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

// GetAircraftWithAccidents
func TestMockStore_GetAccidents_Error(t *testing.T) {
	expectedError := errors.New("query error")
	mockStore := NewMockStore(nil, nil, expectedError)
	_, err := mockStore.GetAccidents()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestMockStore_SaveAircrafts_Error(t *testing.T) {
	simulatedError := errors.New("simulated save error")
	mockStore := NewMockStore(nil, nil, simulatedError)
	err := mockStore.SaveAircrafts([]*models.Aircraft{{RegistrationNumber: "ABCD"}})
	if err != simulatedError {
		t.Errorf("Expected error '%v', got '%v'", simulatedError, err)
	}
}

func TestMockStore_SaveAccidents_Error(t *testing.T) {
	simulatedError := errors.New("simulated save error")
	mockStore := NewMockStore(nil, nil, simulatedError)
	err := mockStore.SaveAccidents([]*models.AircraftAccident{{ID: 3}})
	if err != simulatedError {
		t.Errorf("Expected error '%v', got '%v'", simulatedError, err)
	}
}

// GetAircraftWithAccidents
func TestMockStore_GetAircraftWithAccidents(t *testing.T) {
	expectedAircraft := &models.Aircraft{
		RegistrationNumber: "ABC123",
		ID:                 1,
	}
	expectedAccidents := []*models.AircraftAccident{
		{ID: 1, AircraftID: 1},
		{ID: 2, AircraftID: 1},
	}

	mockStore := NewMockStore([]*models.Aircraft{expectedAircraft}, expectedAccidents, nil)

	aircraftWithAccidents, err := mockStore.GetAircraftWithAccidents("ABC123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if aircraftWithAccidents == nil {
		t.Fatal("Expected non-nil aircraft, got nil")
	}

	if len(aircraftWithAccidents.Accidents) != len(expectedAccidents) {
		t.Fatalf("Expected %d accidents, got %d", len(expectedAccidents), len(aircraftWithAccidents.Accidents))
	}

	for i, accident := range aircraftWithAccidents.Accidents {
		if accident.ID != expectedAccidents[i].ID {
			t.Errorf("Expected accident ID %d, got %d", expectedAccidents[i].ID, accident.ID)
		}
		if accident.AircraftID != expectedAircraft.ID {
			t.Errorf("Expected accident AircraftID %d, got %d", expectedAircraft.ID, accident.AircraftID)
		}
	}
}
