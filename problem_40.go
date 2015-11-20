/*Problem 40 - Champernowne's constant

An irrational decimal fraction is created by concatenating the positive integers:

0.12345678910 1 112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

This is the worst code I've ever written.

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


package main

import (
		"fmt"
		"strconv"
		)

func main() {
    //take the easy way out...

   	concatNum := "";

    for i := 1; len(concatNum) <= 1000000; i++ {
        concatNum += strconv.Itoa( i );
    }

    product := 1; //d1, d10 are one.
    product *= int( concatNum[100-1] - 48 );
    product *= int( concatNum[1000-1] - 48 );
    product *= int( concatNum[10000-1] - 48 );
    product *= int( concatNum[100000-1] - 48 );
    product *= int( concatNum[1000000-1] - 48 );

    fmt.Printf("TOP: %v \n %v \n %v \n %v \n %v \n %v \n %v \n %v \n\n", product, concatNum[1-1] - 48, concatNum[10-1] - 48, concatNum[100-1] - 48, concatNum[1000-1] - 48, concatNum[10000-1] - 48, concatNum[100000-1] - 48, concatNum[1000000-1] - 48 );

}

