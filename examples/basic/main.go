package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pawincpe/golang-key-rotation/pkg/keyrotation"
)

func main() {
	fmt.Println("=== Go Key Rotation Library (Public) - Basic Example ===\n")

	// Check if binary exists
	binaryPath := "../golang-key-rotation-private/build/keyrotation-binary"
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		fmt.Println("❌ Binary not found!")
		fmt.Println("Please build the private project first:")
		fmt.Println("cd ../golang-key-rotation-private && ./build.sh")
		os.Exit(1)
	}

	// Get absolute path to binary
	absPath, err := filepath.Abs(binaryPath)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}

	// Create helper with custom binary path
	helper := keyrotation.NewWithBinaryPath(absPath)
	apiKey := "my-secret-api-key"

	fmt.Println("1. Basic Encryption and Validation:")

	encrypted, err := helper.EncryptApiKey(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("   API Key: %s\n", apiKey)
	fmt.Printf("   Encrypted: %s\n", encrypted)

	isValid, err := helper.ValidateApiKeyToday(apiKey, encrypted)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("   Valid: %t\n\n", isValid)

	fmt.Println("2. Using Specific Date:")
	specificDate := time.Date(2024, 1, 15, 12, 30, 0, 0, time.UTC)

	encryptedWithDate, err := helper.EncryptApiKeyWithDate(apiKey, specificDate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("   Date: %s\n", specificDate.Format("2006-01-02 15:04:05 UTC"))
	fmt.Printf("   Encrypted: %s\n", encryptedWithDate)

	isValidWithDate, err := helper.ValidateApiKey(apiKey, encryptedWithDate, specificDate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("   Valid: %t\n\n", isValidWithDate)

	fmt.Println("3. Using Time Tolerance:")

	encryptedForTolerance, err := helper.EncryptApiKey(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("   Encrypted: %s\n", encryptedForTolerance)

	// Test with different tolerance values
	toleranceValues := []int{0, 1, 5, 10}
	for _, tolerance := range toleranceValues {
		isValid, err := helper.ValidateApiKeyTodayWithTolerance(apiKey, encryptedForTolerance, tolerance)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("   Valid with %d-minute tolerance: %t\n", tolerance, isValid)
	}
	fmt.Println()

	fmt.Println("4. Using Package-level Functions:")

	// Use package-level functions (they will use default binary path)
	encryptedPkg, err := keyrotation.EncryptApiKey(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("   Encrypted: %s\n", encryptedPkg)

	isValidPkg, err := keyrotation.ValidateApiKeyToday(apiKey, encryptedPkg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("   Valid: %t\n\n", isValidPkg)

	fmt.Println("5. Get Date String:")
	dateString := keyrotation.GetDateString(time.Now().UTC())
	fmt.Printf("   Today's date string: %s\n\n", dateString)

	fmt.Println("✅ Example completed successfully!")
	fmt.Println("\n📝 Note: The core logic is protected in the private binary.")
	fmt.Println("   Only the wrapper functions are visible in this public project.")
}
