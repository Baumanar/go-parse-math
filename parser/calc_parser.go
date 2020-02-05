// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 55, 139,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 5, 2, 11, 10, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	97, 10, 3, 3, 3, 3, 3, 5, 3, 101, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	7, 3, 132, 10, 3, 12, 3, 14, 3, 135, 11, 3, 3, 4, 3, 4, 3, 4, 2, 3, 4,
	5, 2, 4, 6, 2, 7, 3, 2, 7, 8, 4, 2, 9, 9, 17, 17, 3, 2, 10, 11, 3, 2, 19,
	20, 3, 2, 21, 22, 2, 179, 2, 8, 3, 2, 2, 2, 4, 100, 3, 2, 2, 2, 6, 136,
	3, 2, 2, 2, 8, 10, 5, 4, 3, 2, 9, 11, 7, 3, 2, 2, 10, 9, 3, 2, 2, 2, 10,
	11, 3, 2, 2, 2, 11, 12, 3, 2, 2, 2, 12, 13, 5, 6, 4, 2, 13, 3, 3, 2, 2,
	2, 14, 15, 8, 3, 1, 2, 15, 16, 7, 22, 2, 2, 16, 101, 5, 4, 3, 45, 17, 18,
	7, 21, 2, 2, 18, 101, 5, 4, 3, 44, 19, 20, 7, 29, 2, 2, 20, 101, 5, 4,
	3, 43, 21, 22, 7, 30, 2, 2, 22, 101, 5, 4, 3, 42, 23, 24, 7, 31, 2, 2,
	24, 101, 5, 4, 3, 41, 25, 26, 7, 32, 2, 2, 26, 27, 7, 4, 2, 2, 27, 28,
	5, 4, 3, 2, 28, 29, 7, 5, 2, 2, 29, 30, 5, 4, 3, 2, 30, 31, 7, 6, 2, 2,
	31, 101, 3, 2, 2, 2, 32, 33, 7, 33, 2, 2, 33, 101, 5, 4, 3, 39, 34, 35,
	7, 34, 2, 2, 35, 101, 5, 4, 3, 38, 36, 37, 7, 35, 2, 2, 37, 101, 5, 4,
	3, 37, 38, 39, 7, 36, 2, 2, 39, 101, 5, 4, 3, 36, 40, 41, 7, 37, 2, 2,
	41, 101, 5, 4, 3, 35, 42, 43, 7, 38, 2, 2, 43, 101, 5, 4, 3, 34, 44, 45,
	7, 39, 2, 2, 45, 101, 5, 4, 3, 33, 46, 47, 7, 40, 2, 2, 47, 101, 5, 4,
	3, 32, 48, 49, 7, 41, 2, 2, 49, 101, 5, 4, 3, 31, 50, 51, 7, 42, 2, 2,
	51, 101, 5, 4, 3, 30, 52, 53, 7, 43, 2, 2, 53, 101, 5, 4, 3, 29, 54, 55,
	7, 44, 2, 2, 55, 101, 5, 4, 3, 28, 56, 57, 7, 45, 2, 2, 57, 58, 7, 4, 2,
	2, 58, 59, 5, 4, 3, 2, 59, 60, 7, 5, 2, 2, 60, 61, 5, 4, 3, 2, 61, 62,
	7, 6, 2, 2, 62, 101, 3, 2, 2, 2, 63, 64, 7, 46, 2, 2, 64, 101, 5, 4, 3,
	26, 65, 66, 7, 47, 2, 2, 66, 101, 5, 4, 3, 25, 67, 68, 7, 48, 2, 2, 68,
	101, 5, 4, 3, 24, 69, 70, 7, 49, 2, 2, 70, 101, 5, 4, 3, 23, 71, 72, 7,
	50, 2, 2, 72, 101, 5, 4, 3, 22, 73, 74, 7, 51, 2, 2, 74, 101, 5, 4, 3,
	21, 75, 76, 7, 52, 2, 2, 76, 101, 5, 4, 3, 20, 77, 78, 7, 27, 2, 2, 78,
	79, 7, 4, 2, 2, 79, 80, 5, 4, 3, 2, 80, 81, 7, 6, 2, 2, 81, 101, 3, 2,
	2, 2, 82, 83, 7, 28, 2, 2, 83, 101, 5, 4, 3, 18, 84, 85, 7, 4, 2, 2, 85,
	86, 5, 4, 3, 2, 86, 87, 7, 6, 2, 2, 87, 101, 3, 2, 2, 2, 88, 89, 7, 4,
	2, 2, 89, 90, 5, 4, 3, 2, 90, 91, 7, 6, 2, 2, 91, 92, 5, 4, 3, 8, 92, 101,
	3, 2, 2, 2, 93, 101, 7, 13, 2, 2, 94, 96, 7, 23, 2, 2, 95, 97, 7, 12, 2,
	2, 96, 95, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 101, 3, 2, 2, 2, 98, 101,
	7, 26, 2, 2, 99, 101, 7, 14, 2, 2, 100, 14, 3, 2, 2, 2, 100, 17, 3, 2,
	2, 2, 100, 19, 3, 2, 2, 2, 100, 21, 3, 2, 2, 2, 100, 23, 3, 2, 2, 2, 100,
	25, 3, 2, 2, 2, 100, 32, 3, 2, 2, 2, 100, 34, 3, 2, 2, 2, 100, 36, 3, 2,
	2, 2, 100, 38, 3, 2, 2, 2, 100, 40, 3, 2, 2, 2, 100, 42, 3, 2, 2, 2, 100,
	44, 3, 2, 2, 2, 100, 46, 3, 2, 2, 2, 100, 48, 3, 2, 2, 2, 100, 50, 3, 2,
	2, 2, 100, 52, 3, 2, 2, 2, 100, 54, 3, 2, 2, 2, 100, 56, 3, 2, 2, 2, 100,
	63, 3, 2, 2, 2, 100, 65, 3, 2, 2, 2, 100, 67, 3, 2, 2, 2, 100, 69, 3, 2,
	2, 2, 100, 71, 3, 2, 2, 2, 100, 73, 3, 2, 2, 2, 100, 75, 3, 2, 2, 2, 100,
	77, 3, 2, 2, 2, 100, 82, 3, 2, 2, 2, 100, 84, 3, 2, 2, 2, 100, 88, 3, 2,
	2, 2, 100, 93, 3, 2, 2, 2, 100, 94, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100,
	99, 3, 2, 2, 2, 101, 133, 3, 2, 2, 2, 102, 103, 12, 17, 2, 2, 103, 104,
	7, 24, 2, 2, 104, 132, 5, 4, 3, 18, 105, 106, 12, 16, 2, 2, 106, 107, 7,
	25, 2, 2, 107, 132, 5, 4, 3, 17, 108, 109, 12, 15, 2, 2, 109, 110, 9, 2,
	2, 2, 110, 132, 5, 4, 3, 16, 111, 112, 12, 14, 2, 2, 112, 113, 9, 3, 2,
	2, 113, 132, 5, 4, 3, 15, 114, 115, 12, 13, 2, 2, 115, 116, 7, 18, 2, 2,
	116, 132, 5, 4, 3, 14, 117, 118, 12, 12, 2, 2, 118, 119, 9, 4, 2, 2, 119,
	132, 5, 4, 3, 13, 120, 121, 12, 11, 2, 2, 121, 122, 9, 5, 2, 2, 122, 132,
	5, 4, 3, 12, 123, 124, 12, 7, 2, 2, 124, 125, 9, 6, 2, 2, 125, 132, 5,
	4, 3, 8, 126, 127, 12, 9, 2, 2, 127, 128, 7, 4, 2, 2, 128, 129, 5, 4, 3,
	2, 129, 130, 7, 6, 2, 2, 130, 132, 3, 2, 2, 2, 131, 102, 3, 2, 2, 2, 131,
	105, 3, 2, 2, 2, 131, 108, 3, 2, 2, 2, 131, 111, 3, 2, 2, 2, 131, 114,
	3, 2, 2, 2, 131, 117, 3, 2, 2, 2, 131, 120, 3, 2, 2, 2, 131, 123, 3, 2,
	2, 2, 131, 126, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2,
	133, 134, 3, 2, 2, 2, 134, 5, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 136, 137,
	7, 2, 2, 3, 137, 7, 3, 2, 2, 2, 7, 10, 96, 100, 131, 133,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'('", "';'", "')'", "'^'", "'**'", "'%'", "'~'", "'//'", "'()'",
	"", "'x'", "", "", "", "", "'*'", "'/'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "NUMBER", "VAR", "FLOAT", "DIGIT",
	"MOD", "WHOLE", "MUL", "DIV", "ADD", "SUB", "PI", "EXPONENT", "NEGEXPONENT",
	"EULER", "SQRT", "SQR", "FLOOR", "CEIL", "ABS", "ROUNDK", "ROUND", "TRUNC",
	"SIN", "COS", "TAN", "COT", "SINH", "COSH", "TANH", "ARCSIN", "ARCCOS",
	"ARCTAN", "ARCTAN2", "ARCCOT", "EXP", "LN", "EEX", "LOG", "RAD", "DEG",
	"WS", "COM", "INVALID",
}

var ruleNames = []string{
	"start", "expression", "compileUnit",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CalcParser struct {
	*antlr.BaseParser
}

func NewCalcParser(input antlr.TokenStream) *CalcParser {
	this := new(CalcParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Calc.g4"

	return this
}

// CalcParser tokens.
const (
	CalcParserEOF         = antlr.TokenEOF
	CalcParserT__0        = 1
	CalcParserT__1        = 2
	CalcParserT__2        = 3
	CalcParserT__3        = 4
	CalcParserT__4        = 5
	CalcParserT__5        = 6
	CalcParserT__6        = 7
	CalcParserT__7        = 8
	CalcParserT__8        = 9
	CalcParserT__9        = 10
	CalcParserNUMBER      = 11
	CalcParserVAR         = 12
	CalcParserFLOAT       = 13
	CalcParserDIGIT       = 14
	CalcParserMOD         = 15
	CalcParserWHOLE       = 16
	CalcParserMUL         = 17
	CalcParserDIV         = 18
	CalcParserADD         = 19
	CalcParserSUB         = 20
	CalcParserPI          = 21
	CalcParserEXPONENT    = 22
	CalcParserNEGEXPONENT = 23
	CalcParserEULER       = 24
	CalcParserSQRT        = 25
	CalcParserSQR         = 26
	CalcParserFLOOR       = 27
	CalcParserCEIL        = 28
	CalcParserABS         = 29
	CalcParserROUNDK      = 30
	CalcParserROUND       = 31
	CalcParserTRUNC       = 32
	CalcParserSIN         = 33
	CalcParserCOS         = 34
	CalcParserTAN         = 35
	CalcParserCOT         = 36
	CalcParserSINH        = 37
	CalcParserCOSH        = 38
	CalcParserTANH        = 39
	CalcParserARCSIN      = 40
	CalcParserARCCOS      = 41
	CalcParserARCTAN      = 42
	CalcParserARCTAN2     = 43
	CalcParserARCCOT      = 44
	CalcParserEXP         = 45
	CalcParserLN          = 46
	CalcParserEEX         = 47
	CalcParserLOG         = 48
	CalcParserRAD         = 49
	CalcParserDEG         = 50
	CalcParserWS          = 51
	CalcParserCOM         = 52
	CalcParserINVALID     = 53
)

// CalcParser rules.
const (
	CalcParserRULE_start       = 0
	CalcParserRULE_expression  = 1
	CalcParserRULE_compileUnit = 2
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) CompileUnit() ICompileUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompileUnitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompileUnitContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *CalcParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.expression(0)
	}
	p.SetState(8)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CalcParserT__0 {
		{
			p.SetState(7)
			p.Match(CalcParserT__0)
		}

	}
	{
		p.SetState(10)
		p.CompileUnit()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TanContext struct {
	*ExpressionContext
}

func NewTanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TanContext {
	var p = new(TanContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TanContext) TAN() antlr.TerminalNode {
	return s.GetToken(CalcParserTAN, 0)
}

func (s *TanContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterTan(s)
	}
}

func (s *TanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitTan(s)
	}
}

type CoshContext struct {
	*ExpressionContext
}

func NewCoshContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoshContext {
	var p = new(CoshContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CoshContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoshContext) COSH() antlr.TerminalNode {
	return s.GetToken(CalcParserCOSH, 0)
}

func (s *CoshContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CoshContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterCosh(s)
	}
}

func (s *CoshContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitCosh(s)
	}
}

type SqRootContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewSqRootContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SqRootContext {
	var p = new(SqRootContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SqRootContext) GetOp() antlr.Token { return s.op }

func (s *SqRootContext) SetOp(v antlr.Token) { s.op = v }

func (s *SqRootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqRootContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SqRootContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SqRootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterSqRoot(s)
	}
}

func (s *SqRootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitSqRoot(s)
	}
}

type VariableContext struct {
	*ExpressionContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) VAR() antlr.TerminalNode {
	return s.GetToken(CalcParserVAR, 0)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitVariable(s)
	}
}

type NegExponentContext struct {
	*ExpressionContext
}

func NewNegExponentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegExponentContext {
	var p = new(NegExponentContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NegExponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegExponentContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *NegExponentContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NegExponentContext) NEGEXPONENT() antlr.TerminalNode {
	return s.GetToken(CalcParserNEGEXPONENT, 0)
}

func (s *NegExponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterNegExponent(s)
	}
}

func (s *NegExponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitNegExponent(s)
	}
}

type ExponentContext struct {
	*ExpressionContext
}

func NewExponentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentContext {
	var p = new(ExponentContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExponentContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExponentContext) EXPONENT() antlr.TerminalNode {
	return s.GetToken(CalcParserEXPONENT, 0)
}

func (s *ExponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterExponent(s)
	}
}

func (s *ExponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitExponent(s)
	}
}

type Arctan2Context struct {
	*ExpressionContext
}

func NewArctan2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Arctan2Context {
	var p = new(Arctan2Context)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *Arctan2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arctan2Context) ARCTAN2() antlr.TerminalNode {
	return s.GetToken(CalcParserARCTAN2, 0)
}

func (s *Arctan2Context) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Arctan2Context) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Arctan2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterArctan2(s)
	}
}

func (s *Arctan2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitArctan2(s)
	}
}

type MulDivContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalcParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(CalcParserDIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

type ArcsinContext struct {
	*ExpressionContext
}

func NewArcsinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArcsinContext {
	var p = new(ArcsinContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ArcsinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArcsinContext) ARCSIN() antlr.TerminalNode {
	return s.GetToken(CalcParserARCSIN, 0)
}

func (s *ArcsinContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArcsinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterArcsin(s)
	}
}

func (s *ArcsinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitArcsin(s)
	}
}

type UnaryPlusContext struct {
	*ExpressionContext
}

func NewUnaryPlusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryPlusContext {
	var p = new(UnaryPlusContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryPlusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryPlusContext) ADD() antlr.TerminalNode {
	return s.GetToken(CalcParserADD, 0)
}

func (s *UnaryPlusContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryPlusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterUnaryPlus(s)
	}
}

func (s *UnaryPlusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitUnaryPlus(s)
	}
}

type ArccotContext struct {
	*ExpressionContext
}

func NewArccotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArccotContext {
	var p = new(ArccotContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ArccotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArccotContext) ARCCOT() antlr.TerminalNode {
	return s.GetToken(CalcParserARCCOT, 0)
}

func (s *ArccotContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArccotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterArccot(s)
	}
}

func (s *ArccotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitArccot(s)
	}
}

type ArccosContext struct {
	*ExpressionContext
}

func NewArccosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArccosContext {
	var p = new(ArccosContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ArccosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArccosContext) ARCCOS() antlr.TerminalNode {
	return s.GetToken(CalcParserARCCOS, 0)
}

func (s *ArccosContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArccosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterArccos(s)
	}
}

func (s *ArccosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitArccos(s)
	}
}

type EulerContext struct {
	*ExpressionContext
}

func NewEulerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EulerContext {
	var p = new(EulerContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EulerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EulerContext) EULER() antlr.TerminalNode {
	return s.GetToken(CalcParserEULER, 0)
}

func (s *EulerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterEuler(s)
	}
}

func (s *EulerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitEuler(s)
	}
}

type ArctanContext struct {
	*ExpressionContext
}

func NewArctanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArctanContext {
	var p = new(ArctanContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ArctanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArctanContext) ARCTAN() antlr.TerminalNode {
	return s.GetToken(CalcParserARCTAN, 0)
}

func (s *ArctanContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArctanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterArctan(s)
	}
}

func (s *ArctanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitArctan(s)
	}
}

type ParenthesisContext struct {
	*ExpressionContext
}

func NewParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisContext {
	var p = new(ParenthesisContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterParenthesis(s)
	}
}

func (s *ParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitParenthesis(s)
	}
}

type AbsContext struct {
	*ExpressionContext
}

func NewAbsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AbsContext {
	var p = new(AbsContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AbsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsContext) ABS() antlr.TerminalNode {
	return s.GetToken(CalcParserABS, 0)
}

func (s *AbsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AbsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterAbs(s)
	}
}

func (s *AbsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitAbs(s)
	}
}

type NumberContext struct {
	*ExpressionContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitNumber(s)
	}
}

type SinhContext struct {
	*ExpressionContext
}

func NewSinhContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SinhContext {
	var p = new(SinhContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SinhContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinhContext) SINH() antlr.TerminalNode {
	return s.GetToken(CalcParserSINH, 0)
}

func (s *SinhContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SinhContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterSinh(s)
	}
}

func (s *SinhContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitSinh(s)
	}
}

type RoundContext struct {
	*ExpressionContext
}

func NewRoundContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoundContext {
	var p = new(RoundContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *RoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundContext) ROUND() antlr.TerminalNode {
	return s.GetToken(CalcParserROUND, 0)
}

func (s *RoundContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterRound(s)
	}
}

func (s *RoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitRound(s)
	}
}

type TruncContext struct {
	*ExpressionContext
}

func NewTruncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TruncContext {
	var p = new(TruncContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TruncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TruncContext) TRUNC() antlr.TerminalNode {
	return s.GetToken(CalcParserTRUNC, 0)
}

func (s *TruncContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TruncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterTrunc(s)
	}
}

func (s *TruncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitTrunc(s)
	}
}

type PiContext struct {
	*ExpressionContext
}

func NewPiContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PiContext {
	var p = new(PiContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PiContext) PI() antlr.TerminalNode {
	return s.GetToken(CalcParserPI, 0)
}

func (s *PiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPi(s)
	}
}

func (s *PiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPi(s)
	}
}

type TanhContext struct {
	*ExpressionContext
}

func NewTanhContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TanhContext {
	var p = new(TanhContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TanhContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TanhContext) TANH() antlr.TerminalNode {
	return s.GetToken(CalcParserTANH, 0)
}

func (s *TanhContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TanhContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterTanh(s)
	}
}

func (s *TanhContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitTanh(s)
	}
}

type FloorContext struct {
	*ExpressionContext
}

func NewFloorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloorContext {
	var p = new(FloorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FloorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloorContext) FLOOR() antlr.TerminalNode {
	return s.GetToken(CalcParserFLOOR, 0)
}

func (s *FloorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FloorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterFloor(s)
	}
}

func (s *FloorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitFloor(s)
	}
}

type LnContext struct {
	*ExpressionContext
}

func NewLnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LnContext {
	var p = new(LnContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LnContext) LN() antlr.TerminalNode {
	return s.GetToken(CalcParserLN, 0)
}

func (s *LnContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterLn(s)
	}
}

func (s *LnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitLn(s)
	}
}

type ModContext struct {
	*ExpressionContext
}

func NewModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModContext {
	var p = new(ModContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ModContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ModContext) MOD() antlr.TerminalNode {
	return s.GetToken(CalcParserMOD, 0)
}

func (s *ModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterMod(s)
	}
}

func (s *ModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitMod(s)
	}
}

type LogContext struct {
	*ExpressionContext
}

func NewLogContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogContext {
	var p = new(LogContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogContext) LOG() antlr.TerminalNode {
	return s.GetToken(CalcParserLOG, 0)
}

func (s *LogContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterLog(s)
	}
}

func (s *LogContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitLog(s)
	}
}

type AddSubContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(CalcParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(CalcParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitAddSub(s)
	}
}

type CosContext struct {
	*ExpressionContext
}

func NewCosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CosContext {
	var p = new(CosContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CosContext) COS() antlr.TerminalNode {
	return s.GetToken(CalcParserCOS, 0)
}

func (s *CosContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterCos(s)
	}
}

func (s *CosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitCos(s)
	}
}

type DegContext struct {
	*ExpressionContext
}

func NewDegContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DegContext {
	var p = new(DegContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DegContext) DEG() antlr.TerminalNode {
	return s.GetToken(CalcParserDEG, 0)
}

func (s *DegContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterDeg(s)
	}
}

func (s *DegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitDeg(s)
	}
}

type SqrtContext struct {
	*ExpressionContext
}

func NewSqrtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SqrtContext {
	var p = new(SqrtContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SqrtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqrtContext) SQRT() antlr.TerminalNode {
	return s.GetToken(CalcParserSQRT, 0)
}

func (s *SqrtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SqrtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterSqrt(s)
	}
}

func (s *SqrtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitSqrt(s)
	}
}

type CotContext struct {
	*ExpressionContext
}

func NewCotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CotContext {
	var p = new(CotContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CotContext) COT() antlr.TerminalNode {
	return s.GetToken(CalcParserCOT, 0)
}

func (s *CotContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterCot(s)
	}
}

func (s *CotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitCot(s)
	}
}

type WholeContext struct {
	*ExpressionContext
}

func NewWholeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WholeContext {
	var p = new(WholeContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *WholeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WholeContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *WholeContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WholeContext) WHOLE() antlr.TerminalNode {
	return s.GetToken(CalcParserWHOLE, 0)
}

func (s *WholeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterWhole(s)
	}
}

func (s *WholeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitWhole(s)
	}
}

type UnaryContext struct {
	*ExpressionContext
}

func NewUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryContext {
	var p = new(UnaryContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) SUB() antlr.TerminalNode {
	return s.GetToken(CalcParserSUB, 0)
}

func (s *UnaryContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitUnary(s)
	}
}

type RadContext struct {
	*ExpressionContext
}

func NewRadContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RadContext {
	var p = new(RadContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *RadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RadContext) RAD() antlr.TerminalNode {
	return s.GetToken(CalcParserRAD, 0)
}

func (s *RadContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterRad(s)
	}
}

func (s *RadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitRad(s)
	}
}

type MultContext struct {
	*ExpressionContext
}

func NewMultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultContext {
	var p = new(MultContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MultContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterMult(s)
	}
}

func (s *MultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitMult(s)
	}
}

type SqrContext struct {
	*ExpressionContext
}

func NewSqrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SqrContext {
	var p = new(SqrContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SqrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqrContext) SQR() antlr.TerminalNode {
	return s.GetToken(CalcParserSQR, 0)
}

func (s *SqrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SqrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterSqr(s)
	}
}

func (s *SqrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitSqr(s)
	}
}

type SinContext struct {
	*ExpressionContext
}

func NewSinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SinContext {
	var p = new(SinContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinContext) SIN() antlr.TerminalNode {
	return s.GetToken(CalcParserSIN, 0)
}

func (s *SinContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterSin(s)
	}
}

func (s *SinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitSin(s)
	}
}

type EexContext struct {
	*ExpressionContext
}

func NewEexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EexContext {
	var p = new(EexContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EexContext) EEX() antlr.TerminalNode {
	return s.GetToken(CalcParserEEX, 0)
}

func (s *EexContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterEex(s)
	}
}

func (s *EexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitEex(s)
	}
}

type PowContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowContext {
	var p = new(PowContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PowContext) GetOp() antlr.Token { return s.op }

func (s *PowContext) SetOp(v antlr.Token) { s.op = v }

func (s *PowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *PowContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPow(s)
	}
}

func (s *PowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPow(s)
	}
}

type CeilContext struct {
	*ExpressionContext
}

func NewCeilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CeilContext {
	var p = new(CeilContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CeilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeilContext) CEIL() antlr.TerminalNode {
	return s.GetToken(CalcParserCEIL, 0)
}

func (s *CeilContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CeilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterCeil(s)
	}
}

func (s *CeilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitCeil(s)
	}
}

type ExpContext struct {
	*ExpressionContext
}

func NewExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpContext {
	var p = new(ExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) EXP() antlr.TerminalNode {
	return s.GetToken(CalcParserEXP, 0)
}

func (s *ExpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitExp(s)
	}
}

type RoundkContext struct {
	*ExpressionContext
}

func NewRoundkContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoundkContext {
	var p = new(RoundkContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *RoundkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundkContext) ROUNDK() antlr.TerminalNode {
	return s.GetToken(CalcParserROUNDK, 0)
}

func (s *RoundkContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *RoundkContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RoundkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterRoundk(s)
	}
}

func (s *RoundkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitRoundk(s)
	}
}

func (p *CalcParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CalcParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CalcParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(13)
			p.Match(CalcParserSUB)
		}
		{
			p.SetState(14)
			p.expression(43)
		}

	case 2:
		localctx = NewUnaryPlusContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)
			p.Match(CalcParserADD)
		}
		{
			p.SetState(16)
			p.expression(42)
		}

	case 3:
		localctx = NewFloorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.Match(CalcParserFLOOR)
		}
		{
			p.SetState(18)
			p.expression(41)
		}

	case 4:
		localctx = NewCeilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)
			p.Match(CalcParserCEIL)
		}
		{
			p.SetState(20)
			p.expression(40)
		}

	case 5:
		localctx = NewAbsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(21)
			p.Match(CalcParserABS)
		}
		{
			p.SetState(22)
			p.expression(39)
		}

	case 6:
		localctx = NewRoundkContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.Match(CalcParserROUNDK)
		}
		{
			p.SetState(24)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(25)
			p.expression(0)
		}
		{
			p.SetState(26)
			p.Match(CalcParserT__2)
		}
		{
			p.SetState(27)
			p.expression(0)
		}
		{
			p.SetState(28)
			p.Match(CalcParserT__3)
		}

	case 7:
		localctx = NewRoundContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			p.Match(CalcParserROUND)
		}
		{
			p.SetState(31)
			p.expression(37)
		}

	case 8:
		localctx = NewTruncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)
			p.Match(CalcParserTRUNC)
		}
		{
			p.SetState(33)
			p.expression(36)
		}

	case 9:
		localctx = NewSinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			p.Match(CalcParserSIN)
		}
		{
			p.SetState(35)
			p.expression(35)
		}

	case 10:
		localctx = NewCosContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(36)
			p.Match(CalcParserCOS)
		}
		{
			p.SetState(37)
			p.expression(34)
		}

	case 11:
		localctx = NewTanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(38)
			p.Match(CalcParserTAN)
		}
		{
			p.SetState(39)
			p.expression(33)
		}

	case 12:
		localctx = NewCotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(40)
			p.Match(CalcParserCOT)
		}
		{
			p.SetState(41)
			p.expression(32)
		}

	case 13:
		localctx = NewSinhContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(42)
			p.Match(CalcParserSINH)
		}
		{
			p.SetState(43)
			p.expression(31)
		}

	case 14:
		localctx = NewCoshContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(44)
			p.Match(CalcParserCOSH)
		}
		{
			p.SetState(45)
			p.expression(30)
		}

	case 15:
		localctx = NewTanhContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(46)
			p.Match(CalcParserTANH)
		}
		{
			p.SetState(47)
			p.expression(29)
		}

	case 16:
		localctx = NewArcsinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.Match(CalcParserARCSIN)
		}
		{
			p.SetState(49)
			p.expression(28)
		}

	case 17:
		localctx = NewArccosContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Match(CalcParserARCCOS)
		}
		{
			p.SetState(51)
			p.expression(27)
		}

	case 18:
		localctx = NewArctanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(CalcParserARCTAN)
		}
		{
			p.SetState(53)
			p.expression(26)
		}

	case 19:
		localctx = NewArctan2Context(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Match(CalcParserARCTAN2)
		}
		{
			p.SetState(55)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(56)
			p.expression(0)
		}
		{
			p.SetState(57)
			p.Match(CalcParserT__2)
		}
		{
			p.SetState(58)
			p.expression(0)
		}
		{
			p.SetState(59)
			p.Match(CalcParserT__3)
		}

	case 20:
		localctx = NewArccotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(CalcParserARCCOT)
		}
		{
			p.SetState(62)
			p.expression(24)
		}

	case 21:
		localctx = NewExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)
			p.Match(CalcParserEXP)
		}
		{
			p.SetState(64)
			p.expression(23)
		}

	case 22:
		localctx = NewLnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Match(CalcParserLN)
		}
		{
			p.SetState(66)
			p.expression(22)
		}

	case 23:
		localctx = NewEexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)
			p.Match(CalcParserEEX)
		}
		{
			p.SetState(68)
			p.expression(21)
		}

	case 24:
		localctx = NewLogContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Match(CalcParserLOG)
		}
		{
			p.SetState(70)
			p.expression(20)
		}

	case 25:
		localctx = NewRadContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(71)
			p.Match(CalcParserRAD)
		}
		{
			p.SetState(72)
			p.expression(19)
		}

	case 26:
		localctx = NewDegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(73)
			p.Match(CalcParserDEG)
		}
		{
			p.SetState(74)
			p.expression(18)
		}

	case 27:
		localctx = NewSqrtContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(75)
			p.Match(CalcParserSQRT)
		}
		{
			p.SetState(76)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(77)
			p.expression(0)
		}
		{
			p.SetState(78)
			p.Match(CalcParserT__3)
		}

	case 28:
		localctx = NewSqrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(80)
			p.Match(CalcParserSQR)
		}
		{
			p.SetState(81)
			p.expression(16)
		}

	case 29:
		localctx = NewParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(82)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(83)
			p.expression(0)
		}
		{
			p.SetState(84)
			p.Match(CalcParserT__3)
		}

	case 30:
		localctx = NewMultContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(86)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(87)
			p.expression(0)
		}
		{
			p.SetState(88)
			p.Match(CalcParserT__3)
		}
		{
			p.SetState(89)
			p.expression(6)
		}

	case 31:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.Match(CalcParserNUMBER)
		}

	case 32:
		localctx = NewPiContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.Match(CalcParserPI)
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(93)
				p.Match(CalcParserT__9)
			}

		}

	case 33:
		localctx = NewEulerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(96)
			p.Match(CalcParserEULER)
		}

	case 34:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(97)
			p.Match(CalcParserVAR)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(100)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(101)
					p.Match(CalcParserEXPONENT)
				}
				{
					p.SetState(102)
					p.expression(16)
				}

			case 2:
				localctx = NewNegExponentContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(104)
					p.Match(CalcParserNEGEXPONENT)
				}
				{
					p.SetState(105)
					p.expression(15)
				}

			case 3:
				localctx = NewPowContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(106)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(107)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*PowContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserT__4 || _la == CalcParserT__5) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*PowContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(108)
					p.expression(14)
				}

			case 4:
				localctx = NewModContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(110)
					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserT__6 || _la == CalcParserMOD) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(111)
					p.expression(13)
				}

			case 5:
				localctx = NewWholeContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(113)
					p.Match(CalcParserWHOLE)
				}
				{
					p.SetState(114)
					p.expression(12)
				}

			case 6:
				localctx = NewSqRootContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(115)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(116)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*SqRootContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserT__7 || _la == CalcParserT__8) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*SqRootContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(117)
					p.expression(11)
				}

			case 7:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(118)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(119)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserMUL || _la == CalcParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(120)
					p.expression(10)
				}

			case 8:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(121)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(122)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserADD || _la == CalcParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(123)
					p.expression(6)
				}

			case 9:
				localctx = NewMultContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(124)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(125)
					p.Match(CalcParserT__1)
				}
				{
					p.SetState(126)
					p.expression(0)
				}
				{
					p.SetState(127)
					p.Match(CalcParserT__3)
				}

			}

		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// ICompileUnitContext is an interface to support dynamic dispatch.
type ICompileUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompileUnitContext differentiates from other interfaces.
	IsCompileUnitContext()
}

type CompileUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompileUnitContext() *CompileUnitContext {
	var p = new(CompileUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_compileUnit
	return p
}

func (*CompileUnitContext) IsCompileUnitContext() {}

func NewCompileUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompileUnitContext {
	var p = new(CompileUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_compileUnit

	return p
}

func (s *CompileUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompileUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalcParserEOF, 0)
}

func (s *CompileUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompileUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompileUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterCompileUnit(s)
	}
}

func (s *CompileUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitCompileUnit(s)
	}
}

func (p *CalcParser) CompileUnit() (localctx ICompileUnitContext) {
	localctx = NewCompileUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcParserRULE_compileUnit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(CalcParserEOF)
	}

	return localctx
}

func (p *CalcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalcParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
