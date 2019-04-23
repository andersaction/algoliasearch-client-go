// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractFacetFilters returns the first found FacetFiltersOption from the
// given variadic arguments or nil otherwise.
func ExtractFacetFilters(opts ...interface{}) *opt.FacetFiltersOption {
	for _, o := range opts {
		if v, ok := o.(*opt.FacetFiltersOption); ok {
			return v
		}
	}
	return nil
}
