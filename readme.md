This assignment was submitted for the GoSchool Go Basics module.

### Assignment

**Background**
It is desired to develop a console application that would serve as a simple shopping list application to store the shopping list of the day.

**Application Feature**
The shopping list application shall contain the following features.

**Main Menu**

1.  Shopping List Menu  
    a. The shopping list menu shall on runtime prompt and allow the user to make a selection.  
    b. Upon selection made, the respective features detailed in “Features” shall be executed.

**Features**

1.  Selection “1” - View Entire Shopping List  
    a. This feature shall display all available items in the shopping list.
2.  Selection “2” - Generate Shopping List Report  
    This feature shall generate the reports based on the available items stored at runtime.
    a. The following types of reports would be generated - Total cost of item by category  
     - Calculate by the summation of all items of the same category  
     - Each item total cost is calculated by Item \* Unit Cost  
     - List of items by category  
     - Each item available will be “grouped” into the relevant category as shown in Figure 6.  

    b. Selecting “3. Main Menu” will return the user to the main menu.

3.  Selection “3” - Add Items information  
    a. This feature adds information for the item of interest.

    - Name of the item
    - Category of the item
    - Unit Cost of the item
    - Quantity of the item

4.  Selection “4” - Modify Existing Items  
    a. This feature allows for the user to change the quantity of the item, regardless of the category.

5.  Selection “5” - Delete Item from shopping list  
    a. This feature deletes the item that is specified by the user.  
    b. If the item does not exist, the user is notified.

6.  Selection “6” - Print current data fields  
    a. This feature allows the display how the data is stored in the application.  
    b. If there is not data available, the user is notified.

7.  Selection “ 7” - Add New Category Name  
    a. This feature allows the user to add new categories to the existing categories.  
    b. If the category exists, the user is notified  
    c. If there is no input from the user, a “no input” is shown and the main menu is shown.

**Advanced Requirements (Optional)**

1.  Other features – Modify and delete category  
    a. Modify category allows the user to modify the category of interest.  
    b. Delete category allows the user delete the category of interest.  
    c. This section is optional and open for each participant to design their own.  
    d. The design of the features should be logical. E.g. deletion of category would delete all stored category as well since it is no longer available.  
    e. Relevant warning should be given to the user as notifications. E.g. warning for nothing to modify or nothing to delete.  
    f. Deletion of category would reshuffle existing indexes of categories that are still available. E.g. deletion of Food at index 1 would make Drinks be reallocated as index 1.

2.  Other features – Storing of shopping lists and retrieving of lists  
    a. Multiple shopping lists can be stored in a slice with allocated indexes using a “save shopping list” option.  
    b. User would be able to retrieve previous shopping indexes using a “retrieve previous shopping list” option and providing an index of interest.  
    c. The existing features should be able to access the retrieved shopping list.

**Development Considerations**
The following are development considerations.

1.  Category  
    a. The categories shall consist of the following stored in a slice:

    - Household
    - Food
    - Drinks

    b. The categories can be entered by the user in plain text.  
    c. The category selected by the user is stored in the application in index form. i.e. Household as 0, Food as 1 and Drinks as 2.

2.  Item information  
    a. The following information of the item shall be stored in a struct

    - Category of item - “int”
    - Quantity of item - “int”
    - Unit Cost of item - “float64”

3.  Item Mapping to Category  
    a. The following information of the item shall be stored in a map
    - Key : Item Name, Value: Item information
4.  Shopping List Menu  
    a. The list menu shall be an infinite loop, each loop requesting for an input from the user unless it is in a feature of choice.

**Assignment Comments**
The comments for the assignment will be based on the following:

1.  Completeness and robustness.
2.  Development design compliance.
