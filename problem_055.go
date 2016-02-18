/* Problem 55 - Lychrel numbers

If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.

Not all numbers produce palindromes so quickly. For example,

349 + 943 = 1292,
1292 + 2921 = 4213
4213 + 3124 = 7337

That is, 349 took three iterations to arrive at a palindrome.

Although no one has proved it yet, it is thought that some numbers, like 196, never produce a palindrome. 
A number that never forms a palindrome through the reverse and add process is called a Lychrel number. 
Due to the theoretical nature of these numbers, and for the purpose of this problem, we shall assume that a number is Lychrel until proven otherwise. 
In addition you are given that for every number below ten-thousand, it will either 
(i) become a palindrome in less than fifty iterations, or, 
(ii) no one, with all the computing power that exists, has managed so far to map it to a palindrome. 

In fact, 10677 is the first number to be shown to require over fifty iterations before producing a palindrome: 4668731596684224866951378664 (53 iterations, 28-digits).

Surprisingly, there are palindromic numbers that are themselves Lychrel numbers; the first example is 4994.

How many Lychrel numbers are there below ten-thousand?

NOTE: Wording was modified slightly on 24 April 2007 to emphasise the theoretical nature of Lychrel numbers.
*/

package main

import (
		"fmt"
		"math"
		"strconv"
)

const LOOP_BOUND = 10000;
const LYCHREL_UPPER_BOUND = 50;

func main() {

	lychrelNums := 0;

	for i := 1; i < LOOP_BOUND; i++ {

		lychrelTest := i;
		isLychrel := true;
		for j := 1; j < LYCHREL_UPPER_BOUND; j++ {

			reversedLychrel, _ := strconv.Atoi( Reverse( strconv.Itoa(lychrelTest) ) );

			lychrelTest += reversedLychrel;

			if isPalindromeNumber( lychrelTest ) { //this check comes after preforming operation because of example in prompt: 4994
				isLychrel = false;
				break;
			}

		}

		if isLychrel {
			lychrelNums++;
		}

	}

	fmt.Printf( "Num of Lychrel Nums: %v\n", lychrelNums );

}


func isPalindromeNumber( palindrome int ) bool {

	palindromeString := strconv.Itoa( palindrome );

	frontPointer := 0;
	backPointer := len( palindromeString ) - 1;

	midPoint := int ( math.Ceil( float64( backPointer ) / 2.0 ) );

	for ; frontPointer < midPoint ; {

		// If the char at the front pointer is not equal to the char at the back pointer, 
		//   we can early out because the number isn't a palindrome.
		if string( [] rune ( palindromeString ) [ frontPointer  ] ) != string( [] rune ( palindromeString ) [ backPointer ] ) {
			return false;
		}

		frontPointer++;
		backPointer--;
	}

	return true;

}


//Totally copypasta'd from SO: http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
//Kind of a bummer there isn't a builtin function in the string library...
func Reverse( s string ) string {
    runes := []rune( s )
    for i, j := 0, len( runes ) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string( runes )
}