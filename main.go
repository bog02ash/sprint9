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
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	max := slices.Max(data)
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	var wg sync.WaitGroup
	finalSlice := make([]int, CHUNKS)
	sizeSrez := len(data) / CHUNKS
	residue := len(data) % CHUNKS
	for i := 0; i < CHUNKS; i++ {
		start := i * sizeSrez
		if i < residue {
			start += i
		} else {
			start += residue
		}
		final := start + sizeSrez
		if i < residue {
			final += 1
		}
		wg.Add(1)
		go func(i, start, final int) {
			defer wg.Done()
			srez := data[start:final]
			finalSlice[i] = maximum(srez)
		}(i, start, final)
	}
	wg.Wait()
	return maximum(finalSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	startTime := time.Now()
	max := maximum(data)
	elapsed := time.Since(startTime)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	startTime = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(startTime)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
