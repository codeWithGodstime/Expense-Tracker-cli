package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

var SliceOfExpenses []Expense

func readFile() []Expense {
	var (
		expenses []Expense
		file     *os.File
		err      error
	)

	filepath := "data.json"
	if file, err = os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644); err != nil {
		log.Fatalf("Failed to open file '%s': %v", filepath, err)
	}

	defer file.Close() // Ensure file is closed
	fileContent, _ := io.ReadAll(file)
	if err := json.Unmarshal(fileContent, &expenses); err != nil {
		log.Fatalf("Failed to decode json %v", err)
	}
	return expenses
}

func saveFile(expenses []Expense) {

	filepath := "./data.json"

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open file '%s': %v", filepath, err)
	}
	defer file.Close() // Ewnsure file is closed

	jsonData, err := json.MarshalIndent(expenses, "", "  ")

	if err != nil {
		_ = fmt.Errorf("failed to marshal SliceOfExpenses %w", err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		_ = fmt.Errorf("failed to write to file: %w", err)
	}
	fmt.Println("Data successfully written to file")
}

// Users can add an expense with a description and amount.
func Add(description string, amount float64) error {
	newIndex := len(SliceOfExpenses) + 1
	e := &Expense{
		ID:          newIndex,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}

	SliceOfExpenses = append(SliceOfExpenses, *e)
	// save to file
	saveFile(SliceOfExpenses)
	return nil
}

// Users can update an expense.
func Update(index int, data ...string) {
	if len(SliceOfExpenses) > index {
		SliceOfExpense := SliceOfExpenses[index]

		if len(data) > 0 {
			description := data[0]
			SliceOfExpense.Description = description
			SliceOfExpenses[index] = SliceOfExpense
			saveFile(SliceOfExpenses)
			fmt.Println("SliceOfExpenses updated successfully", SliceOfExpenses[index])
		}
	} else {
		fmt.Println("Invalid Index")
	}
}

// Users can delete an expense.
func Delete(index int) {
	if len(SliceOfExpenses) > 0 {
		SliceOfExpenses = append(SliceOfExpenses[:index], SliceOfExpenses[index+1:]...)
	}
	saveFile(SliceOfExpenses)
	fmt.Printf("Expense %x has been deleted", index)
}

// Users can view all expenses.
func List() []Expense {
	return nil
}

// Users can view a summary of all expenses.
func Summary() int {
	return 1
}

// Users can view a summary of expenses for a specific month (of current year).

func main() {
	var (
		description string
		amount      string
		id          int
	)

	SliceOfExpenses = readFile()

	// add expense
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	addCommand.StringVar(&description, "desc", "", "Description of the expense to be added")
	addCommand.StringVar(&amount, "amount", "", "Amount you spent ")

	// delete expense
	deleteCommand := flag.NewFlagSet("id", flag.ExitOnError)
	deleteCommand.IntVar(&id, "id", 0, "Unique id of expense you want to delete")

	// list expense
	listCommand := flag.NewFlagSet("id", flag.ExitOnError)

	// // update expense
	// updateCommand := flag.NewFlagSet("id", flag.ContinueOnError)
	// updateCommand.StringVar(&id, "id", "", "Unique id of expense you want to update")

	// Ensure at least one argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add' or 'delete' command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCommand.Parse(os.Args[2:])
		if description == "" || amount == "" {
			fmt.Println("Error: Both -desc and -amount flags are required for 'add' command")
			addCommand.Usage()
			os.Exit(1)
		}

		amount, _ := strconv.ParseFloat(amount, 64)
		if err := Add(description, amount); err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

	case "delete":
		deleteCommand.Parse(os.Args[2:])
		// Delete(id)
	case "list":
		listCommand.Parse(os.Args[2:])
		for _, value := range SliceOfExpenses {
			fmt.Printf("%x  %s  %s  %f\n", value.ID, value.Description, value.Date.Format("02/02/2006"), value.Amount)
		}
	case "summary":
		var summary float64

		for _, value := range SliceOfExpenses {
			summary += value.Amount
		}
		fmt.Print("Overall Summary: $", summary)
	default:
		fmt.Print("Invalid command")
	}

}
