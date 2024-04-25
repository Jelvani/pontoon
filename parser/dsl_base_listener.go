// Code generated from dsl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // dsl

import "github.com/antlr4-go/antlr/v4"

// BasedslListener is a complete listener for a parse tree produced by dslParser.
type BasedslListener struct{}

var _ dslListener = &BasedslListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasedslListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasedslListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasedslListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasedslListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTypesDeclaration is called when production typesDeclaration is entered.
func (s *BasedslListener) EnterTypesDeclaration(ctx *TypesDeclarationContext) {}

// ExitTypesDeclaration is called when production typesDeclaration is exited.
func (s *BasedslListener) ExitTypesDeclaration(ctx *TypesDeclarationContext) {}

// EnterProgram is called when production program is entered.
func (s *BasedslListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasedslListener) ExitProgram(ctx *ProgramContext) {}

// EnterTuple is called when production tuple is entered.
func (s *BasedslListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production tuple is exited.
func (s *BasedslListener) ExitTuple(ctx *TupleContext) {}

// EnterTuplelist is called when production tuplelist is entered.
func (s *BasedslListener) EnterTuplelist(ctx *TuplelistContext) {}

// ExitTuplelist is called when production tuplelist is exited.
func (s *BasedslListener) ExitTuplelist(ctx *TuplelistContext) {}

// EnterDecleration is called when production decleration is entered.
func (s *BasedslListener) EnterDecleration(ctx *DeclerationContext) {}

// ExitDecleration is called when production decleration is exited.
func (s *BasedslListener) ExitDecleration(ctx *DeclerationContext) {}

// EnterActionDeclaration is called when production actionDeclaration is entered.
func (s *BasedslListener) EnterActionDeclaration(ctx *ActionDeclarationContext) {}

// ExitActionDeclaration is called when production actionDeclaration is exited.
func (s *BasedslListener) ExitActionDeclaration(ctx *ActionDeclarationContext) {}

// EnterMonitorDeclaration is called when production monitorDeclaration is entered.
func (s *BasedslListener) EnterMonitorDeclaration(ctx *MonitorDeclarationContext) {}

// ExitMonitorDeclaration is called when production monitorDeclaration is exited.
func (s *BasedslListener) ExitMonitorDeclaration(ctx *MonitorDeclarationContext) {}

// EnterMessageDeclaration is called when production messageDeclaration is entered.
func (s *BasedslListener) EnterMessageDeclaration(ctx *MessageDeclarationContext) {}

// ExitMessageDeclaration is called when production messageDeclaration is exited.
func (s *BasedslListener) ExitMessageDeclaration(ctx *MessageDeclarationContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BasedslListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BasedslListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterBuiltins is called when production builtins is entered.
func (s *BasedslListener) EnterBuiltins(ctx *BuiltinsContext) {}

// ExitBuiltins is called when production builtins is exited.
func (s *BasedslListener) ExitBuiltins(ctx *BuiltinsContext) {}

// EnterFunc is called when production func is entered.
func (s *BasedslListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BasedslListener) ExitFunc(ctx *FuncContext) {}

// EnterArgs is called when production args is entered.
func (s *BasedslListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BasedslListener) ExitArgs(ctx *ArgsContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BasedslListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BasedslListener) ExitStatementList(ctx *StatementListContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BasedslListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BasedslListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterSendExpression is called when production sendExpression is entered.
func (s *BasedslListener) EnterSendExpression(ctx *SendExpressionContext) {}

// ExitSendExpression is called when production sendExpression is exited.
func (s *BasedslListener) ExitSendExpression(ctx *SendExpressionContext) {}

// EnterRecvExpression is called when production recvExpression is entered.
func (s *BasedslListener) EnterRecvExpression(ctx *RecvExpressionContext) {}

// ExitRecvExpression is called when production recvExpression is exited.
func (s *BasedslListener) ExitRecvExpression(ctx *RecvExpressionContext) {}

// EnterBarrier is called when production barrier is entered.
func (s *BasedslListener) EnterBarrier(ctx *BarrierContext) {}

// ExitBarrier is called when production barrier is exited.
func (s *BasedslListener) ExitBarrier(ctx *BarrierContext) {}

// EnterMessage is called when production message is entered.
func (s *BasedslListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BasedslListener) ExitMessage(ctx *MessageContext) {}

// EnterType is called when production type is entered.
func (s *BasedslListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasedslListener) ExitType(ctx *TypeContext) {}

// EnterPrimitive is called when production primitive is entered.
func (s *BasedslListener) EnterPrimitive(ctx *PrimitiveContext) {}

// ExitPrimitive is called when production primitive is exited.
func (s *BasedslListener) ExitPrimitive(ctx *PrimitiveContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasedslListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasedslListener) ExitVariable(ctx *VariableContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasedslListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasedslListener) ExitIdentifier(ctx *IdentifierContext) {}
