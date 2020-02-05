// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterTan is called when entering the Tan production.
	EnterTan(c *TanContext)

	// EnterCosh is called when entering the Cosh production.
	EnterCosh(c *CoshContext)

	// EnterSqRoot is called when entering the SqRoot production.
	EnterSqRoot(c *SqRootContext)

	// EnterVariable is called when entering the Variable production.
	EnterVariable(c *VariableContext)

	// EnterNegExponent is called when entering the NegExponent production.
	EnterNegExponent(c *NegExponentContext)

	// EnterExponent is called when entering the Exponent production.
	EnterExponent(c *ExponentContext)

	// EnterArctan2 is called when entering the Arctan2 production.
	EnterArctan2(c *Arctan2Context)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterArcsin is called when entering the Arcsin production.
	EnterArcsin(c *ArcsinContext)

	// EnterUnaryPlus is called when entering the UnaryPlus production.
	EnterUnaryPlus(c *UnaryPlusContext)

	// EnterArccot is called when entering the Arccot production.
	EnterArccot(c *ArccotContext)

	// EnterArccos is called when entering the Arccos production.
	EnterArccos(c *ArccosContext)

	// EnterEuler is called when entering the Euler production.
	EnterEuler(c *EulerContext)

	// EnterArctan is called when entering the Arctan production.
	EnterArctan(c *ArctanContext)

	// EnterParenthesis is called when entering the Parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterAbs is called when entering the Abs production.
	EnterAbs(c *AbsContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterSinh is called when entering the Sinh production.
	EnterSinh(c *SinhContext)

	// EnterRound is called when entering the Round production.
	EnterRound(c *RoundContext)

	// EnterTrunc is called when entering the Trunc production.
	EnterTrunc(c *TruncContext)

	// EnterPi is called when entering the Pi production.
	EnterPi(c *PiContext)

	// EnterTanh is called when entering the Tanh production.
	EnterTanh(c *TanhContext)

	// EnterFloor is called when entering the Floor production.
	EnterFloor(c *FloorContext)

	// EnterLn is called when entering the Ln production.
	EnterLn(c *LnContext)

	// EnterMod is called when entering the Mod production.
	EnterMod(c *ModContext)

	// EnterLog is called when entering the Log production.
	EnterLog(c *LogContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterCos is called when entering the Cos production.
	EnterCos(c *CosContext)

	// EnterDeg is called when entering the Deg production.
	EnterDeg(c *DegContext)

	// EnterSqrt is called when entering the Sqrt production.
	EnterSqrt(c *SqrtContext)

	// EnterCot is called when entering the Cot production.
	EnterCot(c *CotContext)

	// EnterWhole is called when entering the Whole production.
	EnterWhole(c *WholeContext)

	// EnterUnary is called when entering the Unary production.
	EnterUnary(c *UnaryContext)

	// EnterRad is called when entering the Rad production.
	EnterRad(c *RadContext)

	// EnterMult is called when entering the Mult production.
	EnterMult(c *MultContext)

	// EnterSqr is called when entering the Sqr production.
	EnterSqr(c *SqrContext)

	// EnterSin is called when entering the Sin production.
	EnterSin(c *SinContext)

	// EnterEex is called when entering the Eex production.
	EnterEex(c *EexContext)

	// EnterPow is called when entering the Pow production.
	EnterPow(c *PowContext)

	// EnterCeil is called when entering the Ceil production.
	EnterCeil(c *CeilContext)

	// EnterExp is called when entering the Exp production.
	EnterExp(c *ExpContext)

	// EnterRoundk is called when entering the Roundk production.
	EnterRoundk(c *RoundkContext)

	// EnterCompileUnit is called when entering the compileUnit production.
	EnterCompileUnit(c *CompileUnitContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitTan is called when exiting the Tan production.
	ExitTan(c *TanContext)

	// ExitCosh is called when exiting the Cosh production.
	ExitCosh(c *CoshContext)

	// ExitSqRoot is called when exiting the SqRoot production.
	ExitSqRoot(c *SqRootContext)

	// ExitVariable is called when exiting the Variable production.
	ExitVariable(c *VariableContext)

	// ExitNegExponent is called when exiting the NegExponent production.
	ExitNegExponent(c *NegExponentContext)

	// ExitExponent is called when exiting the Exponent production.
	ExitExponent(c *ExponentContext)

	// ExitArctan2 is called when exiting the Arctan2 production.
	ExitArctan2(c *Arctan2Context)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitArcsin is called when exiting the Arcsin production.
	ExitArcsin(c *ArcsinContext)

	// ExitUnaryPlus is called when exiting the UnaryPlus production.
	ExitUnaryPlus(c *UnaryPlusContext)

	// ExitArccot is called when exiting the Arccot production.
	ExitArccot(c *ArccotContext)

	// ExitArccos is called when exiting the Arccos production.
	ExitArccos(c *ArccosContext)

	// ExitEuler is called when exiting the Euler production.
	ExitEuler(c *EulerContext)

	// ExitArctan is called when exiting the Arctan production.
	ExitArctan(c *ArctanContext)

	// ExitParenthesis is called when exiting the Parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitAbs is called when exiting the Abs production.
	ExitAbs(c *AbsContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitSinh is called when exiting the Sinh production.
	ExitSinh(c *SinhContext)

	// ExitRound is called when exiting the Round production.
	ExitRound(c *RoundContext)

	// ExitTrunc is called when exiting the Trunc production.
	ExitTrunc(c *TruncContext)

	// ExitPi is called when exiting the Pi production.
	ExitPi(c *PiContext)

	// ExitTanh is called when exiting the Tanh production.
	ExitTanh(c *TanhContext)

	// ExitFloor is called when exiting the Floor production.
	ExitFloor(c *FloorContext)

	// ExitLn is called when exiting the Ln production.
	ExitLn(c *LnContext)

	// ExitMod is called when exiting the Mod production.
	ExitMod(c *ModContext)

	// ExitLog is called when exiting the Log production.
	ExitLog(c *LogContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitCos is called when exiting the Cos production.
	ExitCos(c *CosContext)

	// ExitDeg is called when exiting the Deg production.
	ExitDeg(c *DegContext)

	// ExitSqrt is called when exiting the Sqrt production.
	ExitSqrt(c *SqrtContext)

	// ExitCot is called when exiting the Cot production.
	ExitCot(c *CotContext)

	// ExitWhole is called when exiting the Whole production.
	ExitWhole(c *WholeContext)

	// ExitUnary is called when exiting the Unary production.
	ExitUnary(c *UnaryContext)

	// ExitRad is called when exiting the Rad production.
	ExitRad(c *RadContext)

	// ExitMult is called when exiting the Mult production.
	ExitMult(c *MultContext)

	// ExitSqr is called when exiting the Sqr production.
	ExitSqr(c *SqrContext)

	// ExitSin is called when exiting the Sin production.
	ExitSin(c *SinContext)

	// ExitEex is called when exiting the Eex production.
	ExitEex(c *EexContext)

	// ExitPow is called when exiting the Pow production.
	ExitPow(c *PowContext)

	// ExitCeil is called when exiting the Ceil production.
	ExitCeil(c *CeilContext)

	// ExitExp is called when exiting the Exp production.
	ExitExp(c *ExpContext)

	// ExitRoundk is called when exiting the Roundk production.
	ExitRoundk(c *RoundkContext)

	// ExitCompileUnit is called when exiting the compileUnit production.
	ExitCompileUnit(c *CompileUnitContext)
}
