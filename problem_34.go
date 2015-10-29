/*Problem 34 - Digit factorials

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

Work on defining nice bound to stop calculations at.

This brute force solution clearly won't work. I'll have to actually put some thought into it :D

One thought:
iterate by digit rather than just keep going up, or some other way that we can drastically reduce our calculations

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */



package main

import "fmt"

const UPPER_BOUND = 1000000000000 //one trillion b/c #yoloswag :3

func main() {


	var sum int64 = 0;



	var i int64 = 10;
	for ; i < UPPER_BOUND; i++ {
		if digitFactorial( i ) == i {
			fmt.Println( i );
			sum += i;

		}
	}




    fmt.Printf( "Sum of all numbers: %v\n", sum );
}



func digitFactorial( x int64 ) int64 {

	var sum int64 = 0;

	for x != 0 {
		var digit int64 = x % 10;
		sum += factorial( digit );
		x /= 10;
	}

	return sum;

}




//only works for positive ints
func factorial( x int64 ) int64 {

	if x == 0 {
		return 0;
	}
	if x < 0 {
		return -1;
	}

	var product int64 = 1;
	var i int64;
	for i = 1; i <= x; i++ {
		product *= i;
	}

	return product;
}