package stringer

/*
Task 2: Stringers

One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}

A Stringer is a type that can describe itself as a string. The fmt package (and
many others) look for this interface to print values.

Implement the String() method for the Book struct.

A struct

Book{Title: Practical Engineering, Author: Person{FirstName: John, LastName, Doe}, Distributor: Egmont, ReleaseYear: 2018}

should be printed as

"Title: Practical Engineering by Doe, John. Released: 2018, Egmont Distribution."
*/

// Person holds the first name and last name of a person.
type Person struct {
	FirstName string
	LastName  string
}

// Book holds the title, author and other information about a book.
type Book struct {
	Title       string
	Author      Person
	Distributor string
	ReleaseYear int
}

func (b Book) String() string {
	return ""
}
