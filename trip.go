package main

type Trip struct {
	Name       string
	Friends    []Friend
	Activities []Activity
}

func (t *Trip) AddActivities(activities []Activity) {
	t.Activities = append(t.Activities, activities...)
	var fs []Friend
	for _, activity := range t.Activities {
		for _, f := range activity.Friends {
			present := Contains(fs, f)
			if present {
				continue
			} else {
				fs = append(fs, f)
			}
		}
	}
	t.Friends = fs
}

func (t *Trip) GetFriendsFromTrip(firstName string) int {
	var index int
	for i, f := range t.Friends {
		if f.FirstName == firstName {
			index = i
			break
		}
	}
	return index
}
