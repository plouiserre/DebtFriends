package main

import (
	"testing"
)

func TestFriendName(t *testing.T) {
	friend := Friend{
		FirstName: "Harry",
	}
	if friend.FirstName != "Harry" {
		t.Fatalf("TripNameTest fail because Name of Trip is not equal to %s but to %s", "Harry", friend.FirstName)
	}
}
