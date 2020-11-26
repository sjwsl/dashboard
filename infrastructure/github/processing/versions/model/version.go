package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Version struct {
	Main int
	Sub  int
	Fix  int
}

func ParseVersionFromStr(versionStr string) (Version, error) {
	reg := regexp.MustCompile(`^(v?\d+\.\d+\.\d+)$`)
	if !reg.MatchString(versionStr) {
		return Version{}, fmt.Errorf("versionStr do not match the version regexp( we except token offered is valid ," +
			"so there is a fatal err in code )")
	}

	if versionStr[0] == 'v' {
		versionStr = versionStr[1:]
	}
	indexes := strings.Split(versionStr, ".")
	main, _ := strconv.Atoi(indexes[0])
	sub, _ := strconv.Atoi(indexes[1])
	fix, _ := strconv.Atoi(indexes[2])
	version := Version{
		Main: main,
		Sub:  sub,
		Fix:  fix,
	}
	return version, nil
}
