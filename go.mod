module dashboard

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/martian v2.1.0+incompatible
	github.com/machinebox/graphql v0.2.2
	github.com/matryer/is v1.4.0 // indirect
	github.com/onsi/gomega v1.10.3
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/math v0.0.0-20141027224758-f2ed9e40e245
	github.com/shurcooL/githubv4 v0.0.0-20200928013246-d292edc3691b
	github.com/shurcooL/graphql v0.0.0-20200928012149-18c5c3165e3a // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

//replace github.com/PingCAP-QE/libs => ../libs
