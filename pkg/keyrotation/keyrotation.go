package keyrotation

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// KeyRotationHelper provides key rotation functionality by calling the private binary
type KeyRotationHelper struct {
	binaryPath string
}

// New creates a new instance of KeyRotationHelper
func New() *KeyRotationHelper {
	return &KeyRotationHelper{
		binaryPath: "./keyrotation-binary", // Default binary name in current directory
	}
}

// NewWithBinaryPath creates a new instance with custom binary path
func NewWithBinaryPath(binaryPath string) *KeyRotationHelper {
	return &KeyRotationHelper{
		binaryPath: binaryPath,
	}
}

// EncryptApiKey encrypts an API key using SHA256 with the current UTC date
func (k *KeyRotationHelper) EncryptApiKey(apiKey string) (string, error) {
	cmd := exec.Command(k.binaryPath, "encrypt", apiKey)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to encrypt API key: %v", err)
	}

	return strings.TrimSpace(out.String()), nil
}

// EncryptApiKeyWithDate encrypts an API key using SHA256 with a specific UTC date
func (k *KeyRotationHelper) EncryptApiKeyWithDate(apiKey string, utcDateTime time.Time) (string, error) {
	dateStr := utcDateTime.Format("2006-01-02")
	cmd := exec.Command(k.binaryPath, "encrypt-date", apiKey, dateStr)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to encrypt API key with date: %v", err)
	}

	return strings.TrimSpace(out.String()), nil
}

// ValidateApiKey validates if an encrypted API key matches the expected hash for a given date
func (k *KeyRotationHelper) ValidateApiKey(apiKey, encryptedKey string, utcDateTime time.Time) (bool, error) {
	dateStr := utcDateTime.Format("2006-01-02")
	cmd := exec.Command(k.binaryPath, "validate-date", apiKey, encryptedKey, dateStr)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("failed to validate API key: %v", err)
	}

	result := strings.TrimSpace(out.String())
	return result == "true", nil
}

// ValidateApiKeyWithTolerance validates if an encrypted API key matches the expected hash for a given date with time tolerance
func (k *KeyRotationHelper) ValidateApiKeyWithTolerance(apiKey, encryptedKey string, utcDateTime time.Time, toleranceMinutes int) (bool, error) {
	// For now, we'll use the base validation since the binary doesn't support tolerance with specific date
	// In a real implementation, you might want to add this functionality to the binary
	return k.ValidateApiKey(apiKey, encryptedKey, utcDateTime)
}

// ValidateApiKeyToday validates if an encrypted API key matches the expected hash for today (UTC)
func (k *KeyRotationHelper) ValidateApiKeyToday(apiKey, encryptedKey string) (bool, error) {
	cmd := exec.Command(k.binaryPath, "validate", apiKey, encryptedKey)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("failed to validate API key for today: %v", err)
	}

	result := strings.TrimSpace(out.String())
	return result == "true", nil
}

// ValidateApiKeyTodayWithTolerance validates if an encrypted API key matches the expected hash for today (UTC) with time tolerance
func (k *KeyRotationHelper) ValidateApiKeyTodayWithTolerance(apiKey, encryptedKey string, toleranceMinutes int) (bool, error) {
	cmd := exec.Command(k.binaryPath, "validate-tolerance", apiKey, encryptedKey, strconv.Itoa(toleranceMinutes))
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("failed to validate API key with tolerance: %v", err)
	}

	result := strings.TrimSpace(out.String())
	return result == "true", nil
}

// GetDateString gets the date string format used for encryption (yyyyMMdd)
func (k *KeyRotationHelper) GetDateString(utcDateTime time.Time) string {
	return utcDateTime.Format("20060102")
}

// Package-level convenience functions

// EncryptApiKey encrypts an API key using SHA256 with the current UTC date
func EncryptApiKey(apiKey string) (string, error) {
	helper := New()
	return helper.EncryptApiKey(apiKey)
}

// EncryptApiKeyWithDate encrypts an API key using SHA256 with a specific UTC date
func EncryptApiKeyWithDate(apiKey string, utcDateTime time.Time) (string, error) {
	helper := New()
	return helper.EncryptApiKeyWithDate(apiKey, utcDateTime)
}

// ValidateApiKey validates if an encrypted API key matches the expected hash for a given date
func ValidateApiKey(apiKey, encryptedKey string, utcDateTime time.Time) (bool, error) {
	helper := New()
	return helper.ValidateApiKey(apiKey, encryptedKey, utcDateTime)
}

// ValidateApiKeyWithTolerance validates if an encrypted API key matches the expected hash for a given date with time tolerance
func ValidateApiKeyWithTolerance(apiKey, encryptedKey string, utcDateTime time.Time, toleranceMinutes int) (bool, error) {
	helper := New()
	return helper.ValidateApiKeyWithTolerance(apiKey, encryptedKey, utcDateTime, toleranceMinutes)
}

// ValidateApiKeyToday validates if an encrypted API key matches the expected hash for today (UTC)
func ValidateApiKeyToday(apiKey, encryptedKey string) (bool, error) {
	helper := New()
	return helper.ValidateApiKeyToday(apiKey, encryptedKey)
}

// ValidateApiKeyTodayWithTolerance validates if an encrypted API key matches the expected hash for today (UTC) with time tolerance
func ValidateApiKeyTodayWithTolerance(apiKey, encryptedKey string, toleranceMinutes int) (bool, error) {
	helper := New()
	return helper.ValidateApiKeyTodayWithTolerance(apiKey, encryptedKey, toleranceMinutes)
}

// GetDateString gets the date string format used for encryption (yyyyMMdd)
func GetDateString(utcDateTime time.Time) string {
	helper := New()
	return helper.GetDateString(utcDateTime)
}
