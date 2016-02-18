/*Problem 25 - 1000-digit Fibonacci number

The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/

package main

import (
		"fmt"
		"math/big"
)

func main() {

	x := big.NewInt( 1 );
	y := big.NewInt( 1 );

	index := 2
	for ; len( y.String() ) < 1000; index++ {

		// Compute the next Fibonacci number, storing it in a.
		x.Add( x, y );
		// Swap x and y so that y is the next number in the sequence.
		x, y = y, x;

	}

    fmt.Printf( "Index of first term that has 1000 digits: %v\n", index );

}