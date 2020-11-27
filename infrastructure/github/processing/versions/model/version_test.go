package model

import (
	"fmt"
	"testing"
)

func TestParseVersionFromRegularStr(t *testing.T) {
	fmt.Println(ParseVersionFromRegularStr("v4.4.444444444444444444"))
}
