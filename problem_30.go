/*Problem 30 - Digit fifth powers

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + 4^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

1) calculate 0^5, 1^5, 2^5, ..., 9^5 and store those in an array

2) Go from 10000 to 99999 and decompose into 5 numbers of fifth powers.

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */

package main

import 	(
		"fmt"
		"math"
		"strconv"
)

const POWER_TO_TEST = 5; //really only works up to 9 I think?

func main() {

    //populate an array with all the numbers raised to their fifth powers.
	var arr [ 10 ]int;
	for i := 0; i < 10; i++ {
		arr[i] = int( math.Pow( float64( i ), POWER_TO_TEST ) );
	}

	fmt.Println( arr );


	totalSum := 0;

	//This was all operating on the assumption that I went from 10000 to 99999... That was wrong of course...
	i := 1;
	maxBound := 10;
	for j := 1; j < POWER_TO_TEST; j++ {
		i *= 10;
		maxBound *= 10;
	}

	i = 2; //added this and change the for loop before from maxBound to 10000000 to just see if there were more out of my control. There were...
	for ; i < 1000000; i++ { //REFACTOR THIS.

		str := strconv.Itoa( i );

		localSum := 0;
		// for j := 0; j < POWER_TO_TEST; j++ {
		for j := 0; j < len( str ); j++ {


			// fmt.Printf( "str: %v | str[ j ]: %v | j: %v\n", str, str[j], j );

			digit, _ := strconv.Atoi( str[ j : j + 1 ] );
			// fmt.Printf( "Digit: %v | Str: %v\n", digit, str );

			localSum += arr[ digit ];
			// fmt.Printf( "Local: %v\n", arr[ digit ] );

		}
		// fmt.Printf( "Sum: %v | i: %v\n", localSum, i );


		if localSum == i {
			fmt.Printf("Hit: %v | %v | %v\n", localSum, arr, maxBound )
			totalSum += localSum;
		}

	}

	fmt.Printf( "Total Sum: %v\n", totalSum );

}

