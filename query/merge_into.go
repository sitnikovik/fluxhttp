package query

import "net/url"

// MergeInto merges (not replaces) multiple url.Values into a single url.Values.
// The values of the sources are appended to the target.
// If the target already has a value for a key, the values of the source are appended to the target.
// The target is modified in place.
func MergeInto(target url.Values, sources ...url.Values) {
	for _, src := range sources {
		for k, vs := range src {
			target[k] = append(target[k], vs...)
		}
	}
}
