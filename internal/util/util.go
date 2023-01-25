package util

import (
	"github.com/google/uuid"
)

var PageSize int = 10

func GenerateUUId() string {
	return uuid.New().String()
}

func Page(p int) (From int, To int) {
	return (p - 1) * PageSize, p * PageSize
}