package main

import (
	"fmt"
	"os"
	dsl "github.com/jelvani/pontoon/parser"
	antlr "github.com/antlr4-go/antlr/v4"
	"sync"
	"time"
	"reflect"
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


func NodeCtx (filename string, Type string, Id int) {
	defer wg.Done()
	var n Node
	n.Type = Type
	n.Id = Id
	n.Nic = make(chan Message, NETBUF)
	nodePool = append(nodePool,&n)
	var dl dslListener
	dl.init()
	dl.nodeType = n.Type
	dl.thisNode = &n

	is, _ := antlr.NewFileStream(filename)
	lexer := dsl.NewdslLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := dsl.NewdslParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&dl, parser.Program())
	fmt.Println(dl)

}
var wg sync.WaitGroup
func main() {
	// Setup the input
	is, _ := antlr.NewFileStream(os.Args[1])
	// Create the Lexer
	lexer := dsl.NewdslLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := dsl.NewdslParser(stream)

	//entry point into program grammar rule
	

	internet.Nic = make(chan Message, NETBUF)
	var dl dslListener
	dl.init()
	dl.nodeType = "null"
	dl.currentActionType = "null"

	antlr.ParseTreeWalkerDefault.Walk(&dl, parser.TypesDeclaration())

	tree := parser.Program()
	fmt.Println(tree.GetText())
	fooType := reflect.TypeOf(tree)
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)
	}
	go networkHandler()
	wg.Add(1)
	for i, s := range dl.nodeTypes {
		wg.Add(1)
        NodeCtx(os.Args[1], s, i)
	}
	wg.Wait()


}
