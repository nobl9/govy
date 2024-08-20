// Package govyconfig defines configuration functions for govy.
// It also holds internal, shared state for the library.
// The functions defined by govyconfig can be safely called concurrently.
// However, bear in mind that it is still important to take care of the order of calls
// to both govy and govyconfig functions.
//
// For instance calling [govy.For] before [SetNameInferMode] will not have any effect,
// for the [govy.PropertyRules] instance created with [govy.For].
package govyconfig
