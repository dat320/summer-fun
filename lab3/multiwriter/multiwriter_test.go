package multiwriter

import (
	"dat320/lab3/errors"
	"io"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var writeTests = []struct {
	inb     []byte
	inw     []io.Writer
	wantN   []int
	wantErr errors.Errors
}{
	{nil, nil, []int{}, nil},
	{nil, []io.Writer{}, []int{}, nil},
	{nil, []io.Writer{ioutil.Discard}, []int{0}, nil},
	{[]byte(""), []io.Writer{ioutil.Discard}, []int{0}, nil},
	{[]byte("\n"), []io.Writer{ioutil.Discard}, []int{1}, nil},
	{[]byte("DAT320-1\n"), []io.Writer{ioutil.Discard}, []int{9}, nil},
	{[]byte("DAT320-2\n"), []io.Writer{ioutil.Discard}, []int{9}, nil},
	{[]byte("DAT320-3\n"), []io.Writer{ioutil.Discard, ioutil.Discard}, []int{9, 9}, nil},
	{[]byte("DAT320-4\n"), []io.Writer{ioutil.Discard, ioutil.Discard, ioutil.Discard}, []int{9, 9, 9}, nil},
	{
		[]byte("DAT320-5\n"), []io.Writer{ioutil.Discard,
			failureWriter(2),
			ioutil.Discard,
		},
		[]int{9, 2, 9},
		[]error{nil, io.ErrShortWrite, nil},
	},
	{
		[]byte("DAT320-6\n"), []io.Writer{ioutil.Discard,
			failureWriter(2),
			failureWriter(6),
		},
		[]int{9, 2, 6},
		[]error{nil, io.ErrShortWrite, io.ErrShortWrite},
	},
}

func TestWriters(t *testing.T) {
	for _, test := range writeTests {
		n, errs := WriteTo(test.inb, test.inw...)
		if !cmp.Equal(n, test.wantN) {
			t.Errorf("WriteTo(%q, writers...) = %v, want %v", test.inb, n, test.wantN)
		}
		if diff := cmp.Diff(test.wantErr, errs, cmp.Comparer(errorsComparer)); diff != "" {
			t.Errorf("Error(): (-want +got):\n%s", diff)
		}
	}
}
