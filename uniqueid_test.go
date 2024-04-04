package uniqueid

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	iterations := 50
	for i := 0; i < iterations; i++ {
		randomPart := uint64(rand.Intn(10000000000))
		encoded := Encode(randomPart, 62)
		decoded, _ := Decode(encoded, 62)
		fmt.Printf("%d: %s -> %d\n", randomPart, encoded, decoded)
		if randomPart != decoded {
			t.Fatalf("Decoded int doesn't match original int %d: %s -> %d", randomPart, encoded, decoded)
		}
	}
}

// TestGenerate ensures that the Generate function returns a unique ID each time it's called
func TestGenerate(t *testing.T) {
	idMap := make(map[string]bool)
	iterations := 1000000
	for i := 0; i < iterations; i++ {
		id := Generate()
		if _, exists := idMap[id]; exists {
			t.Fatalf("ID is not unique: %s", id)
		}
		idMap[id] = true
	}
}

// TestGenerateFormat checks if the generated IDs are in the expected base62 format
func TestGenerateFormat(t *testing.T) {
	id := Generate()

	fmt.Printf("Unix:  %s\n", GenerateUnixTimestampID())
	fmt.Printf("Mili:  %s\n", GenerateUnixMilliTimestampID())
	fmt.Printf("Micro: %s\n", GenerateUnixMicroTimestampID())
	fmt.Printf("Nano:  %s\n", GenerateUnixNanoTimestampID())
	fmt.Printf("ID:    %s\n", id)

	// Very basic format check: length and allowed characters
	// This might need to be adjusted based on the specifics of your ID format
	if len(id) == 0 {
		t.Fatalf("Generated ID is empty")
	}

	for _, char := range id {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			t.Fatalf("Generated ID contains unexpected characters: %s", id)
		}
	}
}
