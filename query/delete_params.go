package query

import "net/url"

// DeleteParams deletes params from the target by keys.
// The target is modified in place.
func DeleteParams(target url.Values, keys ...string) {
	for _, key := range keys {
		target.Del(key)
	}
}
