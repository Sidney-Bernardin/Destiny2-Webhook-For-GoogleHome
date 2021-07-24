package destinyhome

import (
	"flag"
	"fmt"
	"root/destinyhome/repository"
	"testing"
)

var (
	operation = flag.String("operation", "", "What operation do you want to test.")

	gamertag = flag.String("gamertag", "", "")
	memType  = flag.String("memType", "", "")
	memID    = flag.String("memID", "", "")

	username      = flag.String("username", "", "")
	guardianIndex = flag.String("guardianIndex", "", "")
	bucket        = flag.String("bucket", "", "")
)

func TestOperation(t *testing.T) {

	// Init repo as a mock.
	repo = repository.NewMock(*gamertag, *memType, *memID)

	// Get the operation.
	operationFunc, ok := operations[*operation]
	if !ok {
		t.Fatalf("operation %s not found", *operation)
	}

	// Setup the params.
	params := map[string]string{
		"username":      *username,
		"guardianIndex": *guardianIndex,
		"bucket":        *bucket,
	}

	// Do the operation.
	res, err := operationFunc(params)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(res)
}
