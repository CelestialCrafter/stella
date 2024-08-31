package common

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

const (
	sampleRate   = 44100
	duration     = 2 // seconds
	numSamples   = sampleRate * duration
	maxAmplitude = 32767 // for 16-bit audio
)

func rgbToWave(r, g uint8) (frequency, amplitude float64) {
	// Map RGB to frequency (100-1000 Hz) and amplitude (0-1)
	frequency = 100 + float64(r)/255.0*900
	amplitude = float64(g) / 255.0
	return
}

func generateSineWave(frequency, amplitude float64) []int16 {
	samples := make([]int16, numSamples)
	for i := 0; i < numSamples; i++ {
		t := float64(i) / sampleRate
		sample := amplitude * math.Sin(2*math.Pi*frequency*t)
		samples[i] = int16(sample * maxAmplitude)
	}
	return samples
}

func writeWAVFile(filename string, samples []int16) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write WAV header
	binary.Write(file, binary.LittleEndian, []byte("RIFF"))
	binary.Write(file, binary.LittleEndian, uint32(36+len(samples)*2))
	binary.Write(file, binary.LittleEndian, []byte("WAVE"))
	binary.Write(file, binary.LittleEndian, []byte("fmt "))
	binary.Write(file, binary.LittleEndian, uint32(16))
	binary.Write(file, binary.LittleEndian, uint16(1))
	binary.Write(file, binary.LittleEndian, uint16(1))
	binary.Write(file, binary.LittleEndian, uint32(sampleRate))
	binary.Write(file, binary.LittleEndian, uint32(sampleRate*2))
	binary.Write(file, binary.LittleEndian, uint16(2))
	binary.Write(file, binary.LittleEndian, uint16(16))
	binary.Write(file, binary.LittleEndian, []byte("data"))
	binary.Write(file, binary.LittleEndian, uint32(len(samples)*2))

	// Write audio data
	for _, sample := range samples {
		binary.Write(file, binary.LittleEndian, sample)
	}

	return nil
}

func GenrateSound(planetHash string, r uint8, g uint8) {
	frequency, amplitude := rgbToWave(r, g)
	samples := generateSineWave(frequency, amplitude)

	filename := fmt.Sprintf("sounds/%s.wav", planetHash)
	err := writeWAVFile(filename, samples)
	if err != nil {
		fmt.Println("Error writing WAV file:", err)
		return
	}

}
