/*Problem 9 - Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

Different ways to approach it. I'll try brute forcing and then maybe optimizing later...

NOTE: there are off by ones and bad optimizations, but oh well for now :3

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */




package main

import 	(
		"fmt"
		"math"
)

func main() {

    //why not brute force :3

    var a float64 = 0.0;
    var b float64 = 1.0;
    var c float64 = 1.0;


    for ; a < 333; a++ { //do 333 because a must be less than b which must be less than c, and sum is 1000

    	for b = a + 1; b < 500; b++ {

    		for c = b + 1; c < 1000; c++ {

    			if math.Pow(a, 2) + math.Pow(b, 2) == math.Pow(c, 2) {
    				if (a + b + c) == 1000 {
    					fmt.Printf("Answer: %v | %v %v %v", (a*b*c), a, b, c);
    				}
    			}

    		}


    	}



    }



}


