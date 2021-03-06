package main

import (
	"fmt"
	"log"
)

/**
* This module shows how to:
* 1. Create array of sizes and then get initialize some where
* 2. Create array and initialize them
* 3. Get data on a certain index
* 4. Insert data into a position
* 5. delete data in a position
 */

func arrayCreation() {
	// 1.Creating an array of size 10
	var cars [2]string

	//assign values to the array
	cars[0] = "Ford"        // assigned to position 1
	cars[1] = "Range Rover" // assigned to position 1

	fmt.Println("1. Creating and initializing elsewhere")
	fmt.Println(cars)

	// 2. create and initialize at the same time
	names := [3]string{"Domey", "Benjamin", "Akankobateng"}
	fmt.Println("2. Creating and initializing at the same time")
	fmt.Println(names)

	// 3. Get Data at any position
	fmt.Println("3. Getting Data in any position")
	fmt.Println(names[0])
}

//function that takes an array of size 3
func insertDataIntoGivenArray(names [3]string) {
	// 4. Inserting data into any position :: adding a name to the names
	fmt.Println("3. Inserting Data in any position")
	// first I create a new array with size + 1
	var namesAdded [len(names) + 1]string

	//give the position
	var inputFromUserForPosition int
	fmt.Print("Please enter the position where you want to insert ama :: ")
	_, err := fmt.Scan(&inputFromUserForPosition)
	if err != nil {
		log.Fatalln("An error occured when it tried to collect position of where to insert \"ama\"")
	}

	//now check if its within the array length
	if inputFromUserForPosition < 0 || inputFromUserForPosition > len(namesAdded)-1 {
		log.Fatalln("You entered an invalid position. The length of the array is ", len(namesAdded), ". That being said, please enter a position between 0 and ", len(namesAdded)-1)
	}
	//assign the data to the space
	namesAdded[inputFromUserForPosition] = "ama"

	//now insert into right position
	for i := 0; i < len(namesAdded)-1; i++ {
		// check is there is a space in the slot and then add data
		if namesAdded[i] == "" {
			namesAdded[i] = names[i]
		} else {
			// else move to the next index and insert the data :)
			namesAdded[i+1] = names[i]
		}
	}

	for i := 0; i < len(namesAdded); i++ {
		fmt.Println(i, " - ", namesAdded[i])
	}
}

func removeFromAGivenArrayInAPosition(names [3]string) {
	// 4. Inserting data into any position :: adding a name to the names
	fmt.Println("3. Inserting Data in any position")

	for i := 0; i < len(names); i++ {
		fmt.Println(i, " - ", names[i])
	}
	// first I create a new array with size - 1
	var namesSub [len(names) - 1]string

	// get position to remove from
	var position string
	fmt.Print("Enter the string you'd want to remove from the array")
	_, err := fmt.Scan(&position)
	if err != nil {
		panic("Something happened")
	}

	//find string in array
	getStringPosition := -1
	for i := 0; i < len(names); i++ {
		if names[i] == position {
			getStringPosition = i
		}
	}
	// if string is not ready a, quit the app
	if getStringPosition == -1 {
		panic("Please enter a valid string")
	}

	//reduce the array
	for i := 0; i < len(names)-1; i++ {
		if i < getStringPosition {
			// if its less than the string position a, just add to the next index
			namesSub[i] = names[i]
		} else if i == getStringPosition {
			// if its equal to the string position a, just skip it and choose next value
			namesSub[i] = names[i+1]
		} else if i > getStringPosition {
			// if is greater than the string postion a, that means the currnent would have been assigned
			// so you'd have to take the next one
			namesSub[i] = names[i+1]
		}
	}

	for i := 0; i < len(namesSub); i++ {
		fmt.Println(i, " - ", namesSub[i])
	}
}
