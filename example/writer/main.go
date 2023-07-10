package main

import (
	crytporand "crypto/rand"
	"encoding/binary"
	"math"
	"time"

	"github.com/bclswl0827/sacio"
)

const SAMPLE_RATE = 375
const TIME_SPAN = 60 * 60 * 1

func main() {
	// Generate test data series
	testData, err := getRandomData(SAMPLE_RATE*TIME_SPAN, -100, 100)
	if err != nil {
		panic(err)
	}

	// Init header fields
	var sacData sacio.SACData
	sacData.Init()

	// Set new header fields
	sacData.SetTime(time.Now().UTC().Add(time.Duration(-TIME_SPAN) * time.Second))
	sacData.SetInfo("TEST-NET", "TEST-STA", "TEST-LOC", "BHZ")
	sacData.SetBody(testData, SAMPLE_RATE)

	// Write to file
	err = sacData.Write("./result.sac", sacio.MSBFIRST)
	if err != nil {
		panic(err)
	}
}

func getRandomData(length int, min, max float32) ([]float32, error) {
	array := make([]float32, length)
	rangeSize := float64(max - min)

	for i := 0; i < length; i++ {
		randomBytes := make([]byte, 4)
		_, err := crytporand.Read(randomBytes)
		if err != nil {
			return nil, err
		}

		randomValue := float32(binary.LittleEndian.Uint32(randomBytes))
		randomValue = (randomValue/float32(math.MaxUint32))*float32(rangeSize) + min
		array[i] = randomValue
	}

	return array, nil
}
