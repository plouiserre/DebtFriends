package main

type Activity struct {
	Name    string
	Price   float32
	Friends []Friend
}

func (a *Activity) AddFriends(friends []Friend) {
	a.Friends = append(a.Friends, friends...)
}
