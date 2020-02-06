package utils

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"go-parse-math/parser"
	"math"
	"strconv"
)

type calcListener struct {
	*parser.BaseCalcListener

	stack []func(float64)float64
}

func (l *calcListener) push(i func(float64)float64) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() func(float64)float64 {
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
		l.push(func(x float64)float64{
			return left(x) * right(x)
		})
	case parser.CalcParserDIV:
		l.push(func(x float64)float64{
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
		l.push(func(x float64)float64{
			return left(x) + right(x)
		})
	case parser.CalcParserSUB:
		l.push(func(x float64)float64{
			return left(x) - right(x)
		})
	default:
		panic(fmt.Sprintf("unexpected operation: %s", c.GetOp().GetText()))
	}
}

// ExitMulDiv is called when exiting the MulDiv production.
func (l *calcListener) ExitMult(c *parser.MultContext) {
	right, left := l.pop(), l.pop()

	l.push(func(x float64)float64{
		return left(x) * right(x)
	})
}

// ExitNumber is called when exiting the Number production.
func (l *calcListener) ExitNumber(c *parser.NumberContext) {

	i, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(err.Error())
	}

	l.push(
		func(x float64) float64 {
			return i
	})
}
func (l *calcListener) ExitVariable(c *parser.VariableContext) {

	l.push(
		func(x float64) float64 {
			fmt.Println("Exit variable", x)
			return x
		})
}

func (l *calcListener) ExitSqrt(c *parser.SqrtContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Sqrt(expr(x))
		})
}

func (l *calcListener) ExitLog(c *parser.LogContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Log10(expr(x))
		})
}


func (l *calcListener) ExitLn(c *parser.LnContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Log(expr(x))
		})
}


func (l *calcListener) ExitSin(c *parser.SinContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Sin(expr(x))
		})
}


func (l *calcListener) ExitSinh(c *parser.SinhContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Sinh(expr(x))
		})
}

func (l *calcListener) ExitCos(c *parser.CosContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Cos(expr(x))
		})
}

func (l *calcListener) ExitCosh(c *parser.CoshContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Cosh(expr(x))
		})
}
func (l *calcListener) ExitTan(c *parser.TanContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Tan(expr(x))
		})
}

func (l *calcListener) ExitTanh(c *parser.TanhContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Tanh(expr(x))
		})
}


func (l *calcListener) ExitArcsin(c *parser.ArcsinContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Asin(expr(x))
		})
}

func (l *calcListener) ExitArccos(c *parser.ArccosContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Acos(expr(x))
		})
}

func (l *calcListener) ExitArctan(c *parser.ArctanContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Atan(expr(x))
		})
}

func (l *calcListener) ExitAbs(c *parser.AbsContext) {
	expr := l.pop()
	l.push(
		func(x float64) float64 {
			return math.Abs(expr(x))
		})
}

func (l *calcListener) ExitPi(c *parser.PiContext) {

	l.push(
		func(x float64) float64 {
			return math.Pi
		})
}

func (l *calcListener) ExitEuler(c *parser.EulerContext) {

	l.push(
		func(x float64) float64 {
			return math.E
		})
}

func (l *calcListener) ExitPow(ctx *parser.PowContext) {
	pow := l.pop()
	expr := l.pop()
	l.push(func(x float64)float64{
		return math.Pow(expr(x), pow(x))
	})

}



// calc takes a string expression and returns the evaluated result.
func Calc(input string) func(float64)float64 {
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

