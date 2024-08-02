// Package pkg hold the library code that can be used by external applications (e.g., /pkg/public).
//
// Other projects will import these libraries expecting them to work, so think twice before you put
// something here :-)
// Note that the internal directory is the right way to ensure your private packages are not importable
// because it's enforced by Go compiler.
// The /pkg directory is a good way to explicitly communicate that
// the code in that directory is safe for use by others.
//
// It's also a way to group Go code in one place when your root directory contains lots of non-Go
// components and directories making it easier to run various Go tools.
// This is a common layout pattern, but it's not universally accepted and some in the Go community
// don't recommend it.
//
// The pkg directory origins: The old Go source code used to use pkg for its packages and then various
// Go projects in the community started copying the pattern (see this Brad Fitzpatrick's tweet for more context).
//
// Source: https://github.com/golang-standards/project-layout#pkg
package pkg

// GetWorld is part of the public contract of this library.
func GetWorld() string {
	return "World"
}
