package main

import (
	"testing"
)

func TestDetermineEachFriendsExpenses(t *testing.T) {
	u := Util{}
	activities := u.InitActivities()
	trip := Trip{
		Name: "Hogwarts",
	}
	trip.AddActivities(activities)

	m := ManagerExpense{
		trip: trip,
	}
	m.DetermineFriendsExpenses()

	ExpenseAssert(t, trip.Friends, "Ron", 62.855)
	ExpenseAssert(t, trip.Friends, "Hermione", 173.48)
	ExpenseAssert(t, trip.Friends, "Harry", 111.255)
	ExpenseAssert(t, trip.Friends, "Neville", 221.88)
}

func ExpenseAssert(t *testing.T, fs []Friend, firstName string, expense float64) {
	f := GetSpecificFriend(fs, firstName)
	if f.VirtualExpense != expense {
		t.Fatalf("TestDetermineEachFriendsExpenses fails because expense of %s is %f and not %f", firstName, expense, f.VirtualExpense)
	}
}

func GetSpecificFriend(fs []Friend, firstName string) Friend {
	for _, f := range fs {
		if f.FirstName == firstName {
			return f
		}
	}
	return Friend{}
}
