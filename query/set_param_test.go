package query

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetParam(t *testing.T) {
	t.Parallel()

	t.Run("no panic on nil target", func(t *testing.T) {
		t.Parallel()

		require.NotPanics(t, func() {
			SetParam(nil, "key", "value")
		})
	})

	t.Run("sequential update", func(t *testing.T) {
		t.Parallel()

		vals := url.Values{
			"key1": []string{"value1", "value2"},
			"key2": []string{"value2"},
		}
		expected := url.Values{
			"key1": []string{"value3"},
			"key2": []string{"value4", "value5"},
		}

		SetParam(vals, "key1", "value3")
		SetParam(vals, "key2", "value4", "value5")

		require.Equal(t, expected, vals)
	})
}
