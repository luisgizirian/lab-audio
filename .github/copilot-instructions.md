# Lab-Audio Repository

Lab-Audio is a Go-based audio processing experiments repository that contains utilities for generating Euclidean rhythm patterns and synthesizing audio files.

Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.

## Working Effectively

- Check Go version compatibility:
  - `go version` -- requires Go 1.24.3+, tested with go1.24.7
- Bootstrap and build the repository:
  - `cd cmd/euclidgen`
  - `go mod download` -- downloads dependencies (takes <5 seconds)
  - `go build -v` -- builds the application (takes <10 seconds). NEVER CANCEL.
- Run the audio generation application:
  - `./euclidgen` -- generates euclid.wav file (takes <1 second)
- Format and validate Go code:
  - `go fmt ./...` -- formats Go code according to standard
  - `go vet ./...` -- runs Go static analysis
- Clean build artifacts:
  - `rm -f euclidgen euclid.wav` -- removes generated binary and audio file

## Validation

- ALWAYS manually validate any new code by building and running the euclidgen application.
- ALWAYS verify that euclid.wav is generated successfully after running the application.
- The generated audio file should be:
  - A RIFF WAVE audio file
  - 16-bit mono PCM format
  - 44100 Hz sample rate
  - Approximately 705,644 bytes in size
- ALWAYS run `go fmt ./...` and `go vet ./...` before committing changes.
- Test the application produces expected output: "Generated euclid.wav with Euclidean rhythm!"

## Validation Scenarios

After making changes, ALWAYS test the complete audio generation workflow:
1. Navigate to `cmd/euclidgen`
2. Clean any existing artifacts: `rm -f euclidgen euclid.wav`
3. Build the application: `go build`
4. Run the application: `./euclidgen`
5. Verify the output message appears
6. Confirm euclid.wav file exists and has expected properties: `file euclid.wav`
7. Check file size is reasonable: `wc -c euclid.wav`

## Repository Structure

The repository follows this structure:
- `cmd/euclidgen/` - Main Go application for Euclidean rhythm generation
  - `main.go` - Core application logic with bjorklund algorithm and audio synthesis
  - `go.mod` - Go module definition (module name: euclidgen)
  - `go.sum` - Go dependency checksums
  - `euclidgen` - Generated binary (excluded from git)
  - `euclid.wav` - Generated audio file (excluded from git)
- `.devcontainer/` - VS Code development container configuration
- `.github/` - GitHub configuration and workflows
- `.gitignore` - Git ignore patterns (excludes binaries, audio files, build artifacts)
- `README.md` - Basic repository documentation

## Common Tasks

### Building and Running
```bash
cd cmd/euclidgen
go mod download    # Download dependencies
go build          # Build euclidgen binary
./euclidgen        # Generate euclid.wav audio file
```

### Development Workflow
```bash
cd cmd/euclidgen
go fmt ./...       # Format code
go vet ./...       # Static analysis
go build          # Build application
./euclidgen        # Test execution
```

### Repository Information

#### Repository Root Contents
```
.devcontainer/     - Development container config
.git/             - Git metadata
.github/          - GitHub configuration
.gitignore        - Git ignore patterns
README.md         - Repository documentation
cmd/              - Command-line applications
```

#### euclidgen Directory Contents
```
go.mod            - Go module definition (euclidgen)
go.sum            - Dependency checksums
main.go           - Main application (2229 bytes)
euclidgen*        - Generated binary (~2.3MB, excluded from git)
euclid.wav*       - Generated audio (~705KB, excluded from git)
```

#### Key Dependencies
- `github.com/go-audio/audio v1.0.0` - Audio processing library
- `github.com/go-audio/wav v1.1.0` - WAV file format support
- `github.com/go-audio/riff v1.0.0` - RIFF container format (indirect)

## Application Details

The euclidgen application:
- Implements the Bjorklund algorithm for Euclidean rhythm generation
- Synthesizes drum sounds using mathematical functions
- Generates a 16-step pattern with 6 pulses at 120 BPM
- Outputs a mono 16-bit WAV file at 44.1kHz sample rate
- Uses envelope-shaped sine wave bursts for percussion sounds

## Important Notes

- NEVER CANCEL build or run commands - they complete in under 10 seconds
- The application has no command-line arguments or interactive features
- No existing test suite - manual validation is required
- No CI/CD workflows configured
- Build artifacts (euclidgen binary, euclid.wav) are gitignored
- Development container is configured for Go with audio libraries pre-installed
- All operations are extremely fast (<10 seconds) so no special timeout handling needed
- Always test audio file generation as the primary validation scenario