package main

type Trip struct {
	Name    string
	Friends []Friend
}

func (t *Trip) AddFriends(friends []Friend) {
	t.Friends = append(t.Friends, friends...)
}
