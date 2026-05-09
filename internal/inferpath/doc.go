// Package inferpath provides utilities for inferring relative property paths
// from govy getter expressions.
// The inferred paths do not include a leading `$` segment;
// parent validators and collection rules prepend their own segments later.
package inferpath
