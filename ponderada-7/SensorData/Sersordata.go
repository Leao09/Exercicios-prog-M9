package Sensor
import (
	"math/rand"
)

func SensorData() map[string]int {
	data := map[string]int{
		"NH3_ppm": rand.Intn(400),
		"CO_ppm":  rand.Intn(1000),
		"NO2_ppm": rand.Intn(30),
	}
	return data
}