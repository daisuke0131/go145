package go145_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"go145"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, go145.Analyzer, "a")
}
