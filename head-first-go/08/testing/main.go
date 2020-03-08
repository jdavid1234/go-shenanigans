package main

import (
	"fmt"

	gym "github.com/jdavid1234/go-shenanigans/head-first-go/08/testing/gym_stats"
)

func main() {

	var member1 gym.Member

	member1.Name = "John Smith"
	member1.Email = "jsmith.gmail.com"
	member1.Active = true
	member1.Gender = "male"
	member1.Height = 179.87
	member1.Weight = 80.3
	member1.Plan = "Premium"
	member1.Cost = 150.40
	member1.UpgradeEligible = true
	member1.Street = "10 Estes St"
	member1.City = "Chicago"
	member1.State = "Il"
	member1.PostalCode = "123456"
	member1.Lift = 234
	member1.Bentch = 456
	member1.Situps = 23
	member1.Pushups = 199

	fmt.Println(member1)

}
