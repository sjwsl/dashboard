package util

import (
	"database/sql"
	"fmt"
	"time"

	util2 "github.com/PingCAP-QE/dashboard/infrastructure/github/processing/util"
	model2 "github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/model"
)

func GetIssueClosedTime(closed bool, closeAt *time.Time) sql.NullTime {
	ct := sql.NullTime{}
	if closed {
		ct = sql.NullTime{
			Time:  *closeAt,
			Valid: true,
		}
	}
	return ct
}

func GetIssueClosedWeek(closed bool, closeAt *time.Time) sql.NullTime {
	ct := sql.NullTime{}
	if closed {
		ct = sql.NullTime{
			Time:  util2.ParseDate(*closeAt),
			Valid: true,
		}
	}
	return ct
}

func GenIDFromVersion(version model2.Version) (int, error) {
	if version.Code == model2.Regular {
		if err := CheckVersionSize(&version); err == nil {
			var ID int
			ID += version.Patch
			ID += version.Minor << 10
			ID += version.Major << 20
			return ID, nil
		} else {
			return -1, err
		}
	} else {
		return -1, fmt.Errorf("only regular code version can genid")
	}
}

func CheckVersionSize(version *model2.Version) error {
	if version.Major >= 0 &&
		version.Major < 1024 &&
		version.Minor >= 0 &&
		version.Minor < 1024 &&
		version.Patch >= 0 &&
		version.Patch < 1024 {
		return nil
	} else {
		return fmt.Errorf("version code can only in [0,1024),but get %d,%d,%d",
			version.Major, version.Minor, version.Patch)
	}
}
