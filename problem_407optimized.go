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

    // var countOfMaxVals int64 = 0;
    // var countOfMaxValsTotal int64 = 0;


    //var sum float64 = 0;
    var sum int64 = 0;


    /*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 
	1)   DONE Calculate 1, 2, ..., 10**7 and store in an array. Should be 80mb -> totally manageable.
	1.5) DONE Instead of calculating the powers each time, iterate through the above array and use those numbers instead.

	If 1) doesn't work out super nicely, look into 2) and 3) //NOTE: Definitely does not work nicely :] Too slow.

	2) DONE(with exception of power function) Convert all types to int64.
	3) Implement Power Function for int64. [Look at Knuth's way]

	4) Leo's suggestion of "start at sqrt(n), because you need a = a*a". Might be slightly helpful, but not a large complexity change.

    //TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */

    
    squareArray := calculateSquaresBetweenRange( 1, MAX );

    //fmt.Printf( "%v\n", squareArray );

    // For every n between 1 ≤ n ≤ 10^7
    //var n float64 = 1;
    var n int64 = 1;
    for ; n <= MAX; n++ {


    	

    	//find M(n)
        //var mOfN float64 = 1; //weird name but oh well...
        var mOfN int64 = 1;
        var a int64 = 0;
        for ; a < n; a++ {
            
            // val := squareArray[a] % n;
            // if mOfN < val {
            //     mOfN = val;

            //     if mOfN == ( n - 1 ) { //early out if max value is found
            //         // fmt.Printf("FOUND MAX: %v\n", n);
            //         countOfMaxValsTotal++;
            //         countOfMaxVals++;
            //         break;
            //     }
            // }

            //I originally misread the problem. I have to take the highest a s/t a^2 mod n equals a mod n for all n from 1 <= n <= MAX
            if squareArray[a] % n == a % n {
                mOfN = a;
            }


        }

        //var a float64 = 1;
    	// for ; a < n; a++ {

    	// 	val := math.Mod( math.Pow( float64( a ), 2 ), n ) ;

    	// 	if mOfN < val {
    	// 		mOfN = val;

    	// 		if mOfN == ( n - 1 ) { //early out if max value is found
    	// 			// fmt.Printf("FOUND MAX: %v\n", n);
    	// 			break;
    	// 		}
    	// 	}

    	// }

    	sum += mOfN;

    	// if math.Mod( n, 1000 ) == 0 {
    	// 	fmt.Printf( "Line: %v\n", n );
    	// }



        if n % 1000 == 0 {
            fmt.Printf( "Line: %v\n", n );
            // fmt.Printf( "Line: %v | %v | %v\n", n, countOfMaxValsTotal, countOfMaxVals );
            // countOfMaxVals = 0;
        }

    }

	fmt.Println( sum );

}


//Time to run this is negligible. 80mb for 10 million elements should be fine as well.
func calculateSquaresBetweenRange( start int64, end int64 ) [ ]int64 {

    squareArray := make ( []int64, end - start );

    var i int64 = 0;

    for ; i < end - start; i++ {

        squareArray[ i ] = int64( math.Pow( float64( start + i ), 2 ) );//conversion won't matter even though it's janky...

    }

    fmt.Printf( "Finished at i: %v\n", i );


    return squareArray;

}
