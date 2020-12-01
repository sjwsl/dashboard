package util

import (
	"fmt"
	model2 "github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/model"
	"testing"
)

func TestGenIDFromVersion(t *testing.T) {
	version := model2.Version{
		Major: 1023,
		Minor: 1023,
		Patch: 1023,
		Code:  0,
	}
	ID, err := GenIDFromVersion(version)
	if err != nil {
		panic(err)
	}
	if ID != 0x3FFFFFFF {
		panic(fmt.Errorf("ID error"))
	}
}
