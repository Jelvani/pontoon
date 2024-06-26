package main

import (
	"fmt"
	"os"
	dsl "github.com/jelvani/pontoon/parser"
	antlr "github.com/antlr4-go/antlr/v4"
	"sync"
	"time"
)

type Message struct {
	SourceId  int
	DestId    int
	Payload   string
	Timestamp time.Time
}

type Router struct {
	Nic       chan Message
	neighbors []*Router
}



type Node struct {
	Type     string
	Id       int
	Flops    float32
	Epoch    int
	Nic      chan Message
	StateLog []string
}

var nodePool []*Node
var internet Router
const NETBUF = 1000

func networkHandler () {
	defer wg.Done()
	msg := <- internet.Nic
	fmt.Println(msg)
	for _, n := range nodePool {
		if n.Id == msg.DestId {
			n.Nic <- msg
		}
	}
	
}


func NodeCtx (filename string, Type string, Id int, ruleNames []string) {
	defer wg.Done()
	var n Node
	n.Type = Type
	n.Id = Id
	n.Nic = make(chan Message, NETBUF)
	nodePool = append(nodePool,&n)
	dispatch(tree, ruleNames)
	// var dl dslListener
	// dl.init()
	// dl.nodeType = n.Type
	// dl.thisNode = &n

	// is, _ := antlr.NewFileStream(filename)
	// lexer := dsl.NewdslLexer(is)
	// stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// parser := dsl.NewdslParser(stream)
	//antlr.ParseTreeWalkerDefault.Walk(&dl, parser.Program())

}
var wg sync.WaitGroup

var gs globalState

func main() {
	// Setup the input
	is, _ := antlr.NewFileStream(os.Args[1])
	// Create the Lexer
	lexer := dsl.NewdslLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := dsl.NewdslParser(stream)

	
	

	internet.Nic = make(chan Message, NETBUF)
	//var gs globalState

	//antlr.ParseTreeWalkerDefault.Walk(&dl, parser.TypesDeclaration())

	//entry point into program grammar rule
	tree := parser.Program()
	ruleNames := parser.GetRuleNames()
	fmt.Println(ruleNames)
	initState(tree, ruleNames)
	
	os.Exit(0)

	go networkHandler()
	wg.Add(1)
	for i, s := range gs.nodeTypes {
		wg.Add(1)
        go NodeCtx(os.Args[1], s, i, ruleNames)
	}

	wg.Wait()
}