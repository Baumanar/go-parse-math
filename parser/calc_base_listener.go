// Code generated from ../Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCalcListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCalcListener) ExitStart(ctx *StartContext) {}

// EnterTan is called when production Tan is entered.
func (s *BaseCalcListener) EnterTan(ctx *TanContext) {}

// ExitTan is called when production Tan is exited.
func (s *BaseCalcListener) ExitTan(ctx *TanContext) {}

// EnterCosh is called when production Cosh is entered.
func (s *BaseCalcListener) EnterCosh(ctx *CoshContext) {}

// ExitCosh is called when production Cosh is exited.
func (s *BaseCalcListener) ExitCosh(ctx *CoshContext) {}

// EnterSqRoot is called when production SqRoot is entered.
func (s *BaseCalcListener) EnterSqRoot(ctx *SqRootContext) {}

// ExitSqRoot is called when production SqRoot is exited.
func (s *BaseCalcListener) ExitSqRoot(ctx *SqRootContext) {}

// EnterVariable is called when production Variable is entered.
func (s *BaseCalcListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production Variable is exited.
func (s *BaseCalcListener) ExitVariable(ctx *VariableContext) {}

// EnterNegExponent is called when production NegExponent is entered.
func (s *BaseCalcListener) EnterNegExponent(ctx *NegExponentContext) {}

// ExitNegExponent is called when production NegExponent is exited.
func (s *BaseCalcListener) ExitNegExponent(ctx *NegExponentContext) {}

// EnterExponent is called when production Exponent is entered.
func (s *BaseCalcListener) EnterExponent(ctx *ExponentContext) {}

// ExitExponent is called when production Exponent is exited.
func (s *BaseCalcListener) ExitExponent(ctx *ExponentContext) {}

// EnterArctan2 is called when production Arctan2 is entered.
func (s *BaseCalcListener) EnterArctan2(ctx *Arctan2Context) {}

// ExitArctan2 is called when production Arctan2 is exited.
func (s *BaseCalcListener) ExitArctan2(ctx *Arctan2Context) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseCalcListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseCalcListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterArcsin is called when production Arcsin is entered.
func (s *BaseCalcListener) EnterArcsin(ctx *ArcsinContext) {}

// ExitArcsin is called when production Arcsin is exited.
func (s *BaseCalcListener) ExitArcsin(ctx *ArcsinContext) {}

// EnterUnaryPlus is called when production UnaryPlus is entered.
func (s *BaseCalcListener) EnterUnaryPlus(ctx *UnaryPlusContext) {}

// ExitUnaryPlus is called when production UnaryPlus is exited.
func (s *BaseCalcListener) ExitUnaryPlus(ctx *UnaryPlusContext) {}

// EnterArccot is called when production Arccot is entered.
func (s *BaseCalcListener) EnterArccot(ctx *ArccotContext) {}

// ExitArccot is called when production Arccot is exited.
func (s *BaseCalcListener) ExitArccot(ctx *ArccotContext) {}

// EnterArccos is called when production Arccos is entered.
func (s *BaseCalcListener) EnterArccos(ctx *ArccosContext) {}

// ExitArccos is called when production Arccos is exited.
func (s *BaseCalcListener) ExitArccos(ctx *ArccosContext) {}

// EnterEuler is called when production Euler is entered.
func (s *BaseCalcListener) EnterEuler(ctx *EulerContext) {}

// ExitEuler is called when production Euler is exited.
func (s *BaseCalcListener) ExitEuler(ctx *EulerContext) {}

// EnterArctan is called when production Arctan is entered.
func (s *BaseCalcListener) EnterArctan(ctx *ArctanContext) {}

// ExitArctan is called when production Arctan is exited.
func (s *BaseCalcListener) ExitArctan(ctx *ArctanContext) {}

// EnterParenthesis is called when production Parenthesis is entered.
func (s *BaseCalcListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production Parenthesis is exited.
func (s *BaseCalcListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterAbs is called when production Abs is entered.
func (s *BaseCalcListener) EnterAbs(ctx *AbsContext) {}

// ExitAbs is called when production Abs is exited.
func (s *BaseCalcListener) ExitAbs(ctx *AbsContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseCalcListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseCalcListener) ExitNumber(ctx *NumberContext) {}

// EnterSinh is called when production Sinh is entered.
func (s *BaseCalcListener) EnterSinh(ctx *SinhContext) {}

// ExitSinh is called when production Sinh is exited.
func (s *BaseCalcListener) ExitSinh(ctx *SinhContext) {}

// EnterRound is called when production Round is entered.
func (s *BaseCalcListener) EnterRound(ctx *RoundContext) {}

// ExitRound is called when production Round is exited.
func (s *BaseCalcListener) ExitRound(ctx *RoundContext) {}

// EnterTrunc is called when production Trunc is entered.
func (s *BaseCalcListener) EnterTrunc(ctx *TruncContext) {}

// ExitTrunc is called when production Trunc is exited.
func (s *BaseCalcListener) ExitTrunc(ctx *TruncContext) {}

// EnterPi is called when production Pi is entered.
func (s *BaseCalcListener) EnterPi(ctx *PiContext) {}

// ExitPi is called when production Pi is exited.
func (s *BaseCalcListener) ExitPi(ctx *PiContext) {}

// EnterTanh is called when production Tanh is entered.
func (s *BaseCalcListener) EnterTanh(ctx *TanhContext) {}

// ExitTanh is called when production Tanh is exited.
func (s *BaseCalcListener) ExitTanh(ctx *TanhContext) {}

// EnterFloor is called when production Floor is entered.
func (s *BaseCalcListener) EnterFloor(ctx *FloorContext) {}

// ExitFloor is called when production Floor is exited.
func (s *BaseCalcListener) ExitFloor(ctx *FloorContext) {}

// EnterLn is called when production Ln is entered.
func (s *BaseCalcListener) EnterLn(ctx *LnContext) {}

// ExitLn is called when production Ln is exited.
func (s *BaseCalcListener) ExitLn(ctx *LnContext) {}

// EnterMod is called when production Mod is entered.
func (s *BaseCalcListener) EnterMod(ctx *ModContext) {}

// ExitMod is called when production Mod is exited.
func (s *BaseCalcListener) ExitMod(ctx *ModContext) {}

// EnterLog is called when production Log is entered.
func (s *BaseCalcListener) EnterLog(ctx *LogContext) {}

// ExitLog is called when production Log is exited.
func (s *BaseCalcListener) ExitLog(ctx *LogContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalcListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalcListener) ExitAddSub(ctx *AddSubContext) {}

// EnterCos is called when production Cos is entered.
func (s *BaseCalcListener) EnterCos(ctx *CosContext) {}

// ExitCos is called when production Cos is exited.
func (s *BaseCalcListener) ExitCos(ctx *CosContext) {}

// EnterDeg is called when production Deg is entered.
func (s *BaseCalcListener) EnterDeg(ctx *DegContext) {}

// ExitDeg is called when production Deg is exited.
func (s *BaseCalcListener) ExitDeg(ctx *DegContext) {}

// EnterSqrt is called when production Sqrt is entered.
func (s *BaseCalcListener) EnterSqrt(ctx *SqrtContext) {}

// ExitSqrt is called when production Sqrt is exited.
func (s *BaseCalcListener) ExitSqrt(ctx *SqrtContext) {}

// EnterCot is called when production Cot is entered.
func (s *BaseCalcListener) EnterCot(ctx *CotContext) {}

// ExitCot is called when production Cot is exited.
func (s *BaseCalcListener) ExitCot(ctx *CotContext) {}

// EnterWhole is called when production Whole is entered.
func (s *BaseCalcListener) EnterWhole(ctx *WholeContext) {}

// ExitWhole is called when production Whole is exited.
func (s *BaseCalcListener) ExitWhole(ctx *WholeContext) {}

// EnterUnary is called when production Unary is entered.
func (s *BaseCalcListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production Unary is exited.
func (s *BaseCalcListener) ExitUnary(ctx *UnaryContext) {}

// EnterRad is called when production Rad is entered.
func (s *BaseCalcListener) EnterRad(ctx *RadContext) {}

// ExitRad is called when production Rad is exited.
func (s *BaseCalcListener) ExitRad(ctx *RadContext) {}

// EnterMult is called when production Mult is entered.
func (s *BaseCalcListener) EnterMult(ctx *MultContext) {}

// ExitMult is called when production Mult is exited.
func (s *BaseCalcListener) ExitMult(ctx *MultContext) {}

// EnterSqr is called when production Sqr is entered.
func (s *BaseCalcListener) EnterSqr(ctx *SqrContext) {}

// ExitSqr is called when production Sqr is exited.
func (s *BaseCalcListener) ExitSqr(ctx *SqrContext) {}

// EnterSin is called when production Sin is entered.
func (s *BaseCalcListener) EnterSin(ctx *SinContext) {}

// ExitSin is called when production Sin is exited.
func (s *BaseCalcListener) ExitSin(ctx *SinContext) {}

// EnterEex is called when production Eex is entered.
func (s *BaseCalcListener) EnterEex(ctx *EexContext) {}

// ExitEex is called when production Eex is exited.
func (s *BaseCalcListener) ExitEex(ctx *EexContext) {}

// EnterPow is called when production Pow is entered.
func (s *BaseCalcListener) EnterPow(ctx *PowContext) {}

// ExitPow is called when production Pow is exited.
func (s *BaseCalcListener) ExitPow(ctx *PowContext) {}

// EnterCeil is called when production Ceil is entered.
func (s *BaseCalcListener) EnterCeil(ctx *CeilContext) {}

// ExitCeil is called when production Ceil is exited.
func (s *BaseCalcListener) ExitCeil(ctx *CeilContext) {}

// EnterExp is called when production Exp is entered.
func (s *BaseCalcListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production Exp is exited.
func (s *BaseCalcListener) ExitExp(ctx *ExpContext) {}

// EnterRoundk is called when production Roundk is entered.
func (s *BaseCalcListener) EnterRoundk(ctx *RoundkContext) {}

// ExitRoundk is called when production Roundk is exited.
func (s *BaseCalcListener) ExitRoundk(ctx *RoundkContext) {}

// EnterCompileUnit is called when production compileUnit is entered.
func (s *BaseCalcListener) EnterCompileUnit(ctx *CompileUnitContext) {}

// ExitCompileUnit is called when production compileUnit is exited.
func (s *BaseCalcListener) ExitCompileUnit(ctx *CompileUnitContext) {}
