/*Problem 3 - Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/


package main

import (
		"fmt"
		"container/list"
)

func main() {

	factors := list.New();

	var num int64 = 600851475143;

	//give all largest prime factors because we know that ever positive integer has a single unique prime factorization (fundamental theorem of arithmetic)
	var i int64 = 2;
	for ; num > 1; i++{

		for num % i == 0 {
			factors.PushBack( i );
			num /= i;
		}

	}

	// ////////////Test: print all the values out
	// // Iterate through list and print its contents.
	// for e := factors.Front(); e != nil; e = e.Next() {
	// 	fmt.Println( e.Value )
	// }

    fmt.Printf( "Highest Prime Factor: %v\n", factors.Back().Value );
}