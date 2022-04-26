package main

type Friend struct {
	FirstName string
	Payments  []Payment
}

type FriendLink struct {
	FirstFriend  Friend
	SecondFriend Friend
	Value        float64
}
