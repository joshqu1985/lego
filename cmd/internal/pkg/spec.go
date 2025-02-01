package pkg

import "github.com/antlr4-go/antlr/v4"

type ParseTopContext interface {
	GetStart() antlr.Token
	GetStop() antlr.Token
	GetParser() antlr.Parser
}

type Tree struct {
	Imports  []string
	Options  map[string]string
	Enums    []*EnumNode
	Structs  []*StructNode
	Services []*ServiceNode
}

type EnumNode struct {
	Name     string
	Document string
	Fields   []*EnumField
}

type StructNode struct {
	Name      string
	Document  string
	Fields    []*StructField
	MapFields []*StructMap
	Options   map[string]string
}

type ServiceNode struct {
	Name     string
	Document string
	Methods  []*ServiceMethod
}

type EnumField struct {
	Name    string
	Value   int
	Comment string
}

type StructField struct {
	Repeated bool
	Name     string
	Type     string
	Tag      string
	Comment  string
}

type StructMap struct {
	Name    string
	KeyType string
	ValType string
	Tag     string
	Comment string
}

type ServiceMethod struct {
	Name     string
	Document string
	ReqName  string
	ResName  string
	Options  map[string]string // cmd / uri
}

type OptionConst struct {
	Name     string
	Constant string
}
