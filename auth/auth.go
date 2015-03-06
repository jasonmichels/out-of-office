package auth

import (
    "fmt"
)

type Auth struct {    

}

func (auth *Auth) Attempt(username string, password string) (bool) {
    // validate the username

    // try to pull the user object from the database by username

    // if we get a user validate that the passwords match

    // if they match return true

    // else return false
    fmt.Println(username)
    fmt.Println(password)

    return false
}