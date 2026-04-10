package birdwatcher
import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    total := 0	
	for index := 0; index < len(birdsPerDay); index++ {
		fmt.Println(birdsPerDay[index])
		total = total + birdsPerDay[index]
	}
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	daysinweek := 7
	startofweek := (week - 1) * daysinweek
	endIndex := (week * daysinweek)
	return TotalBirdCount(birdsPerDay[startofweek:endIndex])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for index , num := range birdsPerDay{
        if index%2 == 0{     
            birdsPerDay[index]= num+1
        }
    }
    return birdsPerDay
}
