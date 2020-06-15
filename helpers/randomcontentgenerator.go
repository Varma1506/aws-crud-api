package helpers

import (
	"math/rand"
)

func SetRandomBrand() string {
	brands := []string{"Apple", "Samsung", "Google", "Oneplus"}

	// Generate a random index based on the length of the items slice
	randomIndex := rand.Intn(len(brands))

	return brands[randomIndex]
}

func SetRandomDescription() string {
	descriptions := []string{"Great Phone", "Good Phone", "Best Phone", "Decent Phone"}

	// Generate a random index based on the length of the items slice
	randomIndex := rand.Intn(len(descriptions))

	return descriptions[randomIndex]
}

func SetRandomPrice() float64 {
	descriptions := []float64{899.99, 1099.99, 799.99, 699.99}

	// Generate a random index based on the length of the items slice
	randomIndex := rand.Intn(len(descriptions))

	return descriptions[randomIndex]
}
