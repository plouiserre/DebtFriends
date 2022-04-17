package main

import (
	"testing"
)

func TestTripName(t *testing.T) {
	trip := Trip{
		Name: "holidays",
	}
	if trip.Name != "holidays" {
		t.Fatalf("TripNameTest fail because Name of Trip is not equal to %s but to %s", "holidays", trip.Name)
	}
}

func TestFriendsTrip(t *testing.T) {
	firstFriend := Friend{
		FirstName: "Ron",
	}
	secondFriend := Friend{
		FirstName: "Hermione",
	}
	thirdFriend := Friend{
		FirstName: "Harry",
	}
	var friends []Friend
	friends = append(friends, firstFriend, secondFriend, thirdFriend)

	trip := Trip{
		Name: "holidays",
	}

	trip.AddFriends(friends)

	if len(trip.Friends) != 3 {
		t.Fatalf("Friends of the trip %s have the length %d and %d is waited", trip.Name, len(trip.Friends), 3)
	}
}
