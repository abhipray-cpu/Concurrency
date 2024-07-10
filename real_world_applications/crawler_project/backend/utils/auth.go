package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash of the input password.
//
// Parameters:
// - password: The plaintext password to hash.
//
// Returns:
// - A hashed version of the input password.
// - An error if the hashing process fails.
func HashPassword(password string) (string, error) {
	// GenerateFromPassword returns a hashed password from the given password string.
	// The cost parameter controls the complexity of the hashing process.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Return an empty string and the error if hashing fails.
		return "", err
	}
	// Return the hashed password as a string.
	return string(hashedPassword), nil
}

// ComparePasswords checks if a plaintext password matches a bcrypt hashed password.
//
// Parameters:
// - hashedPassword: The bcrypt hashed password.
// - password: The plaintext password to compare.
//
// Returns:
// - nil if the passwords match.
// - An error if the passwords don't match or if there's another error.
func ComparePasswords(hashedPassword, password string) error {
	// CompareHashAndPassword compares a bcrypt hashed password with its possible
	// plaintext equivalent. Returns nil on success, or an error on failure.
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// If the passwords don't match, or if there's another error, return the error.
		return err
	}
	// If the passwords match, return nil.
	return nil
}

// GenerateJWT creates a JWT (JSON Web Token) for a given user.
// This token can be used for authenticating API requests.
//
// Parameters:
// - userID: The user's ID.
// - userName: The user's name.
//
// Returns:
// - A signed JWT string.
// - An error if the JWT signing process fails.
func GenerateJWT(userID string, userName string) (string, error) {
	var mySigningKey = []byte("secret")      // Use a secret from your environment.
	token := jwt.New(jwt.SigningMethodHS256) // Create a new JWT token using HS256 signing method.
	claims := token.Claims.(jwt.MapClaims)   // Cast the token's claims to a MapClaims object.

	// Set claims for the JWT. These claims include the user's ID, name, and an expiration time.
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["user_name"] = userName
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix() // Set expiration to 30 days from now.

	// Sign the token using the specified secret key.
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
