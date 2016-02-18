/*Problem 14 - Longest Collatz sequence
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import "fmt"

func main() {

    var max int64 = 1000000;

    var largestChainLength int64 = 0;
    var largestChainIndex int64 = 0;

    var i int64 = 1;

    for ; i < max; i++ {

    	var j int64 = i;
    	var k int64 = 0;
    	for j != 1 {

    		if j % 2 == 0 {
    			j = ( j / 2 );
    		} else {
    			j = ( ( 3 * j ) + 1 );
    		}

			k++;

    	}

    	if largestChainLength < k {
    		largestChainLength = k;
    		largestChainIndex = i;
    	}

    }

    fmt.Println( largestChainIndex );

}