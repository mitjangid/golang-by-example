package auth

import "fmt"

// loginWithCredentials is private because it starts with a lowercase letter.
func loginWithCredentials(username, password string) {
	fmt.Println("Login request for:", username)
	fmt.Println("Password length:", len(password))
}

// LoginWithCredentials is exported because it starts with a capital letter.
func LoginWithCredentials(username, password string) {
	loginWithCredentials(username, password)
}
