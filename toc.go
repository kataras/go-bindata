// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// createTOC writes a table of contents file to the given location.
func createTOC(file, pkgname string) error {
	dir, _ := filepath.Split(file)
	file = filepath.Join(dir, "bindata-toc.go")
	code := fmt.Sprintf(`package %s

// Global Table of Contents map. Generated by go-bindata.
// After startup of the program, all generated data files will
// put themselves in this map. The key is the full filename, as
// supplied to go-bindata.
var go_bindata = make(map[string] func() []byte)`, pkgname)

	return ioutil.WriteFile(file, []byte(code), 0600)
}

// writeTOCInit writes the TOC init function for a given data file.
func writeTOCInit(output io.Writer, filename, prefix, funcname string) {
	filename = strings.Replace(filename, prefix, "", 1)
	fmt.Fprintf(output, `

func init() {
	go_bindata[%q] = %s
}
`, filename, funcname)
}