// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractPage returns the first found PageOption from the
// given variadic arguments or nil otherwise.
func ExtractPage(opts ...interface{}) *opt.PageOption {
	for _, o := range opts {
		if v, ok := o.(*opt.PageOption); ok {
			return v
		}
	}
	return nil
}
