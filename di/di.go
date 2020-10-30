package main

import (
    "database/sql"
    "fmt"
    "github.com/PingCAP-QE/libs/di"
    "log"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

var issueDB, diDB *sql.DB

func init() {
    var err error
    diDB, err = sql.Open("mysql", os.Getenv("DI_DSN"))
    if err != nil {
        log.Fatalf("Open diDB: %s", err.Error())
    }
    diDB.SetMaxIdleConns(10)
    diDB.SetMaxOpenConns(10)
    issueDB, err = sql.Open("mysql", os.Getenv("GITHUB_DSN"))
    if err != nil {
        log.Fatalf("Open issueDB: %s", err.Error())
    }
    issueDB.SetMaxIdleConns(10)
    issueDB.SetMaxOpenConns(10)
}

func insertCreatedDI(tx *sql.Tx, startTime, endTime time.Time, di float64) error {
    _, err := tx.Exec(`INSERT INTO CREATED_DI(START_TIME, END_TIME, DI) VALUES(?, ?, ?)`, startTime, endTime, di)
    return err
}

func insertClosedDI(tx *sql.Tx, startTime, endTime time.Time, di float64) error {
    _, err := tx.Exec(`INSERT INTO CLOSED_DI(START_TIME, END_TIME, DI) VALUES(?, ?, ?)`, startTime, endTime, di)
    return err
}

func insertDI(tx *sql.Tx, time time.Time, di float64) error {
    _, err := tx.Exec(`INSERT INTO DI(TIME, DI) VALUES(?, ?)`, time, di)
    return err
}

func main() {
    //f, _ := os.Create("profile")
    //pprof.StartCPUProfile(f)
    //defer pprof.StopCPUProfile()

    start := time.Now()

    tx, err := diDB.Begin()
    if err != nil {
        log.Fatalf("diDB Begin: %s", err.Error())
    }

    startTime := time.Date(2015, 10, 5, 0, 0, 0, 0, time.UTC)

    remainDI, err := di.GetDI(issueDB, startTime)
    if err != nil {
        log.Fatalf("GetDI: %s", err.Error())
    }

    if err := insertDI(tx, startTime, remainDI); err != nil {
        log.Printf("insertDI: %s", err.Error())
    }

    for startTime.Before(time.Now()) {
        endTime := startTime.AddDate(0, 0, 7)

        createdDI, err := di.GetCreatedDIBetweenTime(issueDB, startTime, endTime)
        if err != nil {
            log.Fatalf("GetCreatedDIBetweenTime: %s", err.Error())
        }
        if err := insertCreatedDI(tx, startTime, endTime, createdDI); err != nil {
            log.Printf("insertCreatedDI: %s", err.Error())
        }

        closedDI, err := di.GetClosedDIBetweenTime(issueDB, startTime, endTime)
        if err != nil {
            log.Fatalf("GetClosedDIBetweenTime: %s", err.Error())
        }
        if err := insertClosedDI(tx, startTime, endTime, closedDI); err != nil {
            log.Printf("insertClosedDI: %s", err.Error())
        }

        remainDI += createdDI - closedDI
        if err := insertDI(tx, endTime, remainDI); err != nil {
            log.Printf("insertDI: %s", err.Error())
        }

        startTime = endTime
    }

    if err = tx.Commit(); err != nil {
        log.Fatalf("tx Commit: %s", err.Error())
    }

    fmt.Printf("%v\n%v", start, time.Now())

}
