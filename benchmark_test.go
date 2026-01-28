package main

import (
	"strconv"
	"testing"
)

// Имитируем map[string]int с фильмами
type Movies map[string]int

// Медленный подход: 2 прохода по map
func findMaxSlow(movies Movies) []string {
	maxRating := 0
	// 1-й проход: ищем максимум
	for _, rating := range movies {
		if rating > maxRating {
			maxRating = rating
		}
	}

	// 2-й проход: собираем все с максимумом
	topMovies := []string{}
	for movie, rating := range movies {
		if rating == maxRating {
			topMovies = append(topMovies, movie)
		}
	}
	return topMovies
}

// Быстрый подход: 1 проход по map
func findMaxFast(movies Movies) []string {
	maxRating := 0
	topMovies := []string{}

	// ОДИН проход: ищем max И собираем топ одновременно
	for movie, rating := range movies {
		if rating > maxRating {
			maxRating = rating
			topMovies = []string{movie}
		} else if rating == maxRating {
			topMovies = append(topMovies, movie)
		}
	}
	return topMovies
}

// Генератор тестовых данных (10 000 фильмов)
func generateData(n int) Movies {
	movies := make(Movies)
	for i := 0; i < n; i++ {
		name := "Movie" + strconv.Itoa(i%100)
		rating := (i % 5) + 1
		movies[name] = rating
	}
	return movies
}

func BenchmarkSlow(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMaxSlow(data)
	}
}

func BenchmarkFast(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMaxFast(data)
	}
}
