package utils

import (
	"fmt"
	"time"
)

func NewDateNow() string {
	now := time.Now()
	y, m, d := now.Date()
	return fmt.Sprintf("%v-%v-%v", y, int(m), d)
}