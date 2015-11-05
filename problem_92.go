/*Problem 92 - Square digit chains

A number chain is created by continuously adding the square of the digits in a number 
	to form a new number until it has been seen before.

For example,

44 → 32 → 13 → 10 → 1 → 1
85 → 89 → 145 → 42 → 20 → 4 → 16 → 37 → 58 → 89

Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop. 
What is most amazing is that EVERY starting number will eventually arrive at 1 or 89.

How many starting numbers below ten million will arrive at 89?
*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

Fast, but I can make it better.
Using current solution, I can just memoize the squares so we don't have to do those calculations.

Also, the combinatoric approach makes a lot of sense. zB: 123 is the same result as 132, 213, 231, 312, 321, 1023, etc.
Look into the multinomial coefficient in order to do this.

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


package main

import "fmt"

const MAX = 10000000 //ten million

func main() {


    //NOTE: 9,999,999, the highest number before 10,000,000 has a square digit sum of 567.
    //so the highest we'll ever get is 567.


    // make a map with keys from 0 - 567 and values being either 1 or 89
    // after you get your first number in the chain besides starting number, check to see if it is in map.
    	//if it is in the map, we know the answer and are done for this starting answer.
    	//if it isn't in the map, go to the next number.
    		//if that next number isn't 1 or 89, repeat.
    		//else if that next number is a 1 or 89, then for all previous entries, add them to the map with the approp. value.

    var chainMap map[int]int;
    chainMap = make(map[int]int);

    elemsToAddToChainMap := make( []int, 1 );//can keep this tiny because we won't need it after a while.


    numArrivingAt89 := 0;

    for i := 1; i < MAX; i++ {

        //Get the first number after the starting number
        iSquareSum := digitSquareSum( i );

        elemsToAddToChainMap = elemsToAddToChainMap[:0];

        //for until we find 89 or 1
        for {

	    	//check to see if it is in the map
	    	//if in map, we're done.
            if chainMap[iSquareSum] == 1 || iSquareSum == 1{
                iSquareSum = 1; // force this to be 1 so we can populate our map correctly
                // fmt.Printf( "Found 1: %v\n", i );
                break;
            } else if chainMap[iSquareSum] == 89 || iSquareSum == 89{ //if 89, add to numArrivingAt89. else, nada
                numArrivingAt89++;
                iSquareSum = 89; // force this to be 89 so we can populate our map correctly
                // fmt.Printf( "Found 89: %v\n", i );
                break;
            } else { //else not in map, add this val to a list (use slice instead).
                elemsToAddToChainMap = append(elemsToAddToChainMap, iSquareSum);
                // fmt.Printf( "Have to append: %v\n", iSquareSum );
            }

            //get next number in chain.
            iSquareSum = digitSquareSum( iSquareSum );

        }

    	//if slice has anything in it, populate the chainMap with vals
        for _, val := range elemsToAddToChainMap {
            chainMap[val] = iSquareSum;
        }

    }

    // //////////TESTING: testing digitSquareSum
    // fmt.Printf( "Test 123: %v | Test 4: %v | Test 24: %v | Test 9,999,999: %v | Test 0: %v\n", digitSquareSum( 123 ), digitSquareSum( 4 ), digitSquareSum( 24 ), digitSquareSum( 9999999 ), digitSquareSum( 0 ) )
    // fmt.Printf( "chainMap: %v\n", chainMap );

    fmt.Printf( "Number of starting numbers reaching 89: %v\n", numArrivingAt89 );

}


//get each digit. square that digit. add to sum.
func digitSquareSum( x int ) int {

    sumOfSquareDigits := 0;

    for x != 0 {
        sumOfSquareDigits += ( (x % 10) * (x % 10) );
        x /= 10;
    }

    return sumOfSquareDigits;
}