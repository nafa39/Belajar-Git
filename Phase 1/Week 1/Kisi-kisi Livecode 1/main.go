package main

import (
	"bufio"
	"fmt"
	"kisi-kisi/helpers"
	"os"
	"strings"
)

var books = []helpers.Book{
	{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Available: true},
	{Title: "To Kill A Mockingbird", Author: "Harper Lee", Available: true},
	{Title: "1984", Author: "George Orwell", Available: false},
}

func showAvailableBooks(books []helpers.Book) {
	fmt.Println("\nAvailable Books:")

	for i, book := range books {
		if book.Available {
			fmt.Printf("%d. Title: %s, Author: %s, Availabel: Yes\n", i+1, book.Title, book.Author)
		} else {
			fmt.Printf("%d. Title: %s, Author: %s, Availabel: No\n", i+1, book.Title, book.Author)
		}
	}
}

func rentBook(books []helpers.Book) {
	fmt.Println("Enter the title of the book you want to rent:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	for i, book := range books {
		if book.Title == input {
			if book.Available {
				books[i].Available = false
				fmt.Printf("\nBook \"%s\" has been successfully rented!", input)
			} else {
				fmt.Printf("\nBook \"%s\" is not available for rent.", input)
			}
			return
		}
	}
	fmt.Printf("Book \"%s\" not found", input)

}

func main() {
	fmt.Println("Welcome to the Library Book Rental System!")

	for {
		fmt.Println("Main Menu:")
		fmt.Println("1. Show available books")
		fmt.Println("2. Rent a book")
		fmt.Println("3. Exit program")

		for {
			var choice int
			fmt.Println("\nPlease enter your choice (1/2/3): ")
			fmt.Scanln(&choice)

			// if _, err := fmt.Scanln(&choice); err != nil {
			// 	fmt.Println("Invalid input")
			// 	break
			// }

			switch choice {
			case 1:
				showAvailableBooks(books)
			case 2:
				for {
					rentBook(books)
					fmt.Println("\nDo you want to rent another book? (yes/no):")
					var pilih string
					if _, err := fmt.Scanln(&pilih); err != nil {
						fmt.Println("Invalid input")
					} else if pilih != "yes" {
						fmt.Println("Thank you for using the Library Book Rental System. Goodbye!")
						return
					}
				}
			case 3:
				fmt.Println("Thank you for using the Library Book Rental System. Goodbye!")
				return
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}
	}
}
