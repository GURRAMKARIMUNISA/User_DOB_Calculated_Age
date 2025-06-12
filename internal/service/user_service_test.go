package service_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/GurramKarimunisa/go-user-api/internal/service" // Adjust import path
)

func TestCalculateAge(t *testing.T) {
	// Test case 1: Birthday already passed this year
	dob1 := time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC)
	// Assuming current date is after May 10th
	expectedAge1 := time.Now().Year() - 1990
	if time.Now().Month() < 5 || (time.Now().Month() == 5 && time.Now().Day() < 10) {
		expectedAge1--
	}
	assert.Equal(t, expectedAge1, service.CalculateAge(dob1), "Age should be correct when birthday passed")

	// Test case 2: Birthday yet to come this year
	// Assuming current date is before Nov 15th
	dob2 := time.Date(1985, 11, 15, 0, 0, 0, 0, time.UTC)
	expectedAge2 := time.Now().Year() - 1985
	if time.Now().Month() < 11 || (time.Now().Month() == 11 && time.Now().Day() < 15) {
		expectedAge2--
	}
	assert.Equal(t, expectedAge2, service.CalculateAge(dob2), "Age should be correct when birthday not yet passed")

	// Test case 3: Birthday today
	today := time.Now()
	dob3 := time.Date(today.Year()-30, today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 30, service.CalculateAge(dob3), "Age should be correct on birthday")

	// Test case 4: Leap year baby
	dob4 := time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC)
	expectedAge4 := time.Now().Year() - 2000
	if time.Now().Month() < 2 || (time.Now().Month() == 2 && time.Now().Day() < 29) {
		expectedAge4--
	}
	assert.Equal(t, expectedAge4, service.CalculateAge(dob4), "Age should be correct for leap year baby")
}