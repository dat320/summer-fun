package sequence

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var triangularTests = []struct {
	in, want uint
}{
	{0, 0},
	{1, 1},
	{2, 3},
	{3, 6},
	{4, 10},
	{5, 15},
	{6, 21},
	{7, 28},
	{8, 36},
	{9, 45},
	{10, 55},
	{20, 210},
}

func TestTriangular(t *testing.T) {
	for _, test := range triangularTests {
		if diff := cmp.Diff(test.want, triangular(test.in)); diff != "" {
			t.Errorf("triangular(%d): (-want +got):\n%s", test.in, diff)
		}
	}
}
