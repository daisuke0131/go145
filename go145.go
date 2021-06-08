package go145

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "go145 is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "go145",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		switch expr := n.(type) {
		case ast.Expr:
			if tv, ok := pass.TypesInfo.Types[expr]; ok{
				fmt.Println(expr,pass.Fset.Position(expr.Pos()),tv.Type,tv.Value)
			}
		}
	})

	return nil, nil
}
