# Expense Tracker CLI

This is an application that allow users to add, delete, and view their expenses. The application also provide a summary of the expenses. A solution to a challenge by [Expense Tracker](https://roadmap.sh/projects/expense-tracker)

- [x] Users can add an expense with a description and amount.
- [x] Users can update an expense.
- [x] Users can delete an expense.
- [x] Users can view all expenses.
- [x] Users can view a summary of all expenses.
- [ ] Users can view a summary of expenses for a specific month (of current year).
- [ ] Allow users to set a budget for each month and show a warning when the user exceeds the budget.

### How to Install

1. Clone the github repository

2. Change directory into the clone repository folder on your device
   
   ```
   go build -o tracker.exe
   ```

### Usage

- To Add Expense
  
  ```
  ./tracker add --desc "Groceries shopping" --amount "30.5"
  ```

- To Delete Expense
  
  ```
  ./tracker delete --id 3
  ```

- To Update Expense
  
  ```
  ./tracker delete --id 3
  ```
