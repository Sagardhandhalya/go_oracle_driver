package dgodror

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func ConectOracleGodror() {
	// db, err := sql.Open("godror", "oracle://VINCENNZO:vincenzo123@104.108.154.85:1521/XE")
	db, err := sql.Open("godror", `user="VINCENNZO" password="vincenzo123" connectString="104.108.154.85:1521/XE"
	libDir="/Users/dhandhalyasagar/Downloads/instantclient_19_8"`)

	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected")
}
