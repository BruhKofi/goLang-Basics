package main

import (
	"fmt"
)

var shippingDays = 30
var shippingDaysPtr = new(int)

func main() {
	shippingDayAdjustments(shippingDays)
	fmt.Println("after adjustment", shippingDays)

	shippingDayAdjustmentsPtr(&shippingDays)
	fmt.Println("after adjustment", shippingDays)
	fmt.Println("after adjustment", &shippingDays)

	shipper := shipper{}
	shipper.id = 400
	shipper.name = "FishMan"
	git
	shipperUpdates(&shipper)
	fmt.Println("shipper id after call", shipper.id)
	fmt.Println("shippre name after call", shipper.name)

	*shippingDaysPtr = 55
	shippingDayAdjustmentsPtr(shippingDaysPtr)
	fmt.Println("Ptr after adjustment", *shippingDaysPtr)
}

func shippingDayAdjustments(shippingDays int) {
	shippingDays += 10
}

func shippingDayAdjustmentsPtr(shippingDays *int) {
	*shippingDays += 10
}

func shipperUpdates(s *shipper) {
	s.id += 10
	s.name += "Inc"
}

type shipper struct {
	name string
	id   int
}
