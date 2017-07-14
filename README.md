# go-gender

Package gender provides a few functions to validate a persons gender from a string.

[![GoDoc][1]][2]

[1]: https://godoc.org/github.com/abrander/go-gender?status.svg
[2]: https://godoc.org/github.com/abrander/go-gender

Code Examples
-------------

Validate a gender entered by an end-user.

```go
import "github.com/abrander/go-gender"


func ValidateUser(user User) error {
...
  err := gender.Validate(user.Gender)
  if err != nil {
    return err
  }
...
}

```
