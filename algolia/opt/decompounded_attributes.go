// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

// DecompoundedAttributesOption is a wrapper for an DecompoundedAttributes option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type DecompoundedAttributesOption struct {
	value map[string][]string
}

// DecompoundedAttributes wraps the given value into a DecompoundedAttributesOption.
func DecompoundedAttributes(v map[string][]string) *DecompoundedAttributesOption {
	return &DecompoundedAttributesOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *DecompoundedAttributesOption) Get() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// DecompoundedAttributesOption.
func (o DecompoundedAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// DecompoundedAttributesOption.
func (o *DecompoundedAttributesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = map[string][]string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *DecompoundedAttributesOption) Equal(o2 *DecompoundedAttributesOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, map[string][]string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// DecompoundedAttributesEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func DecompoundedAttributesEqual(o1, o2 *DecompoundedAttributesOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
