// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractOffset returns the first found OffsetOption from the
// given variadic arguments or nil otherwise.
func ExtractOffset(opts ...interface{}) *opt.OffsetOption {
	for _, o := range opts {
		if v, ok := o.(*opt.OffsetOption); ok {
			return v
		}
	}
	return nil
}
