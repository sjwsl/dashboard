package parser

import (
	"fmt"
	"log"
	"testing"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/lexer"

	lex "github.com/timtadh/lexmachine"
)

func TestParser(t *testing.T) {
	s, err := lexer.Lexer.Scanner([]byte(`
#### 5. Affected versions
[:v4.3.10] 3.5.6
`))
	if err != nil {
		log.Fatal(err)
	}
	var tokens []lex.Token
	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if err != nil {
			log.Fatal(err)
		}
		token := tok.(*lex.Token)
		tokens = append(tokens, *token)
		fmt.Printf("%-7v | %-10v | %v:%v-%v:%v\n",
			lexer.Tokens[token.Type],
			string(token.Lexeme),
			token.StartLine,
			token.StartColumn,
			token.EndLine,
			token.EndColumn)
	}
	fmt.Println(Parse(tokens))
}
