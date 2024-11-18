package main

import (
	"fmt"
	"math/rand"
	"time"
)

// BootstrapSample generates bootstrap samples and calculates the statistic (median in this case).
func BootstrapSample(data []float64, nBoot int) ([]float64, float64) {
	bootstrapMedians := make([]float64, nBoot)

	for i := 0; i < nBoot; i++ {
		sample := generateSample(data)
		bootstrapMedians[i] = calculateMedian(sample)
	}

	stdErr := calculateStdErr(bootstrapMedians)
	return bootstrapMedians, stdErr
}

// generateSample generates a bootstrap sample from the input data.
func generateSample(data []float64) []float64 {
	rand.Seed(time.Now().UnixNano())
	n := len(data)
	sample := make([]float64, n)
	for i := 0; i < n; i++ {
		sample[i] = data[rand.Intn(n)]
	}
	return sample
}

// calculateMedian calculates the median of a slice.
func calculateMedian(data []float64) float64 {
	n := len(data)
	copyData := make([]float64, n)
	copy(copyData, data)
	quickSort(copyData)

	if n%2 == 0 {
		return (copyData[n/2-1] + copyData[n/2]) / 2
	}
	return copyData[n/2]
}

// calculateStdErr calculates the standard error from bootstrap results.
func calculateStdErr(data []float64) float64 {
	n := len(data)
	mean := calculateMean(data)
	var variance float64
	for _, val := range data {
		variance += (val - mean) * (val - mean)
	}
	return sqrt(variance / float64(n))
}

// Helper functions for mean and sqrt.
func calculateMean(data []float64) float64 {
	sum := 0.0
	for _, val := range data {
		sum += val
	}
	return sum / float64(len(data))
}

func sqrt(x float64) float64 {
	return x // Replace with actual sqrt implementation or math package
}

// quickSort sorts the data slice in ascending order.
func quickSort(data []float64) {
	// Implement a quick sort or use a standard library sort
}

func main() {
	data := []float64{1.2, 2.4, 3.6, 4.8, 5.0, 6.1, 7.3, 8.4, 9.9}
	nBoot := 1000

	bootstrapMedians, stdErr := BootstrapSample(data, nBoot)
	fmt.Printf("Bootstrap Medians: %v\n", bootstrapMedians)
	fmt.Printf("Standard Error of Median: %.4f\n", stdErr)
}
