package main

import (
	"testing"
)

func TestFriendNameExpense(t *testing.T) {
	friend := Friend{
		FirstName:      "Harry",
		VirtualExpense: 298.23,
	}
	if friend.FirstName != "Harry" {
		t.Fatalf("TestFriendNameExpense fail because his/her firstname is not equal to %s but to %s", "Harry", friend.FirstName)
	}
	if friend.VirtualExpense != 298.23 {
		t.Fatalf("TestFriendNameExpense fail because his/her expense is not equal to %s but to %f", "Harry", friend.VirtualExpense)
	}
}
