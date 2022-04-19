package main

type ManagerExpense struct {
	trip Trip
}

func (m *ManagerExpense) DetermineFriendsExpenses() {
	friends := m.trip.Friends

	for i, friend := range friends {
		for _, activity := range m.trip.Activities {
			isParticipated := IsFriendInActivity(friend, activity)
			if isParticipated {
				nbPlayers := float64(len(activity.Friends))
				m.trip.Friends[i].VirtualExpense += activity.Price / nbPlayers
			}
		}
	}
}

func IsFriendInActivity(friend Friend, activity Activity) bool {
	isParticipated := false
	for _, f := range activity.Friends {
		if f.FirstName == friend.FirstName {
			isParticipated = true
			break
		}
	}
	return isParticipated
}
