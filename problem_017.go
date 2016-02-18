/*Problem 17 - Number letter counts
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters 
  and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/

package main

import "fmt"

func main() {
    fmt.Printf( "hello, world!\n" );


    one, two, six, ten := 3, 3, 3, 3;
    four, five, nine := 4, 4, 4;
    three, seven, eight := 5, 5, 5;

    eleven, twelve := 6, 6;
    fifteen, sixteen := 7, 7;
    thirteen, fourteen, eighteen, nineteen := 8, 8, 8, 8;
    seventeen := 9;

    forty, fifty, sixty := 5, 5, 5;
    twenty, thirty, eighty, ninety := 6, 6, 6, 6;
    seventy := 7;

    hundred := 7;

    one_thousand := 11;

    ten_through_nineteen := ten + eleven + twelve + thirteen + fourteen + fifteen + sixteen + seventeen + eighteen + nineteen;

    one_through_nine := one + two + three + four + five + six + seven + eight + nine;

    one_through_nineteen := one_through_nine + ten_through_nineteen;

    //no classifier for large digit. "zB hundred and 3"
    twenty_through_ninety_nine := 
    ((10 * twenty) + one_through_nine) + 
    ((10 * thirty) + one_through_nine) + 
    ((10 * fourty) + one_through_nine) + 
    ((10 * fifty) + one_through_nine) + 
    ((10 * sixty) + one_through_nine) + 
    ((10 * seventy) + one_through_nine) + 
    ((10 * eighty) + one_through_nine) + 
    ((10 * ninety) + one_through_nine);

    one_through_ninety_nine := one_through_nineteen + twenty_through_ninety_nine;


    //new...
    all_tens_and_ones := one_through_ninety_nine * 10;

    //# letters i.e one two six; amount of items of length 'x' ie 3; and 99 'and's per hundred set. size three for all coincidentally.
    three_sized_hundreds := (3 + hundred) * 100; three_sized_hundreds *= 3; three_sized_hundreds += (3*3*99);

    four_sized_hundreds := (4 + hundred) * 100; four_sized_hundreds *= 3; four_sized_hundreds += (3*3*99);

    five_sized_hundreds := (5 + hundred) * 100; five_sized_hundreds *= 3; five_sized_hundreds += (3*3*99);


    fmt.Println( one_through_nine );
    fmt.Println( ten_through_nineteen );
    fmt.Println( one_through_nineteen );
    fmt.Println( twenty_through_ninety_nine );
    fmt.Println( one_through_ninety_nine );

    //the answer :3
    fmt.Println( all_tens_and_ones + three_sized_hundreds + four_sized_hundreds + five_sized_hundreds + one_thousand);
}



