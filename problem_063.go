/*Problem 63 - Powerful digit counts
The 5-digit number, 16807=7^5, is also a fifth power. Similarly, the 9-digit number, 134217728=8^9, is a ninth power.

How many n-digit positive integers exist which are also an nth power?.
*/

package main

import (
    "fmt"
    "math"
    // "strconv"
    // "unicode/utf8"
)

func main() {

    count := 0;

    var numOfPowerfulDigits float64 = 0;

    var i float64 = 1
    for ; i < 10; i++ { //10 chosen because if you have 10 or more, you will have one digit extra...

        var j float64 = 1
        for ; j < 100; j++ { //100 chosen as upper bound because when you print them off, trend is clearly asymptotic

            var resultant float64 = math.Pow( i, j );

            //check to see if the resultant has j digits
            if resultant >= math.Pow( 10, j - 1) && resultant < math.Pow( 10, j ) {
                // fmt.Printf( "%v %v %v\n", i, j, resultant );

                numOfPowerfulDigits++;

                count += 2;
            }

        }

    }

    fmt.Println( numOfPowerfulDigits );

}
