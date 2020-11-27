package lexer

import (
	"fmt"
	"strings"

	lex "github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

var Literals []string       // The tokens representing literal strings
var Keywords []string       // The keyword tokens
var Tokens []string         // All of the tokens (including literals and keywords)
var TokenIds map[string]int // A map from the token names to their int ids
var Lexer *lex.Lexer        // The lexer object. Use this to construct a Scanner

// Called at package initialization. Creates the lexer and populates token lists.
func init() {
	initTokens()
	var err error
	Lexer, err = initLexer()
	if err != nil {
		panic(err)
	}
}

// The list of tokens.
const (
	COMMENT         int = iota
	AffectedVersion     // #### 5. Affected versions
	FixedVersion        //#### 6. Fixed versions
	VERSION             // v?(\d+\.\d+\.\d+)
	UNRELEASED
	MASTER
	LBRACK // [
	RBRACK // ]
	COLON  // :
)

func initTokens() {
	Tokens = []string{
		"COMMENT",
		"AffectedVersion",
		"FixedVersion",
		"VERSION",
		"UNRELEASED",
		"MASTER",
	}
	Literals = []string{
		"[",
		"]",
		":",
	}
	Tokens = append(Tokens, Literals...)
	TokenIds = map[string]int{
		"COMMENT":         COMMENT,
		"AffectedVersion": AffectedVersion,
		"FixedVersion":    FixedVersion,
		"VERSION":         VERSION,
		"UNRELEASED":      UNRELEASED,
		"MASTER":          MASTER,
		"[":               LBRACK,
		"]":               RBRACK,
		":":               COLON,
	}

	Keywords = []string{
		"#### 5. Affected versions",
		"#### 6. Fixed versions",
	}
}

// a lex.Action function which skips the match.
func skip(*lex.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

// a lex.Action function with constructs a Token of the given token type by
// the token type's name.
func token(name string) lex.Action {
	return func(s *lex.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(TokenIds[name], string(m.Bytes), m), nil
	}
}

// Creates the lexer object and compiles the NFA.
func initLexer() (*lex.Lexer, error) {
	lexer := lex.NewLexer()

	for _, lit := range Literals {
		r := "\\" + strings.Join(strings.Split(lit, ""), "\\")
		lexer.Add([]byte(r), token(lit))
	}
	lexer.Add([]byte(Keywords[0]), token(`AffectedVersion`))
	lexer.Add([]byte(Keywords[1]), token(`FixedVersion`))
	lexer.Add([]byte(`unreleased`), token(`UNRELEASED`))
	lexer.Add([]byte(`master`), token(`MASTER`))

	// ignore the comment
	lexer.Add([]byte(`<!--`),
		func(scan *lex.Scanner, match *machines.Match) (interface{}, error) {
			for tc := scan.TC; tc < len(scan.Text); tc++ {
				if scan.Text[tc] == '-' && tc+1 < len(scan.Text) {
					if scan.Text[tc+1] == '-' && tc+2 < len(scan.Text) {
						if scan.Text[tc+2] == '>' {
							scan.TC = tc + 3
							return nil, nil
						}
					}
				}
			}
			return nil, fmt.Errorf("unclosed comment starting at %d, (%d, %d)",
				match.TC, match.StartLine, match.StartColumn)
		})
	lexer.Add([]byte(`v?\d+\.\d+\.\d+`), token("VERSION"))
	lexer.Add([]byte("(\\s)+"), skip)

	err := lexer.Compile()
	if err != nil {
		return nil, err
	}
	return lexer, nil
}
