package truncate

import (
	"database/sql"
	"log"
)

func AllClear(db *sql.DB) {
	_, err := db.Exec(`truncate table comment;`)
	if err != nil {
		log.Printf("truncate table comment fail: %v", err)
	}

	_, err = db.Exec(`truncate table issue;`)
	if err != nil {
		log.Printf("truncate table issue fail: %v", err)
	}

	_, err = db.Exec(`truncate table issue_label;`)
	if err != nil {
		log.Printf("truncate table issue_label fail: %v", err)
	}

	_, err = db.Exec(`truncate table team_issue;`)
	if err != nil {
		log.Printf("truncate table team_issue fail: %v", err)
	}

	_, err = db.Exec(`truncate table issue_version_affected;`)
	if err != nil {
		log.Printf("truncate table issue_version_affected fail: %v", err)
	}

	_, err = db.Exec(`truncate table issue_version_fixed;`)
	if err != nil {
		log.Printf("truncate table issue_version_fixed faiZ: %v", err)
	}

	_, err = db.Exec(`truncate table label;`)
	if err != nil {
		log.Printf("truncate table label fail: %v", err)
	}

	_, err = db.Exec(`truncate table label_severity_weight;`)
	if err != nil {
		log.Printf("truncate table label_severity_weight fail: %v", err)
	}

	_, err = db.Exec(`truncate table label_sig;`)
	if err != nil {
		log.Printf("truncate table label_sig fail: %v", err)
	}

	_, err = db.Exec(`truncate table repository;`)
	if err != nil {
		log.Printf("truncate table repository fail: %v", err)
	}

	_, err = db.Exec(`truncate table tag;`)
	if err != nil {
		log.Printf("truncate table tag fail: %v", err)
	}

	_, err = db.Exec(`truncate table team;`)
	if err != nil {
		log.Printf("truncate table team fail: %v", err)
	}

	_, err = db.Exec(`truncate table timeline;`)
	if err != nil {
		log.Printf("truncate table timeline fail: %v", err)
	}

	_, err = db.Exec(`truncate table timeline_repository;`)
	if err != nil {
		log.Printf("truncate table timeline_repository fail: %v", err)
	}

	_, err = db.Exec(`truncate table user;`)
	if err != nil {
		log.Printf("truncate table user fail: %v", err)
	}

	_, err = db.Exec(`truncate table user_issue;`)
	if err != nil {
		log.Printf("truncate table user_issue fail: %v", err)
	}

	_, err = db.Exec(`truncate table version;`)
	if err != nil {
		log.Printf("truncate table version fail: %v", err)
	}

	_, err = db.Exec(`truncate table coverage;`)
	if err != nil {
		log.Printf("truncate coverage fail: %v", err)
	}
}
