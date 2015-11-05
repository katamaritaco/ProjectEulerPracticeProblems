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

package main

import "fmt"

const MAX = 10000000 //ten million

func main() {
    fmt.Printf( "hello, world!\n" );


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

    numArrivingAt89 := 0;

    for i := 0; i < MAX; i++ {



    	//for until we find 89 or 1
	    	//get next number in chain.

	    	//check to see if it is in the map
	    	//if in map, we're done.
	    		//if 89, add to numArrivingAt89. else, nada

	    	//else not in map, add this val to a list (use slice instead).




    	//if slice has anything in it, populate the chainMap with vals

    }


    fmt.Printf( "Number of starting numbers reaching 89: %v\n", numArrivingAt89 );

}
