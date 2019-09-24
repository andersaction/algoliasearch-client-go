// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestPaginationLimitedTo(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.PaginationLimitedToOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.PaginationLimitedTo(1000),
		},
		{
			opts:     []interface{}{opt.PaginationLimitedTo(0)},
			expected: opt.PaginationLimitedTo(0),
		},
		{
			opts:     []interface{}{opt.PaginationLimitedTo(1)},
			expected: opt.PaginationLimitedTo(1),
		},
		{
			opts:     []interface{}{opt.PaginationLimitedTo(-42)},
			expected: opt.PaginationLimitedTo(-42),
		},
	} {
		var (
			in  = ExtractPaginationLimitedTo(c.opts...)
			out opt.PaginationLimitedToOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
