package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
    if prepTime > 0 {
        return prepTime * len(layers)
    } else {
        return 2 * len(layers)
    }
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var noodles int
    var sauce float64
    
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += 50
        }
        if layers[i] == "sauce" {
            sauce += 0.2
        }
    }
    return noodles, sauce
}


// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
    if len(friendList) == 0 || len(myList) == 0 {
        return
    }
    myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(q []float64, portions int) []float64 {
    length := len(q)
	newQuantities := make([]float64, length)
    for i := 0; i < len(q); i++ {
        newQuantities[i] = q[i] * (float64(portions) / 2.0)
    }
    return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
