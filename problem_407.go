/*Problem 407 - Idempotents
If we calculate a^2 mod 6 for 0 ≤ a ≤ 5 we get: 0,1,4,3,4,1.

The largest value of a such that a^2 ≡ a mod 6 is 4.
Let's call M(n) the largest value of a < n such that a^2 ≡ a (mod n).
So M(6) = 4.

Find ∑M(n) for 1 ≤ n ≤ 10^7.
*/

package main

import (
    "fmt"
    "math"
)

const MAX = 10000000; // 10 ^ 7;

//Using float64 because the largest possible number, 10 ^ 14 is within the bounds of float64 with no precision loss...
func main() {

    var sum float64 = 0;

    var n float64 = 1;

    //only for me...
    var countOfMaxValsTotal float64 = 0;
    var countOfMaxVals float64 = 0;

    // For every n between 1 ≤ n ≤ 10^7
    for ; n <= MAX; n++ {

        var mOfN float64 = 1; //weird name but oh well...

        var a float64 = 1;

        //find M(n)
        for ; a < n; a++ {

            val := math.Mod( math.Pow( float64( a ), 2 ), n ) ;

            if mOfN < val {
                mOfN = val;

                if mOfN == ( n - 1 ) { //early out if max value is found
                    // fmt.Printf("FOUND MAX: %v\n", n);
                    countOfMaxValsTotal++;
                    countOfMaxVals++;
                    break;
                }
            }

        }

        sum += mOfN;

        if math.Mod( n, 1000 ) == 0 {
            fmt.Printf( "Line: %v %v %v\n", n, countOfMaxVals, countOfMaxValsTotal );
            countOfMaxVals = 0;
        }

    }

    fmt.Println( sum );

}
