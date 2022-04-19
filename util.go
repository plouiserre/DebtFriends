package main

type Util struct {
	FirstFriend  Friend
	SecondFriend Friend
	ThirdFriend  Friend
	FourthFriend Friend
}

func Contains(fs []Friend, f Friend) bool {
	for _, friend := range fs {
		if f == friend {
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
		Name:  "Buy Beers",
		Price: 23.9,
	}

	firstActivity.AddFriends([]Friend{u.FirstFriend, u.SecondFriend})

	secondActivity := Activity{
		Name:  "Rent Magic Brooms",
		Price: 120.7,
	}

	secondActivity.AddFriends([]Friend{u.ThirdFriend, u.FourthFriend})

	thirdActivity := Activity{
		Name:  "Buy some presents",
		Price: 56.98,
	}
	thirdActivity.AddFriends([]Friend{u.FirstFriend, u.ThirdFriend})

	fourthActivity := Activity{
		Name:  "Shopping",
		Price: 278.23,
	}
	fourthActivity.AddFriends([]Friend{u.SecondFriend, u.FourthFriend})

	fifthActivity := Activity{
		Name:  "Bar",
		Price: 89.66,
	}
	fifthActivity.AddFriends([]Friend{u.FirstFriend, u.SecondFriend, u.ThirdFriend, u.FourthFriend})

	activities := []Activity{firstActivity, secondActivity, thirdActivity, fourthActivity, fifthActivity}

	return activities
}
