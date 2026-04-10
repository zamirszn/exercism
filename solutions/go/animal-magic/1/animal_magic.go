package chance
import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	min, max := 1, 20
	return rand.Intn(max-min) + min
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	min, max := .0, 12.0
	return min + rand.Float64() * (max-min)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    items := []string{"ant", "beaver", "cat", "dog", "elephant","fox","giraffe" , "hedgehog" }
	 rand.Shuffle(len(items), func (i, j int){
        items[i], items[j]= items[j] , items[i]
    })
    return items
}
