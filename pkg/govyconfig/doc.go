// Package govyconfig defines global configuration options for govy.
// The functions defined by govyconfig can be safely called concurrently.
// However, it is still important to take care of the order of calls
// to both govy and govyconfig functions.
//
// For instance, calling [govy.For] before [SetInferPathIncludeTestFiles]
// means that property rules created by that constructor call will not consider test files
// when later inferring their relative property paths.
package govyconfig
