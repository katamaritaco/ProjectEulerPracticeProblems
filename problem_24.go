/*Problem 24 - Lexicographic permutations

A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. 
If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

Doesn't work yet. 

1) Work out on paper with smaller example

Work out kinks using general algorithmic though, unless I'm somehow super off track...

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


package main

import (
		"fmt"
		"container/list"
)

const DIGIT_COUNT = 10;
const INDEX_TO_FIND = 1000000;

func main() {

	permutationCounter := INDEX_TO_FIND;
	restAreZeros := false;

	digitsLeft := list.New()
	digitsLeft.PushBack( 0 );
	digitsLeft.PushBack( 1 );
	digitsLeft.PushBack( 2 );
	digitsLeft.PushBack( 3 );
	digitsLeft.PushBack( 4 );
	digitsLeft.PushBack( 5 );
	digitsLeft.PushBack( 6 );
	digitsLeft.PushBack( 7 );
	digitsLeft.PushBack( 8 );
	digitsLeft.PushBack( 9 );


	for i := 1; i <= DIGIT_COUNT; i++ { // look for off by ones...

		if restAreZeros {
			// fmt.Printf( " Index %v: %v \n", i, 0 );
			fmt.Printf("0");
			continue;
		}

		iIterFactorial := factorial( DIGIT_COUNT - i );
		// fmt.Printf( "Index to find: %v Factorial: %v\n", permutationCounter, iIterFactorial );
		
		jIter := permutationCounter;
		for j := 1; j <= DIGIT_COUNT; j++ {

			jIter -= iIterFactorial;
			if jIter < 0 { //off by one? less than or equal?
				// fmt.Printf( " Index %v: %v \n", i, j - 1 );
				fmt.Printf( "%v", j - 1 );


				permutationCounter -= ( iIterFactorial * ( j - 1 ) );
				break;
			} else if jIter == 0 {
				// fmt.Printf( " Index LOL %v: %v\n", i, j );
				fmt.Printf( "%v", j )
				restAreZeros = true;
				break;
			}

		}

	}



	fmt.Println();



}


//only works for positive ints
func factorial( x int ) int {

	if x == 0 {
		return 0;
	}
	if x < 0 {
		return -1;
	}

	product := 1;
	for i := 1; i <= x; i++ {
		product *= i;
	}

	return product;
}