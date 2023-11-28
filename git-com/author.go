//
// git-com/author.go:
//

package main

// Author is the individual contributor
type Author struct {
	// Id is the author's short id.
	// This is used to select members acting as pair programmers.
	// This is 'self' for the author of the current user.
	Id string `yaml:"id"`
	// Name is the author's name.
	Name string `yaml:"name"`
	// Email is the author's email.
	Email string `yaml:"email"`
}

// NewAuthor creates a new author.
func NewAuthor(id, name, email string) Author {
	return Author{
		Id:    id,
		Name:  name,
		Email: email,
	}
}

// Equals returns true if the author is equal to the provided author.
func (a Author) Equals(author Author) bool {
	return a.Email == author.Email
}

// String returns the string representation of the author.
func (a Author) String() string {
	return a.Name + " <" + a.Email + ">"
}

// Authors is a list of authors.
type Authors []Author

// NewAuthors creates a new list of authors.
func NewAuthors() Authors {
	return []Author{}
}

// Add adds an author to the list.
func (a *Authors) Add(author Author) {
	*a = append(*a, author)
}

// Exists returns true if the author exists in the list.
func (a Authors) Exists(author Author) bool {
	for _, a := range a {
		if a.Equals(author) {
			return true
		}
	}
	return false
}

// Find returns the author with the specified id.
func (a Authors) Find(id string) *Author {
	for _, author := range a {
		if author.Id == id {
			return &author
		}
	}
	return nil
}

// FindByEmail returns the author with the specified email.
func (a Authors) FindByEmail(email string) *Author {
	for _, author := range a {
		if author.Email == email {
			return &author
		}
	}
	return nil
}
