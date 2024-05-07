// Generated from /home/alborz/pontoon/dsl.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class dslParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, Wildcard=20, Assign=21, StringLiteral=22, DigitSeq=23, 
		Comment=24, Whitespace=25, Newline=26;
	public static final int
		RULE_typesDeclaration = 0, RULE_program = 1, RULE_tuple = 2, RULE_tuplelist = 3, 
		RULE_decleration = 4, RULE_actionDeclaration = 5, RULE_monitorDeclaration = 6, 
		RULE_messageDeclaration = 7, RULE_varDeclaration = 8, RULE_builtins = 9, 
		RULE_func = 10, RULE_args = 11, RULE_statementList = 12, RULE_assignmentExpression = 13, 
		RULE_sendExpression = 14, RULE_recvExpression = 15, RULE_barrier = 16, 
		RULE_message = 17, RULE_type = 18, RULE_primitive = 19, RULE_variable = 20, 
		RULE_identifier = 21;
	private static String[] makeRuleNames() {
		return new String[] {
			"typesDeclaration", "program", "tuple", "tuplelist", "decleration", "actionDeclaration", 
			"monitorDeclaration", "messageDeclaration", "varDeclaration", "builtins", 
			"func", "args", "statementList", "assignmentExpression", "sendExpression", 
			"recvExpression", "barrier", "message", "type", "primitive", "variable", 
			"identifier"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'types:'", "','", "'('", "')'", "'action'", "'{'", "'}'", "'monitor'", 
			"':'", "'var'", "'rand'", "'tag'", "'store'", "'type'", "'=>'", "'<='", 
			"'.'", "'\"'", "'int'", null, "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, "Wildcard", "Assign", 
			"StringLiteral", "DigitSeq", "Comment", "Whitespace", "Newline"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "dsl.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public dslParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypesDeclarationContext extends ParserRuleContext {
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public TypesDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typesDeclaration; }
	}

	public final TypesDeclarationContext typesDeclaration() throws RecognitionException {
		TypesDeclarationContext _localctx = new TypesDeclarationContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_typesDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			match(T__0);
			setState(45);
			type();
			setState(50);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__1) {
				{
				{
				setState(46);
				match(T__1);
				setState(47);
				type();
				}
				}
				setState(52);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<DeclerationContext> decleration() {
			return getRuleContexts(DeclerationContext.class);
		}
		public DeclerationContext decleration(int i) {
			return getRuleContext(DeclerationContext.class,i);
		}
		public TerminalNode EOF() { return getToken(dslParser.EOF, 0); }
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(53);
			decleration();
			setState(57);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4194594L) != 0)) {
				{
				{
				setState(54);
				decleration();
				}
				}
				setState(59);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(60);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TupleContext extends ParserRuleContext {
		public TuplelistContext tuplelist() {
			return getRuleContext(TuplelistContext.class,0);
		}
		public TupleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple; }
	}

	public final TupleContext tuple() throws RecognitionException {
		TupleContext _localctx = new TupleContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_tuple);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			match(T__2);
			setState(63);
			tuplelist();
			setState(64);
			match(T__3);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TuplelistContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public TuplelistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuplelist; }
	}

	public final TuplelistContext tuplelist() throws RecognitionException {
		TuplelistContext _localctx = new TuplelistContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_tuplelist);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(69); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(66);
				identifier();
				setState(67);
				match(T__1);
				}
				}
				setState(71); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==StringLiteral );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclerationContext extends ParserRuleContext {
		public ActionDeclarationContext actionDeclaration() {
			return getRuleContext(ActionDeclarationContext.class,0);
		}
		public TypesDeclarationContext typesDeclaration() {
			return getRuleContext(TypesDeclarationContext.class,0);
		}
		public MonitorDeclarationContext monitorDeclaration() {
			return getRuleContext(MonitorDeclarationContext.class,0);
		}
		public MessageDeclarationContext messageDeclaration() {
			return getRuleContext(MessageDeclarationContext.class,0);
		}
		public DeclerationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decleration; }
	}

	public final DeclerationContext decleration() throws RecognitionException {
		DeclerationContext _localctx = new DeclerationContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_decleration);
		try {
			setState(77);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__4:
				enterOuterAlt(_localctx, 1);
				{
				setState(73);
				actionDeclaration();
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 2);
				{
				setState(74);
				typesDeclaration();
				}
				break;
			case T__7:
				enterOuterAlt(_localctx, 3);
				{
				setState(75);
				monitorDeclaration();
				}
				break;
			case StringLiteral:
				enterOuterAlt(_localctx, 4);
				{
				setState(76);
				messageDeclaration();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ActionDeclarationContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public ActionDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_actionDeclaration; }
	}

	public final ActionDeclarationContext actionDeclaration() throws RecognitionException {
		ActionDeclarationContext _localctx = new ActionDeclarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_actionDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(79);
			match(T__4);
			setState(80);
			type();
			setState(81);
			identifier();
			setState(82);
			match(T__5);
			setState(83);
			statementList();
			setState(84);
			match(T__6);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MonitorDeclarationContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public MonitorDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_monitorDeclaration; }
	}

	public final MonitorDeclarationContext monitorDeclaration() throws RecognitionException {
		MonitorDeclarationContext _localctx = new MonitorDeclarationContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_monitorDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			match(T__7);
			setState(87);
			type();
			setState(88);
			identifier();
			setState(89);
			match(T__5);
			setState(90);
			statementList();
			setState(91);
			match(T__6);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MessageDeclarationContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public MessageContext message() {
			return getRuleContext(MessageContext.class,0);
		}
		public MessageDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_messageDeclaration; }
	}

	public final MessageDeclarationContext messageDeclaration() throws RecognitionException {
		MessageDeclarationContext _localctx = new MessageDeclarationContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_messageDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(93);
			identifier();
			setState(94);
			match(T__8);
			setState(95);
			message();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarDeclarationContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode Assign() { return getToken(dslParser.Assign, 0); }
		public TupleContext tuple() {
			return getRuleContext(TupleContext.class,0);
		}
		public TerminalNode DigitSeq() { return getToken(dslParser.DigitSeq, 0); }
		public FuncContext func() {
			return getRuleContext(FuncContext.class,0);
		}
		public VarDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDeclaration; }
	}

	public final VarDeclarationContext varDeclaration() throws RecognitionException {
		VarDeclarationContext _localctx = new VarDeclarationContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_varDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(97);
			match(T__9);
			setState(98);
			identifier();
			setState(99);
			match(Assign);
			setState(103);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
				{
				setState(100);
				tuple();
				}
				break;
			case DigitSeq:
				{
				setState(101);
				match(DigitSeq);
				}
				break;
			case T__10:
			case T__11:
			case T__12:
				{
				setState(102);
				func();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BuiltinsContext extends ParserRuleContext {
		public BuiltinsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_builtins; }
	}

	public final BuiltinsContext builtins() throws RecognitionException {
		BuiltinsContext _localctx = new BuiltinsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_builtins);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 14336L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncContext extends ParserRuleContext {
		public BuiltinsContext builtins() {
			return getRuleContext(BuiltinsContext.class,0);
		}
		public ArgsContext args() {
			return getRuleContext(ArgsContext.class,0);
		}
		public FuncContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func; }
	}

	public final FuncContext func() throws RecognitionException {
		FuncContext _localctx = new FuncContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_func);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			builtins();
			setState(108);
			match(T__2);
			setState(109);
			args();
			setState(110);
			match(T__3);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArgsContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ArgsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_args; }
	}

	public final ArgsContext args() throws RecognitionException {
		ArgsContext _localctx = new ArgsContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_args);
		try {
			setState(115);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__13:
				enterOuterAlt(_localctx, 1);
				{
				setState(112);
				match(T__13);
				setState(113);
				type();
				}
				break;
			case StringLiteral:
				enterOuterAlt(_localctx, 2);
				{
				setState(114);
				identifier();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementListContext extends ParserRuleContext {
		public List<AssignmentExpressionContext> assignmentExpression() {
			return getRuleContexts(AssignmentExpressionContext.class);
		}
		public AssignmentExpressionContext assignmentExpression(int i) {
			return getRuleContext(AssignmentExpressionContext.class,i);
		}
		public List<SendExpressionContext> sendExpression() {
			return getRuleContexts(SendExpressionContext.class);
		}
		public SendExpressionContext sendExpression(int i) {
			return getRuleContext(SendExpressionContext.class,i);
		}
		public List<RecvExpressionContext> recvExpression() {
			return getRuleContexts(RecvExpressionContext.class);
		}
		public RecvExpressionContext recvExpression(int i) {
			return getRuleContext(RecvExpressionContext.class,i);
		}
		public List<VarDeclarationContext> varDeclaration() {
			return getRuleContexts(VarDeclarationContext.class);
		}
		public VarDeclarationContext varDeclaration(int i) {
			return getRuleContext(VarDeclarationContext.class,i);
		}
		public List<BarrierContext> barrier() {
			return getRuleContexts(BarrierContext.class);
		}
		public BarrierContext barrier(int i) {
			return getRuleContext(BarrierContext.class,i);
		}
		public StatementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statementList; }
	}

	public final StatementListContext statementList() throws RecognitionException {
		StatementListContext _localctx = new StatementListContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_statementList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(122);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
				case 1:
					{
					setState(117);
					assignmentExpression();
					}
					break;
				case 2:
					{
					setState(118);
					sendExpression();
					}
					break;
				case 3:
					{
					setState(119);
					recvExpression();
					}
					break;
				case 4:
					{
					setState(120);
					varDeclaration();
					}
					break;
				case 5:
					{
					setState(121);
					barrier();
					}
					break;
				}
				}
				}
				setState(126); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 5112832L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentExpressionContext extends ParserRuleContext {
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public TerminalNode Assign() { return getToken(dslParser.Assign, 0); }
		public TerminalNode Wildcard() { return getToken(dslParser.Wildcard, 0); }
		public AssignmentExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentExpression; }
	}

	public final AssignmentExpressionContext assignmentExpression() throws RecognitionException {
		AssignmentExpressionContext _localctx = new AssignmentExpressionContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_assignmentExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			variable();
			setState(129);
			match(Assign);
			setState(130);
			match(Wildcard);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SendExpressionContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public MessageContext message() {
			return getRuleContext(MessageContext.class,0);
		}
		public SendExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sendExpression; }
	}

	public final SendExpressionContext sendExpression() throws RecognitionException {
		SendExpressionContext _localctx = new SendExpressionContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_sendExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__17:
				{
				setState(132);
				message();
				}
				break;
			case StringLiteral:
				{
				setState(133);
				identifier();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(136);
			match(T__14);
			setState(137);
			identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RecvExpressionContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public MessageContext message() {
			return getRuleContext(MessageContext.class,0);
		}
		public RecvExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_recvExpression; }
	}

	public final RecvExpressionContext recvExpression() throws RecognitionException {
		RecvExpressionContext _localctx = new RecvExpressionContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_recvExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__17:
				{
				setState(139);
				message();
				}
				break;
			case StringLiteral:
				{
				setState(140);
				identifier();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(143);
			match(T__15);
			setState(144);
			identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BarrierContext extends ParserRuleContext {
		public BarrierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_barrier; }
	}

	public final BarrierContext barrier() throws RecognitionException {
		BarrierContext _localctx = new BarrierContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_barrier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(146);
			match(T__16);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MessageContext extends ParserRuleContext {
		public TerminalNode StringLiteral() { return getToken(dslParser.StringLiteral, 0); }
		public MessageContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_message; }
	}

	public final MessageContext message() throws RecognitionException {
		MessageContext _localctx = new MessageContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_message);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			match(T__17);
			setState(149);
			match(StringLiteral);
			setState(150);
			match(T__17);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeContext extends ParserRuleContext {
		public TerminalNode StringLiteral() { return getToken(dslParser.StringLiteral, 0); }
		public PrimitiveContext primitive() {
			return getRuleContext(PrimitiveContext.class,0);
		}
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_type);
		try {
			setState(154);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case StringLiteral:
				enterOuterAlt(_localctx, 1);
				{
				setState(152);
				match(StringLiteral);
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 2);
				{
				setState(153);
				primitive();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrimitiveContext extends ParserRuleContext {
		public PrimitiveContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitive; }
	}

	public final PrimitiveContext primitive() throws RecognitionException {
		PrimitiveContext _localctx = new PrimitiveContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_primitive);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(156);
			match(T__18);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VariableContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public VariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable; }
	}

	public final VariableContext variable() throws RecognitionException {
		VariableContext _localctx = new VariableContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_variable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(158);
			type();
			setState(159);
			identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ParserRuleContext {
		public TerminalNode StringLiteral() { return getToken(dslParser.StringLiteral, 0); }
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_identifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			match(StringLiteral);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u001a\u00a4\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001"+
		"\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004"+
		"\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007"+
		"\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b"+
		"\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007"+
		"\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007"+
		"\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007"+
		"\u0015\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0005\u00001\b"+
		"\u0000\n\u0000\f\u00004\t\u0000\u0001\u0001\u0001\u0001\u0005\u00018\b"+
		"\u0001\n\u0001\f\u0001;\t\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0004"+
		"\u0003F\b\u0003\u000b\u0003\f\u0003G\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0003\u0004N\b\u0004\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0003\bh\b\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000bt\b\u000b\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f{\b\f\u0004\f}\b\f\u000b\f\f"+
		"\f~\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0003\u000e"+
		"\u0087\b\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f"+
		"\u0003\u000f\u008e\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010"+
		"\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012"+
		"\u0001\u0012\u0003\u0012\u009b\b\u0012\u0001\u0013\u0001\u0013\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0000\u0000"+
		"\u0016\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018"+
		"\u001a\u001c\u001e \"$&(*\u0000\u0001\u0001\u0000\u000b\r\u009e\u0000"+
		",\u0001\u0000\u0000\u0000\u00025\u0001\u0000\u0000\u0000\u0004>\u0001"+
		"\u0000\u0000\u0000\u0006E\u0001\u0000\u0000\u0000\bM\u0001\u0000\u0000"+
		"\u0000\nO\u0001\u0000\u0000\u0000\fV\u0001\u0000\u0000\u0000\u000e]\u0001"+
		"\u0000\u0000\u0000\u0010a\u0001\u0000\u0000\u0000\u0012i\u0001\u0000\u0000"+
		"\u0000\u0014k\u0001\u0000\u0000\u0000\u0016s\u0001\u0000\u0000\u0000\u0018"+
		"|\u0001\u0000\u0000\u0000\u001a\u0080\u0001\u0000\u0000\u0000\u001c\u0086"+
		"\u0001\u0000\u0000\u0000\u001e\u008d\u0001\u0000\u0000\u0000 \u0092\u0001"+
		"\u0000\u0000\u0000\"\u0094\u0001\u0000\u0000\u0000$\u009a\u0001\u0000"+
		"\u0000\u0000&\u009c\u0001\u0000\u0000\u0000(\u009e\u0001\u0000\u0000\u0000"+
		"*\u00a1\u0001\u0000\u0000\u0000,-\u0005\u0001\u0000\u0000-2\u0003$\u0012"+
		"\u0000./\u0005\u0002\u0000\u0000/1\u0003$\u0012\u00000.\u0001\u0000\u0000"+
		"\u000014\u0001\u0000\u0000\u000020\u0001\u0000\u0000\u000023\u0001\u0000"+
		"\u0000\u00003\u0001\u0001\u0000\u0000\u000042\u0001\u0000\u0000\u0000"+
		"59\u0003\b\u0004\u000068\u0003\b\u0004\u000076\u0001\u0000\u0000\u0000"+
		"8;\u0001\u0000\u0000\u000097\u0001\u0000\u0000\u00009:\u0001\u0000\u0000"+
		"\u0000:<\u0001\u0000\u0000\u0000;9\u0001\u0000\u0000\u0000<=\u0005\u0000"+
		"\u0000\u0001=\u0003\u0001\u0000\u0000\u0000>?\u0005\u0003\u0000\u0000"+
		"?@\u0003\u0006\u0003\u0000@A\u0005\u0004\u0000\u0000A\u0005\u0001\u0000"+
		"\u0000\u0000BC\u0003*\u0015\u0000CD\u0005\u0002\u0000\u0000DF\u0001\u0000"+
		"\u0000\u0000EB\u0001\u0000\u0000\u0000FG\u0001\u0000\u0000\u0000GE\u0001"+
		"\u0000\u0000\u0000GH\u0001\u0000\u0000\u0000H\u0007\u0001\u0000\u0000"+
		"\u0000IN\u0003\n\u0005\u0000JN\u0003\u0000\u0000\u0000KN\u0003\f\u0006"+
		"\u0000LN\u0003\u000e\u0007\u0000MI\u0001\u0000\u0000\u0000MJ\u0001\u0000"+
		"\u0000\u0000MK\u0001\u0000\u0000\u0000ML\u0001\u0000\u0000\u0000N\t\u0001"+
		"\u0000\u0000\u0000OP\u0005\u0005\u0000\u0000PQ\u0003$\u0012\u0000QR\u0003"+
		"*\u0015\u0000RS\u0005\u0006\u0000\u0000ST\u0003\u0018\f\u0000TU\u0005"+
		"\u0007\u0000\u0000U\u000b\u0001\u0000\u0000\u0000VW\u0005\b\u0000\u0000"+
		"WX\u0003$\u0012\u0000XY\u0003*\u0015\u0000YZ\u0005\u0006\u0000\u0000Z"+
		"[\u0003\u0018\f\u0000[\\\u0005\u0007\u0000\u0000\\\r\u0001\u0000\u0000"+
		"\u0000]^\u0003*\u0015\u0000^_\u0005\t\u0000\u0000_`\u0003\"\u0011\u0000"+
		"`\u000f\u0001\u0000\u0000\u0000ab\u0005\n\u0000\u0000bc\u0003*\u0015\u0000"+
		"cg\u0005\u0015\u0000\u0000dh\u0003\u0004\u0002\u0000eh\u0005\u0017\u0000"+
		"\u0000fh\u0003\u0014\n\u0000gd\u0001\u0000\u0000\u0000ge\u0001\u0000\u0000"+
		"\u0000gf\u0001\u0000\u0000\u0000h\u0011\u0001\u0000\u0000\u0000ij\u0007"+
		"\u0000\u0000\u0000j\u0013\u0001\u0000\u0000\u0000kl\u0003\u0012\t\u0000"+
		"lm\u0005\u0003\u0000\u0000mn\u0003\u0016\u000b\u0000no\u0005\u0004\u0000"+
		"\u0000o\u0015\u0001\u0000\u0000\u0000pq\u0005\u000e\u0000\u0000qt\u0003"+
		"$\u0012\u0000rt\u0003*\u0015\u0000sp\u0001\u0000\u0000\u0000sr\u0001\u0000"+
		"\u0000\u0000t\u0017\u0001\u0000\u0000\u0000u{\u0003\u001a\r\u0000v{\u0003"+
		"\u001c\u000e\u0000w{\u0003\u001e\u000f\u0000x{\u0003\u0010\b\u0000y{\u0003"+
		" \u0010\u0000zu\u0001\u0000\u0000\u0000zv\u0001\u0000\u0000\u0000zw\u0001"+
		"\u0000\u0000\u0000zx\u0001\u0000\u0000\u0000zy\u0001\u0000\u0000\u0000"+
		"{}\u0001\u0000\u0000\u0000|z\u0001\u0000\u0000\u0000}~\u0001\u0000\u0000"+
		"\u0000~|\u0001\u0000\u0000\u0000~\u007f\u0001\u0000\u0000\u0000\u007f"+
		"\u0019\u0001\u0000\u0000\u0000\u0080\u0081\u0003(\u0014\u0000\u0081\u0082"+
		"\u0005\u0015\u0000\u0000\u0082\u0083\u0005\u0014\u0000\u0000\u0083\u001b"+
		"\u0001\u0000\u0000\u0000\u0084\u0087\u0003\"\u0011\u0000\u0085\u0087\u0003"+
		"*\u0015\u0000\u0086\u0084\u0001\u0000\u0000\u0000\u0086\u0085\u0001\u0000"+
		"\u0000\u0000\u0087\u0088\u0001\u0000\u0000\u0000\u0088\u0089\u0005\u000f"+
		"\u0000\u0000\u0089\u008a\u0003*\u0015\u0000\u008a\u001d\u0001\u0000\u0000"+
		"\u0000\u008b\u008e\u0003\"\u0011\u0000\u008c\u008e\u0003*\u0015\u0000"+
		"\u008d\u008b\u0001\u0000\u0000\u0000\u008d\u008c\u0001\u0000\u0000\u0000"+
		"\u008e\u008f\u0001\u0000\u0000\u0000\u008f\u0090\u0005\u0010\u0000\u0000"+
		"\u0090\u0091\u0003*\u0015\u0000\u0091\u001f\u0001\u0000\u0000\u0000\u0092"+
		"\u0093\u0005\u0011\u0000\u0000\u0093!\u0001\u0000\u0000\u0000\u0094\u0095"+
		"\u0005\u0012\u0000\u0000\u0095\u0096\u0005\u0016\u0000\u0000\u0096\u0097"+
		"\u0005\u0012\u0000\u0000\u0097#\u0001\u0000\u0000\u0000\u0098\u009b\u0005"+
		"\u0016\u0000\u0000\u0099\u009b\u0003&\u0013\u0000\u009a\u0098\u0001\u0000"+
		"\u0000\u0000\u009a\u0099\u0001\u0000\u0000\u0000\u009b%\u0001\u0000\u0000"+
		"\u0000\u009c\u009d\u0005\u0013\u0000\u0000\u009d\'\u0001\u0000\u0000\u0000"+
		"\u009e\u009f\u0003$\u0012\u0000\u009f\u00a0\u0003*\u0015\u0000\u00a0)"+
		"\u0001\u0000\u0000\u0000\u00a1\u00a2\u0005\u0016\u0000\u0000\u00a2+\u0001"+
		"\u0000\u0000\u0000\u000b29GMgsz~\u0086\u008d\u009a";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}