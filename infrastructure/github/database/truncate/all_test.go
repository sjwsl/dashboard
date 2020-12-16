package truncate

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAllClear(t *testing.T) {
	db, err := sql.Open("mysql", os.Getenv("GITHUB_DSN"))
	if err != nil {
		t.Fatal()
	}
	AllClear(db)
}
