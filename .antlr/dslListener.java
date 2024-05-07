// Generated from /home/alborz/pontoon/dsl.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link dslParser}.
 */
public interface dslListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link dslParser#typesDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterTypesDeclaration(dslParser.TypesDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#typesDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitTypesDeclaration(dslParser.TypesDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(dslParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(dslParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#tuple}.
	 * @param ctx the parse tree
	 */
	void enterTuple(dslParser.TupleContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#tuple}.
	 * @param ctx the parse tree
	 */
	void exitTuple(dslParser.TupleContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#tuplelist}.
	 * @param ctx the parse tree
	 */
	void enterTuplelist(dslParser.TuplelistContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#tuplelist}.
	 * @param ctx the parse tree
	 */
	void exitTuplelist(dslParser.TuplelistContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#decleration}.
	 * @param ctx the parse tree
	 */
	void enterDecleration(dslParser.DeclerationContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#decleration}.
	 * @param ctx the parse tree
	 */
	void exitDecleration(dslParser.DeclerationContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#actionDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterActionDeclaration(dslParser.ActionDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#actionDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitActionDeclaration(dslParser.ActionDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#monitorDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterMonitorDeclaration(dslParser.MonitorDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#monitorDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitMonitorDeclaration(dslParser.MonitorDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#messageDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterMessageDeclaration(dslParser.MessageDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#messageDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitMessageDeclaration(dslParser.MessageDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclaration(dslParser.VarDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclaration(dslParser.VarDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#builtins}.
	 * @param ctx the parse tree
	 */
	void enterBuiltins(dslParser.BuiltinsContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#builtins}.
	 * @param ctx the parse tree
	 */
	void exitBuiltins(dslParser.BuiltinsContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#func}.
	 * @param ctx the parse tree
	 */
	void enterFunc(dslParser.FuncContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#func}.
	 * @param ctx the parse tree
	 */
	void exitFunc(dslParser.FuncContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#args}.
	 * @param ctx the parse tree
	 */
	void enterArgs(dslParser.ArgsContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#args}.
	 * @param ctx the parse tree
	 */
	void exitArgs(dslParser.ArgsContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#statementList}.
	 * @param ctx the parse tree
	 */
	void enterStatementList(dslParser.StatementListContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#statementList}.
	 * @param ctx the parse tree
	 */
	void exitStatementList(dslParser.StatementListContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#assignmentExpression}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentExpression(dslParser.AssignmentExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#assignmentExpression}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentExpression(dslParser.AssignmentExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#sendExpression}.
	 * @param ctx the parse tree
	 */
	void enterSendExpression(dslParser.SendExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#sendExpression}.
	 * @param ctx the parse tree
	 */
	void exitSendExpression(dslParser.SendExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#recvExpression}.
	 * @param ctx the parse tree
	 */
	void enterRecvExpression(dslParser.RecvExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#recvExpression}.
	 * @param ctx the parse tree
	 */
	void exitRecvExpression(dslParser.RecvExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#barrier}.
	 * @param ctx the parse tree
	 */
	void enterBarrier(dslParser.BarrierContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#barrier}.
	 * @param ctx the parse tree
	 */
	void exitBarrier(dslParser.BarrierContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#message}.
	 * @param ctx the parse tree
	 */
	void enterMessage(dslParser.MessageContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#message}.
	 * @param ctx the parse tree
	 */
	void exitMessage(dslParser.MessageContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#type}.
	 * @param ctx the parse tree
	 */
	void enterType(dslParser.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#type}.
	 * @param ctx the parse tree
	 */
	void exitType(dslParser.TypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#primitive}.
	 * @param ctx the parse tree
	 */
	void enterPrimitive(dslParser.PrimitiveContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#primitive}.
	 * @param ctx the parse tree
	 */
	void exitPrimitive(dslParser.PrimitiveContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#variable}.
	 * @param ctx the parse tree
	 */
	void enterVariable(dslParser.VariableContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#variable}.
	 * @param ctx the parse tree
	 */
	void exitVariable(dslParser.VariableContext ctx);
	/**
	 * Enter a parse tree produced by {@link dslParser#identifier}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier(dslParser.IdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link dslParser#identifier}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier(dslParser.IdentifierContext ctx);
}