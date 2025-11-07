package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var normalizedTable string

	if (table < 10) {
        normalizedTable = fmt.Sprintf("00%d", table)
    } else if (table < 100) {
        normalizedTable = fmt.Sprintf("0%d", table)
    } else {
        normalizedTable = fmt.Sprintf("%d", table)
    }
    
	welcome := fmt.Sprintf("Welcome to my party, %s!", name)
    tableAssign := fmt.Sprintf("You have been assigned to table %s.", normalizedTable)
    tableXY := fmt.Sprintf("Your table is %s, exactly %.1f meters from here.", direction, distance)
    nextPerson := fmt.Sprintf("You will be sitting next to %s.", neighbor)

    return fmt.Sprintf("%s\n%s %s\n%s", welcome, tableAssign, tableXY, nextPerson)
}
