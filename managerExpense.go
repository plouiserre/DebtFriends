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

func (m *ManagerExpense) PayDebts() {
	for _, activity := range m.trip.Activities {
		paymaster := activity.Paymaster
		for _, f := range activity.Friends {
			if f.FirstName == paymaster.FirstName {
				continue
			}
			index := m.trip.GetFriendsFromTrip(f.FirstName)
			isNewPaymentNeed := true
			for _, payment := range m.trip.Friends[index].Payments {
				if payment.Recipiant.FirstName == paymaster.FirstName {
					isNewPaymentNeed = false
					payment.Value += activity.Price / float64(len(activity.Friends))
				}
			}
			if isNewPaymentNeed {
				newPayment := Payment{
					Recipiant: paymaster,
					Value:     activity.Price / float64(len(activity.Friends)),
				}
				m.trip.Friends[index].Payments = append(m.trip.Friends[index].Payments, newPayment)
			}

		}
	}
}
