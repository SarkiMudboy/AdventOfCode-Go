package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var total int

func main() {

	// read from the file line by line
	var input io.Reader
	file, err := os.Open("aoc_input.txt")

	if err != nil {
		fmt.Println("An error ocurred opening the file: ", err)
		os.Exit(1)
	}

	defer file.Close()

	input = file

	buf := bufio.NewScanner(input)

	// result of computing each line
	var result int

	//loop through or find the numbers in each line
	for buf.Scan() {

		pat := regexp.MustCompile("[0-9]")
		numbers := pat.FindAllString(buf.Text(), -1)

		first, err := strconv.Atoi(numbers[0])
		second, err := strconv.Atoi(numbers[len(numbers)-1])

		if err != nil {
			fmt.Println("Error: ", err)
		}

		//add them to a total int var (but check for same number)
		if len(numbers) == 1 {
			total += (first * 11)
		} else {
			if second != 0 {
				result = (first * 10) + second

			} else {
				result = (first * 10)
			}
			total += result
			fmt.Println(numbers, result)
		}

	}
	fmt.Println("Total ", total)

}
