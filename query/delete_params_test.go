package query

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteParams(t *testing.T) {
	t.Parallel()

	t.Run("no panic on nil target", func(t *testing.T) {
		t.Parallel()

		require.NotPanics(t, func() {
			DeleteParams(nil, "foo")
		})
	})

	t.Run("delete params", func(t *testing.T) {
		t.Parallel()

		params := url.Values{
			"key1": {"value1"},
			"key2": {"value2"},
			"key3": {"value3"},
		}
		expected := url.Values{
			"key1": {"value1"},
		}
		DeleteParams(params, "key2", "key3")

		require.Equal(t, expected, params)
	})
}
