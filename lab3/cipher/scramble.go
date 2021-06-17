package cipher

/*
Task 4: Scrambling text

In this task the objective is to transform a given text such that all letters of each word
are randomly shuffled except for the first and last letter.

For example, given the word "scramble", the result could be "srmacble" or "sbcamrle",
or any other permutation as long as the first and last letters stay the same.

An entire sentence scrambled like this should still be readable:
"it deosn't mttaer in waht oredr the ltteers in a wrod are,
the olny iprmoetnt tihng is taht the frist and lsat ltteer be at the rghit pclae"
See https://www.mrc-cbu.cam.ac.uk/people/matt.davis/cmabridge/ for more
information and examples.

Implementation:
The task is to implement the scramble function, which takes a text in the form of a string and a seed.
A seed is given so the output from your solution should match the test cases if it is correct.
The seed should be applied at the start of the function.
Remember that the implementation should keep any punctuation and spacing intact, and all numbers should be untouched.

Shuffling the letters and applying the seed can be done using the math/rand package (https://golang.org/pkg/math/rand/).
Use the Shuffle function to ensure you reach the same values as given in the tests (scramble_test.go).
*/

func scramble(text string, seed int64) string {
	// you will need this function:
	tokenize(text)
	return ""
}
