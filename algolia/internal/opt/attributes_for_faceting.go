// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractAttributesForFaceting returns the first found AttributesForFacetingOption from the
// given variadic arguments or nil otherwise.
func ExtractAttributesForFaceting(opts ...interface{}) *opt.AttributesForFacetingOption {
	for _, o := range opts {
		if v, ok := o.(*opt.AttributesForFacetingOption); ok {
			return v
		}
	}
	return nil
}
