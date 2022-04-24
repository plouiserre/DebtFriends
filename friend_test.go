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

func TestFriendLink(t *testing.T) {
	firstFriend := Friend{
		FirstName:      "Harry",
		VirtualExpense: 0,
	}

	secondFriend := Friend{
		FirstName:      "Hermione",
		VirtualExpense: 0,
	}
	friendLink := FriendLink{
		FirstFriend:  firstFriend,
		SecondFriend: secondFriend,
		Value:        98765,
	}
	if friendLink.FirstFriend.FirstName != firstFriend.FirstName {
		t.Fatalf("TestFriendLink fails because firstfriend's firstname is %s and no %s", friendLink.FirstFriend.FirstName, firstFriend.FirstName)
	}
	if friendLink.SecondFriend.FirstName != secondFriend.FirstName {
		t.Fatalf("TestFriendLink fails because secondfriend's firstname is %s and no %s", friendLink.FirstFriend.FirstName, firstFriend.FirstName)
	}
	if friendLink.Value != 98765 {
		t.Fatalf("TestFriendLink fails because Value is %d and not %f", 98765, friendLink.Value)
	}
}
