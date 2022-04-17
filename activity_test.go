package main

import (
	"testing"
)

func TestActivityInit(t *testing.T) {
	activity := Activity{
		Name:  "Drink Beers",
		Price: 23.9,
	}
	if activity.Name != "Drink Beers" {
		t.Fatalf("TestActivityInit fails because Name of the activity is not equal to %s but to %s", "Drink Beers", activity.Name)
	}
	if activity.Price != 23.9 {
		t.Fatalf("TestActivityInit fails because Price of the activitt is not equal to %f but to %f", 23.9, activity.Price)
	}
}

func TestActivityAddFriends(t *testing.T) {
	activity := Activity{
		Name:  "Drink Beers",
		Price: 23.9,
	}

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

	activity.AddFriends(friends)

	if len(activity.Friends) != 3 {
		t.Fatalf("TestActivityAddFriends fails because Friends of the activity %s have the length %d and %d is waited", activity.Name, len(activity.Friends), 3)
	}
}
