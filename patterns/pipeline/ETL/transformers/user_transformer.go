package transformer

import (
	"ETL/data"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// CleanAndHashUser cleans user data and hashes UserID.
func CleanAndHashUser(user *data.User) *data.User {
	// Trim leading and trailing spaces from UserID
	user.CustomerId = strings.TrimSpace(user.CustomerId)

	// Hash UserID for privacy
	hasher := sha256.New()
	hasher.Write([]byte(user.CustomerId))
	user.CustomerId = hex.EncodeToString(hasher.Sum(nil))

	return user
}
