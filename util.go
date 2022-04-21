package main

import "testing"

type Util struct {
	FirstFriend  Friend
	SecondFriend Friend
	ThirdFriend  Friend
	FourthFriend Friend
}

//Because Friend have Payments slice, I cannot compare two struct so I use FirstName like key
func Contains(fs []Friend, f Friend) bool {
	for _, friend := range fs {
		if f.FirstName == friend.FirstName {
			return true
		}
	}
	return false
}

func (u *Util) InitActivities() []Activity {
	u.FirstFriend = Friend{
		FirstName: "Ron",
	}
	u.SecondFriend = Friend{
		FirstName: "Hermione",
	}
	u.ThirdFriend = Friend{
		FirstName: "Harry",
	}
	u.FourthFriend = Friend{
		FirstName: "Neville",
	}

	firstActivity := Activity{
		Name:      "Buy Beers",
		Price:     23.9,
		Paymaster: u.SecondFriend,
	}

	firstActivity.AddFriends([]Friend{u.FirstFriend, u.SecondFriend})

	secondActivity := Activity{
		Name:      "Rent Magic Brooms",
		Price:     120.7,
		Paymaster: u.ThirdFriend,
	}

	secondActivity.AddFriends([]Friend{u.ThirdFriend, u.FourthFriend})

	thirdActivity := Activity{
		Name:      "Buy some presents",
		Price:     56.98,
		Paymaster: u.FirstFriend,
	}
	thirdActivity.AddFriends([]Friend{u.FirstFriend, u.ThirdFriend})

	fourthActivity := Activity{
		Name:      "Shopping",
		Price:     278.23,
		Paymaster: u.FourthFriend,
	}
	fourthActivity.AddFriends([]Friend{u.SecondFriend, u.FourthFriend})

	fifthActivity := Activity{
		Name:      "Bar",
		Price:     89.66,
		Paymaster: u.FirstFriend,
	}
	fifthActivity.AddFriends([]Friend{u.FirstFriend, u.SecondFriend, u.ThirdFriend, u.FourthFriend})

	activities := []Activity{firstActivity, secondActivity, thirdActivity, fourthActivity, fifthActivity}

	return activities
}

func (u *Util) InitPayments() []Payment {
	firstFriend := Friend{
		FirstName: "Ron",
	}
	secondFriend := Friend{
		FirstName: "Hermione",
	}

	firstPayment := Payment{
		Recipiant: firstFriend,
		Value:     87.23,
	}

	secondPayment := Payment{
		Recipiant: firstFriend,
		Value:     66.67,
	}

	thirdPayment := Payment{
		Recipiant: secondFriend,
		Value:     9.83,
	}

	fourthPayment := Payment{
		Recipiant: secondFriend,
		Value:     4.54,
	}

	payments := []Payment{}
	payments = append(payments, firstPayment, secondPayment, thirdPayment, fourthPayment)

	return payments
}

func (u *Util) AssertPayments(payments []Payment, t *testing.T) {
	if len(payments) != 4 {
		t.Fatalf("TestInitPayment fails because length of the payment is %d and not %d", 4, len(payments))
	}

	if payments[0].Recipiant.FirstName != "Ron" {
		t.Fatalf("TestInitPayment fails because the name's recipants of the payment is %s and not %s", "Ron", payments[0].Recipiant.FirstName)
	}

	if payments[0].Value != 87.23 {
		t.Fatalf("TestInitPayment fails because the value of the payment is %f and not %f", 87.23, payments[0].Value)
	}

	if payments[1].Recipiant.FirstName != "Ron" {
		t.Fatalf("TestInitPayment fails because the name's recipants of the payment is %s and not %s", "Ron", payments[1].Recipiant.FirstName)
	}

	if payments[1].Value != 66.67 {
		t.Fatalf("TestInitPayment fails because the value of the payment is %f and not %f", 66.67, payments[1].Value)
	}

	if payments[2].Recipiant.FirstName != "Hermione" {
		t.Fatalf("TestInitPayment fails because the name's recipants of the payment is %s and not %s", "Hermione", payments[2].Recipiant.FirstName)
	}

	if payments[2].Value != 9.83 {
		t.Fatalf("TestInitPayment fails because the value of the payment is %f and not %f", 9.83, payments[2].Value)
	}

	if payments[3].Recipiant.FirstName != "Hermione" {
		t.Fatalf("TestInitPayment fails because the name's recipants of the payment is %s and not %s", "Hermione", payments[3].Recipiant.FirstName)
	}

	if payments[3].Value != 4.54 {
		t.Fatalf("TestInitPayment fails because the value of the payment is %f and not %f", 4.54, payments[3].Value)
	}
}
