// Example script showing how to generate multiple rhythm patterns
// This demonstrates programmatic use of the Euclidean rhythm generator
// to create a collection of different rhythmic patterns.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"math"
)

// Copy the necessary functions from main.go
func bjorklund(steps, pulses int) []int {
	if pulses == 0 {
		return make([]int, steps)
	}
	if pulses == steps {
		out := make([]int, steps)
		for i := range out {
			out[i] = 1
		}
		return out
	}
	groups := make([][]int, steps)
	for i := 0; i < steps; i++ {
		if i < pulses {
			groups[i] = []int{1}
		} else {
			groups[i] = []int{0}
		}
	}
	for {
		count := 0
		for i := 0; i < len(groups)-1; i++ {
			if len(groups[i]) == 1 && len(groups[len(groups)-1]) == 1 && 
			   groups[i][0] != groups[len(groups)-1][0] {
				groups[i] = append(groups[i], groups[len(groups)-1][0])
				groups = groups[:len(groups)-1]
				count++
			}
		}
		if count == 0 {
			break
		}
	}
	pattern := make([]int, 0, steps)
	for _, g := range groups {
		pattern = append(pattern, g...)
	}
	return pattern
}

func synthDrum(sampleRate, lengthMs int, freq float64) []int {
	samples := sampleRate * lengthMs / 1000
	buf := make([]int, samples)
	
	for i := 0; i < samples; i++ {
		amp := 0.5 * math.Exp(-4*float64(i)/float64(samples))
		phase := 2 * math.Pi * freq * float64(i) / float64(sampleRate)
		buf[i] = int(amp * 32767 * math.Sin(phase))
	}
	return buf
}

// RhythmConfig holds parameters for generating a rhythm
type RhythmConfig struct {
	Name       string
	Steps      int
	Pulses     int
	BPM        int
	DrumFreq   float64
	OutputFile string
}

// generateRhythm creates a WAV file for the given rhythm configuration
func generateRhythm(config RhythmConfig) error {
	sampleRate := 44100
	drumLengthMs := 80
	beatMs := 60000 / config.BPM
	
	// Generate pattern and drum sound
	pattern := bjorklund(config.Steps, config.Pulses)
	drum := synthDrum(sampleRate, drumLengthMs, config.DrumFreq)
	
	// Create audio buffer
	totalSamples := sampleRate * config.Steps * beatMs / 1000
	out := make([]int, totalSamples)
	
	// Place drum hits
	for i, v := range pattern {
		if v == 1 {
			start := i * sampleRate * beatMs / 1000
			for j := 0; j < len(drum) && start+j < len(out); j++ {
				out[start+j] += drum[j]
			}
		}
	}
	
	// Create output file
	f, err := os.Create(config.OutputFile)
	if err != nil {
		return err
	}
	defer f.Close()
	
	// Encode to WAV
	buf := &audio.IntBuffer{
		Data:           out,
		Format:         &audio.Format{SampleRate: sampleRate, NumChannels: 1},
		SourceBitDepth: 16,
	}
	
	enc := wav.NewEncoder(f, sampleRate, 16, 1, 1)
	if err := enc.Write(buf); err != nil {
		return err
	}
	if err := enc.Close(); err != nil {
		return err
	}
	
	// Display pattern
	fmt.Printf("Generated %s: ", config.Name)
	for _, v := range pattern {
		if v == 1 {
			fmt.Print("X")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Printf(" -> %s\n", config.OutputFile)
	
	return nil
}

func main() {
	fmt.Println("🎵 EUCLIDEAN RHYTHM COLLECTION GENERATOR")
	fmt.Println("========================================")
	fmt.Println("Generating a collection of famous Euclidean rhythms...")
	fmt.Println()
	
	// Create examples directory
	outputDir := "examples"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}
	
	// Collection of famous rhythms
	rhythms := []RhythmConfig{
		{
			Name:       "Cuban Tresillo",
			Steps:      8,
			Pulses:     3,
			BPM:        120,
			DrumFreq:   180.0,
			OutputFile: filepath.Join(outputDir, "cuban_tresillo.wav"),
		},
		{
			Name:       "Turkish Aksak",
			Steps:      8,
			Pulses:     5,
			BPM:        100,
			DrumFreq:   200.0,
			OutputFile: filepath.Join(outputDir, "turkish_aksak.wav"),
		},
		{
			Name:       "West African",
			Steps:      12,
			Pulses:     5,
			BPM:        110,
			DrumFreq:   160.0,
			OutputFile: filepath.Join(outputDir, "west_african.wav"),
		},
		{
			Name:       "Bossa Nova",
			Steps:      16,
			Pulses:     6,
			BPM:        120,
			DrumFreq:   180.0,
			OutputFile: filepath.Join(outputDir, "bossa_nova.wav"),
		},
		{
			Name:       "Minimalist",
			Steps:      5,
			Pulses:     2,
			BPM:        90,
			DrumFreq:   220.0,
			OutputFile: filepath.Join(outputDir, "minimalist.wav"),
		},
	}
	
	// Generate all rhythms
	for _, rhythm := range rhythms {
		if err := generateRhythm(rhythm); err != nil {
			fmt.Printf("Error generating %s: %v\n", rhythm.Name, err)
		}
	}
	
	fmt.Printf("\n✓ Generated %d rhythm examples in '%s/' directory\n", len(rhythms), outputDir)
	fmt.Println("\nListen to the different patterns and compare their feels!")
	fmt.Println("Each pattern uses slightly different tempo and pitch for character.")
}