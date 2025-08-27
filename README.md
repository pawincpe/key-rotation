# Go Key Rotation Library (Public)

A secure key rotation library for Go with SHA256-based API key encryption using UTC timezone. This is the public wrapper that calls the protected binary implementation.

## 🔒 Security Architecture

This library uses a **two-project architecture** to protect the core encryption logic:

- **Public Project** (this repository): Contains the wrapper functions and public API
- **Private Project**: Contains the core logic compiled into a binary

The core encryption and validation logic is **never exposed** in the public repository.

## Project Structure

```
golang-key-rotation-public/
├── pkg/keyrotation/              # Public wrapper package
│   ├── keyrotation.go            # Wrapper implementation
│   └── keyrotation_test.go       # Tests
├── examples/                     # Usage examples
├── docs/                         # Documentation
├── scripts/                      # Build scripts
├── go.mod                        # Go module file
└── README.md                     # This file
```

## Quick Start

### Prerequisites

1. **Build the Private Binary**: First, you need to build the private project binary
2. **Place Binary**: Copy the binary to your project or set the path

### Installation

```bash
go get github.com/pawincpe/golang-key-rotation
```

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/pawincpe/golang-key-rotation/pkg/keyrotation"
)

func main() {
    // Your API key
    apiKey := "my-secret-api-key"
    
    // Encrypt the API key
    encrypted, err := keyrotation.EncryptApiKey(apiKey)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Encrypted: %s\n", encrypted)
    
    // Validate the encrypted key
    isValid, err := keyrotation.ValidateApiKeyToday(apiKey, encrypted)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Valid: %t\n", isValid)
}
```

## API Reference

### Package-level Functions

```go
// Encrypt API key with current UTC date
func EncryptApiKey(apiKey string) (string, error)

// Encrypt API key with specific UTC date
func EncryptApiKeyWithDate(apiKey string, utcDateTime time.Time) (string, error)

// Validate API key for specific date
func ValidateApiKey(apiKey, encryptedKey string, utcDateTime time.Time) (bool, error)

// Validate API key with time tolerance
func ValidateApiKeyWithTolerance(apiKey, encryptedKey string, utcDateTime time.Time, toleranceMinutes int) (bool, error)

// Validate API key for today (UTC)
func ValidateApiKeyToday(apiKey, encryptedKey string) (bool, error)

// Validate API key for today with time tolerance
func ValidateApiKeyTodayWithTolerance(apiKey, encryptedKey string, toleranceMinutes int) (bool, error)

// Get date string format used for encryption (yyyyMMdd)
func GetDateString(utcDateTime time.Time) string
```

### Struct-based API

```go
// Create a new KeyRotationHelper instance
helper := keyrotation.New()

// Create with custom binary path
helper := keyrotation.NewWithBinaryPath("/path/to/keyrotation-binary")

// Use instance methods
encrypted, err := helper.EncryptApiKey(apiKey)
isValid, err := helper.ValidateApiKeyTodayWithTolerance(apiKey, encrypted, 5)
```

## Examples

### Basic Encryption and Validation

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/pawincpe/golang-key-rotation/pkg/keyrotation"
)

func main() {
    apiKey := "my-secret-api-key"
    
    // Encrypt with current date
    encrypted, err := keyrotation.EncryptApiKey(apiKey)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Encrypted: %s\n", encrypted)
    
    // Validate
    isValid, err := keyrotation.ValidateApiKeyToday(apiKey, encrypted)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Valid: %t\n", isValid)
}
```

### Using Custom Binary Path

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/pawincpe/golang-key-rotation/pkg/keyrotation"
)

func main() {
    // Create helper with custom binary path
    helper := keyrotation.NewWithBinaryPath("/path/to/keyrotation-binary")
    
    apiKey := "my-secret-api-key"
    
    encrypted, err := helper.EncryptApiKey(apiKey)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Encrypted: %s\n", encrypted)
}
```

## Security Features

- ✅ **Protected Core Logic**: Encryption logic is compiled into binary
- ✅ **SHA256 Encryption**: Uses secure SHA256 algorithm
- ✅ **UTC Timezone**: Consistent across all timezones
- ✅ **Time Tolerance**: Handle clock differences (±minutes)
- ✅ **Thread Safe**: Safe for concurrent use
- ✅ **Go Idiomatic**: Follows Go best practices

## Setup Instructions

### 1. Build Private Binary

```bash
# Clone and build the private project
git clone <private-repo-url> golang-key-rotation-private
cd golang-key-rotation-private
./build.sh
```

### 2. Use in Your Project

```bash
# Install the public library
go get github.com/pawincpe/golang-key-rotation

# Copy binary to your project or set path
cp golang-key-rotation-private/build/keyrotation-binary /usr/local/bin/
```

### 3. Run Examples

```bash
# Run the example
cd golang-key-rotation-public/examples/basic
go run main.go
```

## Testing

### Run Tests

```bash
go test ./pkg/keyrotation
```

**Note**: Tests require the private binary to be available.

### Test with Custom Binary Path

```go
helper := keyrotation.NewWithBinaryPath("/path/to/keyrotation-binary")
// Use helper for testing
```

## Binary Commands

The private binary supports these commands:

```bash
# Encrypt API key with current date
keyrotation-binary encrypt <apikey>

# Encrypt API key with specific date
keyrotation-binary encrypt-date <apikey> <date>

# Validate API key for today
keyrotation-binary validate <apikey> <encrypted>

# Validate API key for specific date
keyrotation-binary validate-date <apikey> <encrypted> <date>

# Validate with tolerance
keyrotation-binary validate-tolerance <apikey> <encrypted> <tolerance>
```

## Migration from .NET

If migrating from the .NET KeyRotation library:

1. **Same Logic**: Identical encryption formula and behavior
2. **Same Date Format**: Uses `yyyyMMdd` format
3. **Same Tolerance**: Supports time tolerance functionality
4. **Protected Implementation**: Core logic is now protected in binary

## Troubleshooting

### Binary Not Found

```bash
# Error: exec: "keyrotation-binary": executable file not found in $PATH

# Solution: Set custom binary path
helper := keyrotation.NewWithBinaryPath("/full/path/to/keyrotation-binary")
```

### Permission Denied

```bash
# Error: permission denied

# Solution: Make binary executable
chmod +x /path/to/keyrotation-binary
```

### Build Issues

```bash
# Ensure private project is built first
cd golang-key-rotation-private
./build.sh
```

## Support

For issues, questions, or feature requests:

- [GitHub Issues](https://github.com/pawincpe/golang-key-rotation/issues)
- [Documentation](docs/README.md)
- [Examples](examples/)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**🔒 Your encryption logic is now protected!** 🎉
