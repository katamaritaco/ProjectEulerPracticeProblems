/*Problem 20 - Factorial digit sum

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import (
		"fmt"
		"math/big"
		"strconv"
)


func main() {

    product := big.NewInt( 1 );

    //calculate 100!
    var i int64 = 1;
    for ; i <= 100; i++ {
    	product.Mul( product, big.NewInt(i) );
    }

    //add the digits together
    sum := 0;
    for i := 0; i < len( product.String() ); i++ {
    	index, _ := strconv.Atoi( string( product.String()[i] ) );
    	sum += index;
    }

    fmt.Println( "Answer: %v\n", sum );

}