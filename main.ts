/**
 * Author: Bianca Bertinato
 * Version: 1.0.0
 * Date: 2025-11-17
 * This program calculates if contestant is eligible for pie contest
 */

// Weight limit
const MIN_WEIGHT = 77.0;
const MAX_WEIGHT = 105.0;

// Input
const weight = Number(prompt("How much do you weigh?") || "0");

// Decision
if (weight >= MIN_WEIGHT && weight <= MAX_WEIGHT) {
  console.log("You may enter the contest.");
} else {
  console.log("You cannot enter the contest.");
}

console.log("\nDone.");
