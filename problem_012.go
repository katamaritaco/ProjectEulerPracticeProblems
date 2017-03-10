/*Problem 12 - Highly divisible triangular number
The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

Let us list the factors of the first seven triangle numbers:

 1: 1
 3: 1,3
 6: 1,2,3,6
10: 1,2,5,10
15: 1,3,5,15
21: 1,3,7,21
28: 1,2,4,7,14,28
We can see that 28 is the first triangle number to have over five divisors.

What is the value of the first triangle number to have over five hundred divisors?

*/

package main

import (
	"fmt"
	"math/big"
)

const CONST_MAX_NUM_DIVISORS = 100;

func checkDivisors( runningSum *big.Int ) int {

    increment := big.NewInt(1)
    var sumCutoff big.Int;
    var tmpMod big.Int;
    numDivisors := 0;


    for counter, sumCutoff := big.NewInt( 1 ), sumCutoff.Div( runningSum, big.NewInt(2) ); counter.Cmp( sumCutoff ) <= 0 ; counter.Add( counter, increment ) {
        tmpMod.Mod( runningSum, counter );
        if tmpMod.Cmp( big.NewInt(0) ) == 0 {
            numDivisors++;
        }
    }

    // Since the number itself is a divisor, we'll add that here since we are only checking for num/2 to save time.
    numDivisors++;

    return numDivisors;
}


func main() {

    runningSum := big.NewInt( 0 );
    increment := big.NewInt( 1 );

    isOverFiveHundredDivisors := false;

    for !isOverFiveHundredDivisors {

        runningSum.Add( runningSum, increment );
        increment.Add( increment, big.NewInt(1) );

        if checkDivisors( runningSum ) > CONST_MAX_NUM_DIVISORS {
            isOverFiveHundredDivisors = true;
        }
    }

    fmt.Printf( "Answer: %v\n", runningSum );
}
