package pkg

import "github.com/antlr4-go/antlr/v4"

const (
	StringType  = "type"
	StringDst   = "dst"
	StringProto = "proto"
	StringForm  = "form"
	StringAuth  = "auth"
	StringDB    = "db"
	StringHost  = "host"
)

type (
	ParseTopContext interface {
		GetStart() antlr.Token
		GetStop() antlr.Token
		GetParser() antlr.Parser
	}

	Tree struct {
		Package  string
		Imports  []string
		Options  map[string]string
		Enums    []*EnumNode
		Structs  []*StructNode
		Services []*ServiceNode
	}

	EnumNode struct {
		Name     string
		Document string
		Fields   []*EnumField
	}

	StructNode struct {
		Options   map[string]string
		Name      string
		Document  string
		Fields    []*StructField
		MapFields []*StructMap
	}

	ServiceNode struct {
		Name     string
		Document string
		Methods  []*ServiceMethod
	}

	EnumField struct {
		Name    string
		Comment string
		Value   int
	}

	StructField struct {
		Name       string
		Type       string
		Tag        string
		Comment    string
		TypeName   string
		Repeated   bool
		TypeObject bool
	}

	StructMap struct {
		Name    string
		KeyType string
		ValType string
		Tag     string
		Comment string
	}

	ServiceMethod struct {
		Options  map[string]string
		Name     string
		Document string
		ReqName  string
		ResName  string
	}

	OptionConst struct {
		Name     string
		Constant string
	}
)
