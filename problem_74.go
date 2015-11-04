/*Problem 74 - Digit factorial chains

The number 145 is well known for the property that the sum of the factorial of its digits is equal to 145:

1! + 4! + 5! = 1 + 24 + 120 = 145

Perhaps less well known is 169, in that it produces the longest chain of numbers that link back to 169; it turns out that there are only three such loops that exist:

169 → 363601 → 1454 → 169
871 → 45361 → 871
872 → 45362 → 872

It is not difficult to prove that EVERY starting number will eventually get stuck in a loop. For example,

69 → 363600 → 1454 → 169 → 363601 (→ 1454)
78 → 45360 → 871 → 45361 (→ 871)
540 → 145 (→ 145)

Starting with 69 produces a chain of five non-repeating terms, but the longest non-repeating chain with a starting number below one million is sixty terms.

How many chains, with a starting number below one million, contain exactly sixty non-repeating terms?
*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

can create a lookup table for 0! - 9! instead of calculating each time.

convert from map to array. would that be faster? Complexity says no..

Finishes in about 6 seconds. Good enough for now.

keep track of the chain sizes for a given number. if you come across it again, you'll know the size...

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


package main

import (
		"fmt"
)

const MAX = 1000000
const CHAIN_SIZE_MAX = 60;

func main() {

    // //////////Test DigitFactorialSum
    // testDigitFactorialSum := 169;
    // fmt.Printf("%v = %v\n", testDigitFactorialSum, digitFactorialSum( testDigitFactorialSum ) );
    // testDigitFactorialSum = 363601;
    // fmt.Printf("%v = %v\n", testDigitFactorialSum, digitFactorialSum( testDigitFactorialSum ) );
    // testDigitFactorialSum = 1454;
    // fmt.Printf("%v = %v\n", testDigitFactorialSum, digitFactorialSum( testDigitFactorialSum ) );

    nonRepeatingChains := 0;
    var chainMap map[int]int;

    for i := 10; i < MAX; i++ { //see if bounds should be changed to 1 or whatnot. 10 is first digital factorial. Probably doesn't matter for problem anyways,
    	
    	chainMap = make(map[int]int); // should I keep remaking or use 'init()'?
    	j, intermediateVal := 1, i;

    	for ; j < CHAIN_SIZE_MAX; j++ { //Problem states that "the longest non-repeating chain with a starting number below one million is sixty terms", so we check all.

    		intermediateVal = digitFactorialSum( intermediateVal );

    		//if intermediateVal is a repeat, break.
    		if _, doesKeyExist := chainMap[ intermediateVal ]; doesKeyExist { 
    			break;
			}

			chainMap[ intermediateVal ] = 1;

    	}

    	if j == CHAIN_SIZE_MAX {
    		nonRepeatingChains++;
    		// fmt.Printf("Found Non Repeat: %v\n", i);
    	}

    }

    fmt.Printf( "Total num of non-repeating chains: %v\n", nonRepeatingChains );

}


func digitFactorialSum( x int ) int {

	if x == 0 {
		return 1;
	}

	var sum int = 0;

	for x != 0 {
		var digit int = x % 10;
		sum += factorial( digit );
		x /= 10;
	}

	return sum;

}


//only works for positive ints
func factorial( x int ) int {

	if x == 0 {
		return 1;
	}
	if x < 0 {
		return -1;
	}

	product := 1;
	for i := 1; i <= x; i++ {
		product *= i;
	}

	return product;
}