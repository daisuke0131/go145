package main

import (
	"go145"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(go145.Analyzer) }
