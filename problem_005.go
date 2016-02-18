/*Problem 5 - Smallest multiples
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import "fmt"

const MAX_DIVISOR = 20;


func main() {

    //There may be a more elegant solution, but I'll go with this.
    //Try every number in increments of 20. zB: 20, 40, 60, ...

    var dividend int64 = int64( MAX_DIVISOR );

    for ; ; dividend += int64( MAX_DIVISOR ) { //until we find our answer. I suppose could set upper bound of maxint64, but we'll see if it is needed...
		
    	if smallestMultiple( dividend ) != 0 {
    		break;
    	}

    }

    fmt.Println( "End of search" );

}

func smallestMultiple( dividend int64 ) int64 {

	divisor := 1;

   	for ; divisor < MAX_DIVISOR; divisor++ {

		//fmt.Printf("%v %v \n",divisor, dividend);

   		if dividend % int64( divisor ) != 0 {
   			break;
   		}

   	}

   	if divisor == MAX_DIVISOR {

   		fmt.Printf( "Found smallest multiple: %v\n\n" , dividend );
   		return dividend;

   	}


	 return 0;
}