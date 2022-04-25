package main

import "fmt"

func main(){
  // Main loop with n as control
  var n uint
  for n < 10 {

    fmt.Println("")
    fmt.Println("Shopping List Application")
    fmt.Println("=========================")
    fmt.Println("1. View entire shopping list")
    fmt.Println("2. Generate shopping list report")
    fmt.Println("3. Add item")
    fmt.Println("4. Modify item")
    fmt.Println("5. Delete item")
    fmt.Println("6. Print current data")
    fmt.Println("7. Add new category name")
    fmt.Println("8. Modify category")
    fmt.Println("9. Delete category")
    fmt.Println("10. Exit")
    fmt.Printf("Select your choice: ")
    fmt.Scanln(&input)

    switch input {
    
      // View entire shopping list
      case 1:
        fmt.Println("")
        fmt.Println("Shopping List Contents")
        fmt.Println("======================")
        for key, element := range itemMap {
          fmt.Printf("Category: %v - Item: %v  Quantity: %v  Unit Cost: $%.2f\n", categories[element.category], key, element.quantity, element.cost)
        }
        
      // Generate shopping list report
      case 2:
        fmt.Println("")
        fmt.Println("Generate Report")
        fmt.Println("===============")
        fmt.Println("1. Total cost of each category")
        fmt.Println("2. List of items by category")
        fmt.Println("3. Main Menu")
        fmt.Printf("Choose your report: ")
        fmt.Scanln(&reportInput)

        switch reportInput {

        // Total cost of each category
        case 1:
          fmt.Println("")
          fmt.Println("Total cost by category")
          fmt.Println("======================")
          for index, elementOutside := range categories {
            sum := 0.00
            for _, elementInside := range itemMap {
              if index == int(elementInside.category) {
                sum = sum + (float64(elementInside.quantity) * elementInside.cost)
              }
            }
            fmt.Printf("%v cost: $%.2f\n", elementOutside, sum)
          }

        // List of items by category
        case 2:
          fmt.Println("")
          fmt.Println("List of items by category")
          fmt.Println("=========================")
          for index, elementOutside := range categories {
            for key, elementInside := range itemMap {
              if index == int(elementInside.category) {
                fmt.Printf("Category: %v - Item: %v  Quantity: %v  Unit Cost: $%.2f\n", elementOutside, key, elementInside.quantity, elementInside.cost)
              }
            }
          }

        // Back to Main Menu
        case 3:
          break
        }

      // Add items information
      case 3:
        var itemName, itemCategory string
        var itemUnits uint
        var itemCost float64
        
        fmt.Println("")
        fmt.Println("Add item")
        fmt.Println("========")
        fmt.Println("What is the name of your item?")
        fmt.Scanln(&itemName)
        fmt.Println("What category does it belong to?")
        fmt.Scanln(&itemCategory)
        fmt.Println("How many units are there?")
        fmt.Scanln(&itemUnits)
        fmt.Println("How much does it cost per unit?")
        fmt.Scanln(&itemCost)

        // Search through slice to find index value of category
        var cat uint
        for index, element := range categories {
          if itemCategory == element {
            cat = uint(index)
          }
        }

        itemMap[itemName] = item{cat, itemUnits, itemCost}

      // Modify existing items
      case 4:
        var itemName, newName, newCategory string
        var newQuantity uint
        var newCost float64

        fmt.Println("")
        fmt.Println("Modify item")
        fmt.Println("===========")
        fmt.Printf("Which item do you wish to modify? ")
        fmt.Scanln(&itemName)

        fmt.Printf("Category: %v - Item: %v  Quantity: %v  Unit Cost: $%.2f\n", categories[itemMap[itemName].category], itemName, itemMap[itemName].quantity, itemMap[itemName].cost)

        fmt.Println("")
        fmt.Println("Enter new name. Enter for no change.")
        fmt.Scanln(&newName)
        fmt.Println("Enter new category. Enter for no change.")
        fmt.Scanln(&newCategory)
        fmt.Println("Enter new quantity. Enter for no change.")
        fmt.Scanln(&newQuantity)
        fmt.Println("Enter new unit cost. Enter for no change.")
        fmt.Scanln(&newCost)

        // Change name
        if newName == "" {
          fmt.Println("No changes to name made.")
        } else {
          itemMap[newName] = item{itemMap[itemName].category, itemMap[itemName].quantity, itemMap[itemName].cost}
          delete(itemMap, itemName)
          // Make sure all other fields are updating the right item
          itemName = newName
        }

        if newCategory == "" {
          fmt.Println("No changes to category made.")
        } else {
          // Search through slice to find index value of category
          var cat uint
          for index, element := range categories {
            if newCategory == element {
              cat = uint(index)
            }
          }

          itemMap[itemName] = item{cat, itemMap[itemName].quantity, itemMap[itemName].cost}
        }

        if newQuantity == 0 {
          fmt.Println("No changes to quantity made.")
        } else {
          itemMap[itemName] = item{itemMap[itemName].category, newQuantity, itemMap[itemName].cost}
        }

        if newCost == 0.0 {
          fmt.Println("No changes to cost made.")
        } else {
          itemMap[itemName] = item{itemMap[itemName].category, itemMap[itemName].quantity, newCost}
        }

      // Delete item
      case 5:
        var itemDelete string
        
        fmt.Println("")
        fmt.Println("Delete item")
        fmt.Println("===========")
        fmt.Printf("Enter item name to delete: ")
        fmt.Scanln(&itemDelete)

        if _, found := itemMap[itemDelete]; found {
          delete(itemMap, itemDelete)
          fmt.Printf("Deleted %v\n", itemDelete)
        } else {
          fmt.Println("Item not found. Nothing to delete!")
        }

      // Print current data fields
      case 6:
        fmt.Println("")
        fmt.Println("Print current data")
        fmt.Println("==================")

        if len(itemMap) == 0 {
          fmt.Println("No data found!")
        } else {
          for key, element := range itemMap {
            fmt.Printf("%v - %v\n", key, element)
          }
        }

      // Add new category
      case 7:
        var newCat string

        fmt.Println("")
        fmt.Println("Add new category")
        fmt.Println("================")
        fmt.Printf("What is the new category name to add? ")
        fmt.Scanln(&newCat)

        if len(newCat) == 0 {
          fmt.Println("No input found!")
        } else {

          // Search through slice to find index value of category
          var catIndex int = -1
          for index, element := range categories {
            if newCat == element {
              catIndex = index
            }
          }

          if catIndex == -1 {
            categories = append(categories, newCat)
            fmt.Printf("New category: %v added at index %v\n", newCat, len(categories) - 1)
          } else {
            fmt.Printf("Category: %v already exists at index %v!\n", newCat, catIndex)
          }

        }

      // Modify existing category
      case 8:
        modifyCategory()

      // Delete existing category
      case 9:
        deleteCategory()

      // Exit
      case 10:
        n = input
        break

    }
  }
}
