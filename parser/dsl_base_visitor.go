// Code generated from dsl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // dsl

import "github.com/antlr4-go/antlr/v4"

type BasedslVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasedslVisitor) VisitTypesDeclaration(ctx *TypesDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitTuplelist(ctx *TuplelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitDecleration(ctx *DeclerationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitActionDeclaration(ctx *ActionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitMonitorDeclaration(ctx *MonitorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitMessageDeclaration(ctx *MessageDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitVarDeclaration(ctx *VarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitBuiltins(ctx *BuiltinsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitFunc(ctx *FuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitSendExpression(ctx *SendExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitRecvExpression(ctx *RecvExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitBarrier(ctx *BarrierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitMessage(ctx *MessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitPrimitive(ctx *PrimitiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
