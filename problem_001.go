/*Problem 1 - Multiples of 3 and 5
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import "fmt"

func main() {

    max := 1000;
    sum := 0;

    //add multiples of 3 and 5
    for i := 0; i < max; i++ {
    	if i % 3 == 0 {
    		sum += i;
    		continue;
    	}
    	if i % 5 == 0 {
    		sum += i;
    	}
    }

    fmt.Println( sum );

}
