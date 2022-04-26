package main

import (
	"testing"
)

func TestFriendNameExpense(t *testing.T) {
	friend := InitFriend(nil)

	if friend.FirstName != "Harry" {
		t.Fatalf("TestFriendNameExpense fail because his/her firstname is not equal to %s but to %s", "Harry", friend.FirstName)
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
		FirstName: "Harry",
	}
	if len(payments) > 0 {
		friend.Payments = payments
	}
	return friend
}

func TestFriendLink(t *testing.T) {
	firstFriend := Friend{
		FirstName: "Harry",
	}

	secondFriend := Friend{
		FirstName: "Hermione",
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
