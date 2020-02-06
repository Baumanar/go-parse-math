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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 56, 185,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 5, 2, 11, 10, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 143, 10, 3,
	3, 3, 3, 3, 5, 3, 147, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 178,
	10, 3, 12, 3, 14, 3, 181, 11, 3, 3, 4, 3, 4, 3, 4, 2, 3, 4, 5, 2, 4, 6,
	2, 7, 3, 2, 7, 8, 4, 2, 9, 9, 17, 17, 3, 2, 10, 11, 3, 2, 19, 20, 3, 2,
	21, 22, 2, 226, 2, 8, 3, 2, 2, 2, 4, 146, 3, 2, 2, 2, 6, 182, 3, 2, 2,
	2, 8, 10, 5, 4, 3, 2, 9, 11, 7, 3, 2, 2, 10, 9, 3, 2, 2, 2, 10, 11, 3,
	2, 2, 2, 11, 12, 3, 2, 2, 2, 12, 13, 5, 6, 4, 2, 13, 3, 3, 2, 2, 2, 14,
	15, 8, 3, 1, 2, 15, 16, 7, 22, 2, 2, 16, 147, 5, 4, 3, 46, 17, 18, 7, 21,
	2, 2, 18, 147, 5, 4, 3, 45, 19, 20, 7, 30, 2, 2, 20, 147, 5, 4, 3, 44,
	21, 22, 7, 31, 2, 2, 22, 147, 5, 4, 3, 43, 23, 24, 7, 32, 2, 2, 24, 147,
	5, 4, 3, 42, 25, 26, 7, 23, 2, 2, 26, 27, 5, 4, 3, 2, 27, 28, 7, 23, 2,
	2, 28, 147, 3, 2, 2, 2, 29, 30, 7, 33, 2, 2, 30, 31, 7, 4, 2, 2, 31, 32,
	5, 4, 3, 2, 32, 33, 7, 5, 2, 2, 33, 34, 5, 4, 3, 2, 34, 35, 7, 6, 2, 2,
	35, 147, 3, 2, 2, 2, 36, 37, 7, 34, 2, 2, 37, 147, 5, 4, 3, 39, 38, 39,
	7, 35, 2, 2, 39, 147, 5, 4, 3, 38, 40, 41, 7, 36, 2, 2, 41, 42, 7, 4, 2,
	2, 42, 43, 5, 4, 3, 2, 43, 44, 7, 6, 2, 2, 44, 147, 3, 2, 2, 2, 45, 46,
	7, 37, 2, 2, 46, 47, 7, 4, 2, 2, 47, 48, 5, 4, 3, 2, 48, 49, 7, 6, 2, 2,
	49, 147, 3, 2, 2, 2, 50, 51, 7, 38, 2, 2, 51, 52, 7, 4, 2, 2, 52, 53, 5,
	4, 3, 2, 53, 54, 7, 6, 2, 2, 54, 147, 3, 2, 2, 2, 55, 56, 7, 39, 2, 2,
	56, 57, 7, 4, 2, 2, 57, 58, 5, 4, 3, 2, 58, 59, 7, 6, 2, 2, 59, 147, 3,
	2, 2, 2, 60, 61, 7, 40, 2, 2, 61, 62, 7, 4, 2, 2, 62, 63, 5, 4, 3, 2, 63,
	64, 7, 6, 2, 2, 64, 147, 3, 2, 2, 2, 65, 66, 7, 41, 2, 2, 66, 67, 7, 4,
	2, 2, 67, 68, 5, 4, 3, 2, 68, 69, 7, 6, 2, 2, 69, 147, 3, 2, 2, 2, 70,
	71, 7, 42, 2, 2, 71, 72, 7, 4, 2, 2, 72, 73, 5, 4, 3, 2, 73, 74, 7, 6,
	2, 2, 74, 147, 3, 2, 2, 2, 75, 76, 7, 43, 2, 2, 76, 77, 7, 4, 2, 2, 77,
	78, 5, 4, 3, 2, 78, 79, 7, 6, 2, 2, 79, 147, 3, 2, 2, 2, 80, 81, 7, 44,
	2, 2, 81, 82, 7, 4, 2, 2, 82, 83, 5, 4, 3, 2, 83, 84, 7, 6, 2, 2, 84, 147,
	3, 2, 2, 2, 85, 86, 7, 45, 2, 2, 86, 87, 7, 4, 2, 2, 87, 88, 5, 4, 3, 2,
	88, 89, 7, 6, 2, 2, 89, 147, 3, 2, 2, 2, 90, 91, 7, 46, 2, 2, 91, 92, 7,
	4, 2, 2, 92, 93, 5, 4, 3, 2, 93, 94, 7, 5, 2, 2, 94, 95, 5, 4, 3, 2, 95,
	96, 7, 6, 2, 2, 96, 147, 3, 2, 2, 2, 97, 98, 7, 47, 2, 2, 98, 99, 7, 4,
	2, 2, 99, 100, 5, 4, 3, 2, 100, 101, 7, 6, 2, 2, 101, 147, 3, 2, 2, 2,
	102, 103, 7, 48, 2, 2, 103, 147, 5, 4, 3, 25, 104, 105, 7, 49, 2, 2, 105,
	106, 7, 4, 2, 2, 106, 107, 5, 4, 3, 2, 107, 108, 7, 6, 2, 2, 108, 147,
	3, 2, 2, 2, 109, 110, 7, 50, 2, 2, 110, 147, 5, 4, 3, 23, 111, 112, 7,
	51, 2, 2, 112, 113, 7, 4, 2, 2, 113, 114, 5, 4, 3, 2, 114, 115, 7, 6, 2,
	2, 115, 147, 3, 2, 2, 2, 116, 117, 7, 52, 2, 2, 117, 147, 5, 4, 3, 21,
	118, 119, 7, 53, 2, 2, 119, 147, 5, 4, 3, 20, 120, 121, 7, 28, 2, 2, 121,
	122, 7, 4, 2, 2, 122, 123, 5, 4, 3, 2, 123, 124, 7, 6, 2, 2, 124, 147,
	3, 2, 2, 2, 125, 126, 7, 29, 2, 2, 126, 127, 7, 4, 2, 2, 127, 128, 5, 4,
	3, 2, 128, 129, 7, 6, 2, 2, 129, 147, 3, 2, 2, 2, 130, 131, 7, 4, 2, 2,
	131, 132, 5, 4, 3, 2, 132, 133, 7, 6, 2, 2, 133, 147, 3, 2, 2, 2, 134,
	135, 7, 4, 2, 2, 135, 136, 5, 4, 3, 2, 136, 137, 7, 6, 2, 2, 137, 138,
	5, 4, 3, 8, 138, 147, 3, 2, 2, 2, 139, 147, 7, 13, 2, 2, 140, 142, 7, 24,
	2, 2, 141, 143, 7, 12, 2, 2, 142, 141, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2,
	143, 147, 3, 2, 2, 2, 144, 147, 7, 27, 2, 2, 145, 147, 7, 14, 2, 2, 146,
	14, 3, 2, 2, 2, 146, 17, 3, 2, 2, 2, 146, 19, 3, 2, 2, 2, 146, 21, 3, 2,
	2, 2, 146, 23, 3, 2, 2, 2, 146, 25, 3, 2, 2, 2, 146, 29, 3, 2, 2, 2, 146,
	36, 3, 2, 2, 2, 146, 38, 3, 2, 2, 2, 146, 40, 3, 2, 2, 2, 146, 45, 3, 2,
	2, 2, 146, 50, 3, 2, 2, 2, 146, 55, 3, 2, 2, 2, 146, 60, 3, 2, 2, 2, 146,
	65, 3, 2, 2, 2, 146, 70, 3, 2, 2, 2, 146, 75, 3, 2, 2, 2, 146, 80, 3, 2,
	2, 2, 146, 85, 3, 2, 2, 2, 146, 90, 3, 2, 2, 2, 146, 97, 3, 2, 2, 2, 146,
	102, 3, 2, 2, 2, 146, 104, 3, 2, 2, 2, 146, 109, 3, 2, 2, 2, 146, 111,
	3, 2, 2, 2, 146, 116, 3, 2, 2, 2, 146, 118, 3, 2, 2, 2, 146, 120, 3, 2,
	2, 2, 146, 125, 3, 2, 2, 2, 146, 130, 3, 2, 2, 2, 146, 134, 3, 2, 2, 2,
	146, 139, 3, 2, 2, 2, 146, 140, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146,
	145, 3, 2, 2, 2, 147, 179, 3, 2, 2, 2, 148, 149, 12, 17, 2, 2, 149, 150,
	7, 25, 2, 2, 150, 178, 5, 4, 3, 18, 151, 152, 12, 16, 2, 2, 152, 153, 7,
	26, 2, 2, 153, 178, 5, 4, 3, 17, 154, 155, 12, 15, 2, 2, 155, 156, 9, 2,
	2, 2, 156, 178, 5, 4, 3, 16, 157, 158, 12, 14, 2, 2, 158, 159, 9, 3, 2,
	2, 159, 178, 5, 4, 3, 15, 160, 161, 12, 13, 2, 2, 161, 162, 7, 18, 2, 2,
	162, 178, 5, 4, 3, 14, 163, 164, 12, 12, 2, 2, 164, 165, 9, 4, 2, 2, 165,
	178, 5, 4, 3, 13, 166, 167, 12, 11, 2, 2, 167, 168, 9, 5, 2, 2, 168, 178,
	5, 4, 3, 12, 169, 170, 12, 7, 2, 2, 170, 171, 9, 6, 2, 2, 171, 178, 5,
	4, 3, 8, 172, 173, 12, 9, 2, 2, 173, 174, 7, 4, 2, 2, 174, 175, 5, 4, 3,
	2, 175, 176, 7, 6, 2, 2, 176, 178, 3, 2, 2, 2, 177, 148, 3, 2, 2, 2, 177,
	151, 3, 2, 2, 2, 177, 154, 3, 2, 2, 2, 177, 157, 3, 2, 2, 2, 177, 160,
	3, 2, 2, 2, 177, 163, 3, 2, 2, 2, 177, 166, 3, 2, 2, 2, 177, 169, 3, 2,
	2, 2, 177, 172, 3, 2, 2, 2, 178, 181, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2,
	179, 180, 3, 2, 2, 2, 180, 5, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 182, 183,
	7, 2, 2, 3, 183, 7, 3, 2, 2, 2, 7, 10, 142, 146, 177, 179,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'('", "';'", "')'", "'^'", "'**'", "'%'", "'~'", "'//'", "'()'",
	"", "'x'", "", "", "", "", "'*'", "'/'", "'+'", "'-'", "'|'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "NUMBER", "VAR", "FLOAT", "DIGIT",
	"MOD", "WHOLE", "MUL", "DIV", "ADD", "SUB", "ABSLINE", "PI", "EXPONENT",
	"NEGEXPONENT", "EULER", "SQRT", "SQR", "FLOOR", "CEIL", "ABS", "ROUNDK",
	"ROUND", "TRUNC", "SIN", "COS", "TAN", "COT", "SINH", "COSH", "TANH", "ARCSIN",
	"ARCCOS", "ARCTAN", "ARCTAN2", "ARCCOT", "EXP", "LN", "EEX", "LOG", "RAD",
	"DEG", "WS", "COM", "INVALID",
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
	CalcParserABSLINE     = 21
	CalcParserPI          = 22
	CalcParserEXPONENT    = 23
	CalcParserNEGEXPONENT = 24
	CalcParserEULER       = 25
	CalcParserSQRT        = 26
	CalcParserSQR         = 27
	CalcParserFLOOR       = 28
	CalcParserCEIL        = 29
	CalcParserABS         = 30
	CalcParserROUNDK      = 31
	CalcParserROUND       = 32
	CalcParserTRUNC       = 33
	CalcParserSIN         = 34
	CalcParserCOS         = 35
	CalcParserTAN         = 36
	CalcParserCOT         = 37
	CalcParserSINH        = 38
	CalcParserCOSH        = 39
	CalcParserTANH        = 40
	CalcParserARCSIN      = 41
	CalcParserARCCOS      = 42
	CalcParserARCTAN      = 43
	CalcParserARCTAN2     = 44
	CalcParserARCCOT      = 45
	CalcParserEXP         = 46
	CalcParserLN          = 47
	CalcParserEEX         = 48
	CalcParserLOG         = 49
	CalcParserRAD         = 50
	CalcParserDEG         = 51
	CalcParserWS          = 52
	CalcParserCOM         = 53
	CalcParserINVALID     = 54
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

func (s *AbsContext) AllABSLINE() []antlr.TerminalNode {
	return s.GetTokens(CalcParserABSLINE)
}

func (s *AbsContext) ABSLINE(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserABSLINE, i)
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
	p.SetState(144)
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
			p.expression(44)
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
			p.expression(43)
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
			p.expression(42)
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
			p.expression(41)
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
			p.expression(40)
		}

	case 6:
		localctx = NewAbsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.Match(CalcParserABSLINE)
		}
		{
			p.SetState(24)
			p.expression(0)
		}
		{
			p.SetState(25)
			p.Match(CalcParserABSLINE)
		}

	case 7:
		localctx = NewRoundkContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)
			p.Match(CalcParserROUNDK)
		}
		{
			p.SetState(28)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(29)
			p.expression(0)
		}
		{
			p.SetState(30)
			p.Match(CalcParserT__2)
		}
		{
			p.SetState(31)
			p.expression(0)
		}
		{
			p.SetState(32)
			p.Match(CalcParserT__3)
		}

	case 8:
		localctx = NewRoundContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			p.Match(CalcParserROUND)
		}
		{
			p.SetState(35)
			p.expression(37)
		}

	case 9:
		localctx = NewTruncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(36)
			p.Match(CalcParserTRUNC)
		}
		{
			p.SetState(37)
			p.expression(36)
		}

	case 10:
		localctx = NewSinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(38)
			p.Match(CalcParserSIN)
		}
		{
			p.SetState(39)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(40)
			p.expression(0)
		}
		{
			p.SetState(41)
			p.Match(CalcParserT__3)
		}

	case 11:
		localctx = NewCosContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(43)
			p.Match(CalcParserCOS)
		}
		{
			p.SetState(44)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(45)
			p.expression(0)
		}
		{
			p.SetState(46)
			p.Match(CalcParserT__3)
		}

	case 12:
		localctx = NewTanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.Match(CalcParserTAN)
		}
		{
			p.SetState(49)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(50)
			p.expression(0)
		}
		{
			p.SetState(51)
			p.Match(CalcParserT__3)
		}

	case 13:
		localctx = NewCotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			p.Match(CalcParserCOT)
		}
		{
			p.SetState(54)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(55)
			p.expression(0)
		}
		{
			p.SetState(56)
			p.Match(CalcParserT__3)
		}

	case 14:
		localctx = NewSinhContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.Match(CalcParserSINH)
		}
		{
			p.SetState(59)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(60)
			p.expression(0)
		}
		{
			p.SetState(61)
			p.Match(CalcParserT__3)
		}

	case 15:
		localctx = NewCoshContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)
			p.Match(CalcParserCOSH)
		}
		{
			p.SetState(64)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(65)
			p.expression(0)
		}
		{
			p.SetState(66)
			p.Match(CalcParserT__3)
		}

	case 16:
		localctx = NewTanhContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(CalcParserTANH)
		}
		{
			p.SetState(69)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(70)
			p.expression(0)
		}
		{
			p.SetState(71)
			p.Match(CalcParserT__3)
		}

	case 17:
		localctx = NewArcsinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(73)
			p.Match(CalcParserARCSIN)
		}
		{
			p.SetState(74)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(75)
			p.expression(0)
		}
		{
			p.SetState(76)
			p.Match(CalcParserT__3)
		}

	case 18:
		localctx = NewArccosContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(78)
			p.Match(CalcParserARCCOS)
		}
		{
			p.SetState(79)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(80)
			p.expression(0)
		}
		{
			p.SetState(81)
			p.Match(CalcParserT__3)
		}

	case 19:
		localctx = NewArctanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(83)
			p.Match(CalcParserARCTAN)
		}
		{
			p.SetState(84)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(85)
			p.expression(0)
		}
		{
			p.SetState(86)
			p.Match(CalcParserT__3)
		}

	case 20:
		localctx = NewArctan2Context(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Match(CalcParserARCTAN2)
		}
		{
			p.SetState(89)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(90)
			p.expression(0)
		}
		{
			p.SetState(91)
			p.Match(CalcParserT__2)
		}
		{
			p.SetState(92)
			p.expression(0)
		}
		{
			p.SetState(93)
			p.Match(CalcParserT__3)
		}

	case 21:
		localctx = NewArccotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(95)
			p.Match(CalcParserARCCOT)
		}
		{
			p.SetState(96)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(97)
			p.expression(0)
		}
		{
			p.SetState(98)
			p.Match(CalcParserT__3)
		}

	case 22:
		localctx = NewExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.Match(CalcParserEXP)
		}
		{
			p.SetState(101)
			p.expression(23)
		}

	case 23:
		localctx = NewLnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(102)
			p.Match(CalcParserLN)
		}
		{
			p.SetState(103)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(104)
			p.expression(0)
		}
		{
			p.SetState(105)
			p.Match(CalcParserT__3)
		}

	case 24:
		localctx = NewEexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(107)
			p.Match(CalcParserEEX)
		}
		{
			p.SetState(108)
			p.expression(21)
		}

	case 25:
		localctx = NewLogContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(109)
			p.Match(CalcParserLOG)
		}
		{
			p.SetState(110)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(111)
			p.expression(0)
		}
		{
			p.SetState(112)
			p.Match(CalcParserT__3)
		}

	case 26:
		localctx = NewRadContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(114)
			p.Match(CalcParserRAD)
		}
		{
			p.SetState(115)
			p.expression(19)
		}

	case 27:
		localctx = NewDegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(116)
			p.Match(CalcParserDEG)
		}
		{
			p.SetState(117)
			p.expression(18)
		}

	case 28:
		localctx = NewSqrtContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(118)
			p.Match(CalcParserSQRT)
		}
		{
			p.SetState(119)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(120)
			p.expression(0)
		}
		{
			p.SetState(121)
			p.Match(CalcParserT__3)
		}

	case 29:
		localctx = NewSqrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(123)
			p.Match(CalcParserSQR)
		}
		{
			p.SetState(124)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(125)
			p.expression(0)
		}
		{
			p.SetState(126)
			p.Match(CalcParserT__3)
		}

	case 30:
		localctx = NewParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(128)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(129)
			p.expression(0)
		}
		{
			p.SetState(130)
			p.Match(CalcParserT__3)
		}

	case 31:
		localctx = NewMultContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(132)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(133)
			p.expression(0)
		}
		{
			p.SetState(134)
			p.Match(CalcParserT__3)
		}
		{
			p.SetState(135)
			p.expression(6)
		}

	case 32:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(137)
			p.Match(CalcParserNUMBER)
		}

	case 33:
		localctx = NewPiContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(138)
			p.Match(CalcParserPI)
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(139)
				p.Match(CalcParserT__9)
			}

		}

	case 34:
		localctx = NewEulerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(142)
			p.Match(CalcParserEULER)
		}

	case 35:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(143)
			p.Match(CalcParserVAR)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(147)
					p.Match(CalcParserEXPONENT)
				}
				{
					p.SetState(148)
					p.expression(16)
				}

			case 2:
				localctx = NewNegExponentContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(150)
					p.Match(CalcParserNEGEXPONENT)
				}
				{
					p.SetState(151)
					p.expression(15)
				}

			case 3:
				localctx = NewPowContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(153)

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
					p.SetState(154)
					p.expression(14)
				}

			case 4:
				localctx = NewModContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(156)
					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserT__6 || _la == CalcParserMOD) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(157)
					p.expression(13)
				}

			case 5:
				localctx = NewWholeContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(159)
					p.Match(CalcParserWHOLE)
				}
				{
					p.SetState(160)
					p.expression(12)
				}

			case 6:
				localctx = NewSqRootContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(161)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(162)

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
					p.SetState(163)
					p.expression(11)
				}

			case 7:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(164)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(165)

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
					p.SetState(166)
					p.expression(10)
				}

			case 8:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(167)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(168)

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
					p.SetState(169)
					p.expression(6)
				}

			case 9:
				localctx = NewMultContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expression)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(171)
					p.Match(CalcParserT__1)
				}
				{
					p.SetState(172)
					p.expression(0)
				}
				{
					p.SetState(173)
					p.Match(CalcParserT__3)
				}

			}

		}
		p.SetState(179)
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
		p.SetState(180)
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
