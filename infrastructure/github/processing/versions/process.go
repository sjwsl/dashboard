package versions

import (
	"fmt"
	"strings"

	lex "github.com/timtadh/lexmachine"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/lexer"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/model"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/parser"
)

const affectedVersionTemplate string = "#### 5. Affected versions"

func ScanTokensFromBody(body string) ([]lex.Token, error) {
	var strSlice string
	if idx := strings.Index(body, affectedVersionTemplate); idx != -1 {
		strSlice = body[idx:]

		scanner, err := lexer.Lexer.Scanner([]byte(strSlice))
		if err != nil {
			return nil, err
		}

		var tokens []lex.Token
		for tok, err, eof := scanner.Next(); !eof; tok, err, eof = scanner.Next() {
			if err != nil {
				return nil, err
			}
			token := tok.(*lex.Token)
			tokens = append(tokens, *token)
		}
		return tokens, nil
	}
	return nil, fmt.Errorf("no affected version title in body")
}

func ParseVersionFromTokens(tokens []lex.Token) ([]model.Version, []model.Version, error) {
	var affectedVersionIdx, fixedVersionIdx int

	if affectedVersionIdx = parser.FindFirstToken(tokens, lexer.AffectedVersion); affectedVersionIdx != 0 {
		return nil, nil, fmt.Errorf("parse err invalid token array")
	}
	if fixedVersionIdx = parser.FindFirstToken(tokens, lexer.FixedVersion); fixedVersionIdx == -1 {
		return nil, nil, fmt.Errorf("parse err invalid token array")
	}

	affectedVersions, err := parser.Parse(tokens[affectedVersionIdx:fixedVersionIdx])
	if err != nil {
		return nil, nil, err
	}
	fixedVersions, err := parser.Parse(tokens[fixedVersionIdx:])
	if err != nil {
		return nil, nil, err
	}
	return affectedVersions, fixedVersions, nil
}
