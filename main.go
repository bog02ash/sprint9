package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	var data []int
	if size <= 0 {
		return nil
	}
	for i := 0; i < size; i++ {
		number := rand.Intn(size)
		data = append(data, number)
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 || len(data) == 1 {
		return 0
	}
	max := slices.Max(data)
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	var wg sync.WaitGroup
	var finalSlice []int
	for i := 0; i < CHUNKS; i++ {
		sizeSrez := len(data) / CHUNKS
		initialIndex := i * sizeSrez
		finalIndex := initialIndex + sizeSrez
		srez := data[initialIndex:finalIndex]
		wg.Add(1)
		go func() {
			defer wg.Done()
			max := slices.Max(srez)
			finalSlice = append(finalSlice, max)
		}()
	}
	wg.Wait()
	result := slices.Max(finalSlice)
	return result
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	startTime := time.Now()
	max := maximum(data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	elapsed := duration.Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	startTime = time.Now()
	max = maxChunks(data)
	endTime = time.Now()
	duration = endTime.Sub(startTime)
	elapsed = duration.Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
