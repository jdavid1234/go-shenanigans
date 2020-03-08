package main

import "fmt"

func main() {

	//define map

	emails := make(map[string]string)

	//assign key values

	emails["bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["name_1"] = "name_1@gmail.com"
	emails["name_2"] = "name_2@gmail.com"
	emails["name_3"] = "name_3@gmail.com"
	emails["name_4"] = "name_4@gmail.com"
	emails["name_5"] = "name_5@gmail.com"
	emails["name_6"] = "name_6@gmail.com"
	emails["name_7"] = "name_7@gmail.com"
	emails["name_8"] = "name_8@gmail.com"
	emails["name_9"] = "name_9@gmail.com"
	emails["name_10"] = "name_10@gmail.com"
	emails["name_11"] = "name_11@gmail.com"
	emails["name_12"] = "name_12@gmail.com"
	emails["name_13"] = "name_13@gmail.com"
	emails["name_14"] = "name_14@gmail.com"
	emails["name_15"] = "name_15@gmail.com"
	emails["name_16"] = "name_16@gmail.com"
	emails["name_17"] = "name_17@gmail.com"
	emails["name_18"] = "name_18@gmail.com"
	emails["name_19"] = "name_19@gmail.com"
	emails["name_20"] = "name_20@gmail.com"
	emails["name_21"] = "name_21@gmail.com"
	emails["name_22"] = "name_22@gmail.com"
	emails["name_23"] = "name_23@gmail.com"
	emails["name_24"] = "name_24@gmail.com"
	emails["name_25"] = "name_25@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["name_23"])

	//or you can make map and asggin key values at the same time:
	//emails := map[string]string{"Bob": "bob@gmail.com", "sharon": "sharon@gmail.com"}

	//delete from map
	delete(emails, "bob")

}
