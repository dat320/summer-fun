package stringer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var stringerBookTests = []struct {
	in   Book
	want string
}{
	{Book{
		Title: "Practical Engineering",
		Author: Person{
			FirstName: "John",
			LastName:  "Doe",
		},
		Distributor: "Egmont",
		ReleaseYear: 2018,
	}, "Title: Practical Engineering by Doe, John. Released: 2018, Egmont Distribution."},
	{Book{
		Title: "Impractical Engineering",
		Author: Person{
			FirstName: "Leeroy",
			LastName:  "Jenkins",
		},
		Distributor: "Comedic",
		ReleaseYear: 2014,
	}, "Title: Impractical Engineering by Jenkins, Leeroy. Released: 2014, Comedic Distribution."},
}

func TestStringerBook(t *testing.T) {
	for _, test := range stringerBookTests {
		if diff := cmp.Diff(test.want, test.in.String()); diff != "" {
			t.Errorf("String(%q): (-want +got):\n%s", test.in, diff)
		}
	}
}
