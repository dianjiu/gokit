package assertions

import "testing"

func Test_AssertEqual(t *testing.T) {
	t.Parallel()
	AssertEqual(nil, []string{}, []string{})
	AssertEqual(t, []string{}, []string{})
}
