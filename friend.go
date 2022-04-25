package main

type Friend struct {
	FirstName      string
	VirtualExpense float64
	Payments       []Payment
}

type FriendLink struct {
	FirstFriend  Friend
	SecondFriend Friend
	Value        float64
}
