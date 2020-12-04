package insert

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

var db *sql.DB

var c config.Config

func init() {
	c = config.GetConfig("../../../../config.toml")
	fmt.Println(c)
	var err error
	db, err = sql.Open("mysql", c.DatabaseConfig.DatabaseUrl)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		log.Panicf("open database fail:%v", err)
	}
	fmt.Println("connect success")
}

func TestTeam(t *testing.T) {
	Team(db, &c)
}
