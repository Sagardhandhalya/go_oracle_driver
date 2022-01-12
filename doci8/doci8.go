package doci8

import (
	"fmt"
	//_ "github.com/mattn/go-oci8"
)

func ConectOracleOci8() {
	// db, err := sql.Open("godror", "oracle://VINCENNZO:vincenzo123@104.108.154.85:1521/XE")
	// db, err := sql.Open("godror", `user="VINCENNZO" password="vincenzo123" connectString="104.108.154.85:1521/XE"
	// libDir="/Users/dhandhalyasagar/Downloads/instantclient_19_10"`)

	// if err != nil {
	// 	panic(err.Error())
	// }
	// err = db.Ping()
	// if err != nil {
	// 	panic(err.Error())
	// }
	fmt.Println("connected")
}
