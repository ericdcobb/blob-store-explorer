package explore

import (
	"fmt"
	"strconv"
	"time"
)

// IsBefore ...
func IsBefore(userInput string, blobStoreTimeStamp string) bool {
	userInputTime, tm := convertTimes(userInput, blobStoreTimeStamp)

	return tm.Before(userInputTime)
}

// IsAfter ...
func IsAfter(userInput string, blobStoreTimeStamp string) bool {
	userInputTime, tm := convertTimes(userInput, blobStoreTimeStamp)

	return tm.After(userInputTime)
}

func convertTimes(userInput string, blobStoreTimeStamp string) (time.Time, time.Time) {
	blobStoreCreatedTime := BlobCreatedTime(blobStoreTimeStamp)

	userInputTime, userInputErr := time.Parse(time.RFC3339, userInput)
	if userInputErr != nil {

		printTimeExample()
		panic(userInputErr)
	}
	return userInputTime, blobStoreCreatedTime
}

// BlobCreatedTime ...
func BlobCreatedTime(blobStoreTimeStamp string) time.Time {
	i, err := strconv.ParseInt(blobStoreTimeStamp, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(i/1000, 0)
}

func printTimeExample() {
	fmt.Println("Time Should be in RFC3339 format:")
	fmt.Println(time.Now().Format(time.RFC3339))
}
