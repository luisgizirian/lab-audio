# Euclidean Rhythm Generator (euclidgen)

A Go program that generates Euclidean rhythm patterns and synthesizes them into audio files.

## What are Euclidean Rhythms?

Euclidean rhythms distribute pulses as evenly as possible across time intervals, creating naturally pleasing rhythmic patterns found in music traditions worldwide. The algorithm was originally described by Euclid for finding the greatest common divisor, but it has fascinating applications in music.

## Quick Start

### Using the Helper Script (Recommended)
```bash
# Show available commands
./run.sh help

# Generate default rhythm
./run.sh generate

# See rhythm demonstrations  
./run.sh demo

# Generate example collection
./run.sh examples
```

### Direct Go Commands
```bash
# Build and run the main generator (16 steps, 6 pulses, 120 BPM)
go build main.go
./main

# Or run directly
go run main.go
```

This generates `euclid.wav` - a rhythmic pattern with the visualization:
```
Pattern: X.X.X.X.X.X..... (X=hit, .=rest)
```

## Exploring Different Patterns

### Interactive Demo
See various rhythm patterns from world music:
```bash
go run demo.go
```

### Generate Multiple Examples
Create a collection of famous rhythms as separate audio files:
```bash
go run examples.go
```
This creates an `examples/` directory with WAV files for different cultural patterns.

### Custom Patterns
To try different patterns, edit the values in `main.go`:
```go
steps := 8     // Change to 8 for shorter patterns
pulses := 3    // Change to 3 for Cuban tresillo: X.X.X...
bpm := 140     // Change tempo
```

## Famous Patterns to Try

| Steps | Pulses | Pattern | Style |
|-------|--------|---------|-------|
| 8 | 3 | `X.X.X...` | Cuban Tresillo |
| 8 | 5 | `X.X.X.XX` | Turkish Aksak |
| 12 | 5 | `X.X.X.X.X...` | West African |
| 16 | 6 | `X.X.X.X.X.X.....` | Bossa Nova (default) |

## Files

- `main.go` - Main generator program
- `demo.go` - Interactive demonstration of various patterns  
- `examples.go` - Batch generator for multiple rhythm examples
- `run.sh` - Helper script for easy building and running
- `README.md` - This documentation
- `euclid.wav` - Generated audio output (created when you run main.go)
- `examples/` - Directory with example rhythm files (created by examples.go)
- `go.mod` - Go module dependencies

## Audio Output

- **Format**: WAV, 16-bit, mono
- **Sample Rate**: 44.1 kHz (CD quality)
- **Default Duration**: ~8 seconds (depends on pattern and BPM)
- **Drum Sound**: Synthesized sine wave with exponential decay

The program provides detailed output showing the pattern, audio parameters, and file information.