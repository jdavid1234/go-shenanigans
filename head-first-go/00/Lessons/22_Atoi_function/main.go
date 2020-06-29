

//the Args function will take all the input (arguments) from our command line, and put into an empty slice called arguments 
arguments := make([]string)
arguments = os.Args[1:]

//we also need to preapre a and empty slice which will receive the numbers after they have been converted to integers 
finalNumbers := make([]int)

//the idea here is to range throguh our strings (get rid of the index values in the process) and put this into a variable called 'arguement' 
//as the Atoi function converts the contents of our argumment variable into integers, it splits each value to 2 sections; number and errors 
//if the user inputs a '6' into the command line, a 6 int will be passed into the 'n'umber' variable, and a nil will be passed into the 'err' variable 
//if the user inputs a letter into the command line; the Atoi function will pass error contents into the err variable
//we need an if statement in order to log this potential error, perfrom the 'Fatal' function on the err variable, and stop the program. 
​


	for _, argument := range arguments {
​
		number, err := strconv.Atoi(argument)
​
		if err != nil {
			log.Fatal(err)
		}
​
// the idea here is to range through our command line arguments, convert them to integers ('number'),
//and on each iteration of this range we would append each number to our empty finalNumbers slice.
//(append takes in the list you want to append to (finalNumbers) and the value you want to append to it (number))
		finalNumbers = append(finalNumbers, number)
​
	}
​
	fmt.Println(finalNumbers)
