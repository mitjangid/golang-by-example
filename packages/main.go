package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/mitjangid/golang-by-example/packages/auth"
	"github.com/mitjangid/golang-by-example/packages/user"
)

func main() {
	start := time.Now()

	// Import local packages with the module path from go.mod.
	auth.LoginWithCredentials("amitjangir", "strong-password")
	session := auth.NewSession("amitjangir")
	fmt.Println(session)

	currentUser := user.User{
		Email: "mitjangid@gmail.com",
		Name:  "Amit K. Jangir",
		Phone: "+91-9876543210",
		ID:    108,
	}
	fmt.Println(currentUser)

	// Third-party packages are downloaded with go get and tracked in go.mod.
	// Example: go get github.com/fatih/color
	color.Cyan("Prints text in cyan.")
	color.Blue("Prints %s in blue.", "text")
	color.Red("We have red")
	color.Magenta("And many others...")

	// go.mod records module requirements. go.sum stores checksums that verify
	// downloaded dependencies. Run go mod tidy after changing imports.

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
