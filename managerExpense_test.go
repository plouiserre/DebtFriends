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
	if ronPayments[0].Recipiant.FirstName != "Hermione" || ronPayments[0].Value != 11.95 {
		t.Fatalf("TestRefundFriends fails because  Ron must paid 11.95 to Hermione and not %f to %s ", ronPayments[0].Value, ronPayments[0].Recipiant.FirstName)
	}

	hermionePayment := m.trip.Friends[1].Payments
	AssertPaymentTimes(hermionePayment, 2, t, "Hermione")
	if hermionePayment[0].Recipiant.FirstName != "Neville" || hermionePayment[0].Value != 139.115 {
		t.Fatalf("TestRefundFriends fails because  Hermione must paid 139.115 to Neville and not %f to %s ", hermionePayment[0].Value, hermionePayment[0].Recipiant.FirstName)
	}
	if hermionePayment[1].Recipiant.FirstName != "Ron" || hermionePayment[1].Value != 22.415 {
		t.Fatalf("TestRefundFriends fails because  Hermione must paid 22.415 to Ron and not %f to %s ", hermionePayment[1].Value, hermionePayment[1].Recipiant.FirstName)
	}

	harryPayment := m.trip.Friends[2].Payments
	AssertPaymentTimes(m.trip.Friends[2].Payments, 1, t, "Harry")
	if harryPayment[0].Recipiant.FirstName != "Ron" || harryPayment[0].Value != 50.905 {
		t.Fatalf("TestRefundFriends fails because  Harry must paid 50.905 to Ron and not %f to %s ", harryPayment[0].Value, harryPayment[0].Recipiant.FirstName)
	}

	nevillePayment := m.trip.Friends[3].Payments
	AssertPaymentTimes(m.trip.Friends[3].Payments, 2, t, "Neville")
	if nevillePayment[0].Recipiant.FirstName != "Harry" || nevillePayment[0].Value != 60.35 {
		t.Fatalf("TestRefundFriends fails because  Neville must paid 60.35 to Harry and not %f to %s ", nevillePayment[0].Value, nevillePayment[0].Recipiant.FirstName)
	}
	if nevillePayment[1].Recipiant.FirstName != "Ron" || nevillePayment[1].Value != 22.415 {
		t.Fatalf("TestRefundFriends fails because  Neville must paid 22.415 to Ron and not %f to %s ", nevillePayment[1].Value, nevillePayment[1].Recipiant.FirstName)
	}
}

func TestCreateFriendsLinks(t *testing.T) {
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

	m.CreateFriendsLinks()

	FriendLinksCount := len(m.FriendLinks)

	if FriendLinksCount != 5 {
		t.Fatalf("TestRefundFriends fails because  there is 5 FriendsLinks and not %d ", FriendLinksCount)
	}

	AssertCreationFriendsLinks(t, m.FriendLinks[0], u.FirstFriend, u.SecondFriend, "The first link between Ron and Hermione is failing", -10.465)

	AssertCreationFriendsLinks(t, m.FriendLinks[1], u.SecondFriend, u.FourthFriend, "The second link between Hermione and Neville is failing", 139.115)

	AssertCreationFriendsLinks(t, m.FriendLinks[2], u.ThirdFriend, u.FirstFriend, "The third link between Harry and Ron is failing", 50.905)

	AssertCreationFriendsLinks(t, m.FriendLinks[3], u.FourthFriend, u.FirstFriend, "The fourth link between Neville and Ron is failing", 22.415)

	AssertCreationFriendsLinks(t, m.FriendLinks[4], u.FourthFriend, u.ThirdFriend, "The fifth link between Neville and Harry is failing", 60.35)
}

func AssertCreationFriendsLinks(t *testing.T, fL FriendLink, firstFriend Friend, secondFriend Friend, messageError string, value float64) {
	isOk := (fL.FirstFriend.FirstName == firstFriend.FirstName && fL.SecondFriend.FirstName == secondFriend.FirstName && fL.Value == value)
	if !isOk {
		t.Fatalf(messageError)
	}
}

/*func TestMaximizeExpense(t *testing.T) {
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

	m.MaximizeExpense()

	ronPayments := m.trip.Friends[0].Payments
	AssertPaymentTimes(ronPayments, 0, t, "Ron")

	hermionePayment := m.trip.Friends[1].Payments
	AssertPaymentTimes(hermionePayment, 2, t, "Hermione")
	if hermionePayment[0].Recipiant.FirstName != "Neville" && hermionePayment[0].Value != 139.115 {
		t.Fatalf("TestCreateFriendLinks fails because  Hermione must paid 139.115 to Neville and not %f to %s ", hermionePayment[0].Value, hermionePayment[0].Recipiant.FirstName)
	}
	if hermionePayment[1].Recipiant.FirstName != "Ron" && hermionePayment[1].Value != 10.465 {
		t.Fatalf("TestCreateFriendLinks fails because  Hermione must paid 10.465 to Ron and not %f to %s ", hermionePayment[1].Value, hermionePayment[1].Recipiant.FirstName)
	}

	harryPayment := m.trip.Friends[2].Payments
	AssertPaymentTimes(m.trip.Friends[2].Payments, 1, t, "Harry")
	if harryPayment[0].Recipiant.FirstName != "Ron" && harryPayment[0].Value != 50.905 {
		t.Fatalf("TestCreateFriendLinks fails because  Harry must paid 50.905 to Ron and not %f to %s ", ronPayments[0].Value, ronPayments[0].Recipiant.FirstName)
	}

	nevillePayment := m.trip.Friends[3].Payments
	AssertPaymentTimes(m.trip.Friends[3].Payments, 2, t, "Neville")
	if nevillePayment[0].Recipiant.FirstName != "Harry" && nevillePayment[0].Value != 139.115 {
		t.Fatalf("TestCreateFriendLinks fails because  Neville must paid 60.35 to Harry and not %f to %s ", nevillePayment[0].Value, nevillePayment[0].Recipiant.FirstName)
	}
	if nevillePayment[1].Recipiant.FirstName != "Ron" && nevillePayment[1].Value != 22.415 {
		t.Fatalf("TestCreateFriendLinks fails because  Neville must paid 22.415 to Ron and not %f to %s ", nevillePayment[1].Value, nevillePayment[1].Recipiant.FirstName)
	}
}*/

//TODO reprendre cette méthode
func AssertPaymentTimes(payments []Payment, numberPaymentWaited int, t *testing.T, friendName string) {
	if len(payments) != numberPaymentWaited {
		t.Fatalf("TestRefundFriends fails because  %s must paid %d friend(s) and not %d ", friendName, numberPaymentWaited, len(payments))
	}
}
