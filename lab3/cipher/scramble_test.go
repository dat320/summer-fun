package cipher

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var scrambleTests = []struct {
	in, want string
}{
	{"", ""},
	{"The quick brown fox jumps over the lazy dog", "The quick bowrn fox jmups over the lzay dog"},
	{"it doesn't matter in what order the letters in a word are, the only important thing is that the first and last letter be at the right place.", "it deson't metatr in waht oredr the lettres in a wrod are, the only inmproatt thing is taht the frist and last letetr be at the right plcae."},
	{"1234567890 0987654321", "1234567890 0987654321"},
	{"Testing punctuation, (parentheses), and spacing   ...", "Tsteing pioutacuntn, (paerhnsetes), and spncaig   ..."},
	{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus vestibulum vestibulum elementum. Suspendisse at tellus eget magna faucibus molestie. Aliquam erat volutpat. Vestibulum mattis mauris nisi, eu dapibus diam pretium vitae.",
		"Lorem isupm dloor sit amet, csttunoecer andiipcsig elit. Vmiuvas vsetiuublm viebtlusum eenuetlmm. Sessndupsie at tuells eget mgnaa fubiacus moetsile. Auqialm erat vulpaott. Vutsuilebm mtatis mriuas nsii, eu dbuaips daim priutem viate."},
}

func TestScramble(t *testing.T) {
	for _, test := range scrambleTests {
		got := scramble(test.in, 100)
		if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("scramble(%q): (-want +got):\n%s", test.in, diff)
		}
	}
}
