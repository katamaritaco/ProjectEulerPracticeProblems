/*Problem43 - Sub-string divisibility

The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

d2d3d4=406 is divisible by 2
d3d4d5=063 is divisible by 3
d4d5d6=635 is divisible by 5
d5d6d7=357 is divisible by 7
d6d7d8=572 is divisible by 11
d7d8d9=728 is divisible by 13
d8d9d10=289 is divisible by 17
Find the sum of all 0 to 9 pandigital numbers with this property.

*/

/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

This is super slow, so I'll need to optimize... Still gets the correct answer as is.

1) Since the last three digits need to be divisible by 17, we can increase by 17 rather than 1 in our for loop.

2) Likewise, there are a lot more of these little tricks we can do -> zB adding 1000000 (six zeros) to i when we get a d4 that is odd (ex 1405000000), since we know that d4 must be even.
	With these tricks, we'll have to manage them carefully...

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


package main

import (
		"fmt"
		"math"
		"strconv"
)

func main() {

	var sum int64 = 0;
	per := 1;

	for i := 1000000000; i < 10000000000; i++ { // nine zeros, ten zeros

		if isPandigital( i ) {
			if isFollowProperty ( i ) {
				sum += int64( i );
			}
		}

		//Only for seeing how far we are...
		if i % 100000000 == 0 {
			fmt.Printf( "%v percent done\n", per )
			per++;
		}


	}

	fmt.Printf( "Sum: %v\n", sum );

	//////////Testing: see if functions work correctly
    fmt.Printf( "%v | %v\n", isPandigital( 1406357289 ), isFollowProperty( 1406357289 ) );
}


func isFollowProperty( x int ) bool {

	str := strconv.Itoa( x );

	//d2d3d4=406 is divisible by 2
	// fmt.Printf( "2: %v\n", str[1:4] )
	x, _ = strconv.Atoi( str[1:4] );
	if x % 2 != 0 {
		return false;
	}

	//d3d4d5=063 is divisible by 3
	// fmt.Printf( "3: %v\n", str[2:5] )
	x, _ = strconv.Atoi( str[2:5] );
	if x % 3 != 0 {
		return false;
	}

	//d4d5d6=635 is divisible by 5
	// fmt.Printf( "5: %v\n", str[3:6] )		
	x, _ = strconv.Atoi( str[3:6] );
	if x % 5 != 0 {
		return false;
	}

	//d5d6d7=357 is divisible by 7
	// fmt.Printf( "7: %v\n", str[4:7] )
	x, _ = strconv.Atoi( str[4:7] );
	if x % 7 != 0 {
		return false;
	}

	//d6d7d8=572 is divisible by 11
	// fmt.Printf( "11: %v\n", str[5:8] )
	x, _ = strconv.Atoi( str[5:8] );
	if x % 11 != 0 {
		return false;
	}

	//d7d8d9=728 is divisible by 13
	// fmt.Printf( "13: %v\n", str[6:9] )
	x, _ = strconv.Atoi( str[6:9] );
	if x % 13 != 0 {
		return false;
	}

	//d8d9d10=289 is divisible by 17
	// fmt.Printf( "17: %v\n", str[7:10] )
	x, _ = strconv.Atoi( str[7:10] );
	if x % 17 != 0 {
		return false;
	}

	return true;

}


//will only take a number that is 10 digits long.
func isPandigital( x int ) bool {

	var arr [ 10 ]int;

	//let's divide our number by 10^9, 10^8, ..., 10^1 and see if we have that already.
	for i := 10; i > 0; i-- {

		divisor := int( math.Pow( 10, float64( i - 1 ) ) );

		index := x / divisor;

		// fmt.Printf("  PanDigital i: %v | divisor: %v | x: %v | index: %v\n", i, divisor, x, index );

		if arr[ index ] == 0 {
			arr[ index ] = 1;

			x -= ( index * divisor ); //don't forget to subtract the base number!

		} else {
			return false;
		}

	}

	return true;
}