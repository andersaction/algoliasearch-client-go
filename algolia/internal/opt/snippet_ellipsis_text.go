// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractSnippetEllipsisText returns the first found SnippetEllipsisTextOption from the
// given variadic arguments or nil otherwise.
func ExtractSnippetEllipsisText(opts ...interface{}) *opt.SnippetEllipsisTextOption {
	for _, o := range opts {
		if v, ok := o.(*opt.SnippetEllipsisTextOption); ok {
			return v
		}
	}
	return nil
}
