package util

import (
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	"io/ioutil"
	"os"
	. "reflect"
	"strings"
)

// ReadFile Readfile with path return bytes
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

// IdCompletenessProof check the unique id with totalCount.
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

// NameCompletenessProof check the unique names with totalCount.
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

// NotEmptyStrInQuery check string fields in complex struct named  "name" || "title" || "url" || "login" are not empty.
func NotEmptyStrInQuery(v Value, fieldName string) bool {
	name := strings.ToLower(fieldName)
	switch v.Kind() {
	case Bool,
		Int, Int8, Int16, Int32, Int64,
		Uint, Uint8, Uint16, Uint32, Uint64, Uintptr,
		Float32, Float64,
		Complex64, Complex128,
		Chan, Func, Interface, Map, UnsafePointer:
		return true
	case Array:
		for i := 0; i < v.Len(); i++ {
			if !NotEmptyStrInQuery(v.Index(i), "") {
				return false
			}
		}
		return true
	case Slice:
		if v.IsNil() {
			return true
		} else {
			for i := 0; i < v.Len(); i++ {
				if !NotEmptyStrInQuery(v.Index(i), "") {
					return false
				}
			}
		}
	case Struct:
		for i := 0; i < v.NumField(); i++ {
			if !NotEmptyStrInQuery(v.Field(i), v.Type().Field(i).Name) {
				fmt.Printf("err Field : %v\n", v.Type().Field(i).Name)
				return false
			}
		}
		return true
	case Ptr:
		if v.Elem().Kind() == String || v.IsNil() {
			return true
		} else {
			return NotEmptyStrInQuery(v.Elem(), fieldName)
		}
	case String:
		if v.Len() == 0 &&
			(name == "name" || name == "title" || name == "url" || name == "login") {
			fmt.Printf("err empty string : %v, %v\n", v, fieldName)
			return false
		}
		return true
	default:
		// This should never happens, but will act as a safeguard for
		// later, as a default value doesn't makes sense here.
		panic(&ValueError{Method: "utile.NotEmptyStrInQuery catch error kind", Kind: v.Kind()})
	}
	return true
}

// QueryCompletenessSpec check completeness of issue numbers & tag names.
func QueryCompletenessSpec(totalData *model.Query) {
	nums := make([]int, len(totalData.Repository.Issues.Nodes))
	for i, _ := range nums {
		nums[i] = totalData.Repository.Issues.Nodes[i].Number
	}
	err := IdCompletenessProof(totalData.Repository.Issues.TotalCount, nums)
	if err != nil {
		panic(err)
	}

	nums = make([]int, len(totalData.Repository.AssignableUsers.Nodes))
	for i, _ := range nums {
		nums[i] = *totalData.Repository.AssignableUsers.Nodes[i].DatabaseID
	}
	err = IdCompletenessProof(totalData.Repository.AssignableUsers.TotalCount, nums)
	if err != nil {
		panic(err)
	}

	names := make([]string, len(totalData.Repository.Labels.Nodes))
	for i, _ := range nums {
		names[i] = totalData.Repository.Labels.Nodes[i].Name
	}
	err = NameCompletenessProof(totalData.Repository.Labels.TotalCount, names)
	if err != nil {
		panic(err)
	}

	names = make([]string, len(totalData.Repository.Refs.Nodes))
	for i, _ := range names {
		names[i] = totalData.Repository.Refs.Nodes[i].Name
	}
	err = NameCompletenessProof(totalData.Repository.Refs.TotalCount, names)
	if err != nil {
		panic(err)
	}
}

// QueryDataInvalidSpec check if data is invalid, because of no name or other important fields.
func QueryDataInvalidSpec(totalData *model.Query) {
	if !NotEmptyStrInQuery(ValueOf(totalData), "") {
		panic("invalid data leak")
	}
}
