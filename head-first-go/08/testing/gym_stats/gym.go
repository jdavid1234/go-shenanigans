//this is my first package where I make structs

package gym

type Member struct {
	member []Memberships
	Name   string
	Email  string
	Active bool
	PersonalInfo
	Deal
	Address
	Stats
}

type PersonalInfo struct {
	Gender string
	Height float64
	Weight float64
}

type Deal struct {
	Plan            string
	Cost            float64
	UpgradeEligible bool
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Stats struct {
	Lift    float64
	Bentch  float64
	Situps  float64
	Pushups float64
}
