package main

var input, reportInput uint

var categories = []string{"Household", "Food", "Drinks"}

type item struct {
	category uint
	quantity uint
	cost     float64
}

var itemMap = make(map[string]item)

func init() {
	itemMap["Fork"] = item{0, 4, 3.00}
	itemMap["Plates"] = item{0, 4, 3.00}
	itemMap["Cups"] = item{0, 5, 3.00}
	itemMap["Bread"] = item{1, 2, 2.00}
	itemMap["Cake"] = item{1, 3, 1.00}
	itemMap["Coke"] = item{2, 5, 2.00}
	itemMap["Sprite"] = item{2, 5, 2.00}
}