// Package main holds executable functions for this project.
//
// If you have more than one executable, a directory name for each application should match
// the name of the executable you want to have (e.g., /cmd/app).
//
// Don't put a lot of code in the application directory.
// If you think the code can be imported and used in other projects, then it should live in
// the /pkg directory.
// If the code is not reusable or if you don't want others to reuse it, put that code in the
// /internal directory.
// You'll be surprised what others will do, so be explicit about your intentions!
//
// It's common to have a small main function that imports and invokes the code from the /internal
// and /pkg directories and nothing else.
//
// Source: https://github.com/golang-standards/project-layout#cmd
package main

import (
	"fmt"

	"github.com/nobl9/your-module-name/internal"
	"github.com/nobl9/your-module-name/pkg"
)

// main is the entry point for your module's executable.
func main() {
	fmt.Println(internal.GetHello(), pkg.GetWorld())
}
