/*
Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

For example:

    For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
    For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
    For a mass of 1969, the fuel required is 654.
    For a mass of 100756, the fuel required is 33583.

The Fuel Counter-Upper needs to know the total fuel requirement. To find it, individually calculate the fuel needed for the mass of each module (your puzzle input), then add together all the fuel values.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code - Day 1")

	if len(os.Args) == 1 {
		fmt.Println("Please provide an input file.")
		os.Exit(1)
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fuel := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		fuel += total(mass)
	}

	fmt.Printf("Total amount of fuel needed: %d\n", fuel)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// calc takes the mass of a "module" as input and returns the fuel required.
// This is calculated by taking the mass, dividing by 3, rounding down, and then
// subtracting 2.
func calc(mass int) int {
	s1 := mass / 3
	fuel := s1 - 2
	return int(fuel)
}

// total takes the mass and continues calculating fuel costs including the fuel
// needed to carry the fuel
func total(mass int) int {
	total := 0
	fuel := calc(mass)
	for fuel > 0 {
		total += fuel
		fuel = calc(fuel)
	}
	return total
}
