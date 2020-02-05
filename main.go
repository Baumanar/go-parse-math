// example1.go
package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"./parser"
	"./utils"
	)

func main() {
	// Setup the input
	line := "x * x+2*x"
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

	fmt.Printf("The answer is: %d\n", utils.Calc(line)(2))

}