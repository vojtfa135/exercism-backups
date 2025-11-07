package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if (kind == "car" || kind == "truck") {
        return true
    }
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    recomMsg := func(recomOption string) string {
        return fmt.Sprintf("%s is clearly the better choice.", recomOption)
    }
	if (option1 < option2) {
        return recomMsg(option1)
    }
    return recomMsg(option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if (age >= 10) {
        return 0.5 * originalPrice
    } else if (age >= 3 && age < 10) {
        return 0.7 * originalPrice
    }
    return 0.8 * originalPrice
}
