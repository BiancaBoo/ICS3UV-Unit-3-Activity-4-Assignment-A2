/**
 * Author: Bianca Bertinato
 * Version: 1.0.0
 * Date: 2025-11-17
 * This program calculates if contestant is eligible for pie contest
 */

package main

import (
	"fmt"
)

func main() {
	// Constants
	const MIN_WEIGHT float64 = 77.0
	const MAX_WEIGHT float64 = 105.0

	// Input
	var weight float64

	fmt.Print("How much do you weigh? ")
	fmt.Scan(&weight)

	// Decision
	if weight >= MIN_WEIGHT && weight <= MAX_WEIGHT {
		fmt.Println("You may enter the contest.")
	} else {
		fmt.Println("You cannot enter the contest.")
	}

	fmt.Println("\nDone.")
}