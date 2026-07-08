package auth

func internalLoginExample() {
	// Files in the same package can call each other without an import prefix.
	LoginWithCredentials("amitjangir", "strong-password")
}
