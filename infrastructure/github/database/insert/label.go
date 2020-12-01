package insert

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	"strings"
)

func Label(db *sql.Tx, repo *model.Repository, label *model.Label) {
	_, err := db.Exec(`
insert into label (name,repository_id) values (?,?);`,
		label.Name, repo.DatabaseID)
	if err != nil {
		fmt.Println("Insert fail while insert into label (name,repository_id)", err)
	}
}

func LabelSeverityWeight(db *sql.Tx, weight config.LabelSeverityWeight) {
	_, err := db.Exec(`
insert into label_severity_weight (label_id,weight)
select label.id, ?
from label where label.name = ?;`,
		weight.Weight, weight.LabelName)
	if err != nil {
		fmt.Println("Insert fail while insert into label_severity_weight (label_id,weight)", err)
	}
}

func LabelSig(db *sql.Tx, label *model.Label) {
	if strings.Contains(label.Name, "sig") {
		_, err := db.Exec(`
insert into label_sig (label_name)
values (?)`, label.Name)
		if err != nil {
			fmt.Println("Insert fail while insert into label_sig (label_id,label_name)", err)
		}
	}
}
