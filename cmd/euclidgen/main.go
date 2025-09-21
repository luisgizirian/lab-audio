// Package main implements a Euclidean rhythm generator that creates rhythmic patterns
// using the Bjorklund algorithm and synthesizes them into audio files.
//
// Euclidean rhythms distribute a given number of pulses as evenly as possible
// across a specified number of time steps, creating naturally pleasing rhythmic
// patterns found in many musical traditions worldwide.
package main

import (
	"fmt"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"math"
	"os"
)

// bjorklund generates a Euclidean rhythm pattern using the Bjorklund algorithm.
//
// The algorithm distributes 'pulses' as evenly as possible across 'steps' time intervals.
// It returns a slice where 1 represents a drum hit and 0 represents silence.
//
// Examples:
//   - bjorklund(8, 3) returns [1,0,0,1,0,0,1,0] (Cuban tresillo)
//   - bjorklund(16, 6) returns [1,0,1,0,1,0,1,0,1,0,1,0,0,0,0,0]
//
// Parameters:
//   - steps: total number of time intervals (must be positive)
//   - pulses: number of drum hits to distribute (must be <= steps)
//
// Returns:
//   - []int: pattern where 1=hit, 0=rest
func bjorklund(steps, pulses int) []int {
	// Handle edge cases
	if pulses == 0 {
		return make([]int, steps) // all zeros
	}
	if pulses == steps {
		out := make([]int, steps)
		for i := range out {
			out[i] = 1 // all ones
		}
		return out
	}

	// Initialize groups: first 'pulses' groups get 1, rest get 0
	groups := make([][]int, steps)
	for i := 0; i < steps; i++ {
		if i < pulses {
			groups[i] = []int{1}
		} else {
			groups[i] = []int{0}
		}
	}

	// Apply Bjorklund algorithm: repeatedly pair different groups
	for {
		count := 0
		for i := 0; i < len(groups)-1; i++ {
			// If current group and last group are different and both single elements
			if len(groups[i]) == 1 && len(groups[len(groups)-1]) == 1 && 
			   groups[i][0] != groups[len(groups)-1][0] {
				// Combine them
				groups[i] = append(groups[i], groups[len(groups)-1][0])
				groups = groups[:len(groups)-1]
				count++
			}
		}
		if count == 0 {
			break // No more pairs to combine
		}
	}

	// Flatten groups into final pattern
	pattern := make([]int, 0, steps)
	for _, g := range groups {
		pattern = append(pattern, g...)
	}
	return pattern
}

// synthDrum generates a short percussive sound using a sine wave with exponential decay.
//
// The function creates a drum-like sound by:
// 1. Generating a sine wave at the specified frequency
// 2. Applying an exponential decay envelope for percussive character
// 3. Scaling to 16-bit audio range
//
// Parameters:
//   - sampleRate: audio sample rate in Hz (e.g., 44100)
//   - lengthMs: duration of the drum sound in milliseconds
//   - freq: fundamental frequency of the drum in Hz
//
// Returns:
//   - []int: audio samples as 16-bit integers
func synthDrum(sampleRate, lengthMs int, freq float64) []int {
	samples := sampleRate * lengthMs / 1000
	buf := make([]int, samples)
	
	for i := 0; i < samples; i++ {
		// Create exponential decay envelope (0.5 amplitude, -4 decay rate)
		amp := 0.5 * math.Exp(-4*float64(i)/float64(samples))
		
		// Generate sine wave and scale to 16-bit range (-32767 to 32767)
		phase := 2 * math.Pi * freq * float64(i) / float64(sampleRate)
		buf[i] = int(amp * 32767 * math.Sin(phase))
	}
	return buf
}

func main() {
	// === Rhythm Parameters ===
	steps := 16    // Number of time intervals in the pattern
	pulses := 6    // Number of drum hits to distribute
	
	// === Audio Parameters ===
	sampleRate := 44100  // Standard CD quality sample rate
	bpm := 120          // Beats per minute
	
	// === Synthesis Parameters ===
	drumLengthMs := 80  // Duration of each drum hit in milliseconds
	drumFreq := 180.0   // Fundamental frequency of drum sound in Hz
	
	// Calculate timing
	beatMs := 60000 / bpm  // Milliseconds per beat
	
	fmt.Printf("=== Euclidean Rhythm Generator ===\n")
	fmt.Printf("Generating pattern: %d steps, %d pulses\n", steps, pulses)
	fmt.Printf("Tempo: %d BPM\n", bpm)
	fmt.Printf("Audio: %d Hz, %d-bit\n", sampleRate, 16)
	
	// Generate the rhythmic pattern
	pattern := bjorklund(steps, pulses)
	
	// Display the pattern visually
	fmt.Printf("Pattern: ")
	for _, v := range pattern {
		if v == 1 {
			fmt.Print("X")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Printf(" (X=hit, .=rest)\n")
	
	// Generate drum sound
	drum := synthDrum(sampleRate, drumLengthMs, drumFreq)
	
	// Calculate total audio length and create output buffer
	totalSamples := sampleRate * steps * beatMs / 1000
	out := make([]int, totalSamples)
	
	// Place drum hits according to the pattern
	for i, v := range pattern {
		if v == 1 {
			// Calculate start position for this beat
			start := i * sampleRate * beatMs / 1000
			
			// Add drum sound to output buffer
			for j := 0; j < len(drum) && start+j < len(out); j++ {
				out[start+j] += drum[j]
			}
		}
	}
	
	// Prepare audio buffer for encoding
	buf := &audio.IntBuffer{
		Data:           out,
		Format:         &audio.Format{SampleRate: sampleRate, NumChannels: 1},
		SourceBitDepth: 16,
	}
	
	// Create output file
	outputFile := "euclid.wav"
	f, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", outputFile, err)
		os.Exit(1)
	}
	defer f.Close()
	
	// Encode to WAV format
	enc := wav.NewEncoder(f, sampleRate, 16, 1, 1)
	err = enc.Write(buf)
	if err != nil {
		fmt.Printf("Error writing audio data: %v\n", err)
		os.Exit(1)
	}
	
	err = enc.Close()
	if err != nil {
		fmt.Printf("Error finalizing audio file: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("✓ Generated '%s' with Euclidean rhythm!\n", outputFile)
	fmt.Printf("Duration: %.1f seconds\n", float64(len(out))/float64(sampleRate))
	fmt.Printf("File size: %.1f KB\n", float64(len(out)*2)/1024) // 16-bit = 2 bytes per sample
}
