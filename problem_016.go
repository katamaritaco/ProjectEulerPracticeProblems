/*Problem 16 - Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?

*/

package main

import (
		"fmt"
		"math/big"
		"strconv"
)

const POWER = 1000;

func main() {

    //1) calculate 2^1000

    two := big.NewInt( 2 );
    num := big.NewInt( 2 );

    for i := 1; i < POWER; i++ { //init at 1 because num is already 2.
    	num.Mul( num, two );
    }

    //2) add the digits together.

    str := num.String()
    sum := 0;

    for i := 0; i < len( str ); i++ {

    	sumStr, _ := strconv.Atoi( str[ i : i + 1 ] );

    	sum += sumStr;

    }

    fmt.Printf( "Power Digit Sum: %v\n", sum );

}

