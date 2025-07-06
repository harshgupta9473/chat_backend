package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
)

func GenerateChatIDForUsers(user1, user2 string) string {
	users := []string{user1, user2}
	sort.Strings(users)
	hash := sha256.Sum256([]byte(users[0] + ":" + users[1]))
	return hex.EncodeToString(hash[:])
}
