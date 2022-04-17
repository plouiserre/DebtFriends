package main

import "fmt"

func main() {
	trip := Trip{
		Name: "Hogwarts",
	}
	fmt.Printf("trip %s created \n", trip.Name)

	firstFriend := Friend{
		FirstName: "Ron",
	}
	secondFriend := Friend{
		FirstName: "Hermione",
	}
	thirdFriend := Friend{
		FirstName: "Harry",
	}
	var friends []Friend
	friends = append(friends, firstFriend, secondFriend, thirdFriend)

	trip.AddFriends(friends)

	for _, friend := range trip.Friends {
		fmt.Printf("%s participate to the trip %s \n", friend.FirstName, trip.Name)
	}

	firstActivity := Activity{
		Name:  "Drink Beers",
		Price: 23.9,
	}

	firstActivity.AddFriends(friends)

	secondActivity := Activity{
		Name:  "Rent Magic Brooms",
		Price: 120.7,
	}

	secondActivity.AddFriends(friends)

	var activities []Activity
	activities = append(activities, firstActivity, secondActivity)

	trip.AddActivities(activities)

	for _, activity := range trip.Activities {
		fmt.Printf("%s is an activity to the trip %s and cost %f \n", activity.Name, trip.Name, activity.Price)
	}
}
