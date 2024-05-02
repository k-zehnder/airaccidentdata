// MockStore tests validate the mock implementation used for isolating and testing
// database-related logic in controllers and other components without a real database connection.
package store

import (
	"errors"
	"testing"

	"github.com/computers33333/airaccidentdata/internal/models"
)

// Aircraft-related tests

// TestMockStore_GetAircrafts tests the GetAircrafts method.
func TestMockStore_GetAircrafts(t *testing.T) {
	expectedAircrafts := []*models.Aircraft{{RegistrationNumber: "ABC123"}}
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

// TestMockStore_SaveAircrafts tests the SaveAircrafts method of the MockStore.
func TestMockStore_SaveAircrafts(t *testing.T) {
	mockStore := NewMockStore(nil, nil, nil)
	aircrafts := []*models.Aircraft{{RegistrationNumber: "XYZ789"}}

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

// TestMockStore_GetAircrafts_Error tests error handling in the GetAircrafts method of the MockStore.
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

// Accident-related tests

// TestMockStore_GetAccidents tests the GetAccidents method of the MockStore.
func TestMockStore_GetAccidents(t *testing.T) {
	expectedAccidents := []*models.Accident{{ID: 1}}
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

// TestMockStore_SaveAccidents tests the SaveAccidents method of the MockStore.
func TestMockStore_SaveAccidents(t *testing.T) {
	mockStore := NewMockStore(nil, nil, nil)
	accidents := []*models.Accident{{ID: 2}}

	err := mockStore.SaveAccidents(accidents)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(mockStore.Accidents) != len(accidents) {
		t.Fatalf("Expected %d accidents, got %d", len(accidents), len(mockStore.Accidents))
	}
	for i, accident := range mockStore.Accidents {
		if accident.ID != accidents[i].ID {
			t.Fatalf("Expected ID %d, got %d", accidents[i].ID, accident.ID)
		}
	}
}

// TestMockStore_GetAccidentById tests fetching a single accident by its ID.
func TestMockStore_GetAccidentById(t *testing.T) {
	expectedAccidents := []*models.Accident{
		{ID: 1, AircraftID: 1},
		{ID: 2, AircraftID: 2},
	}
	mockStore := NewMockStore(nil, expectedAccidents, nil)

	tests := []struct {
		name     string
		id       int
		wantErr  bool
		expected *models.Accident
	}{
		{"Accident exists", 1, false, &models.Accident{ID: 1, AircraftID: 1}},
		{"Accident does not exist", 3, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accident, err := mockStore.GetAccidentById(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccidentById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && accident.ID != tt.expected.ID {
				t.Errorf("GetAccidentById() got = %v, want %v", accident.ID, tt.expected.ID)
			}
		})
	}
}

// TestMockStore_GetAccidents_Error tests error handling in the GetAccidents method of the MockStore.
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

// TestMockStore_SaveAccidents_Error tests error handling in the SaveAccidents method of the MockStore.
func TestMockStore_SaveAccidents_Error(t *testing.T) {
	simulatedError := errors.New("simulated save error")
	mockStore := NewMockStore(nil, nil, simulatedError)
	err := mockStore.SaveAccidents([]*models.Accident{{ID: 3}})
	if err != simulatedError {
		t.Errorf("Expected error '%v', got '%v'", simulatedError, err)
	}
}
