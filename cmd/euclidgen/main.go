package main

import (
	"fmt"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"math"
	"os"
)

// bjorklund generates a Euclidean rhythm pattern
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
			if len(groups[i]) == 1 && len(groups[len(groups)-1]) == 1 && groups[i][0] != groups[len(groups)-1][0] {
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

// synthDrum returns a short percussive sound (simple sine burst)
func synthDrum(sampleRate, lengthMs int, freq float64) []int {
	samples := sampleRate * lengthMs / 1000
	buf := make([]int, samples)
	for i := 0; i < samples; i++ {
		amp := 0.5 * math.Exp(-4*float64(i)/float64(samples)) // envelope
		buf[i] = int(amp * 32767 * math.Sin(2*math.Pi*freq*float64(i)/float64(sampleRate)))
	}
	return buf
}

func main() {
	steps := 16
	pulses := 6
	sampleRate := 44100
	bpm := 120
	beatMs := 60000 / bpm
	drum := synthDrum(sampleRate, 80, 180.0)
	pattern := bjorklund(steps, pulses)
	totalSamples := sampleRate * steps * beatMs / 1000
	out := make([]int, totalSamples)
	for i, v := range pattern {
		if v == 1 {
			start := i * sampleRate * beatMs / 1000
			for j := 0; j < len(drum) && start+j < len(out); j++ {
				out[start+j] += drum[j]
			}
		}
	}
	buf := &audio.IntBuffer{
		Data:           out,
		Format:         &audio.Format{SampleRate: sampleRate, NumChannels: 1},
		SourceBitDepth: 16,
	}
	f, err := os.Create("euclid.wav")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	enc := wav.NewEncoder(f, sampleRate, 16, 1, 1)
	err = enc.Write(buf)
	if err != nil {
		panic(err)
	}
	err = enc.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Generated euclid.wav with Euclidean rhythm!")
}
