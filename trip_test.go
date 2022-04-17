package main

import (
	"testing"
)

func TripNameTest(t *testing.T) {
	trip := Trip{
		Name: "holidays",
	}
	if trip.Name != "holidays" {
		t.Fatalf("TripNameTest fail because Name of Trip is not equal to %s but to %s", "holidays", trip.Name)
	}
}
