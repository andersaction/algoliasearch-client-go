// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

// QueryOption is a wrapper for an Query option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type QueryOption struct {
	value string
}

// Query wraps the given value into a QueryOption.
func Query(v string) *QueryOption {
	return &QueryOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *QueryOption) Get() string {
	if o == nil {
		return ""
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// QueryOption.
func (o QueryOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// QueryOption.
func (o *QueryOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = ""
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *QueryOption) Equal(o2 *QueryOption) bool {
	if o2 == nil {
		return o.value == ""
	}
	return o.value == o2.value
}

// QueryEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func QueryEqual(o1, o2 *QueryOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
