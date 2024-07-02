package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func UniqueId() string {
	timestamp := time.Now().UnixNano() // Get the current timestamp in nanoseconds
	randomNum := rand.Intn(1000000)    // Generate a random number between 0 and 999999

	// Combine the timestamp and random number to form a unique ID
	uniqueId := strconv.FormatInt(timestamp, 10) + strconv.Itoa(randomNum)
	return uniqueId
}
