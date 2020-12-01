package model

import (
	"fmt"
	"testing"
)

func TestParseVersionFromRegularStr(t *testing.T) {
	fmt.Println(ParseVersionFromRegularStrMustHaveV("v4.4.444444"))
}
