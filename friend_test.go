package main

import (
	"testing"
)

func TestFriendNameExpense(t *testing.T) {
	friend := InitFriend(nil)

	if friend.FirstName != "Harry" {
		t.Fatalf("TestFriendNameExpense fail because his/her firstname is not equal to %s but to %s", "Harry", friend.FirstName)
	}
	if friend.VirtualExpense != 298.23 {
		t.Fatalf("TestFriendNameExpense fail because his/her expense is not equal to %s but to %f", "Harry", friend.VirtualExpense)
	}
}

func TestFriendPayments(t *testing.T) {
	u := Util{}
	payments := u.InitPayments()
	friend := InitFriend(payments)

	u.AssertPayments(friend.Payments, t)
}

func InitFriend(payments []Payment) Friend {
	friend := Friend{
		FirstName:      "Harry",
		VirtualExpense: 298.23,
	}
	if len(payments) > 0 {
		friend.Payments = payments
	}
	return friend
}
