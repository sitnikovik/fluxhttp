package query

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeInto(t *testing.T) {
	t.Parallel()

	t.Run("no panic on nil target and sources", func(t *testing.T) {
		t.Parallel()

		require.NotPanics(t, func() {
			MergeInto(nil, nil, nil)
		})
	})

	t.Run("ok merged", func(t *testing.T) {
		t.Parallel()

		target := url.Values{
			"key1": []string{"value1"},
			"key2": []string{"value2"},
		}
		src1 := url.Values{
			"key1": []string{"value3"},
			"key3": []string{"value4"},
		}
		src2 := url.Values{
			"key2": []string{"value5"},
			"key4": []string{"value6"},
		}
		var src3 url.Values

		MergeInto(target, src1, src2, src3)

		require.Equal(t, url.Values{
			"key1": []string{"value1", "value3"},
			"key2": []string{"value2", "value5"},
			"key3": []string{"value4"},
			"key4": []string{"value6"},
		}, target)
	})
}
