/*Problem 4 - Largest palindrome product
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
	"math"
)

func main() {

	//3-digit numbers: less than 1000.
    maxDigitNumber := 1000;
    largestPalindrome := 0;

    for x := 100; x < maxDigitNumber; x++ {

    	for y := 100; y < maxDigitNumber; y++ {

    		//store the product of x and y into product
    		product := x * y;

    		if isPalindromeNumber( product ) {

    			if largestPalindrome < product {

    				largestPalindrome = product;

    			}

    		}

    	}

    }

    fmt.Println( largestPalindrome );

}


func isPalindromeNumber( palindrome int ) bool {

	palindromeString := strconv.Itoa( palindrome );

	frontPointer := 0;
	backPointer := utf8.RuneCountInString( palindromeString ) - 1;

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