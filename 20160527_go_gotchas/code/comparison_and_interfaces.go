package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//START COMP OMIT
	var f *os.File = nil
	var a io.Reader = f

	fmt.Printf("Interface value %v\n", a == nil)

	f2 := a.(*os.File)
	fmt.Printf("Type-asserted value %v\n", f2 == nil)
	//END COMP OMIT
}
