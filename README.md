# Lab Audio 🎵

An experimental audio processing laboratory featuring algorithmic rhythm generation and audio synthesis utilities written in Go.

## What's Inside

This repository explores the fascinating world of **Euclidean rhythms** - mathematically distributed rhythmic patterns that appear in music traditions across many cultures. The main utility generates these patterns and synthesizes them into audio files.

### 🥁 Euclidean Rhythm Generator (`euclidgen`)

Euclidean rhythms are created by distributing a given number of pulses as evenly as possible across a specified number of steps. This creates naturally occurring rhythmic patterns found in various musical traditions:

- **Cuban clave** (3 pulses in 8 steps): `X..X..X.`
- **Turkish aksak** (5 pulses in 8 steps): `X.X.X.X.`
- **West African polyrhythms** and many more!

The generator uses the **Bjorklund algorithm** to compute these patterns and synthesizes them using a simple drum sound.

## Repository Structure

```
lab-audio/
├── README.md              # This documentation
├── cmd/
│   └── euclidgen/         # Euclidean rhythm generator
│       ├── main.go        # Main application code
│       ├── go.mod         # Go module definition
│       └── euclid.wav     # Generated audio output
```

## Getting Started

### Prerequisites

- Go 1.24+ installed on your system
- Audio player capable of playing WAV files

### Quick Start

1. **Clone the repository:**
   ```bash
   git clone https://github.com/luisgizirian/lab-audio.git
   cd lab-audio
   ```

2. **Build and run the Euclidean rhythm generator:**
   ```bash
   cd cmd/euclidgen
   
   # Easy way (using helper script)
   ./run.sh generate
   
   # Or direct Go commands
   go run main.go
   ```

3. **Explore different patterns:**
   ```bash
   # See rhythm demonstrations from world music
   ./run.sh demo
   
   # Generate a collection of example rhythms
   ./run.sh examples
   ```

4. **Listen to the generated rhythm:**
   ```bash
   # The program generates 'euclid.wav' - play it with your preferred audio player
   # On Linux: aplay euclid.wav
   # On macOS: afplay euclid.wav
   # On Windows: start euclid.wav
   ```

### Current Default Settings

The generator currently produces:
- **16 steps** with **6 pulses** (pattern: `X.X.X.X.X.X.....`)
- **120 BPM** tempo
- **44.1kHz** sample rate
- **80ms** drum hits at **180Hz** frequency

### Program Features

- **✨ Rich documentation** with mathematical insights and cultural context
- **🎭 Interactive demos** showing patterns from different musical traditions  
- **📚 Example generator** creating collections of famous rhythms
- **🔧 Helper script** for easy building and running
- **📊 Visual pattern display** showing rhythms as X (hit) and . (rest)
- **🎵 High-quality audio** output in standard WAV format
- **💻 Well-documented code** with comprehensive comments

## Example Outputs

Here are some classic Euclidean rhythm patterns you can experiment with:

| Steps | Pulses | Pattern | Musical Example |
|-------|--------|---------|-----------------|
| 8 | 3 | `X..X..X.` | Cuban tresillo |
| 8 | 5 | `X.X.X.XX` | Turkish aksak |
| 16 | 6 | `X.X.X.X.X.X....` | Current default |
| 12 | 5 | `X..X.X..X.X.` | West African pattern |

## Understanding Euclidean Rhythms

Euclidean rhythms solve the problem: "How do you distribute k pulses among n time intervals as evenly as possible?" This mathematical approach to rhythm generation creates patterns that feel natural and musically pleasing because they maximize the temporal distance between pulses.

The algorithm was originally described by Euclid for finding the greatest common divisor of two numbers, but it turns out to have fascinating applications in music theory and rhythm generation.

## Future Improvements

Ideas for enhancing this project:
- [ ] Command-line arguments for customizing steps, pulses, BPM, and sound parameters
- [ ] Multiple drum sounds and instruments
- [ ] Polyrhythmic combinations (multiple patterns playing simultaneously)
- [ ] Real-time pattern visualization
- [ ] MIDI file export
- [ ] Interactive web interface
- [ ] Pattern library with famous rhythms from different cultures

## Contributing

Feel free to fork this repository and experiment with:
- Different synthesis algorithms
- New rhythm generation techniques
- Audio effects and processing
- Visualization components
- Additional musical applications

## License

MIT License - feel free to use this code for your own audio experiments!
