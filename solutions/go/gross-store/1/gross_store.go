package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	val, exists := units[unit]
    if !exists {
        return false
    }
    _, exists2 := bill[item]
    if exists2 {
        bill[item] += val
    } else {
        bill[item] = val
    }
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	val1, exists := units[unit]
    val2, exists2 := bill[item]
    if !exists || !exists2 || val1 > val2 {
        return false
    }
    if val1 == val2 {
        delete(bill, item)
    } else {
        bill[item] -= val1
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, exists := bill[item]
    if exists {
        return val, true
    } else {
        return 0, false
    }
}
