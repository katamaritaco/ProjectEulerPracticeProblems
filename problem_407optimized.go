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


    /*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 
	1)   Calculate 1, 2, ..., 10**7 and store in an array. Should be 80mb -> totally manageable.
	1.5) Instead of calculating the powers each time, iterate through the above array and use those numbers instead.

	If 1) doesn't work out super nicely, look into 2) and 3)

	2) Convert all types to int64.
	3) Implement Power Function for int64. [Look at Knuth's way]

	4) Leo's suggestion of "start at sqrt(n), because you need a = a*a". Could be helpful if I'm misreading the question. Could give false answers if he is misreading the question.

    //TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


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
    				break;
    			}
    		}

    	}

    	sum += mOfN;

    	if math.Mod( n, 1000 ) == 0 {
    		fmt.Printf( "Line: %v\n", n );
    	}

    }

	fmt.Println( sum );

}
