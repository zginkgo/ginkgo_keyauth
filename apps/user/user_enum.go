// Code generated by github.com/infraboard/ginkgo_mcube
// DO NOT EDIT

package user

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseDescribeByFromString Parse DescribeBy from string
func ParseDescribeByFromString(str string) (DescribeBy, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := DescribeBy_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown DescribeBy: %s", str)
	}

	return DescribeBy(v), nil
}

// Equal type compare
func (t DescribeBy) Equal(target DescribeBy) bool {
	return t == target
}

// IsIn todo
func (t DescribeBy) IsIn(targets ...DescribeBy) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t DescribeBy) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *DescribeBy) UnmarshalJSON(b []byte) error {
	ins, err := ParseDescribeByFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
