package main

import "fmt"

func modifyCategory() {
	var modCat, newCatName string
	
	fmt.Println("")
	fmt.Println("Modify category")
	fmt.Println("===============")
	fmt.Printf("Which category do you wish to modify? ")
	fmt.Scanln(&modCat)

	indexCat := searchCategory(modCat)
	if indexCat == -1 {
		fmt.Println("Category not found.")
	} else {
		fmt.Printf("You have selected Category: %v. What would you like the new name to be? ", categories[indexCat])
		fmt.Scanln(&newCatName)
		
		categories[indexCat] = newCatName
		fmt.Printf("Category name has been changed to %v.\n", newCatName)
	}
}

func deleteCategory() {
	var delCat string
	
	fmt.Println("")
	fmt.Println("Delete category")
	fmt.Println("===============")
	fmt.Printf("Which category do you wish to delete? ")
	fmt.Scanln(&delCat)

	// Check if category exists
	indexCat := searchCategory(delCat)
	if indexCat == -1 {
		fmt.Println("Category not found.")
	} else {
		
		// Delete from map and reorder categories
		for key, value := range itemMap {
			if categories[value.category] == delCat {
				delete(itemMap, key)
			}
			if int(value.category) > indexCat {
				itemMap[key] = item{value.category - 1, value.quantity, value.cost}
			}
		}

		// Delete from slice
		copy(categories[indexCat:], categories[indexCat + 1:])
		categories = categories[:len(categories) - 1]

		fmt.Printf("Category: %v deleted. Remaining categories are %v\n", delCat, categories)

	}

}

func searchCategory(category string) (index int) {
	index = -1
	for i, element := range categories {
		if category == element {
			index = i
		}
	}
	return index
}