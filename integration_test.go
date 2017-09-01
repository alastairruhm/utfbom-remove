package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	shutil "github.com/termie/go-shutil"
	"github.com/urfave/cli"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestIntegration(t *testing.T) { TestingT(t) }

type AppRunSuite struct {
	dir string
	app cli.App
}

var _ = Suite(&AppRunSuite{})

func (s *AppRunSuite) SetUpSuite(c *C) {
	s.dir = filepath.Join(c.MkDir(), "test") // TODO: c.MkDir() will be deleted TearDownSuite ?
	s.app = NewApp()
	err := shutil.CopyTree("testdata", s.dir, nil)
	if err != nil {
		c.Fatalf("%v", err)
	}
}

func (s *AppRunSuite) TestAppRunOutputVersion(c *C) {
	buf := new(bytes.Buffer)
	s.app.Writer = buf
	err := s.app.Run([]string{"utfbom-remove", "-v"})
	if err != nil {
		c.Fatal(err)
	}
	output := buf.String()
	c.Assert(output, Contains, s.app.Version)
}

func (s *AppRunSuite) TestAppRunCheckBom(c *C) {
	buf := new(bytes.Buffer)
	s.app.Writer = buf
	caseDir := filepath.Join(s.dir, "test000")
	err := s.app.Run([]string{"utfbom-remove", "--check-only", "--path", caseDir})
	if err != nil {
		c.Fatal(err)
	}
	output := buf.String()
	c.Assert(output, Contains, "bom.txt")
}

// -----------------------------------------------------------------------
// Equals checker.

type containsChecker struct {
	*CheckerInfo
}

// The Equals checker verifies that the obtained value is equal to
// the expected value, according to usual Go semantics for ==.
//
// For example:
//
//     c.Assert(value, Equals, 42)
//
var Contains Checker = &containsChecker{
	&CheckerInfo{Name: "Contains", Params: []string{"obtained", "expected"}},
}

func (checker *containsChecker) Check(params []interface{}, names []string) (result bool, error string) {
	return strings.Contains(params[0].(string), params[1].(string)), ""
}
