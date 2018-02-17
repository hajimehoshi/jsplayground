package imports_test

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"testing"

	. "github.com/hajimehoshi/jsplayground/imports"
)

const testfile1 = `

package main

import (
	"fmt"
)

func main() {
	println("Hello World")
}

`

const testfile2 = `
package main

func main() {
	fmt.Println("Hello World")
}
`

func TestRemove(t *testing.T) {
	fset := &token.FileSet{}
	f, err := parser.ParseFile(fset, "test1.go", testfile1, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = FixImports(fset, f); err != nil {
		t.Fatal(err)
	}

	buf := &bytes.Buffer{}
	err = printer.Fprint(buf, fset, f)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(buf.String())
}

func TestAdd(t *testing.T) {
	fset := &token.FileSet{}
	f, err := parser.ParseFile(fset, "test2.go", testfile2, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = FixImports(fset, f); err != nil {
		t.Fatal(err)
	}

	buf := &bytes.Buffer{}
	err = printer.Fprint(buf, fset, f)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(buf.String())
}
