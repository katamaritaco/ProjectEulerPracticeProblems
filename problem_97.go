/*Problem 97 - Large non-Mersenne prime

The first known prime found to exceed one million digits was discovered in 1999, and is a Mersenne prime of the form 26972593−1; it contains exactly 2,098,960 digits. Subsequently other Mersenne primes, of the form 2p−1, have been found which contain more digits.

However, in 2004 there was found a massive non-Mersenne prime which contains 2,357,207 digits: 28433×27830457+1.

Find the last ten digits of this prime number.
*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

Can probably optimize by not iterating 1 < 7830457. Can do squares and mult. together. zB: 2**19 = 2**16 * 2**2 * 2. Can find the right powers from binary conversion.

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


package main

import "fmt"

func main() {

    prime := 2;

    for i := 1; i < 7830457; i++ {
    	prime *= 2;
    	prime %= 10000000000;
    }
    fmt.Printf( "After Loop:  %v\n", prime );

    prime *= 28433;
    prime %= 10000000000;
    fmt.Printf( "After 28433: %v\n", prime );

    prime++;
    prime %= 10000000000;
    fmt.Printf( "After Add 1: %v\n", prime );

    fmt.Printf( "Massive non-Mersenne Prime: %v\n", prime );

}

