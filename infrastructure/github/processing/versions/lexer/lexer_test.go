package lexer

import (
	"fmt"
	"log"
	"testing"

	lex "github.com/timtadh/lexmachine"
)

func TestLexer(t *testing.T) {
	s, err := Lexer.Scanner([]byte(`
#### 5. Affected versions
<!-- a -->
<!-- a	 --sd
	csd- -asdsad 	
sd-->
unreleased master

[v4.3.1:v3.4.1] 3.5.6666666666666666666666666666666
#### 6. Fixed versions
`))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Type    | Lexeme     | Position")
	fmt.Println("--------+------------+------------")
	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if err != nil {
			log.Fatal(err)
		}
		token := tok.(*lex.Token)
		fmt.Printf("%-7v | %-10v | %v:%v-%v:%v\n",
			Tokens[token.Type],
			string(token.Lexeme),
			token.StartLine,
			token.StartColumn,
			token.EndLine,
			token.EndColumn)
	}
}
