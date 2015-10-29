//hello.go base project to start from

//not the greatest way of benchmarking programs in windows:
// SET startTime=%time% && go run XXX.go && echo %startTime% && echo %time%

package main

import "fmt"

func main() {
    fmt.Printf( "hello, world!\n" );
}

