package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
    if avgPrepTime == 0 {
        avgPrepTime = 2
    }
    return len(layers) * avgPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var (
        noodlesQuantity int
        sauceQuantity float64
	)

    for _, layer := range layers {
        if layer == "noodles" {
			noodlesQuantity += 50
        }
        if layer == "sauce" {
            sauceQuantity += 0.2
        }
    }

    return noodlesQuantity, sauceQuantity
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    myList[len(myList) - 1] = friendList[len(friendList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numPortions int) []float64 {
    var scaledQuantities []float64
    scaleCoef := float64(numPortions) / 2.0

    for _, num := range quantities {
        scaledQuantities = append(scaledQuantities, num * scaleCoef)
    }
    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
