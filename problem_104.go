/*Problem 104 - Pandigital Fibonacci ends
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
It turns out that F541, which contains 113 digits, is the first Fibonacci number for which the last nine digits are 1-9 pandigital (contain all the digits 1 to 9, but not necessarily in order). 
And F2749, which contains 575 digits, is the first Fibonacci number for which the first nine digits are 1-9 pandigital.

Given that Fk is the first Fibonacci number for which the first nine digits AND the last nine digits are 1-9 pandigital, find k.
*/

/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

This is super slow, so I'll need to optimize... Still gets correct answer though...

Can also convert pandigital function to use bit shifts...

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */

package main

import (
		"fmt"
		"math/big"
		"math"
		"strconv"
)

const NUM_DIGITS_TO_IGNORE = 10;

//NOTE: for how I am solving, the first number that has first nine digits be pandigital is at index 2745, 4 less than what the above says...

func main() {

	x := big.NewInt( 1 );
	y := big.NewInt( 2 );

	dummy := big.NewInt( 0 );

	zero := big.NewInt( 0 );

	digitsToIgnoreSum := big.NewInt( 1 );
	for i := 1; i < NUM_DIGITS_TO_IGNORE; i++ {
		digitsToIgnoreSum.Mul( digitsToIgnoreSum, big.NewInt( 10 ) );
	}

	for i := 0; ; i++ {

		if i == math.MaxInt32 { //MAX INT
			fmt.Println( "Increase your loop..." );
			break;
		}

		// fmt.Printf( "%v   %v\n", x, y );

		// Compute the next Fibonacci number, storing it in a.
		x.Add( x, y )
		// Swap a and b so that b is the next number in the sequence.
		x, y = y, x;

		// fmt.Printf( "%v   %v\n", x, y );



		// //////////Testing: find the first number where LAST 9 digits are pandigital
		// if i == 537 {
		// 	fmt.Printf( "Line Break: %v %v %v\n", x, y, i );
		// 	break;
		// }

		// //////////Testing: find the first number where FIRST 9 digits are pandigital
		// if i == 2745 {
		// 	fmt.Printf( "Line Break: %v %v %v\n", x, y, i );
		// 	break;
		// }

		//if we don't have 575 digits, continue. This is because the first number with the first numbers that are pandigital.
		if zero.Cmp( dummy.Div( y, digitsToIgnoreSum ) ) == 0 {
			// fmt.Printf( "Comparison: %v %v %v\n", x, y, i );
			continue;
		}

		// //since we know the first number that has initial pandigital params starts here.
		// if i < 2745 {
		// 	continue;
		// }



		//if z is pandigital for first 9 digits
		//Could do this by converting to string. Could do by dividing by digits of z - 9. Probably more...

		arg, _ := strconv.Atoi( y.String()[ 0 : 9 ] );
		if isPandigital( arg ) {

			//if z is pandigital for last 9 digits
			arg, _ = strconv.Atoi( y.String()[ len( y.String() ) - 9 : ] );
			if isPandigital( arg ) {
				//Don't forget to add four to this answer, since my counts are offsetted :3
				fmt.Printf( "Found Initial and Final Pandigital Index at: %v\n\n%v\n", y, i );
				break;
			}
		}

		// //////////Testing: find the numbers where the FIRST 9 digits are pandigital.
		// arg0, _ := strconv.Atoi( y.String()[ 0 : 9 ] );
		// if isPandigital( arg0 ) {
		// 	fmt.Printf( "Found Initial Pandigital At: %v | %v\n", i, arg0 );
		// }

		// //////////Testing: find the numbers where the LAST 9 digits are pandigital.
		// arg1, _ := strconv.Atoi( y.String()[ len( y.String() ) - 9 : ] );
		// if isPandigital( arg1 ) {
		// 	fmt.Printf( "Found Final Pandigital At: %v | %v\n", i, arg1 );
		// }

	}

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