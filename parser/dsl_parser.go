// Code generated from dsl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // dsl

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type dslParser struct {
	*antlr.BaseParser
}

var DslParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dslParserInit() {
	staticData := &DslParserStaticData
	staticData.LiteralNames = []string{
		"", "'types:'", "'('", "')'", "'action'", "'{'", "'}'", "'monitor'",
		"':'", "'var'", "'rand'", "'tag'", "'store'", "'type'", "'=>'", "'<='",
		"'.'", "'\"'", "'int'", "", "'='", "", "", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "Wildcard", "Assign", "StringLiteral", "DigitSeq", "Comma",
		"Comment", "Whitespace", "Newline",
	}
	staticData.RuleNames = []string{
		"typesDeclaration", "program", "tuple", "tuplelist", "decleration",
		"actionDeclaration", "monitorDeclaration", "messageDeclaration", "varDeclaration",
		"builtins", "func", "args", "statementList", "assignmentExpression",
		"sendExpression", "recvExpression", "barrier", "message", "type", "primitive",
		"variable", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 164, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 49, 8, 0, 10, 0, 12, 0, 52, 9,
		0, 1, 1, 1, 1, 5, 1, 56, 8, 1, 10, 1, 12, 1, 59, 9, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 4, 3, 70, 8, 3, 11, 3, 12, 3, 71, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 78, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 104, 8, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3, 11, 116, 8, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 123, 8, 12, 4, 12, 125, 8, 12, 11,
		12, 12, 12, 126, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 135,
		8, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15, 142, 8, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 3, 18,
		155, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 0,
		0, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 0, 1, 1, 0, 10, 12, 158, 0, 44, 1, 0, 0, 0, 2, 53, 1, 0,
		0, 0, 4, 62, 1, 0, 0, 0, 6, 69, 1, 0, 0, 0, 8, 77, 1, 0, 0, 0, 10, 79,
		1, 0, 0, 0, 12, 86, 1, 0, 0, 0, 14, 93, 1, 0, 0, 0, 16, 97, 1, 0, 0, 0,
		18, 105, 1, 0, 0, 0, 20, 107, 1, 0, 0, 0, 22, 115, 1, 0, 0, 0, 24, 124,
		1, 0, 0, 0, 26, 128, 1, 0, 0, 0, 28, 134, 1, 0, 0, 0, 30, 141, 1, 0, 0,
		0, 32, 146, 1, 0, 0, 0, 34, 148, 1, 0, 0, 0, 36, 154, 1, 0, 0, 0, 38, 156,
		1, 0, 0, 0, 40, 158, 1, 0, 0, 0, 42, 161, 1, 0, 0, 0, 44, 45, 5, 1, 0,
		0, 45, 50, 3, 36, 18, 0, 46, 47, 5, 23, 0, 0, 47, 49, 3, 36, 18, 0, 48,
		46, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0,
		0, 51, 1, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 57, 3, 8, 4, 0, 54, 56, 3,
		8, 4, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 0, 0,
		1, 61, 3, 1, 0, 0, 0, 62, 63, 5, 2, 0, 0, 63, 64, 3, 6, 3, 0, 64, 65, 5,
		3, 0, 0, 65, 5, 1, 0, 0, 0, 66, 67, 3, 42, 21, 0, 67, 68, 5, 23, 0, 0,
		68, 70, 1, 0, 0, 0, 69, 66, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69, 1,
		0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 7, 1, 0, 0, 0, 73, 78, 3, 10, 5, 0, 74,
		78, 3, 0, 0, 0, 75, 78, 3, 12, 6, 0, 76, 78, 3, 14, 7, 0, 77, 73, 1, 0,
		0, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 9,
		1, 0, 0, 0, 79, 80, 5, 4, 0, 0, 80, 81, 3, 36, 18, 0, 81, 82, 3, 42, 21,
		0, 82, 83, 5, 5, 0, 0, 83, 84, 3, 24, 12, 0, 84, 85, 5, 6, 0, 0, 85, 11,
		1, 0, 0, 0, 86, 87, 5, 7, 0, 0, 87, 88, 3, 36, 18, 0, 88, 89, 3, 42, 21,
		0, 89, 90, 5, 5, 0, 0, 90, 91, 3, 24, 12, 0, 91, 92, 5, 6, 0, 0, 92, 13,
		1, 0, 0, 0, 93, 94, 3, 42, 21, 0, 94, 95, 5, 8, 0, 0, 95, 96, 3, 34, 17,
		0, 96, 15, 1, 0, 0, 0, 97, 98, 5, 9, 0, 0, 98, 99, 3, 42, 21, 0, 99, 103,
		5, 20, 0, 0, 100, 104, 3, 4, 2, 0, 101, 104, 5, 22, 0, 0, 102, 104, 3,
		20, 10, 0, 103, 100, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 102, 1, 0,
		0, 0, 104, 17, 1, 0, 0, 0, 105, 106, 7, 0, 0, 0, 106, 19, 1, 0, 0, 0, 107,
		108, 3, 18, 9, 0, 108, 109, 5, 2, 0, 0, 109, 110, 3, 22, 11, 0, 110, 111,
		5, 3, 0, 0, 111, 21, 1, 0, 0, 0, 112, 113, 5, 13, 0, 0, 113, 116, 3, 36,
		18, 0, 114, 116, 3, 42, 21, 0, 115, 112, 1, 0, 0, 0, 115, 114, 1, 0, 0,
		0, 116, 23, 1, 0, 0, 0, 117, 123, 3, 26, 13, 0, 118, 123, 3, 28, 14, 0,
		119, 123, 3, 30, 15, 0, 120, 123, 3, 16, 8, 0, 121, 123, 3, 32, 16, 0,
		122, 117, 1, 0, 0, 0, 122, 118, 1, 0, 0, 0, 122, 119, 1, 0, 0, 0, 122,
		120, 1, 0, 0, 0, 122, 121, 1, 0, 0, 0, 123, 125, 1, 0, 0, 0, 124, 122,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0,
		0, 0, 127, 25, 1, 0, 0, 0, 128, 129, 3, 40, 20, 0, 129, 130, 5, 20, 0,
		0, 130, 131, 5, 19, 0, 0, 131, 27, 1, 0, 0, 0, 132, 135, 3, 34, 17, 0,
		133, 135, 3, 42, 21, 0, 134, 132, 1, 0, 0, 0, 134, 133, 1, 0, 0, 0, 135,
		136, 1, 0, 0, 0, 136, 137, 5, 14, 0, 0, 137, 138, 3, 42, 21, 0, 138, 29,
		1, 0, 0, 0, 139, 142, 3, 34, 17, 0, 140, 142, 3, 42, 21, 0, 141, 139, 1,
		0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 5, 15, 0,
		0, 144, 145, 3, 42, 21, 0, 145, 31, 1, 0, 0, 0, 146, 147, 5, 16, 0, 0,
		147, 33, 1, 0, 0, 0, 148, 149, 5, 17, 0, 0, 149, 150, 5, 21, 0, 0, 150,
		151, 5, 17, 0, 0, 151, 35, 1, 0, 0, 0, 152, 155, 5, 21, 0, 0, 153, 155,
		3, 38, 19, 0, 154, 152, 1, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155, 37, 1, 0,
		0, 0, 156, 157, 5, 18, 0, 0, 157, 39, 1, 0, 0, 0, 158, 159, 3, 36, 18,
		0, 159, 160, 3, 42, 21, 0, 160, 41, 1, 0, 0, 0, 161, 162, 5, 21, 0, 0,
		162, 43, 1, 0, 0, 0, 11, 50, 57, 71, 77, 103, 115, 122, 126, 134, 141,
		154,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// dslParserInit initializes any static state used to implement dslParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewdslParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DslParserInit() {
	staticData := &DslParserStaticData
	staticData.once.Do(dslParserInit)
}

// NewdslParser produces a new parser instance for the optional input antlr.TokenStream.
func NewdslParser(input antlr.TokenStream) *dslParser {
	DslParserInit()
	this := new(dslParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DslParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "dsl.g4"

	return this
}

// dslParser tokens.
const (
	dslParserEOF           = antlr.TokenEOF
	dslParserT__0          = 1
	dslParserT__1          = 2
	dslParserT__2          = 3
	dslParserT__3          = 4
	dslParserT__4          = 5
	dslParserT__5          = 6
	dslParserT__6          = 7
	dslParserT__7          = 8
	dslParserT__8          = 9
	dslParserT__9          = 10
	dslParserT__10         = 11
	dslParserT__11         = 12
	dslParserT__12         = 13
	dslParserT__13         = 14
	dslParserT__14         = 15
	dslParserT__15         = 16
	dslParserT__16         = 17
	dslParserT__17         = 18
	dslParserWildcard      = 19
	dslParserAssign        = 20
	dslParserStringLiteral = 21
	dslParserDigitSeq      = 22
	dslParserComma         = 23
	dslParserComment       = 24
	dslParserWhitespace    = 25
	dslParserNewline       = 26
)

// dslParser rules.
const (
	dslParserRULE_typesDeclaration     = 0
	dslParserRULE_program              = 1
	dslParserRULE_tuple                = 2
	dslParserRULE_tuplelist            = 3
	dslParserRULE_decleration          = 4
	dslParserRULE_actionDeclaration    = 5
	dslParserRULE_monitorDeclaration   = 6
	dslParserRULE_messageDeclaration   = 7
	dslParserRULE_varDeclaration       = 8
	dslParserRULE_builtins             = 9
	dslParserRULE_func                 = 10
	dslParserRULE_args                 = 11
	dslParserRULE_statementList        = 12
	dslParserRULE_assignmentExpression = 13
	dslParserRULE_sendExpression       = 14
	dslParserRULE_recvExpression       = 15
	dslParserRULE_barrier              = 16
	dslParserRULE_message              = 17
	dslParserRULE_type                 = 18
	dslParserRULE_primitive            = 19
	dslParserRULE_variable             = 20
	dslParserRULE_identifier           = 21
)

// ITypesDeclarationContext is an interface to support dynamic dispatch.
type ITypesDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsTypesDeclarationContext differentiates from other interfaces.
	IsTypesDeclarationContext()
}

type TypesDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesDeclarationContext() *TypesDeclarationContext {
	var p = new(TypesDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_typesDeclaration
	return p
}

func InitEmptyTypesDeclarationContext(p *TypesDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_typesDeclaration
}

func (*TypesDeclarationContext) IsTypesDeclarationContext() {}

func NewTypesDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesDeclarationContext {
	var p = new(TypesDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_typesDeclaration

	return p
}

func (s *TypesDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesDeclarationContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *TypesDeclarationContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypesDeclarationContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(dslParserComma)
}

func (s *TypesDeclarationContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(dslParserComma, i)
}

func (s *TypesDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterTypesDeclaration(s)
	}
}

func (s *TypesDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitTypesDeclaration(s)
	}
}

func (p *dslParser) TypesDeclaration() (localctx ITypesDeclarationContext) {
	localctx = NewTypesDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, dslParserRULE_typesDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(dslParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Type_()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == dslParserComma {
		{
			p.SetState(46)
			p.Match(dslParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Type_()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDecleration() []IDeclerationContext
	Decleration(i int) IDeclerationContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllDecleration() []IDeclerationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclerationContext); ok {
			len++
		}
	}

	tst := make([]IDeclerationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclerationContext); ok {
			tst[i] = t.(IDeclerationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Decleration(i int) IDeclerationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclerationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclerationContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(dslParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *dslParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, dslParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Decleration()
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097298) != 0 {
		{
			p.SetState(54)
			p.Decleration()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
		p.Match(dslParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITupleContext is an interface to support dynamic dispatch.
type ITupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tuplelist() ITuplelistContext

	// IsTupleContext differentiates from other interfaces.
	IsTupleContext()
}

type TupleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTupleContext() *TupleContext {
	var p = new(TupleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_tuple
	return p
}

func InitEmptyTupleContext(p *TupleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_tuple
}

func (*TupleContext) IsTupleContext() {}

func NewTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleContext {
	var p = new(TupleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_tuple

	return p
}

func (s *TupleContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleContext) Tuplelist() ITuplelistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITuplelistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITuplelistContext)
}

func (s *TupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterTuple(s)
	}
}

func (s *TupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitTuple(s)
	}
}

func (p *dslParser) Tuple() (localctx ITupleContext) {
	localctx = NewTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, dslParserRULE_tuple)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(dslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Tuplelist()
	}
	{
		p.SetState(64)
		p.Match(dslParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITuplelistContext is an interface to support dynamic dispatch.
type ITuplelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsTuplelistContext differentiates from other interfaces.
	IsTuplelistContext()
}

type TuplelistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTuplelistContext() *TuplelistContext {
	var p = new(TuplelistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_tuplelist
	return p
}

func InitEmptyTuplelistContext(p *TuplelistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_tuplelist
}

func (*TuplelistContext) IsTuplelistContext() {}

func NewTuplelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TuplelistContext {
	var p = new(TuplelistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_tuplelist

	return p
}

func (s *TuplelistContext) GetParser() antlr.Parser { return s.parser }

func (s *TuplelistContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *TuplelistContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *TuplelistContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(dslParserComma)
}

func (s *TuplelistContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(dslParserComma, i)
}

func (s *TuplelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TuplelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TuplelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterTuplelist(s)
	}
}

func (s *TuplelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitTuplelist(s)
	}
}

func (p *dslParser) Tuplelist() (localctx ITuplelistContext) {
	localctx = NewTuplelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, dslParserRULE_tuplelist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == dslParserStringLiteral {
		{
			p.SetState(66)
			p.Identifier()
		}
		{
			p.SetState(67)
			p.Match(dslParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclerationContext is an interface to support dynamic dispatch.
type IDeclerationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ActionDeclaration() IActionDeclarationContext
	TypesDeclaration() ITypesDeclarationContext
	MonitorDeclaration() IMonitorDeclarationContext
	MessageDeclaration() IMessageDeclarationContext

	// IsDeclerationContext differentiates from other interfaces.
	IsDeclerationContext()
}

type DeclerationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclerationContext() *DeclerationContext {
	var p = new(DeclerationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_decleration
	return p
}

func InitEmptyDeclerationContext(p *DeclerationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_decleration
}

func (*DeclerationContext) IsDeclerationContext() {}

func NewDeclerationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclerationContext {
	var p = new(DeclerationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_decleration

	return p
}

func (s *DeclerationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclerationContext) ActionDeclaration() IActionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionDeclarationContext)
}

func (s *DeclerationContext) TypesDeclaration() ITypesDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesDeclarationContext)
}

func (s *DeclerationContext) MonitorDeclaration() IMonitorDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMonitorDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMonitorDeclarationContext)
}

func (s *DeclerationContext) MessageDeclaration() IMessageDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageDeclarationContext)
}

func (s *DeclerationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclerationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclerationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterDecleration(s)
	}
}

func (s *DeclerationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitDecleration(s)
	}
}

func (p *dslParser) Decleration() (localctx IDeclerationContext) {
	localctx = NewDeclerationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, dslParserRULE_decleration)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case dslParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.ActionDeclaration()
		}

	case dslParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.TypesDeclaration()
		}

	case dslParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.MonitorDeclaration()
		}

	case dslParserStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.MessageDeclaration()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IActionDeclarationContext is an interface to support dynamic dispatch.
type IActionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Identifier() IIdentifierContext
	StatementList() IStatementListContext

	// IsActionDeclarationContext differentiates from other interfaces.
	IsActionDeclarationContext()
}

type ActionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionDeclarationContext() *ActionDeclarationContext {
	var p = new(ActionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_actionDeclaration
	return p
}

func InitEmptyActionDeclarationContext(p *ActionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_actionDeclaration
}

func (*ActionDeclarationContext) IsActionDeclarationContext() {}

func NewActionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionDeclarationContext {
	var p = new(ActionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_actionDeclaration

	return p
}

func (s *ActionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ActionDeclarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ActionDeclarationContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *ActionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterActionDeclaration(s)
	}
}

func (s *ActionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitActionDeclaration(s)
	}
}

func (p *dslParser) ActionDeclaration() (localctx IActionDeclarationContext) {
	localctx = NewActionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, dslParserRULE_actionDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(dslParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Type_()
	}
	{
		p.SetState(81)
		p.Identifier()
	}
	{
		p.SetState(82)
		p.Match(dslParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.StatementList()
	}
	{
		p.SetState(84)
		p.Match(dslParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMonitorDeclarationContext is an interface to support dynamic dispatch.
type IMonitorDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Identifier() IIdentifierContext
	StatementList() IStatementListContext

	// IsMonitorDeclarationContext differentiates from other interfaces.
	IsMonitorDeclarationContext()
}

type MonitorDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonitorDeclarationContext() *MonitorDeclarationContext {
	var p = new(MonitorDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_monitorDeclaration
	return p
}

func InitEmptyMonitorDeclarationContext(p *MonitorDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_monitorDeclaration
}

func (*MonitorDeclarationContext) IsMonitorDeclarationContext() {}

func NewMonitorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonitorDeclarationContext {
	var p = new(MonitorDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_monitorDeclaration

	return p
}

func (s *MonitorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MonitorDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MonitorDeclarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MonitorDeclarationContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *MonitorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonitorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonitorDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterMonitorDeclaration(s)
	}
}

func (s *MonitorDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitMonitorDeclaration(s)
	}
}

func (p *dslParser) MonitorDeclaration() (localctx IMonitorDeclarationContext) {
	localctx = NewMonitorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, dslParserRULE_monitorDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(dslParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Type_()
	}
	{
		p.SetState(88)
		p.Identifier()
	}
	{
		p.SetState(89)
		p.Match(dslParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.StatementList()
	}
	{
		p.SetState(91)
		p.Match(dslParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageDeclarationContext is an interface to support dynamic dispatch.
type IMessageDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Message() IMessageContext

	// IsMessageDeclarationContext differentiates from other interfaces.
	IsMessageDeclarationContext()
}

type MessageDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageDeclarationContext() *MessageDeclarationContext {
	var p = new(MessageDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_messageDeclaration
	return p
}

func InitEmptyMessageDeclarationContext(p *MessageDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_messageDeclaration
}

func (*MessageDeclarationContext) IsMessageDeclarationContext() {}

func NewMessageDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageDeclarationContext {
	var p = new(MessageDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_messageDeclaration

	return p
}

func (s *MessageDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageDeclarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MessageDeclarationContext) Message() IMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *MessageDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterMessageDeclaration(s)
	}
}

func (s *MessageDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitMessageDeclaration(s)
	}
}

func (p *dslParser) MessageDeclaration() (localctx IMessageDeclarationContext) {
	localctx = NewMessageDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, dslParserRULE_messageDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Identifier()
	}
	{
		p.SetState(94)
		p.Match(dslParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Message()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Assign() antlr.TerminalNode
	Tuple() ITupleContext
	DigitSeq() antlr.TerminalNode
	Func_() IFuncContext

	// IsVarDeclarationContext differentiates from other interfaces.
	IsVarDeclarationContext()
}

type VarDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_varDeclaration
	return p
}

func InitEmptyVarDeclarationContext(p *VarDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_varDeclaration
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_varDeclaration

	return p
}

func (s *VarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *VarDeclarationContext) Assign() antlr.TerminalNode {
	return s.GetToken(dslParserAssign, 0)
}

func (s *VarDeclarationContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *VarDeclarationContext) DigitSeq() antlr.TerminalNode {
	return s.GetToken(dslParserDigitSeq, 0)
}

func (s *VarDeclarationContext) Func_() IFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *VarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitVarDeclaration(s)
	}
}

func (p *dslParser) VarDeclaration() (localctx IVarDeclarationContext) {
	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, dslParserRULE_varDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(dslParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Identifier()
	}
	{
		p.SetState(99)
		p.Match(dslParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case dslParserT__1:
		{
			p.SetState(100)
			p.Tuple()
		}

	case dslParserDigitSeq:
		{
			p.SetState(101)
			p.Match(dslParserDigitSeq)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case dslParserT__9, dslParserT__10, dslParserT__11:
		{
			p.SetState(102)
			p.Func_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBuiltinsContext is an interface to support dynamic dispatch.
type IBuiltinsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBuiltinsContext differentiates from other interfaces.
	IsBuiltinsContext()
}

type BuiltinsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinsContext() *BuiltinsContext {
	var p = new(BuiltinsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_builtins
	return p
}

func InitEmptyBuiltinsContext(p *BuiltinsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_builtins
}

func (*BuiltinsContext) IsBuiltinsContext() {}

func NewBuiltinsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinsContext {
	var p = new(BuiltinsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_builtins

	return p
}

func (s *BuiltinsContext) GetParser() antlr.Parser { return s.parser }
func (s *BuiltinsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltinsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterBuiltins(s)
	}
}

func (s *BuiltinsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitBuiltins(s)
	}
}

func (p *dslParser) Builtins() (localctx IBuiltinsContext) {
	localctx = NewBuiltinsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, dslParserRULE_builtins)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7168) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Builtins() IBuiltinsContext
	Args() IArgsContext

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_func
	return p
}

func InitEmptyFuncContext(p *FuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_func
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) Builtins() IBuiltinsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltinsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltinsContext)
}

func (s *FuncContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (p *dslParser) Func_() (localctx IFuncContext) {
	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, dslParserRULE_func)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Builtins()
	}
	{
		p.SetState(108)
		p.Match(dslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Args()
	}
	{
		p.SetState(110)
		p.Match(dslParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Identifier() IIdentifierContext

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_args
	return p
}

func InitEmptyArgsContext(p *ArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_args
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ArgsContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *dslParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, dslParserRULE_args)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case dslParserT__12:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(dslParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Type_()
		}

	case dslParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignmentExpression() []IAssignmentExpressionContext
	AssignmentExpression(i int) IAssignmentExpressionContext
	AllSendExpression() []ISendExpressionContext
	SendExpression(i int) ISendExpressionContext
	AllRecvExpression() []IRecvExpressionContext
	RecvExpression(i int) IRecvExpressionContext
	AllVarDeclaration() []IVarDeclarationContext
	VarDeclaration(i int) IVarDeclarationContext
	AllBarrier() []IBarrierContext
	Barrier(i int) IBarrierContext

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllAssignmentExpression() []IAssignmentExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentExpressionContext); ok {
			tst[i] = t.(IAssignmentExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) AssignmentExpression(i int) IAssignmentExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *StatementListContext) AllSendExpression() []ISendExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISendExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISendExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISendExpressionContext); ok {
			tst[i] = t.(ISendExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) SendExpression(i int) ISendExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISendExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISendExpressionContext)
}

func (s *StatementListContext) AllRecvExpression() []IRecvExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRecvExpressionContext); ok {
			len++
		}
	}

	tst := make([]IRecvExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRecvExpressionContext); ok {
			tst[i] = t.(IRecvExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) RecvExpression(i int) IRecvExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecvExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecvExpressionContext)
}

func (s *StatementListContext) AllVarDeclaration() []IVarDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclarationContext); ok {
			tst[i] = t.(IVarDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) VarDeclaration(i int) IVarDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclarationContext)
}

func (s *StatementListContext) AllBarrier() []IBarrierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBarrierContext); ok {
			len++
		}
	}

	tst := make([]IBarrierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBarrierContext); ok {
			tst[i] = t.(IBarrierContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) Barrier(i int) IBarrierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBarrierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBarrierContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (p *dslParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, dslParserRULE_statementList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2556416) != 0) {
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(117)
				p.AssignmentExpression()
			}

		case 2:
			{
				p.SetState(118)
				p.SendExpression()
			}

		case 3:
			{
				p.SetState(119)
				p.RecvExpression()
			}

		case 4:
			{
				p.SetState(120)
				p.VarDeclaration()
			}

		case 5:
			{
				p.SetState(121)
				p.Barrier()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentExpressionContext is an interface to support dynamic dispatch.
type IAssignmentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	Assign() antlr.TerminalNode
	Wildcard() antlr.TerminalNode

	// IsAssignmentExpressionContext differentiates from other interfaces.
	IsAssignmentExpressionContext()
}

type AssignmentExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentExpressionContext() *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_assignmentExpression
	return p
}

func InitEmptyAssignmentExpressionContext(p *AssignmentExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_assignmentExpression
}

func (*AssignmentExpressionContext) IsAssignmentExpressionContext() {}

func NewAssignmentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_assignmentExpression

	return p
}

func (s *AssignmentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentExpressionContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentExpressionContext) Assign() antlr.TerminalNode {
	return s.GetToken(dslParserAssign, 0)
}

func (s *AssignmentExpressionContext) Wildcard() antlr.TerminalNode {
	return s.GetToken(dslParserWildcard, 0)
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

func (p *dslParser) AssignmentExpression() (localctx IAssignmentExpressionContext) {
	localctx = NewAssignmentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, dslParserRULE_assignmentExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Variable()
	}
	{
		p.SetState(129)
		p.Match(dslParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(dslParserWildcard)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISendExpressionContext is an interface to support dynamic dispatch.
type ISendExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	Message() IMessageContext

	// IsSendExpressionContext differentiates from other interfaces.
	IsSendExpressionContext()
}

type SendExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendExpressionContext() *SendExpressionContext {
	var p = new(SendExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_sendExpression
	return p
}

func InitEmptySendExpressionContext(p *SendExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_sendExpression
}

func (*SendExpressionContext) IsSendExpressionContext() {}

func NewSendExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendExpressionContext {
	var p = new(SendExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_sendExpression

	return p
}

func (s *SendExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SendExpressionContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *SendExpressionContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SendExpressionContext) Message() IMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *SendExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterSendExpression(s)
	}
}

func (s *SendExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitSendExpression(s)
	}
}

func (p *dslParser) SendExpression() (localctx ISendExpressionContext) {
	localctx = NewSendExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, dslParserRULE_sendExpression)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case dslParserT__16:
		{
			p.SetState(132)
			p.Message()
		}

	case dslParserStringLiteral:
		{
			p.SetState(133)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(136)
		p.Match(dslParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Identifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRecvExpressionContext is an interface to support dynamic dispatch.
type IRecvExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	Message() IMessageContext

	// IsRecvExpressionContext differentiates from other interfaces.
	IsRecvExpressionContext()
}

type RecvExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecvExpressionContext() *RecvExpressionContext {
	var p = new(RecvExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_recvExpression
	return p
}

func InitEmptyRecvExpressionContext(p *RecvExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_recvExpression
}

func (*RecvExpressionContext) IsRecvExpressionContext() {}

func NewRecvExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecvExpressionContext {
	var p = new(RecvExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_recvExpression

	return p
}

func (s *RecvExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RecvExpressionContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RecvExpressionContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RecvExpressionContext) Message() IMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *RecvExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecvExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecvExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterRecvExpression(s)
	}
}

func (s *RecvExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitRecvExpression(s)
	}
}

func (p *dslParser) RecvExpression() (localctx IRecvExpressionContext) {
	localctx = NewRecvExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, dslParserRULE_recvExpression)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case dslParserT__16:
		{
			p.SetState(139)
			p.Message()
		}

	case dslParserStringLiteral:
		{
			p.SetState(140)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(143)
		p.Match(dslParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Identifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBarrierContext is an interface to support dynamic dispatch.
type IBarrierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBarrierContext differentiates from other interfaces.
	IsBarrierContext()
}

type BarrierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBarrierContext() *BarrierContext {
	var p = new(BarrierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_barrier
	return p
}

func InitEmptyBarrierContext(p *BarrierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_barrier
}

func (*BarrierContext) IsBarrierContext() {}

func NewBarrierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BarrierContext {
	var p = new(BarrierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_barrier

	return p
}

func (s *BarrierContext) GetParser() antlr.Parser { return s.parser }
func (s *BarrierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BarrierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BarrierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterBarrier(s)
	}
}

func (s *BarrierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitBarrier(s)
	}
}

func (p *dslParser) Barrier() (localctx IBarrierContext) {
	localctx = NewBarrierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, dslParserRULE_barrier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(dslParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringLiteral() antlr.TerminalNode

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_message
	return p
}

func InitEmptyMessageContext(p *MessageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_message
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(dslParserStringLiteral, 0)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitMessage(s)
	}
}

func (p *dslParser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, dslParserRULE_message)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(dslParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(dslParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(dslParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	Primitive() IPrimitiveContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(dslParserStringLiteral, 0)
}

func (s *TypeContext) Primitive() IPrimitiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *dslParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, dslParserRULE_type)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case dslParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(dslParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case dslParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Primitive()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitiveContext is an interface to support dynamic dispatch.
type IPrimitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimitiveContext differentiates from other interfaces.
	IsPrimitiveContext()
}

type PrimitiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveContext() *PrimitiveContext {
	var p = new(PrimitiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_primitive
	return p
}

func InitEmptyPrimitiveContext(p *PrimitiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_primitive
}

func (*PrimitiveContext) IsPrimitiveContext() {}

func NewPrimitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveContext {
	var p = new(PrimitiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_primitive

	return p
}

func (s *PrimitiveContext) GetParser() antlr.Parser { return s.parser }
func (s *PrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterPrimitive(s)
	}
}

func (s *PrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitPrimitive(s)
	}
}

func (p *dslParser) Primitive() (localctx IPrimitiveContext) {
	localctx = NewPrimitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, dslParserRULE_primitive)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(dslParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Identifier() IIdentifierContext

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VariableContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *dslParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, dslParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Type_()
	}
	{
		p.SetState(159)
		p.Identifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringLiteral() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = dslParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(dslParserStringLiteral, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *dslParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, dslParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(dslParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
