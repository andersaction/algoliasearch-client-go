// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

// ExplainOption is a wrapper for an Explain option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type ExplainOption struct {
	value []string
}

// Explain wraps the given value into a ExplainOption.
func Explain(v ...string) *ExplainOption {
	return &ExplainOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *ExplainOption) Get() []string {
	if o == nil {
		return []string{}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// ExplainOption.
func (o ExplainOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// ExplainOption.
func (o *ExplainOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *ExplainOption) Equal(o2 *ExplainOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// ExplainEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func ExplainEqual(o1, o2 *ExplainOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
