// Package govyconfig defines global configuration options for govy.
// The functions defined by govyconfig can be safely called concurrently.
// However, it is still important to take care of the order of calls
// to both govy and govyconfig functions.
//
// For instance calling [govy.For] before [SetInferNameIncludeTestFiles]
// will not have any effect on the [govy.PropertyRules] instance created with [govy.For].
package govyconfig
