/*Problem 32 - Pandigital products

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

Works and is pretty fast (a second or so)

Can totally speed up the program by optimizing the bounds in the loops.

Can fix the pandigital detection function to use bit shifts

Can probably clean up all the conversion stuff, but I focused on simplicity.

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


package main

import (
		"fmt"
		"strconv"
		"strings"
		"math"
)

func main() {

	//treat the hashmap as a set because golang doesn't have explicit set, and I don't want to make a struct for it...
	set := make( map[int]bool )


    for i := 1; i < 10000; i++ { //can totally reduce bounds but won't matter for speed


    	for j := 1; j < 10000; j++ { //can totally reduce bounds but won't matter for speed

    		product := i * j;

 			nums := []string{ strconv.Itoa( i ), strconv.Itoa( j ), strconv.Itoa( product ) };

 			concatNums := strings.Join( nums, "" ); //may be able to use a different method that is faster.

 			if len( concatNums ) < 9 {
 				continue;
 			} else if len( concatNums ) > 9 {
 				break;
 			}


 			x, _ := strconv.Atoi( concatNums );

    		if isPandigital( x ) {
    			set[product] = true;
    		}

    	}

    }


    var sum int64 = 0;

	for key/*, value*/ := range set {
		sum += int64( key );
    	// fmt.Println("Key:", key, "Value:", value)
	}


    fmt.Printf( "Sum of Pandigital products: %v\n", sum );


}


//will only take a number that is 9 digits long, testing numbers between 1 - 9.
func isPandigital( x int ) bool {

	var arr [ 10 ]int;

	//let's divide our number by 10^9, 10^8, ..., 10^1 and see if we have that already.
	for i := 9; i > 0; i-- {

		divisor := int( math.Pow( 10, float64( i - 1 ) ) );

		index := x / divisor;

		// fmt.Printf("  PanDigital i: %v | divisor: %v | x: %v | index: %v\n", i, divisor, x, index );

		//pandigital is 1-9, no zeros allowed.
		if index == 0 {
			return false;
		} else if arr[ index ] == 0 {
			arr[ index ] = 1;

			x -= ( index * divisor ); //don't forget to subtract the base number!

		} else {
			return false;
		}

	}

	return true;
}