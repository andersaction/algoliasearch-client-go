// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractMinimumAroundRadius returns the first found MinimumAroundRadiusOption from the
// given variadic arguments or nil otherwise.
func ExtractMinimumAroundRadius(opts ...interface{}) *opt.MinimumAroundRadiusOption {
	for _, o := range opts {
		if v, ok := o.(*opt.MinimumAroundRadiusOption); ok {
			return v
		}
	}
	return nil
}
