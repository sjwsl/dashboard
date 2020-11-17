package util

import (
	"fmt"
	"testing"
)

func TestIdCompletenessProof(t *testing.T) {
	simples := []struct {
		in1    int
		in2    []int
		expect error
	}{
		{
			in1:    0,
			in2:    []int{},
			expect: nil,
		},
		{
			in1:    1,
			in2:    []int{3},
			expect: nil,
		},
		{
			in1:    5,
			in2:    []int{1, 2, 3, 4, 5},
			expect: nil,
		},
		{
			in1:    5,
			in2:    []int{1, 2, 3, 3, 5},
			expect: fmt.Errorf("fatal : duplicate id 3. "),
		},
	}

	for _, simple := range simples {
		out := IdCompletenessProof(simple.in1, simple.in2)
		if out != simple.expect {
			if out != nil && simple.expect != nil && out.Error() != simple.expect.Error() {
				panic(fmt.Errorf("fatal : expect: (%d, %d ; out :%v) real : (out :%v)",
					simple.in1, simple.in2, simple.expect, out))
			}
		}
	}
}
