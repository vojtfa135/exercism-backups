package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdSum := 0
    for _, birdCount := range birdsPerDay {
        birdSum += birdCount
    }
    return birdSum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    startIdx := 7 * (week - 1)
    endIdx := 7 + 7 * (week - 1)
    return TotalBirdCount(birdsPerDay[startIdx:endIdx])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for idx, birdCount := range birdsPerDay {
        if idx % 2 == 0 {
            birdsPerDay[idx] = birdCount + 1
        }
    }
    return birdsPerDay
}
