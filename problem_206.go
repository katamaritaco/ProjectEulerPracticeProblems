/*Problem 206 - Concealed Square

Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0,
where each “_” is a single digit.
*/

package main

import (
		"fmt"
		"math/big"
)


/*TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO 

Took about 7ish minutes...

Can we use something like binary search? We know that number must be between 1 trillion and 10 trilion guaranteed, because it needs to be 19 zeros.
	We can cut that lower AND upper bound down further for sure. 

Might be able to convert to int64... Look into that as well to speed up operations

//TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO */


const STARTING_INDEX = 1000000000;

func main() {

	zero := big.NewInt( 0 );
	increm := big.NewInt( 1 );
    square := big.NewInt( 2 );

    // // 1_2_3_4_5_6_7_8_9_0 //1 amd 18 zeros

    i := big.NewInt( STARTING_INDEX ); 
    
    product := big.NewInt( 3 );

    for {
        product.Exp( i, square, zero );

        if isConcealedSquare( product ) {
        	fmt.Printf( "%v squared is %v\n", i, product );
        	break;
        }

        if len( product.String() ) > 19 {
        	fmt.Printf( "Whoops, somehow we have a number that is too big. Num: %v at index %v\n", product, i );
        	break;
        }

        i.Add( i, increm );
    }

    // //////////TESTING: ConcealedSquare
    // // 1020304050607080900
    // testIsConcealedSquare := big.NewInt ( 1 );
    // fmt.Println(testIsConcealedSquare.String());
    // testIsConcealedSquare.SetString("1020304050607080900", 10);
    // fmt.Println(testIsConcealedSquare.String());
    // isConcealedSquare(testIsConcealedSquare)

}

func isConcealedSquare( i *big.Int ) bool {

	str := i.String();

	if len( str ) != 19 {
		return false;
	}

	// fmt.Printf("19 chars: %v | %v %v %v %v %v %v %v %v %v %v\n", i, str[0:1], str[2:3],str[4:5] ,str[6:7] ,str[8:9] ,str[10:11] ,str[12:13] ,str[14:15] ,str[16:17] ,str[18:19]);

	if str[0:1] == "1" && 
	   str[2:3] == "2" && 
	   str[4:5] == "3" && 
	   str[6:7] == "4" && 
	   str[8:9] == "5" && 
	 str[10:11] == "6" &&  
	 str[12:13] == "7" && 
	 str[14:15] == "8" && 
	 str[16:17] == "9" && 
	 str[18:19] == "0"{
	 	// fmt.Printf("19 chars: %v\n", i]);
	 	return true;
	} 
	return false;
}