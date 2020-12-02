package truncate

import (
	"database/sql"
	"log"
)

func AllClear(db *sql.DB) {
	_, err := db.Exec(`truncate table comment;`)
	if err != nil {
		log.Fatalf("truncate table comment fail")
	}

	_, err = db.Exec(`truncate table issue;`)
	if err != nil {
		log.Fatalf("truncate table issue fail")
	}

	_, err = db.Exec(`truncate table issue_label;`)
	if err != nil {
		log.Fatalf("truncate table issue_label fail")
	}

	//_, err = db.Exec(`truncate table issue_team;`)
	//if err != nil {
	//	log.Fatalf("truncate table issue_team fail")
	//}

	_, err = db.Exec(`truncate table issue_version_affected;`)
	if err != nil {
		log.Fatalf("truncate table issue_version_affected fail")
	}

	_, err = db.Exec(`truncate table issue_version_fixed;`)
	if err != nil {
		log.Fatalf("truncate table issue_version_fixed fail")
	}

	_, err = db.Exec(`truncate table label;`)
	if err != nil {
		log.Fatalf("truncate table label fail")
	}

	_, err = db.Exec(`truncate table label_severity_weight;`)
	if err != nil {
		log.Fatalf("truncate table label_severity_weight fail")
	}

	_, err = db.Exec(`truncate table label_sig;`)
	if err != nil {
		log.Fatalf("truncate table label_sig fail")
	}

	_, err = db.Exec(`truncate table repository;`)
	if err != nil {
		log.Fatalf("truncate table repository fail")
	}

	_, err = db.Exec(`truncate table tag;`)
	if err != nil {
		log.Fatalf("truncate table tag fail")
	}

	//_, err = db.Exec(`truncate table team;`)
	//if err != nil {
	//	log.Fatalf("truncate table team fail")
	//}
	//
	//_, err = db.Exec(`truncate table team_label;`)
	//if err != nil {
	//	log.Fatalf("truncate table team_label fail")
	//}

	_, err = db.Exec(`truncate table timeline;`)
	if err != nil {
		log.Fatalf("truncate table timeline fail")
	}

	_, err = db.Exec(`truncate table timeline_repository;`)
	if err != nil {
		log.Fatalf("truncate table timeline_repository fail")
	}

	_, err = db.Exec(`truncate table user;`)
	if err != nil {
		log.Fatalf("truncate table user fail")
	}

	_, err = db.Exec(`truncate table user_issue;`)
	if err != nil {
		log.Fatalf("truncate table user_issue fail")
	}

	_, err = db.Exec(`truncate table version;`)
	if err != nil {
		log.Fatalf("truncate table version fail")
	}
}
