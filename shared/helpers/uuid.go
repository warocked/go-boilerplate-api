package helpers

import (
	"github.com/google/uuid"
)

// GenerateUUID generates a new UUID v4
func GenerateUUID() string {
	return uuid.New().String()
}

// GenerateUUIDBytes generates a new UUID v4 as bytes
func GenerateUUIDBytes() []byte {
	id := uuid.New()
	return id[:]
}

// ParseUUID parses a UUID string and returns an error if invalid
func ParseUUID(s string) error {
	_, err := uuid.Parse(s)
	return err
}

// IsValidUUID checks if a string is a valid UUID
func IsValidUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
