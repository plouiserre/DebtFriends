package main

import "fmt"

func main() {
	trip := Trip{
		Name: "Hogwarts",
	}
	fmt.Printf("trip %s created \n", trip.Name)

	u := Util{}

	activities := u.InitActivity()

	trip.AddActivities(activities)

	for _, activity := range trip.Activities {
		fmt.Printf("%s is an activity to the trip %s and cost %f \n", activity.Name, trip.Name, activity.Price)
	}

	for _, friend := range trip.Friends {
		fmt.Printf("%s is a friend of the trip\n", friend.FirstName)
	}
}
