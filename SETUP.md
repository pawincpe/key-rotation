# Setup Guide - Two-Project Architecture

This guide explains how to set up and use the two-project architecture for protecting the core encryption logic.

## 🏗️ Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                    PUBLIC PROJECT                          │
│  ┌─────────────────────────────────────────────────────┐   │
│  │              golang-key-rotation-public             │   │
│  │  ┌─────────────────────────────────────────────┐   │   │
│  │  │           pkg/keyrotation/                  │   │   │
│  │  │  ┌─────────────────────────────────────┐   │   │   │
│  │  │  │         Wrapper Functions           │   │   │   │
│  │  │  │  - EncryptApiKey()                  │   │   │   │
│  │  │  │  - ValidateApiKey()                 │   │   │   │
│  │  │  │  - Calls private binary             │   │   │   │
│  │  │  └─────────────────────────────────────┘   │   │   │
│  │  └─────────────────────────────────────────────┘   │   │
│  └─────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                              │
                              │ calls
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   PRIVATE PROJECT                          │
│  ┌─────────────────────────────────────────────────────┐   │
│  │            golang-key-rotation-private              │   │
│  │  ┌─────────────────────────────────────────────┐   │   │
│  │  │           pkg/core/                        │   │   │
│  │  │  ┌─────────────────────────────────────┐   │   │   │
│  │  │  │         Core Logic                  │   │   │   │
│  │  │  │  - SHA256 encryption                │   │   │   │
│  │  │  │  - Date-based validation            │   │   │   │
│  │  │  │  - Time tolerance logic             │   │   │   │
│  │  │  └─────────────────────────────────────┘   │   │   │
│  │  └─────────────────────────────────────────────┘   │   │
│  │  ┌─────────────────────────────────────────────┐   │   │
│  │  │           cmd/build/                       │   │   │
│  │  │  ┌─────────────────────────────────────┐   │   │   │
│  │  │  │         Binary Builder              │   │   │   │
│  │  │  │  - Compiles core logic              │   │   │
│  │  │  │  - Creates executable               │   │   │   │
│  │  │  └─────────────────────────────────────┘   │   │   │
│  │  └─────────────────────────────────────────────┘   │   │
│  └─────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

## 📋 Prerequisites

- Go 1.16 or later
- Git
- Access to both repositories

## 🚀 Step-by-Step Setup

### Step 1: Clone Both Projects

```bash
# Clone private project (keep this private!)
git clone <private-repo-url> golang-key-rotation-private
cd golang-key-rotation-private

# Clone public project
cd ..
git clone <public-repo-url> golang-key-rotation-public
```

### Step 2: Build Private Binary

```bash
cd golang-key-rotation-private

# Run build script
./build.sh

# Verify binary was created
ls -la build/keyrotation-binary
```

Expected output:
```
[INFO] Build completed successfully!
[INFO] Binary location: build/keyrotation-binary
[INFO] Binary size: 2.8M
```

### Step 3: Set Up Public Project

```bash
cd ../golang-key-rotation-public

# Copy binary to public project
cp ../golang-key-rotation-private/build/keyrotation-binary .

# Make binary executable
chmod +x keyrotation-binary
```

### Step 4: Test the Setup

```bash
# Test the binary directly
./keyrotation-binary encrypt "test-key"

# Test the public library
go test ./pkg/keyrotation

# Run example
go run examples/basic/main.go
```

## 🔧 Configuration Options

### Option 1: Local Binary (Recommended for Development)

```bash
# Copy binary to public project directory
cp ../golang-key-rotation-private/build/keyrotation-binary .

# Use default wrapper (looks for ./keyrotation-binary)
helper := keyrotation.New()
```

### Option 2: System-wide Binary

```bash
# Copy to system PATH (requires sudo)
sudo cp ../golang-key-rotation-private/build/keyrotation-binary /usr/local/bin/

# Use default wrapper
helper := keyrotation.New()
```

### Option 3: Custom Binary Path

```go
// Use custom path
helper := keyrotation.NewWithBinaryPath("/path/to/keyrotation-binary")
```

## 🧪 Testing

### Test Private Project

```bash
cd golang-key-rotation-private

# Run core tests
go test ./pkg/core

# Test binary
./build/keyrotation-binary encrypt "test-key"
./build/keyrotation-binary validate "test-key" "encrypted-hash"
```

### Test Public Project

```bash
cd golang-key-rotation-public

# Run wrapper tests
go test ./pkg/keyrotation

# Run with binary available
go test -v ./pkg/keyrotation
```

## 📦 Deployment

### Development Environment

```bash
# 1. Build private binary
cd golang-key-rotation-private
./build.sh

# 2. Copy to public project
cp build/keyrotation-binary ../golang-key-rotation-public/

# 3. Use in your application
cd ../golang-key-rotation-public
go mod tidy
go run your-app.go
```

### Production Environment

```bash
# 1. Build private binary on secure machine
cd golang-key-rotation-private
./build.sh

# 2. Copy binary to production server
scp build/keyrotation-binary user@server:/usr/local/bin/

# 3. Install public library
go get github.com/pawincpe/golang-key-rotation

# 4. Use in production code
helper := keyrotation.New() // Will find binary in PATH
```

### Docker Deployment

```dockerfile
# Dockerfile for your application
FROM golang:1.21-alpine

# Copy private binary
COPY keyrotation-binary /usr/local/bin/
RUN chmod +x /usr/local/bin/keyrotation-binary

# Install public library
RUN go get github.com/pawincpe/golang-key-rotation

# Copy your application
COPY . /app
WORKDIR /app

# Build and run
RUN go build -o main .
CMD ["./main"]
```

## 🔒 Security Considerations

### Private Project Security

1. **Repository Access**: Keep private repository access limited
2. **Binary Distribution**: Distribute binary securely
3. **Code Review**: Review all changes to core logic
4. **Version Control**: Tag releases for binary distribution

### Public Project Security

1. **No Core Logic**: Public project contains no encryption logic
2. **Wrapper Only**: Only wrapper functions are visible
3. **Binary Calls**: All encryption happens in external binary
4. **Error Handling**: Secure error messages (no information leakage)

## 🐛 Troubleshooting

### Binary Not Found

```bash
# Error: exec: "keyrotation-binary": executable file not found in $PATH

# Solutions:
# 1. Copy binary to current directory
cp ../golang-key-rotation-private/build/keyrotation-binary .

# 2. Use absolute path
helper := keyrotation.NewWithBinaryPath("/full/path/to/keyrotation-binary")

# 3. Add to PATH
export PATH=$PATH:/path/to/binary/directory
```

### Permission Denied

```bash
# Error: permission denied

# Solution: Make executable
chmod +x keyrotation-binary
```

### Build Failures

```bash
# Private project build fails
cd golang-key-rotation-private
go clean -cache -testcache
./build.sh

# Public project tests fail
cd ../golang-key-rotation-public
go clean -cache -testcache
go test ./pkg/keyrotation
```

### Version Mismatch

```bash
# If binary and wrapper are incompatible
# 1. Rebuild private project
cd golang-key-rotation-private
./build.sh

# 2. Update binary in public project
cp build/keyrotation-binary ../golang-key-rotation-public/
```

## 📚 Best Practices

### Development Workflow

1. **Private Changes**: Make changes to core logic in private project
2. **Test Core**: Test changes in private project first
3. **Build Binary**: Build new binary version
4. **Update Public**: Copy new binary to public project
5. **Test Integration**: Test public project with new binary

### Version Management

```bash
# Tag private project releases
cd golang-key-rotation-private
git tag v1.0.0
git push origin v1.0.0

# Tag public project releases
cd ../golang-key-rotation-public
git tag v1.0.0
git push origin v1.0.0
```

### Binary Distribution

```bash
# Create release package
cd golang-key-rotation-private
tar -czf keyrotation-binary-v1.0.0.tar.gz build/keyrotation-binary

# Distribute securely
# - Use signed releases
# - Verify checksums
# - Secure download links
```

## 🎯 Success Criteria

Your setup is successful when:

- ✅ Private project builds without errors
- ✅ Binary is created and executable
- ✅ Public project tests pass
- ✅ Examples run successfully
- ✅ No core logic is visible in public project
- ✅ Encryption/validation works correctly

## 📞 Support

For setup issues:

1. Check this guide first
2. Verify all prerequisites are met
3. Check binary permissions and paths
4. Review error messages carefully
5. Contact support with specific error details

---

**🔒 Your encryption logic is now protected and ready to use!** 🎉
