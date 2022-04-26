package main

import "fmt"

func main() {
	trip := Trip{
		Name: "Hogwarts",
	}
	fmt.Printf("trip %s created \n", trip.Name)

	u := Util{}

	activities := u.InitActivities()

	trip.AddActivities(activities)

	m := ManagerExpense{
		trip: trip,
	}

	for _, activity := range trip.Activities {
		fmt.Printf("%s is an activity to the trip %s and cost %f \n", activity.Name, trip.Name, activity.Price)
	}

	m.PayDebts()

	PrintExpensesFriend(trip.Friends)

	m.CreateFriendsLinks()

	for _, friendLink := range m.FriendLinks {
		fmt.Printf("%s as first friend %s as second friend and %.2f as value \n", friendLink.FirstFriend.FirstName, friendLink.SecondFriend.FirstName, friendLink.Value)
	}

	m.ExpensesOptimized()

	fmt.Println("|-------------Expenses optimized-------------|")

	PrintExpensesFriend(trip.Friends)

}

func PrintExpensesFriend(friends []Friend) {
	for _, friend := range friends {
		fmt.Printf("%s is a friend of the trip\n", friend.FirstName)

		for _, payment := range friend.Payments {
			fmt.Printf("%s must paid %.2f to %s\n", friend.FirstName, payment.Value, payment.Recipiant.FirstName)
		}
	}
}
