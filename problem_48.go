/*Problem 48 - Self powers

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
*/


package main

import (
        "fmt"
        "math/big"
)

const MAX = 1000;

func main() {
    fmt.Printf( "hello, world!\n" );

    sum := big.NewInt( 0 );

    zero := big.NewInt( 0 );

    var i int64;
    for i = 1; i <= MAX; i++ {
        num := big.NewInt( i );
        sum.Add( sum, num.Exp( num, num, zero ) );
    }

    fmt.Printf( "Last ten digits: %v\n", sum.String()[ len( sum.String() ) - 10 : len( sum.String() ) ] );//todo make this last ten

}

