package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"go-parse-math/parser"
	"go-parse-math/utils"
)

func main() {
	// Setup the input
	line := "ln(e^10)"
	is := antlr.NewInputStream(line)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}

	fmt.Printf("The answer is: %f\n", utils.Calc(line)(9))

}