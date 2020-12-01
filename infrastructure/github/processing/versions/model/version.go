package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type VersionCode int

const (
	Regular    VersionCode = iota // (v?\d{1,8}\.\d{1,8}\.\d{1,8})$
	Unreleased                    // unreleased
	Master                        // master
)

type Version struct {
	Major int
	Minor int
	Patch int
	Code  VersionCode
}

// versionStr => RegularVersion
//if match ^(v?\d{1,8}\.\d{1,8}\.\d{1,8})$ :
//	return Version{...,Code: regular}
//else:
//	return err
func ParseVersionFromRegularStr(versionStr string) (Version, error) {
	// Check if versionStr is a regular version string
	reg := regexp.MustCompile(`^(v?\d{1,8}\.\d{1,8}\.\d{1,8})$`)
	if !reg.MatchString(versionStr) {
		return Version{}, fmt.Errorf("versionStr do not match the version regexp( " +
			`^(v?\d{1,8}\.\d{1,8}\.\d{1,8})$` + " )")
	}

	// Fill info into version struct
	if versionStr[0] == 'v' {
		versionStr = versionStr[1:]
	}
	indexes := strings.Split(versionStr, ".")
	major, _ := strconv.Atoi(indexes[0])
	minor, _ := strconv.Atoi(indexes[1])
	patch, _ := strconv.Atoi(indexes[2])
	version := Version{
		Major: major,
		Minor: minor,
		Patch: patch,
		Code:  Regular,
	}
	return version, nil
}

// versionStr => RegularVersion
//if match ^(v\d{1,8}\.\d{1,8}\.\d{1,8})$ :
//	return Version{...,Code: regular}
//else:
//	return err
func ParseVersionFromRegularStrMustHaveV(versionStr string) (Version, error) {
	// Check if versionStr is a regular version string
	reg := regexp.MustCompile(`^(v\d{1,8}\.\d{1,8}\.\d{1,8})$`)
	if !reg.MatchString(versionStr) {
		return Version{}, fmt.Errorf("versionStr do not match the version regexp( " +
			`^(v\d{1,8}\.\d{1,8}\.\d{1,8})$` + " )")
	}

	// Fill info into version struct
	if versionStr[0] == 'v' {
		versionStr = versionStr[1:]
	}
	indexes := strings.Split(versionStr, ".")
	major, _ := strconv.Atoi(indexes[0])
	minor, _ := strconv.Atoi(indexes[1])
	patch, _ := strconv.Atoi(indexes[2])
	version := Version{
		Major: major,
		Minor: minor,
		Patch: patch,
		Code:  Regular,
	}
	return version, nil
}
