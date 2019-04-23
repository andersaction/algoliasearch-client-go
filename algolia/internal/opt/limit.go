// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractLimit returns the first found LimitOption from the
// given variadic arguments or nil otherwise.
func ExtractLimit(opts ...interface{}) *opt.LimitOption {
	for _, o := range opts {
		if v, ok := o.(*opt.LimitOption); ok {
			return v
		}
	}
	return nil
}
