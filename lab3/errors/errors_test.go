package errors

import (
	. "io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var errTests = []struct {
	in   Errors
	want string
}{
	{nil, "(0 errors)"},
	{[]error{}, "(0 errors)"},
	{[]error{ErrClosedPipe}, wantClosePipe},
	{[]error{ErrClosedPipe, ErrClosedPipe}, wantClosePipe + " (and 1 other error)"},
	{[]error{ErrClosedPipe, ErrClosedPipe, ErrClosedPipe}, wantClosePipe + " (and 2 other errors)"},
	{[]error{nil}, "(0 errors)"},
	{[]error{ErrClosedPipe, nil}, wantClosePipe},
	{[]error{nil, ErrClosedPipe}, wantClosePipe},
	{[]error{ErrClosedPipe, ErrClosedPipe, nil}, wantClosePipe + " (and 1 other error)"},
	{[]error{ErrClosedPipe, ErrClosedPipe, nil, nil}, wantClosePipe + " (and 1 other error)"},
	{[]error{ErrClosedPipe, ErrClosedPipe, nil, nil, nil}, wantClosePipe + " (and 1 other error)"},
	{[]error{nil, ErrClosedPipe, nil, ErrClosedPipe}, wantClosePipe + " (and 1 other error)"},
	{[]error{nil, nil, ErrClosedPipe, ErrClosedPipe}, wantClosePipe + " (and 1 other error)"},
	{[]error{ErrClosedPipe, nil, nil, ErrClosedPipe, ErrClosedPipe}, wantClosePipe + " (and 2 other errors)"},
	{[]error{ErrClosedPipe, ErrClosedPipe, nil, ErrClosedPipe, nil}, wantClosePipe + " (and 2 other errors)"},
	{[]error{nil, nil, nil, nil, ErrClosedPipe, ErrClosedPipe, ErrClosedPipe}, wantClosePipe + " (and 2 other errors)"},
}

func TestErrors(t *testing.T) {
	for i, test := range errTests {
		if diff := cmp.Diff(test.want, test.in.Error()); diff != "" {
			t.Errorf("Test %2d: Error(): (-want +got):\n%s", i, diff)
		}
	}
}
