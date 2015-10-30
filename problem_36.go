/*Problem 36 - Double-base palindromes

The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

Implement the isPalindrome function.

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


package main

import (
		"fmt"
		"strconv"
)

const MAX = 1000000;

func main() {
    fmt.Printf( "hello, world!\n" );

    //Easy way to do it is make a palindrome detector. loop through all numbers from 1 - 1,000,000; 
    //Take each palindrome in base ten, convert to binary, then check.

    sum := 0;

    for i := 0; i < MAX; i++ {

    	if isPalindrome( strconv.Itoa( i ) ) && isPalindrome( strconv.FormatInt( i, 2 ) ) {
    		sum += i;
    	}

    }

    fmt.Printf( "Sum: %v\n", sum );

}

//TODO
func isPalindrome( s string ) bool {

	return false;
}
