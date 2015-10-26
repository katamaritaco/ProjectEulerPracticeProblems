/*Problem 407 - Idempotents
If we calculate a^2 mod 6 for 0 ≤ a ≤ 5 we get: 0,1,4,3,4,1.

The largest value of a such that a^2 ≡ a mod 6 is 4.
Let's call M(n) the largest value of a < n such that a^2 ≡ a (mod n).
So M(6) = 4.

Find ∑M(n) for 1 ≤ n ≤ 10^7.

NOTE: If we calculate a^2 mod 8 for 0 ≤ a ≤ 5 we get: 0,1,4,1,0,1,4,2. -> answer is 0, because no indices match up.
        Doesn't go by highest value, goes by highest index where a squared is the same as a mod n.

*/


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

1)   DONE Calculate 1, 2, ..., 10**7 and store in an array. Should be 80mb -> totally manageable.
1.5) DONE Instead of calculating the powers each time, iterate through the above array and use those numbers instead.

If 1) doesn't work out super nicely, look into 2) and 3) //NOTE: Definitely does not work nicely :] Too slow.

2) DONE(with exception of power function) Convert all types to int64.
3) Implement Power Function for int64. [Look at Knuth's way] -> wouldn't be that useful since finding powers is fast (maybe done like this under the hood anyways?). Might be fun though.

4) Leo's suggestion of "start at sqrt(n), because you need a = a*a". Might be slightly helpful, but not a large complexity change. Seems like a complicated change.

5) Experiment with splitting the problem up
    5a) Goroutines -> learn more about them
    5b) Create a program that splits up problem into chunks of 10,000 or whatnot. Get sum for each part. Add together at the end.



UPDATE: This solves the problem, but is off by one. I am investigating :3

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


package main

import (
	"fmt"
	"math"
)


const MAX = 10000000; // 10 ^ 7;


func main() {

    var sum uint64 = 0;
    var numDuds int64 = 0; //how many vals of n we have with highest mOfN to be the first index.
    var numDudsPerLine int64 = 0;
    
    squareArray := calculateSquaresBetweenRange( 0, MAX );


    // For every n between 1 ≤ n ≤ 10^7
    var n int64 = 1;
    for ; n <= MAX; n++ {

        //find M(n)
        var mOfN int64 = 1; // 1^2 = 1 mod n is always true.
        var index int64 = n;
        for ; index > 0; index-- { //can stop at index > 1, but want to count amount of 'duds'

            //I have to take the highest a s/t a^2 mod n equals a mod n for all n from 1 <= n <= MAX
            //Ex 
            if squareArray[index] % n == index {
                // fmt.Printf("found a: %v | at n: %v \n", squareArray[index], n);
                mOfN = index;
                break; //early out because we found the highest index since we're counting down.
            }


        }


        /////////////////////////////////testing with the given test case in the problem statement
        // if n == 6 {
        //     fmt.Printf( "n is 6; index: %v | sq[index]: %v | line: %v\n", index, squareArray[index], n)
        //     break;
        // }


        //for us to track numDuds for curiousity's sake
        if index == 1 {
            numDuds++;
            numDudsPerLine++;
        }


    	sum += uint64( mOfN );


        ////////////////////////////////testing with some data
        if n % 10000 == 0 {
            fmt.Printf( "Ln: %v Duds: %v DudsLn: %v PerDuds: %v Sum: %v\n", n, numDuds, numDudsPerLine , ( numDuds * 100 / n ), sum );
            numDudsPerLine = 0;
        }

    }

	fmt.Printf( "Final Sum: %v\n", sum );

}


//Time to run this is negligible. 80mb for 10 million elements should be fine as well.
func calculateSquaresBetweenRange( start int64, end int64 ) [ ]int64 {

    squareArray := make ( []int64, ( end - start + 1 ) );//plus 1 since end is inclusive

    var i int64 = 0;

    for ; i <= ( end - start ) ; i++ {

        squareArray[ i ] = int64( math.Pow( float64( start + i ), 2 ) );//conversion won't matter even though it's janky...

    }

    fmt.Printf( "Calculated squares ending at i: %v\n", i - 1 );

    /////////////////////////////////////test first 10 elems
    // for j := 0; j < 10; j++ {
    //     fmt.Printf( "ARRAY: %v  %v\n", j, squareArray[j] );
    // }

    ////////////////////////////////////test out of bounds
    // fmt.Println( squareArray[0]);
    // fmt.Println( squareArray[1]);
    // fmt.Println( squareArray[2]);
    // fmt.Println( squareArray[MAX-1]);
    // fmt.Println( squareArray[MAX]);

    return squareArray;

}
