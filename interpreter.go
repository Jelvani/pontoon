package main

import (
	"math/rand"
	"strconv"
	"time"
	dsl "github.com/jelvani/pontoon/parser"
	"strings"
)

type dslListener struct {
	*dsl.BasedslListener
	variables map[string]value
	nodeType string
	nodeTypes []string
	currentActionType string
	thisNode *Node
	
}

type dslVisitor struct {
	*dsl.BasedslVisitor
}

type value struct {
	type_ string
	val   string
	validNodeIds []*Node
}

const MAXRAND = 1000

//on action exit we flush the instruction buffer
func (s *dslListener) EnterActionDeclaration(ctx *dsl.ActionDeclarationContext) {
	if s.nodeType != "null" {
		s.currentActionType = ctx.Type_().GetText()
	}
	

}

func (s *dslListener) ExitAssignmentExpression(ctx *dsl.AssignmentExpressionContext) {
	if s.nodeType != s.currentActionType {
		return
	}
	var v value
	v.type_ = ctx.Variable().Type_().GetText()
	
	v.val = ctx.Wildcard().GetText()
	var t = strings.Split(v.val, "_")
	i,err := strconv.Atoi(t[1])
	if err != nil {
        // ... handle error
        panic(err)
    }
	var validNodeIds []*Node
	for _,n := range nodePool {
		if i <= 0 {
			break
		}
		if n.Type == v.type_ {
			validNodeIds = append(validNodeIds,n)
		}
		i = i-1

	}
	v.validNodeIds = validNodeIds
	s.variables[ctx.Variable().Identifier().GetText()] = v

}

// evaluate function
func VisitFunc(ctx *dsl.FuncContext) int {

	return 0
}

func (s *dslListener) ExitVarDeclaration(ctx *dsl.VarDeclarationContext) {
	if s.nodeType != s.currentActionType {
		return
	}

	var v value
	if ctx.Tuple() != nil {
		v.val = ctx.Tuple().GetText()
	} else if ctx.Func_() != nil {
		if ctx.Func_().Builtins().GetText() == "rand" {
			if ctx.Func_().Args().Type_().GetText() == "int" {
				rand.Seed(time.Now().UnixNano())
				v.val = strconv.Itoa(rand.Intn(MAXRAND))
			}
		}

	}

	s.variables[ctx.Identifier().GetText()] = v
}

func (s *dslListener) init() {
	s.variables = make(map[string]value)
}

func (s *dslListener) ExitTypesDeclaration(ctx *dsl.TypesDeclarationContext) {

	if s.nodeType != s.currentActionType {
		return
	}

	stringSlice := strings.Split(ctx.GetText()[6:], ",")
	s.nodeTypes = stringSlice[:]
}

func (s *dslListener) ExitSendExpression(ctx *dsl.SendExpressionContext) {
	var sourceData = ctx.Identifier(0).GetText()
	var destNodeIdent = ctx.Identifier(1).GetText()


	for _,n := range s.variables[destNodeIdent].validNodeIds {
		// for each node bound to identifer, create message
		var m Message
		m.SourceId = s.thisNode.Id
		m.DestId = n.Id
		m.Payload = sourceData
		m.Timestamp = time.Now()

		// send message to global router network
		internet.Nic <- m
		
	}
}