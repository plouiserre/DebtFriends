package main

import (
	"testing"
)

//TODO factoriser la partie InitTrip

func TestTripName(t *testing.T) {
	trip := Trip{
		Name: "holidays",
	}
	if trip.Name != "holidays" {
		t.Fatalf("TripNameTest fail because Name of Trip is not equal to %s but to %s", "holidays", trip.Name)
	}
}

func TestActivityTrip(t *testing.T) {
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
		Name: "Hogwarts",
	}

	firstActivity := Activity{
		Name:  "Drink Beers",
		Price: 23.9,
	}

	firstActivity.AddFriends(friends)

	secondActivity := Activity{
		Name:  "Rent Magic Brooms",
		Price: 120.7,
	}

	secondActivity.AddFriends(friends)

	var activities []Activity
	activities = append(activities, firstActivity, secondActivity)

	trip.AddActivities(activities)

	if len(trip.Activities) != 2 {
		t.Fatalf("TestActivityTrip fails because Activities of the trip %s are %d and %d is waited", trip.Name, len(trip.Activities), 2)
	}
}

func TestTripDeterminedFriends(t *testing.T) {
	u := Util{}
	trip := Trip{
		Name: "Hogwarts",
	}

	activities := u.InitActivities()

	trip.AddActivities(activities)

	if len(trip.Friends) != 4 {
		t.Fatalf("TestTripDeterminedFriends fails because Friends of the trip %s are %d and %d is waited", trip.Name, len(trip.Friends), 4)
	}
	CheckFriendsPresent(t, trip.Friends, u.FirstFriend)
	CheckFriendsPresent(t, trip.Friends, u.SecondFriend)
	CheckFriendsPresent(t, trip.Friends, u.ThirdFriend)
	CheckFriendsPresent(t, trip.Friends, u.FourthFriend)

}

func CheckFriendsPresent(t *testing.T, fs []Friend, f Friend) {
	friendPresent := Contains(fs, f)
	if !friendPresent {
		t.Fatalf("TestTripDeterminedFriends fails because %s is missing", f.FirstName)
	}
}

//TODO revoir le nom
func TestTripPaymasterByActivities(t *testing.T) {
	u := Util{}
	trip := Trip{
		Name: "Hogwarts",
	}

	activities := u.InitActivities()

	trip.AddActivities(activities)

	ActivitiesByPaymaster := trip.GetActivitiesByPaymaster()

	ronActivities := ActivitiesByPaymaster[u.FirstFriend]
	hermioneActivities := ActivitiesByPaymaster[u.SecondFriend]
	nevilleActivities := ActivitiesByPaymaster[u.FourthFriend]

	ronActivityCount := len(ronActivities)
	hermioneActivityCount := len(hermioneActivities)
	nevilleActivityCount := len(nevilleActivities)

	if ronActivityCount != 1 {
		t.Fatalf("TestTripPaymasterByActivities fails because Ron has paid %d activity/ies and not %d", 1, ronActivityCount)
	}

	if hermioneActivityCount != 2 {
		t.Fatalf("TestTripPaymasterByActivities fails because Hermione has paid %d activity/ies and not %d", 2, hermioneActivityCount)
	}

	if nevilleActivityCount != 2 {
		t.Fatalf("TestTripPaymasterByActivities fails because Neville has paid %d activity/ies and not %d", 2, nevilleActivityCount)
	}

	AssertSpecificActivity(t, ronActivities, "Bar", "Ron")
	AssertSpecificActivity(t, hermioneActivities, "Buy Beers", "Hermione")
	AssertSpecificActivity(t, hermioneActivities, "Rent Magic Brooms", "Hermione")
	AssertSpecificActivity(t, nevilleActivities, "Buy some presents", "Neville")
	AssertSpecificActivity(t, nevilleActivities, "Shopping", "Neville")

}

func AssertSpecificActivity(t *testing.T, activities []Activity, activityName string, friendName string) {
	isOk := false
	for _, activity := range activities {
		if activity.Name == activityName {
			isOk = true
			break
		}
	}
	if !isOk {
		t.Fatalf("TestTripPaymasterByActivities fails because %s has paid this activity %s and GetActivitiesByPaymaster doesn't indicate", friendName, activityName)
	}
}
