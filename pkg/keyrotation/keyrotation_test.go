package keyrotation

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestKeyRotationHelper_WithBinaryPath(t *testing.T) {
	// Skip if binary doesn't exist
	binaryPath := "../golang-key-rotation-private/build/keyrotation-binary"
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		t.Skip("Binary not found, skipping test")
	}

	// Use absolute path
	absPath, err := filepath.Abs(binaryPath)
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	helper := NewWithBinaryPath(absPath)
	testApiKey := "testApiKey123"

	// Test encryption
	encrypted, err := helper.EncryptApiKey(testApiKey)
	if err != nil {
		t.Fatalf("EncryptApiKey failed: %v", err)
	}

	if encrypted == "" {
		t.Error("Expected non-empty encrypted result")
	}

	// Test validation
	isValid, err := helper.ValidateApiKeyToday(testApiKey, encrypted)
	if err != nil {
		t.Fatalf("ValidateApiKeyToday failed: %v", err)
	}

	if !isValid {
		t.Error("Expected validation to succeed")
	}
}

func TestKeyRotationHelper_WithSpecificDate(t *testing.T) {
	// Skip if binary doesn't exist
	binaryPath := "../golang-key-rotation-private/build/keyrotation-binary"
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		t.Skip("Binary not found, skipping test")
	}

	absPath, err := filepath.Abs(binaryPath)
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	helper := NewWithBinaryPath(absPath)
	testApiKey := "testApiKey123"
	testDate := time.Date(2024, 1, 15, 12, 30, 0, 0, time.UTC)

	// Test encryption with specific date
	encrypted, err := helper.EncryptApiKeyWithDate(testApiKey, testDate)
	if err != nil {
		t.Fatalf("EncryptApiKeyWithDate failed: %v", err)
	}

	if encrypted == "" {
		t.Error("Expected non-empty encrypted result")
	}

	// Test validation with specific date
	isValid, err := helper.ValidateApiKey(testApiKey, encrypted, testDate)
	if err != nil {
		t.Fatalf("ValidateApiKey failed: %v", err)
	}

	if !isValid {
		t.Error("Expected validation to succeed")
	}
}

func TestKeyRotationHelper_WithTolerance(t *testing.T) {
	// Skip if binary doesn't exist
	binaryPath := "../golang-key-rotation-private/build/keyrotation-binary"
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		t.Skip("Binary not found, skipping test")
	}

	absPath, err := filepath.Abs(binaryPath)
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	helper := NewWithBinaryPath(absPath)
	testApiKey := "testApiKey123"

	// Test encryption
	encrypted, err := helper.EncryptApiKey(testApiKey)
	if err != nil {
		t.Fatalf("EncryptApiKey failed: %v", err)
	}

	// Test validation with tolerance
	isValid, err := helper.ValidateApiKeyTodayWithTolerance(testApiKey, encrypted, 5)
	if err != nil {
		t.Fatalf("ValidateApiKeyTodayWithTolerance failed: %v", err)
	}

	if !isValid {
		t.Error("Expected validation to succeed with tolerance")
	}
}

func TestKeyRotationHelper_GetDateString(t *testing.T) {
	helper := New()
	testDate := time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC)

	result := helper.GetDateString(testDate)
	expected := "20240115"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// Package-level function tests

func TestEncryptApiKey(t *testing.T) {
	// Skip if binary doesn't exist
	binaryPath := "../golang-key-rotation-private/build/keyrotation-binary"
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		t.Skip("Binary not found, skipping test")
	}

	// Set custom binary path for testing
	originalHelper := New()
	originalHelper.binaryPath = filepath.Join("..", "golang-key-rotation-private", "build", "keyrotation-binary")

	testApiKey := "testApiKey123"

	result, err := originalHelper.EncryptApiKey(testApiKey)
	if err != nil {
		t.Fatalf("EncryptApiKey failed: %v", err)
	}

	if result == "" {
		t.Error("Expected non-empty result")
	}
}

func TestValidateApiKeyToday(t *testing.T) {
	// Skip if binary doesn't exist
	binaryPath := "../golang-key-rotation-private/build/keyrotation-binary"
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		t.Skip("Binary not found, skipping test")
	}

	// Set custom binary path for testing
	originalHelper := New()
	originalHelper.binaryPath = filepath.Join("..", "golang-key-rotation-private", "build", "keyrotation-binary")

	testApiKey := "testApiKey123"

	encrypted, err := originalHelper.EncryptApiKey(testApiKey)
	if err != nil {
		t.Fatalf("EncryptApiKey failed: %v", err)
	}

	result, err := originalHelper.ValidateApiKeyToday(testApiKey, encrypted)
	if err != nil {
		t.Fatalf("ValidateApiKeyToday failed: %v", err)
	}

	if !result {
		t.Error("Expected validation to succeed")
	}
}

func TestGetDateString(t *testing.T) {
	testDate := time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC)

	result := GetDateString(testDate)
	expected := "20240115"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
