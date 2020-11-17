package model

type Calculate struct {
	bench []struct {
	}
	IssueAffectedVersions []struct {
		IssueDatabaseId int
		TagId           int
	}
	IssueFixedVersions []struct {
		IssueDatabaseId int
		TagId           int
	}
}
