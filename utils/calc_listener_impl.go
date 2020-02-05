package utils

import (
	"../parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type calcListener struct {
	*parser.BaseCalcListener

	stack []func(int)int
}

func (l *calcListener) push(i func(int)int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() func(int)int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Pop the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

// ExitMulDiv is called when exiting the MulDiv production.
func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType()  {
	case parser.CalcParserMUL:
		l.push(func(x int)int{
			return left(x) * right(x)
		})
	case parser.CalcParserDIV:
		l.push(func(x int)int{
			return left(x) / right(x)
		})
	default:
		panic(fmt.Sprintf("unexpected operation: %s", c.GetOp().GetText()))
	}
}

// ExitAddSub is called when exiting the AddSub production.
func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(func(x int)int{
			return left(x) + right(x)
		})
	case parser.CalcParserSUB:
		l.push(func(x int)int{
			return left(x) - right(x)
		})
	default:
		panic(fmt.Sprintf("unexpected operation: %s", c.GetOp().GetText()))
	}
}

// ExitNumber is called when exiting the Number production.
func (l *calcListener) ExitNumber(c *parser.NumberContext) {

	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(
		func(x int) int {
			return i
	})
}
func (l *calcListener) ExitVariable(c *parser.VariableContext) {

	l.push(
		func(x int) int {
			fmt.Println("here", x)
			return x
		})
}
// calc takes a string expression and returns the evaluated result.
func Calc(input string) func(int)int {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.pop()
}

