package query

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetInto(t *testing.T) {
	t.Run("no panic on nil target", func(t *testing.T) {
		t.Parallel()

		require.NotPanics(t, func() {
			SetInto(nil, nil)
		})
	})

	t.Run("no panic on nil source", func(t *testing.T) {
		t.Parallel()

		require.NotPanics(t, func() {
			SetInto(url.Values{}, nil)
		})
	})

	t.Run("sequential update", func(t *testing.T) {
		t.Parallel()

		params := url.Values{
			"key1": {"old"},
			"key2": {"value2"},
		}
		expected := url.Values{
			"key1": {"new"},
			"key2": {"value2_new"},
			"key3": {"value3"},
		}
		SetInto(params, url.Values{
			"key1": {"new"},
			"key2": {"value2_new"},
			"key3": {"value3"},
		})

		require.Equal(t, expected, params)
	})
}
