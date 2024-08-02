// Package internal holds private application and library code.
//
// This is the code you don't want others importing in their applications or libraries.
// Note that this layout pattern is enforced by the Go compiler itself.
//
// You can optionally add a bit of extra structure to your internal packages to separate
// your shared and non-shared internal code.
// It's not required (especially for smaller projects), but it's nice to have visual
// clues showing the intended package use.
// Your actual application code can go in the /internal/app directory (e.g., /internal/app/app)
// and the code shared by those apps in the /internal/pkg directory (e.g., /internal/pkg/private).
//
// Source: https://github.com/golang-standards/project-layout#internal
package internal

// GetHello is not accessible outside this module.
// It is a private function, and thus not part of the public contract.
func GetHello() string {
	return "Hello"
}
