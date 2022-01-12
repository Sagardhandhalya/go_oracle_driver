package main

import (
	"fmt"

	"github.com/sagarsearce/oraDriver/dgodror"
	"github.com/sagarsearce/oraDriver/doci8"
)

func main() {
	fmt.Println("Main function done")
	fmt.Println("Dricer Option 1 : godror v:1.14 c:user c :")
	dgodror.ConectOracleGodror()
	doci8.ConectOracleOci8()
}
