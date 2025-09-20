// Demo script showing various Euclidean rhythm patterns
// This script demonstrates the mathematical beauty of Euclidean rhythms
// by generating and displaying various famous patterns from world music.
//
// Run with: go run demo.go

package main

import (
	"fmt"
	"strings"
)

// bjorklund generates a Euclidean rhythm pattern (copied from main.go for standalone demo)
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

// formatPattern converts a pattern to visual representation
func formatPattern(pattern []int) string {
	var result strings.Builder
	for _, v := range pattern {
		if v == 1 {
			result.WriteString("X")
		} else {
			result.WriteString(".")
		}
	}
	return result.String()
}

// rhythmExample represents a famous rhythm pattern
type rhythmExample struct {
	name        string
	steps       int
	pulses      int
	description string
	origin      string
}

func main() {
	fmt.Println("🎵 EUCLIDEAN RHYTHMS DEMONSTRATION 🎵")
	fmt.Println("=====================================")
	fmt.Println("Exploring rhythmic patterns from around the world using mathematics!")
	fmt.Println("")

	// Famous Euclidean rhythm examples
	examples := []rhythmExample{
		{
			name:        "Cuban Tresillo",
			steps:       8,
			pulses:      3,
			description: "The fundamental rhythm of Cuban music",
			origin:      "Cuba, Latin America",
		},
		{
			name:        "Turkish Aksak",
			steps:       8,
			pulses:      5,
			description: "Asymmetrical rhythm common in Turkish folk music",
			origin:      "Turkey, Eastern Europe",
		},
		{
			name:        "West African Polyrhythm",
			steps:       12,
			pulses:      5,
			description: "Complex polyrhythmic pattern",
			origin:      "West Africa",
		},
		{
			name:        "Flamenco Bulería",
			steps:       12,
			pulses:      7,
			description: "Fast-paced rhythm in flamenco music",
			origin:      "Spain",
		},
		{
			name:        "Brazilian Bossa Nova",
			steps:       16,
			pulses:      6,
			description: "Smooth, syncopated rhythm (current default)",
			origin:      "Brazil",
		},
		{
			name:        "Indian Classical Tala",
			steps:       7,
			pulses:      3,
			description: "Asymmetrical cycle in Indian classical music",
			origin:      "India",
		},
		{
			name:        "Minimalist Pattern",
			steps:       5,
			pulses:      2,
			description: "Simple, hypnotic pattern",
			origin:      "Modern/Minimal Music",
		},
		{
			name:        "Dense Polyrhythm",
			steps:       16,
			pulses:      11,
			description: "Complex, dense rhythmic texture",
			origin:      "Contemporary/Experimental",
		},
	}

	for i, example := range examples {
		pattern := bjorklund(example.steps, example.pulses)
		visual := formatPattern(pattern)
		
		fmt.Printf("%d. %s (%d/%d)\n", i+1, example.name, example.pulses, example.steps)
		fmt.Printf("   Pattern: %s\n", visual)
		fmt.Printf("   Origin:  %s\n", example.origin)
		fmt.Printf("   Notes:   %s\n", example.description)
		fmt.Printf("   Density: %.1f%% (%.1f pulses per beat)\n", 
			float64(example.pulses)/float64(example.steps)*100,
			float64(example.pulses*4)/float64(example.steps))
		fmt.Println()
	}

	fmt.Println("🔬 MATHEMATICAL INSIGHTS")
	fmt.Println("=========================")
	fmt.Println("• Euclidean rhythms maximize the temporal distance between pulses")
	fmt.Println("• They solve: 'distribute k pulses among n intervals as evenly as possible'")
	fmt.Println("• The Bjorklund algorithm is related to Euclid's algorithm for GCD")
	fmt.Println("• These patterns naturally emerge in traditional music worldwide")
	fmt.Println("• Musicians often discover them intuitively without knowing the mathematics")
	fmt.Println("")
	
	fmt.Println("🎛️  EXPERIMENT IDEAS")
	fmt.Println("====================")
	fmt.Println("Try modifying the main program with different values:")
	fmt.Println("• Change steps/pulses ratio for different feels")
	fmt.Println("• Adjust BPM for different tempos")
	fmt.Println("• Modify drum frequency for different pitches")
	fmt.Println("• Layer multiple patterns for polyrhythms")
	fmt.Println("")
	
	fmt.Println("To generate any of these patterns as audio:")
	fmt.Println("1. Edit the 'steps' and 'pulses' values in main.go")
	fmt.Println("2. Run: go run main.go")
	fmt.Println("3. Listen to the generated euclid.wav file")
}