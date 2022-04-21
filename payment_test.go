package main

import (
	"testing"
)

func TestInitPayment(t *testing.T) {
	u := Util{}
	payments := u.InitPayments()

	u.AssertPayments(payments, t)
}
