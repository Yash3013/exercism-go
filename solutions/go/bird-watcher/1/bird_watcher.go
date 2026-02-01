package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	// panic("Please implement the TotalBirdCount() function")
    total := 0
    for _,cnt := range birdsPerDay{
        total += cnt
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// panic("Please implement the BirdsInWeek() function")
    total := 0
    st := (week-1)*7
    ed := st+7
    for i:=st; i<ed; i++{
        total += birdsPerDay[i]
    }
    return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// panic("Please implement the FixBirdCountLog() function")
    for i:=0; i<len(birdsPerDay); i+=2{
        birdsPerDay[i]++
    }
    return birdsPerDay
}
