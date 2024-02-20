package store

import (
	"testing"

	"github.com/computers33333/airaccidentdata/internal/models"
)

// TestMockStore_Getaccidents validates the retrieval of incidents from MockStore.
// It ensures the method returns the correct set of accidents and checks for the absence of errors.
func TestMockStore_GetAccidents(t *testing.T) {
	// Prepare expected accidents for the mock store.
	expectedAccidents := []*models.AircraftAccident{
		{RegistrationNumber: "1234"},
	}
	mockStore := NewMockStore(expectedAccidents, nil)

	// Call Getaccidents to retrieve incidents from the mock store.
	accidents, err := mockStore.GetAccidents()

	// Ensure no error is returned.
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify that the number of accidents returned matches the expected number.
	if len(accidents) != len(expectedAccidents) {
		t.Fatalf("Expected %d accidents, got %d", len(expectedAccidents), len(accidents))
	}

	// Check if the returned accidents match the expected ones.
	for i, incident := range accidents {
		if incident.RegistrationNumber != expectedAccidents[i].RegistrationNumber {
			t.Errorf("Expected title %s, got %s", expectedAccidents[i].RegistrationNumber, incident.RegistrationNumber)
		}
	}
}

// TODO
// TestMockStore_Saveaccidents checks the functionality of adding incidents to MockStore.
// It tests whether the accidents are correctly saved in the store.
func TestMockStore_SaveAccidents(t *testing.T) {
	// Initialize a mock store without any pre-existing accidents
	mockStore := NewMockStore(nil, nil)

	// Define the accidentts to be saved in the mock store.
	accidents := []*models.AircraftAccident{
		{RegistrationNumber: "1234A"},
		{RegistrationNumber: "1234B"},
	}

	// Attempt to save accidents in the mock store.
	err := mockStore.SaveAccidents(accidents)

	// Ensure that no error is returned from SaveAccidents.
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify that the mockStore contains the expected number of accidents.
	if len(mockStore.Accidents) != len(accidents) {
		t.Fatalf("Expected %d accidents, got %d", len(accidents), len(mockStore.Accidents))
	}

	// Verify that the mock store now contains the expected number of accidents.
	for i, accident := range accidents {
		if accident.RegistrationNumber != accidents[i].RegistrationNumber {
			t.Fatalf("Expected registration number %s, got %s", accidents[i].RegistrationNumber, accident.RegistrationNumber)
		}
	}
}
