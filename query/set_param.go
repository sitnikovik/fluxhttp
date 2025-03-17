package query

import "net/url"

// SetParam sets (or replaces) the values of the target with the values of the source.
// If the target already has a value for a key, the specified values replace the target ones.
// Updating is done sequentially and the target is modified in place.
func SetParam(target url.Values, key string, values ...string) {
	if target == nil {
		return
	}

	target.Del(key)
	for _, v := range values {
		target.Add(key, v)
	}
}
