package main

type Activity struct {
	Name      string
	Price     float64
	Friends   []Friend
	Paymaster Friend
}

func (a *Activity) AddFriends(friends []Friend) {
	a.Friends = append(a.Friends, friends...)
}
