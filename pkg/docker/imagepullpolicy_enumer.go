// Code generated by "enumer -type=ImagePullPolicy -trimprefix=ImagePullPolicy --json"; DO NOT EDIT.

//
package docker

import (
	"encoding/json"
	"fmt"
)

const _ImagePullPolicyName = "AlwaysIfNotPresentNever"

var _ImagePullPolicyIndex = [...]uint8{0, 6, 18, 23}

func (i ImagePullPolicy) String() string {
	if i < 0 || i >= ImagePullPolicy(len(_ImagePullPolicyIndex)-1) {
		return fmt.Sprintf("ImagePullPolicy(%d)", i)
	}
	return _ImagePullPolicyName[_ImagePullPolicyIndex[i]:_ImagePullPolicyIndex[i+1]]
}

var _ImagePullPolicyValues = []ImagePullPolicy{0, 1, 2}

var _ImagePullPolicyNameToValueMap = map[string]ImagePullPolicy{
	_ImagePullPolicyName[0:6]:   0,
	_ImagePullPolicyName[6:18]:  1,
	_ImagePullPolicyName[18:23]: 2,
}

// ImagePullPolicyString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ImagePullPolicyString(s string) (ImagePullPolicy, error) {
	if val, ok := _ImagePullPolicyNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ImagePullPolicy values", s)
}

// ImagePullPolicyValues returns all values of the enum
func ImagePullPolicyValues() []ImagePullPolicy {
	return _ImagePullPolicyValues
}

// IsAImagePullPolicy returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ImagePullPolicy) IsAImagePullPolicy() bool {
	for _, v := range _ImagePullPolicyValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for ImagePullPolicy
func (i ImagePullPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ImagePullPolicy
func (i *ImagePullPolicy) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ImagePullPolicy should be a string, got %s", data)
	}

	var err error
	*i, err = ImagePullPolicyString(s)
	return err
}
