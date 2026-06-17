package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elms := make([]int, size)

	for i := 0; i < size; i++ {
		elms[i] = rand.Int()
	}

	return elms
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) < CHUNKS*2 || data == nil {
		return maximum(data)
	}

	maxValues := make([]int, 8)
	var chunkSize = len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		begin := i * chunkSize
		end := begin + chunkSize

		if end > len(data) || i == CHUNKS-1 {
			end = len(data)
		}

		maxValues = append(maxValues, maximum(data[begin:end]))
	}

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elms := generateRandomElements(SIZE)

	now := time.Now()
	fmt.Println("Ищем максимальное значение в один поток")
	max := maximum(elms)
	elapsed := time.Since(now)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	now = time.Now()
	max = maxChunks(elms)
	elapsed = time.Since(now)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())
}
