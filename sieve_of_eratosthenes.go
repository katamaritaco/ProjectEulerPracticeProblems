/* Sieve of Eratosthenes


For finding all the small primes, say all those less than 10,000,000,000; one of the most efficient ways is by using the Sieve of Eratosthenes (ca 240 BC):

Make a list of all the integers less than or equal to n (greater than one) and strike out the multiples of all primes less than or equal to the square root of n, 
then the numbers that are left are the primes. (See also our glossary page.)

For example, to find all the odd primes less than or equal to 100 we first list the odd numbers from 3 to 100 (why even list the evens?)  
The first number is 3 so it is the first odd prime--cross out all of its multiples.  Now the first number left is 5, the second odd prime--cross out all of its multiples.  

Repeat with 7 and then since the first number left, 11, is larger than the square root of 100, all of the numbers left are primes.

*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

NOT DONE AT ALL....

Probably really slow with the list interface, but oh well...

Yeah, Need to convert this to slices for sure...

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


package main

import (
		"fmt"
		"container/list"
		"math"
)

const NUMS_LIST = 100; //ten million

func main() {
    fmt.Printf( "hello, world!\n" );

    nums := list.New();

    //init list containing odd ints from 3 ... NUMS_LIST
    for i := 3; i <= NUMS_LIST; i += 2 {

    	nums.PushBack( i );

    }


    //remove all multiples of the primes
    bound := math.Sqrt( NUMS_LIST );
    fmt.Println( bound );

    isInBounds := true;


	for e := nums.Front(); e != nil; e = e.Next() {
		// do something with e.Value
	}




}