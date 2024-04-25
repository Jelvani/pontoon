// Code generated from dsl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // dsl

import "github.com/antlr4-go/antlr/v4"

// dslListener is a complete listener for a parse tree produced by dslParser.
type dslListener interface {
	antlr.ParseTreeListener

	// EnterTypesDeclaration is called when entering the typesDeclaration production.
	EnterTypesDeclaration(c *TypesDeclarationContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTuple is called when entering the tuple production.
	EnterTuple(c *TupleContext)

	// EnterTuplelist is called when entering the tuplelist production.
	EnterTuplelist(c *TuplelistContext)

	// EnterDecleration is called when entering the decleration production.
	EnterDecleration(c *DeclerationContext)

	// EnterActionDeclaration is called when entering the actionDeclaration production.
	EnterActionDeclaration(c *ActionDeclarationContext)

	// EnterMonitorDeclaration is called when entering the monitorDeclaration production.
	EnterMonitorDeclaration(c *MonitorDeclarationContext)

	// EnterMessageDeclaration is called when entering the messageDeclaration production.
	EnterMessageDeclaration(c *MessageDeclarationContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterBuiltins is called when entering the builtins production.
	EnterBuiltins(c *BuiltinsContext)

	// EnterFunc is called when entering the func production.
	EnterFunc(c *FuncContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterSendExpression is called when entering the sendExpression production.
	EnterSendExpression(c *SendExpressionContext)

	// EnterRecvExpression is called when entering the recvExpression production.
	EnterRecvExpression(c *RecvExpressionContext)

	// EnterBarrier is called when entering the barrier production.
	EnterBarrier(c *BarrierContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterPrimitive is called when entering the primitive production.
	EnterPrimitive(c *PrimitiveContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitTypesDeclaration is called when exiting the typesDeclaration production.
	ExitTypesDeclaration(c *TypesDeclarationContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTuple is called when exiting the tuple production.
	ExitTuple(c *TupleContext)

	// ExitTuplelist is called when exiting the tuplelist production.
	ExitTuplelist(c *TuplelistContext)

	// ExitDecleration is called when exiting the decleration production.
	ExitDecleration(c *DeclerationContext)

	// ExitActionDeclaration is called when exiting the actionDeclaration production.
	ExitActionDeclaration(c *ActionDeclarationContext)

	// ExitMonitorDeclaration is called when exiting the monitorDeclaration production.
	ExitMonitorDeclaration(c *MonitorDeclarationContext)

	// ExitMessageDeclaration is called when exiting the messageDeclaration production.
	ExitMessageDeclaration(c *MessageDeclarationContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitBuiltins is called when exiting the builtins production.
	ExitBuiltins(c *BuiltinsContext)

	// ExitFunc is called when exiting the func production.
	ExitFunc(c *FuncContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitSendExpression is called when exiting the sendExpression production.
	ExitSendExpression(c *SendExpressionContext)

	// ExitRecvExpression is called when exiting the recvExpression production.
	ExitRecvExpression(c *RecvExpressionContext)

	// ExitBarrier is called when exiting the barrier production.
	ExitBarrier(c *BarrierContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitPrimitive is called when exiting the primitive production.
	ExitPrimitive(c *PrimitiveContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
