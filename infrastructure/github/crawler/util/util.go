package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

//func VersionToInt(version string) int {
//	versionDefinition := regexp.MustCompile(`\d+\.\d+\.\d+|master|unplanned|unplaned`)
//	version = versionDefinition.FindString(version)
//	intDefinition := regexp.MustCompile(`\d+`)
//	versionIndexs := intDefinition.FindStringSubmatch(version)
//	if len(versionIndexs) != 3 {
//		log.Errorf("Error version type")
//	}
//	versionIndex := make([]int,3)
//	for i := 0; i < 3; i++ {
//		versionIndex[i] := int()
//	}
//}

func IdCompletenessProof(totalCount int, ids []int) error {
	if len(ids) != totalCount {
		return fmt.Errorf("fatal : lack ids total count %d, ids length %d. ", totalCount, len(ids))
	}

	set := make(map[interface{}]bool)
	for _, id := range ids {
		if set[id] {
			return fmt.Errorf("fatal : duplicate id %d. ", id)
		} else {
			set[id] = true
		}
	}
	return nil
}

func NameCompletenessProof(totalCount int, names []string) error {
	if len(names) != totalCount {
		return fmt.Errorf("fatal : lack names total count %d, names length %d. ", totalCount, len(names))
	}

	set := make(map[string]bool)
	for _, name := range names {
		if set[name] {
			return fmt.Errorf("fatal : duplicate name %s. ", name)
		} else {
			set[name] = true
		}
	}
	return nil
}
