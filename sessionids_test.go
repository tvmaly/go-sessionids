package sessionids

import (
	"fmt"
	"testing"
)

func TestGenerateSessionId(t *testing.T) {

	session := GenerateSessionId(40)

	if session == "" {
		t.Fatalf("Error empty string returned")
	}

	length := len(session)

	fmt.Printf("GenerateSessionId returned lenght: %v and value: %s\n", length, session)

}
