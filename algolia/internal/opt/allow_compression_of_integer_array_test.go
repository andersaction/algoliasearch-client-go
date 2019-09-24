// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAllowCompressionOfIntegerArray(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.AllowCompressionOfIntegerArrayOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AllowCompressionOfIntegerArray(false),
		},
		{
			opts:     []interface{}{opt.AllowCompressionOfIntegerArray(true)},
			expected: opt.AllowCompressionOfIntegerArray(true),
		},
		{
			opts:     []interface{}{opt.AllowCompressionOfIntegerArray(false)},
			expected: opt.AllowCompressionOfIntegerArray(false),
		},
	} {
		var (
			in  = ExtractAllowCompressionOfIntegerArray(c.opts...)
			out opt.AllowCompressionOfIntegerArrayOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
