// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestLength(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.LengthOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Length(0),
		},
		{
			opts:     []interface{}{opt.Length(0)},
			expected: opt.Length(0),
		},
		{
			opts:     []interface{}{opt.Length(1)},
			expected: opt.Length(1),
		},
		{
			opts:     []interface{}{opt.Length(-42)},
			expected: opt.Length(-42),
		},
	} {
		var (
			in  = ExtractLength(c.opts...)
			out opt.LengthOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
