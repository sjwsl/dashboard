package parser

import (
	"fmt"

	"github.com/timtadh/lexmachine"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/lexer"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/model"
)

// Tokens => ( AffectedVersion | FixedVersion ) (IndexVersion | VERSION | COMMENT | MASTER | UNRELEASED)*
func Parse(tokens []lexmachine.Token) ([]model.Version, error) {
	var versions []model.Version
	if len(tokens) < 2 {
		return nil, nil
	}
	if tokens[0].Type != lexer.AffectedVersion &&
		tokens[0].Type != lexer.FixedVersion {
		return nil, fmt.Errorf("parse err: duplicate or position err of Affected Version title")
	}
	pos := 1
	for {
		switch tokens[pos].Type {
		case lexer.COMMENT:
			pos++
		case lexer.DOT:
			pos++
		case lexer.LBRACK:
			versionsIn, length, err := parseIndexVersion(tokens[pos:])
			if err != nil {
				return nil, err
			}
			versions = append(versions, versionsIn...)
			pos += length
		case lexer.COLON, lexer.RBRACK:
			return nil, fmt.Errorf("parse err: have %s in wrong position ", string(tokens[pos].Lexeme))
		case lexer.VERSION:
			version, err := model.ParseVersionFromRegularStr(string(tokens[pos].Lexeme))
			if err != nil {
				return nil, err
			}
			versions = append(versions, version)
			pos++
		case lexer.MASTER:
			versions = append(versions, model.Version{
				Code: model.Master,
			})
			pos++
		case lexer.UNRELEASED:
			versions = append(versions, model.Version{
				Code: model.Unreleased,
			})
			pos++
		default:
			panic(fmt.Errorf("I do not know what happen but it must be wrong ,"+
				" get %v when switch tokens[pos].Type", tokens[pos].Type))
		}

		if pos >= len(tokens) {
			return versions, nil
		}
	}
}

// IndexVersion => ( [VERSION : VERSION] | [ : VERSION] | [VERSION])
func parseIndexVersion(tokens []lexmachine.Token) ([]model.Version, int, error) {
	//[ v : v ]
	//0 1 2 3 4 5
	if FindFirstToken(tokens, lexer.RBRACK) == 4 &&
		tokens[3].Type == lexer.VERSION &&
		tokens[2].Type == lexer.COLON &&
		tokens[1].Type == lexer.VERSION {
		firstVersion, err := model.ParseVersionFromRegularStr(string(tokens[1].Lexeme))
		if err != nil {
			return nil, 0, err
		}
		lastVersion, err := model.ParseVersionFromRegularStr(string(tokens[3].Lexeme))
		if err != nil {
			return nil, 0, err
		}
		if firstVersion.Major == lastVersion.Major &&
			firstVersion.Minor == lastVersion.Minor &&
			firstVersion.Patch <= lastVersion.Patch {
			versions := make([]model.Version, lastVersion.Patch-firstVersion.Patch+1)
			for i, _ := range versions {
				versions[i] = model.Version{
					Major: firstVersion.Major,
					Minor: firstVersion.Minor,
					Patch: firstVersion.Patch + i,
				}
			}
			return versions, 5, nil
		} else {
			return nil, 0, fmt.Errorf("err version slice, first 2 index in version must be equal, firstVersion: %v, lastVersion: %v", firstVersion, lastVersion)
		}
	}

	//[ : v ]
	//0 1 2 3 4
	if FindFirstToken(tokens, lexer.RBRACK) == 3 &&
		tokens[2].Type == lexer.VERSION &&
		tokens[1].Type == lexer.COLON {
		lastVersion, err := model.ParseVersionFromRegularStr(string(tokens[2].Lexeme))
		if err != nil {
			return nil, 0, err
		}
		versions := make([]model.Version, lastVersion.Patch+1)
		for i, _ := range versions {
			versions[i] = model.Version{
				Major: lastVersion.Major,
				Minor: lastVersion.Minor,
				Patch: i,
			}
		}
		return versions, 4, nil
	}

	//[ id ]
	//0 1  2 3
	if FindFirstToken(tokens, lexer.RBRACK) == 2 &&
		tokens[1].Type == lexer.VERSION {
		version, err := model.ParseVersionFromRegularStr(string(tokens[1].Lexeme))
		if err != nil {
			return nil, 0, err
		}
		return []model.Version{version}, 3, nil
	}
	return nil, 0, fmt.Errorf("parse err: err tokens, %v", tokens)
}

func FindFirstToken(tokens []lexmachine.Token, tokenType int) int {
	for i, token := range tokens {
		if token.Type == tokenType {
			return i
		}
	}
	return -1
}
