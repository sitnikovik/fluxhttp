package query

import (
	"net/url"
)

// SetInto sets (or replaces) the values of the target with the values of the source.
// If the target already has a value for a key, the values of the source replace the target.
// Updating is done sequentially and the target is modified in place.
func SetInto(target url.Values, source url.Values) {
	if target == nil {
		return
	}

	for k, vv := range source {
		target.Del(k)
		for _, v := range vv {
			target.Add(k, v)
		}
	}
}
