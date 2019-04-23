// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractFacetingAfterDistinct returns the first found FacetingAfterDistinctOption from the
// given variadic arguments or nil otherwise.
func ExtractFacetingAfterDistinct(opts ...interface{}) *opt.FacetingAfterDistinctOption {
	for _, o := range opts {
		if v, ok := o.(*opt.FacetingAfterDistinctOption); ok {
			return v
		}
	}
	return nil
}
