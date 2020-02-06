// example1.go
package main

import (
	"./parser"
	"./utils"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	// Setup the input
	line := "|-3|"
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

	fmt.Printf("The answer is: %f\n", utils.Calc(line)(0.5))

}