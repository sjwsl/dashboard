package process

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
)

func ProcessAll(db *sql.DB) {
	fmt.Println("Process tables...")
	file, err := ioutil.ReadFile("./infrastructure/github/database/process/process.sql")

	if err != nil {
		fmt.Printf("ReadFile error: %v\n", err)
		return
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			fmt.Printf("%v error: %v", request, err)
		}
	}
}
