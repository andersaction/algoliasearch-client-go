// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

// RankingOption is a wrapper for an Ranking option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type RankingOption struct {
	value []string
}

// Ranking wraps the given value into a RankingOption.
func Ranking(v ...string) *RankingOption {
	return &RankingOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *RankingOption) Get() []string {
	if o == nil {
		return []string{"typo", "geo", "words", "filters", "proximity", "attribute", "exact", "custom"}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// RankingOption.
func (o RankingOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// RankingOption.
func (o *RankingOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{"typo", "geo", "words", "filters", "proximity", "attribute", "exact", "custom"}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *RankingOption) Equal(o2 *RankingOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, []string{"typo", "geo", "words", "filters", "proximity", "attribute", "exact", "custom"})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// RankingEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func RankingEqual(o1, o2 *RankingOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
