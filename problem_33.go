/*Problem 33 - Digit cancelling fractions

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/



package main

import "fmt"

func main() {
    fmt.Printf( "hello, world!\n" );

    a := 10;
    b := 10;

    //brute force at first
    //for all a/b where a < b, until b=100;
    //check to see if a digit matches in both a and b
    	//for example, a = 14 and b = 71. a '1' matches.
    	//if this is the case, then check 4 / 7, or the fraction wihtout the matches.
    		//either reduce this (look into programming this), or compare value (be weary of small mismatches)
    			//if the two fractions are equivalent, save the fraction.


    //when we have all four fractions, we can find the denominator.
}

