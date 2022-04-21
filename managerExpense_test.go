package main

import (
	"testing"
)

//TODO facto à faire
func TestDetermineEachFriendsExpenses(t *testing.T) {
	u := Util{}
	activities := u.InitActivities()
	trip := Trip{
		Name: "Hogwarts",
	}
	trip.AddActivities(activities)

	m := ManagerExpense{
		trip: trip,
	}
	m.DetermineFriendsExpenses()

	ExpenseAssert(t, trip.Friends, "Ron", 62.855)
	ExpenseAssert(t, trip.Friends, "Hermione", 173.48)
	ExpenseAssert(t, trip.Friends, "Harry", 111.255)
	ExpenseAssert(t, trip.Friends, "Neville", 221.88)
}

func ExpenseAssert(t *testing.T, fs []Friend, firstName string, expense float64) {
	f := GetSpecificFriend(fs, firstName)
	if f.VirtualExpense != expense {
		t.Fatalf("TestDetermineEachFriendsExpenses fails because expense of %s is %f and not %f", firstName, expense, f.VirtualExpense)
	}
}

func GetSpecificFriend(fs []Friend, firstName string) Friend {
	for _, f := range fs {
		if f.FirstName == firstName {
			return f
		}
	}
	return Friend{}
}

func TestRefundFriends(t *testing.T) {
	u := Util{}
	activities := u.InitActivities()
	trip := Trip{
		Name: "Hogwarts",
	}
	trip.AddActivities(activities)

	m := ManagerExpense{
		trip: trip,
	}

	m.PayDebts()

	ronPayments := m.trip.Friends[0].Payments
	AssertPaymentTimes(ronPayments, 1, t, "Ron")
	if ronPayments[0].Recipiant.FirstName != "Hermione" && ronPayments[0].Value != 11.95 {
		t.Fatalf("TestRefundFriends fails because  Ron must paid 11.95 to Hermione and not %f to %s ", ronPayments[0].Value, ronPayments[0].Recipiant.FirstName)
	}

	hermionePayment := m.trip.Friends[1].Payments
	AssertPaymentTimes(hermionePayment, 2, t, "Hermione")
	if hermionePayment[0].Recipiant.FirstName != "Neville" && hermionePayment[0].Value != 139.115 {
		t.Fatalf("TestRefundFriends fails because  Hermione must paid 139.115 to Neville and not %f to %s ", hermionePayment[0].Value, hermionePayment[0].Recipiant.FirstName)
	}
	if hermionePayment[1].Recipiant.FirstName != "Ron" && hermionePayment[1].Value != 22.415 {
		t.Fatalf("TestRefundFriends fails because  Hermione must paid 22.415 to Ron and not %f to %s ", hermionePayment[1].Value, hermionePayment[1].Recipiant.FirstName)
	}

	harryPayment := m.trip.Friends[2].Payments
	AssertPaymentTimes(m.trip.Friends[2].Payments, 1, t, "Harry")
	if harryPayment[0].Recipiant.FirstName != "Ron" && harryPayment[0].Value != 50.905 {
		t.Fatalf("TestRefundFriends fails because  Harry must paid 50.905 to Ron and not %f to %s ", ronPayments[0].Value, ronPayments[0].Recipiant.FirstName)
	}

	nevillePayment := m.trip.Friends[3].Payments
	AssertPaymentTimes(m.trip.Friends[3].Payments, 2, t, "Neville")
	if nevillePayment[0].Recipiant.FirstName != "Harry" && nevillePayment[0].Value != 139.115 {
		t.Fatalf("TestRefundFriends fails because  Neville must paid 60.35 to Harry and not %f to %s ", nevillePayment[0].Value, nevillePayment[0].Recipiant.FirstName)
	}
	if nevillePayment[1].Recipiant.FirstName != "Ron" && nevillePayment[1].Value != 22.415 {
		t.Fatalf("TestRefundFriends fails because  Neville must paid 22.415 to Ron and not %f to %s ", nevillePayment[1].Value, nevillePayment[1].Recipiant.FirstName)
	}
}

func AssertPaymentTimes(payments []Payment, numberPaymentWaited int, t *testing.T, friendName string) {
	if len(payments) != numberPaymentWaited {
		t.Fatalf("TestRefundFriends fails because  %s must paid %d friend(s) and not %d ", friendName, numberPaymentWaited, len(payments))
	}
}
