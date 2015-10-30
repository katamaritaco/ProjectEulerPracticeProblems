/*Problem 21 - Amicable numbers

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. 
The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

import "fmt"

const MAX = 10000;

func main() {

    totalSum := 0;

    for i := 1; i < MAX; i++ { // 1 won't count

    	divisorSum := sumOfProperDivisors( i );

    	if divisorSum != i && i == sumOfProperDivisors( divisorSum ) {
    		totalSum += i;
    	}

    }

    fmt.Printf( "Sum of amicable numbers: %v\n", totalSum );

}


func sumOfProperDivisors( num int ) int {
	sum := 0;

	for i := 1; i < num; i++ {
		if num % i == 0 {
			sum += i;
		}
	}

	return sum;
}