/*Problem 56 -Powerful digit sum

A googol (10^100) is a massive number: one followed by one-hundred zeros; 100^100 is almost unimaginably large: one followed by two-hundred zeros. 
Despite their size, the sum of the digits in each number is only 1.

Considering natural numbers of the form, a^b, where a, b < 100, what is the maximum digital sum?
*/

package main

import (
		"fmt"
		"math/big"
		"strconv"
)

const MAX = 100;

func main() {

    fmt.Printf( "hello, world!\n" );


    maxDigitSum := 0;
    zero := big.NewInt( 0 );
    num := big.NewInt( 0 );

    var a int64 = 1;
    for ; a < MAX; a++{

    	var b int64 = 1;
    	for ; b < MAX; b++{
    		num.Exp( big.NewInt( a ), big.NewInt( b ), zero );

    		localSum := 0;
		    for i := 0; i < len( num.String() ); i++ {
		    	index, _ := strconv.Atoi( string( num.String()[i] ) );
		    	localSum += index;
		    }

		    // fmt.Printf( "A: %v | B: %v | num: %v | loc: %v\n", a, b, num, localSum)

		    if maxDigitSum < localSum {
		    	// fmt.Printf( "maxDigitSum < localSum: %v | %v\n", maxDigitSum, localSum );
		    	maxDigitSum = localSum;
		    }

    	}

    }

    fmt.Printf( "Max Digital Sum: %v\n", maxDigitSum );

}