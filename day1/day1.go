package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)
func main(){
    part2()
}

func part1() {

	//Need to read in the txt file

	//Opening the .txt file
	file, ferr := os.Open("real.txt")
	if ferr != nil {
		//TOOD: Learn what a panic is
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	//Creating an array for the first numbers
	//Creating a slice where I can use append to add number
	firstNumbers := make([]int, 0)

	//Creating an array for the second numbers
	secondNumbers := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		//Get the first number and put it in one array
		firstNum, err := strconv.Atoi(line[0:5])
		if err != nil {
			panic(err)
		}

		//Get the second number and put it in the second array
		secondNum, err := strconv.Atoi(line[len(line)-5 :])
		if err != nil {
			panic(err)
		}

		firstNumbers = append(firstNumbers, firstNum)
		secondNumbers = append(secondNumbers, secondNum)
	}

	//Sort the first array

	slices.Sort(firstNumbers)

	//Sort the second array

	slices.Sort(secondNumbers)

	//Create a sum variable
	sum := 0

	//Add the absolute value of the first value minus the second value

	for i := 0; i < len(firstNumbers); i++ {
		sum = sum + (int)(math.Abs(float64(firstNumbers[i])-float64(secondNumbers[i])))
	}

	//print the sum of the absolute values
	fmt.Printf("The sum of the distances is: %d\n", sum)

}

func part2(){


	//Need to read in the txt file

	//Opening the .txt file
	file, ferr := os.Open("real.txt")
	if ferr != nil {
		//TOOD: Learn what a panic is
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	//Creating an array for the first numbers
	//Creating a slice where I can use append to add number
	firstNumbers := make([]int, 0)

	//Creating a map for the second numbers
    secondNumbersCount := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		//Get the first number and put it in one array
		firstNum, err := strconv.Atoi(line[0:5])
		if err != nil {
			panic(err)
		}

		//Get the second number and put it in the second array
		secondNum, err := strconv.Atoi(line[len(line)-5 :])
		if err != nil {
			panic(err)
		}

		firstNumbers = append(firstNumbers, firstNum)

        secondNumbersCount[secondNum] = secondNumbersCount[secondNum] + 1

		//secondNumbers = append(secondNumbers, secondNum)
	}

	//Create a sum variable
	sum := 0

	//Add the absolute value of the first value minus the second value

	for i := 0; i < len(firstNumbers); i++ {
        sum = sum + (firstNumbers[i] * secondNumbersCount[firstNumbers[i]])
	}

	//print the sum of the absolute values
	fmt.Printf("The sum of the distances is: %d\n", sum)

}
