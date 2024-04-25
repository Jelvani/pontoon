// Code generated from dsl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // dsl

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by dslParser.
type dslVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by dslParser#typesDeclaration.
	VisitTypesDeclaration(ctx *TypesDeclarationContext) interface{}

	// Visit a parse tree produced by dslParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by dslParser#tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by dslParser#tuplelist.
	VisitTuplelist(ctx *TuplelistContext) interface{}

	// Visit a parse tree produced by dslParser#decleration.
	VisitDecleration(ctx *DeclerationContext) interface{}

	// Visit a parse tree produced by dslParser#actionDeclaration.
	VisitActionDeclaration(ctx *ActionDeclarationContext) interface{}

	// Visit a parse tree produced by dslParser#monitorDeclaration.
	VisitMonitorDeclaration(ctx *MonitorDeclarationContext) interface{}

	// Visit a parse tree produced by dslParser#messageDeclaration.
	VisitMessageDeclaration(ctx *MessageDeclarationContext) interface{}

	// Visit a parse tree produced by dslParser#varDeclaration.
	VisitVarDeclaration(ctx *VarDeclarationContext) interface{}

	// Visit a parse tree produced by dslParser#builtins.
	VisitBuiltins(ctx *BuiltinsContext) interface{}

	// Visit a parse tree produced by dslParser#func.
	VisitFunc(ctx *FuncContext) interface{}

	// Visit a parse tree produced by dslParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by dslParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by dslParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by dslParser#sendExpression.
	VisitSendExpression(ctx *SendExpressionContext) interface{}

	// Visit a parse tree produced by dslParser#recvExpression.
	VisitRecvExpression(ctx *RecvExpressionContext) interface{}

	// Visit a parse tree produced by dslParser#barrier.
	VisitBarrier(ctx *BarrierContext) interface{}

	// Visit a parse tree produced by dslParser#message.
	VisitMessage(ctx *MessageContext) interface{}

	// Visit a parse tree produced by dslParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by dslParser#primitive.
	VisitPrimitive(ctx *PrimitiveContext) interface{}

	// Visit a parse tree produced by dslParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by dslParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
