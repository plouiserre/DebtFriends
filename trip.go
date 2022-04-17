package main

type Trip struct {
	Name       string
	Friends    []Friend
	Activities []Activity
}

func (t *Trip) AddFriends(friends []Friend) {
	t.Friends = append(t.Friends, friends...)
}

func (t *Trip) AddActivities(activities []Activity) {
	t.Activities = append(t.Activities, activities...)
}
