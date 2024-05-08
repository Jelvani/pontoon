package main

import (
	_ "math/rand"
	_ "strconv"
	_ "time"
	_ "github.com/jelvani/pontoon/parser"
	_ "strings"
	"fmt"
	antlr "github.com/antlr4-go/antlr/v4"
	"os"
	_ "reflect"
)

type globalState struct {
	variables map[string]value
	nodeType string
	nodeTypes []string
	currentActionType string
	thisNode *Node
	
}


type value struct {
	type_ string
	val   string
	validNodeIds []*Node
}

const MAXRAND = 1000

func dispatch (tree antlr.ParseTree, ruleNames []string) {

	children := tree.GetChildren()
	for _, child := range children {
		switch c := child.(type) {
			case antlr.RuleContext:
				//idea: https://github.com/huahuayu/go-dynamic-call
				var rule string = ruleNames[c.GetRuleIndex()]

				fmt.Println(rule)

				switch rule {
					case "typesDeclaration":
						fmt.Println("types!!")
						typesDeclaration(c)
				}
				dispatch(c,ruleNames)
				
		}
	}
	
}

func typesDeclaration(node antlr.ParseTree) {
	children := node.GetChildren()
	fmt.Println("called type!")
	for _, child := range children {
		switch c := child.(type) {
			case antlr.ParseTree:
				nodeType := c.GetText()
				if nodeType != "," {
					gs.nodeTypes = append(gs.nodeTypes, nodeType)
				}
		}

	}

	for i, s := range gs.nodeTypes {
		wg.Add(1)
        go NodeCtx(os.Args[1], s, i)
	}
}



// //on action exit we flush the instruction buffer
// func WalkActionDeclaration(tree antlr.ParseTree) {
	

// }

// func WalkAssignmentExpression(tree antlr.ParseTree) {
// 	if s.nodeType != s.currentActionType {
// 		return
// 	}
// 	var v value
// 	v.type_ = ctx.Variable().Type_().GetText()
	
// 	v.val = ctx.Wildcard().GetText()
// 	var t = strings.Split(v.val, "_")
// 	i,err := strconv.Atoi(t[1])
// 	if err != nil {
//         // ... handle error
//         panic(err)
//     }
// 	var validNodeIds []*Node
// 	for _,n := range nodePool {
// 		if i <= 0 {
// 			break
// 		}
// 		if n.Type == v.type_ {
// 			validNodeIds = append(validNodeIds,n)
// 		}
// 		i = i-1

// 	}
// 	v.validNodeIds = validNodeIds
// 	s.variables[ctx.Variable().Identifier().GetText()] = v

// }


// func WalkVarDeclaration(tree antlr.ParseTree) {
// 	if s.nodeType != s.currentActionType {
// 		return
// 	}

// 	var v value
// 	if ctx.Tuple() != nil {
// 		v.val = ctx.Tuple().GetText()
// 	} else if ctx.Func_() != nil {
// 		if ctx.Func_().Builtins().GetText() == "rand" {
// 			if ctx.Func_().Args().Type_().GetText() == "int" {
// 				rand.Seed(time.Now().UnixNano())
// 				v.val = strconv.Itoa(rand.Intn(MAXRAND))
// 			}
// 		}

// 	}

// 	s.variables[ctx.Identifier().GetText()] = v
// }

// func init() {
// 	s.variables = make(map[string]value)
// }

// func WalkTypesDeclaration(tree antlr.ParseTree) {

// 	if s.nodeType != s.currentActionType {
// 		return
// 	}

// 	stringSlice := strings.Split(ctx.GetText()[6:], ",")
// 	s.nodeTypes = stringSlice[:]
// }

// func WalkSendExpression(tree antlr.ParseTree) {
// 	var sourceData = ctx.Identifier(0).GetText()
// 	var destNodeIdent = ctx.Identifier(1).GetText()


// 	for _,n := range s.variables[destNodeIdent].validNodeIds {
// 		// for each node bound to identifer, create message
// 		var m Message
// 		m.SourceId = s.thisNode.Id
// 		m.DestId = n.Id
// 		m.Payload = sourceData
// 		m.Timestamp = time.Now()

// 		// send message to global router network
// 		internet.Nic <- m
		
// 	}
// }