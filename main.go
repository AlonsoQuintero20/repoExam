package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Student struct {
	Name   string
	Score1 int
	Score2 int
	Score3 int
}

func main() {
	var bandera = false
	for bandera == false {
		election := 0

		fmt.Print("\nSelect the option that you wish:\n")
		fmt.Print(`
		[1] - Create a list of Students.
		[2] - Open the file.
		[3] - Exit ` + "\n")
		fmt.Print("---->")
		fmt.Scanln(&election)

		switch {
		case election == 1:
			newFile()
		case election == 2:
			readFile()
		case election == 3:
			os.Exit(3)

		}
	}
}

func newFile() {

	var students Student

	fmt.Print("Write the first name: ")
	fmt.Scanln(&students.Name)
	fmt.Print("Write the First grade: ")
	fmt.Scanln(&students.Score1)
	fmt.Print("Write the Second grade: ")
	fmt.Scanln(&students.Score2)
	fmt.Print("Write the Third grade: ")
	fmt.Scanln(&students.Score3)

	mydata := []byte("Student: " + students.Name + " has ratings: " + strconv.Itoa(students.Score1) + ", " + strconv.Itoa(students.Score2) + ", " + strconv.Itoa(students.Score3))

	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("myfile.txt", mydata, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("myfile.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}

func readFile() {
	data, err := ioutil.ReadFile("myfile.txt")
	if err != nil {
		panic(err)
	}
	// Should output: `The file contains: Hello, world!`
	fmt.Println("The file contains: " + string(data))
}
