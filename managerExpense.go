package main

type ManagerExpense struct {
	trip        Trip
	FriendLinks []FriendLink
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
			friend := &m.trip.Friends[index]
			isNewPaymentNeed := true
			for _, payment := range friend.Payments {
				paymentIndex := 0
				if payment.Recipiant.FirstName == paymaster.FirstName {
					isNewPaymentNeed = false
					payment.Value += activity.Price / float64(len(activity.Friends))
					friend.Payments[paymentIndex].Value = payment.Value
				}
				paymentIndex += 1
			}
			if isNewPaymentNeed {
				newPayment := Payment{
					Recipiant: paymaster,
					Value:     activity.Price / float64(len(activity.Friends)),
				}
				friend.Payments = append(friend.Payments, newPayment)
			}

		}
	}
}

func (m *ManagerExpense) CreateFriendsLinks() {
	for _, f := range m.trip.Friends {
		for _, fr := range m.trip.Friends {
			var paymentValue float64
			for _, payment := range f.Payments {
				if payment.Recipiant.FirstName != fr.FirstName {
					continue
				}
				paymentValue = payment.Value
			}
			if paymentValue == 0 {
				continue
			}
			entryExisting := false
			i := 0
			for _, fl := range m.FriendLinks {
				isFriendLinkInThisOrder := fl.FirstFriend.FirstName == f.FirstName && fl.SecondFriend.FirstName == fr.FirstName
				isFriendLinkInOppositeOrder := fl.SecondFriend.FirstName == f.FirstName && fl.FirstFriend.FirstName == fr.FirstName

				if isFriendLinkInThisOrder {
					m.FriendLinks[i].Value += paymentValue
					entryExisting = true
				} else if isFriendLinkInOppositeOrder {
					m.FriendLinks[i].Value -= paymentValue
					entryExisting = true
				}
				i += 1
			}
			if !entryExisting && paymentValue > 0 {
				newFriendLink := FriendLink{
					FirstFriend:  f,
					SecondFriend: fr,
					Value:        paymentValue,
				}
				m.FriendLinks = append(m.FriendLinks, newFriendLink)
			}
		}
	}
}

func (m *ManagerExpense) ExpensesOptimized() {
	for _, fl := range m.FriendLinks {
		for i, f := range m.trip.Friends {
			if f.FirstName != fl.FirstFriend.FirstName {
				continue
			}
			for j, p := range f.Payments {
				if p.Recipiant.FirstName != fl.SecondFriend.FirstName {
					continue
				}
				if fl.Value < 0 {
					payments := m.trip.Friends[i].Payments
					payments = append(payments[:j], payments[j+1:]...)
					m.trip.Friends[i].Payments = payments
					continue
				}
				f.Payments[j].Value = fl.Value
			}
		}
	}
}
