package main

import "fmt"

func main() {
  var publisher string = "Marvel Comics"
  var writer string = "Stan Lee"
  var artist string = "Jack Kirby"
  var title string = "Amazing Fantasy #15"
  var year int = 1962
  var pageNumber int = 12
  var grade float32 = 9.4

  // Print information using formatted strings
  fmt.Println("Title:", title)
  fmt.Println("Published by:", publisher)
  fmt.Println("Written by:", writer)
  fmt.Println("Drawn by:", artist)
  fmt.Println("Year:", year)
  fmt.Println("Pages:", pageNumber)
  fmt.Printf("Grade: %.2f\n", grade) // Use format specifier for float output
  // Define comic book data (second comic)
  publisher2 := "DizzyBooks Publishing Inc."
  writer2 := "Ryan N. Shawn"
  artist2 := "Phoebe Paperclips"
  title2 := "Epic Vol. 1"
  year2 := 2013
  pageNumber2 := 160
  grade2 := 9.0

  // Print information for the second comic
  fmt.Println("**Comic Book 2:**")
  fmt.Println("Title:", title2)
  fmt.Println("Published by:", publisher2)
  fmt.Println("Written by:", writer2)
  fmt.Println("Drawn by:", artist2)
  fmt.Println("Year:", year2)
  fmt.Println("Pages:", pageNumber2)
  fmt.Printf("Grade: %.2f\n", grade2)

// Define comic book data (third comic)
	publisher3 := "Image Comics"
	writer3 := "Robert Kirkman"
	artist3 := "Charlie Adlard"
	title3 := "The Walking Dead"
	year3 := 2003
	pageNumber3 := 32
	grade3 := 9.2

	// Print information for the third comic
	fmt.Println("**Comic Book 3:**")
	fmt.Println("Title:", title3)
	fmt.Println("Published by:", publisher3)
	fmt.Println("Written by:", writer3)
	fmt.Println("Drawn by:", artist3)
	fmt.Println("Year:", year3)
	fmt.Println("Pages:", pageNumber3)
	fmt.Printf("Grade: %.2f\n", grade3)

	// Define comic book data (fourth comic)
	publisher4 := "IDW Publishing"
	writer4 := "Kevin Eastman"
	artist4 := "Peter Laird"
	title4 := "Teenage Mutant Ninja Turtles"
	year4 := 1984
	pageNumber4 := 32
	grade4 := 9.0

	// Print information for the fourth comic
	fmt.Println("**Comic Book 4:**")
	fmt.Println("Title:", title4)
	fmt.Println("Published by:", publisher4)
	fmt.Println("Written by:", writer4)
	fmt.Println("Drawn by:", artist4)
	fmt.Println("Year:", year4)
	fmt.Println("Pages:", pageNumber4)
	fmt.Printf("Grade: %.2f\n", grade4)

	// Define comic book data (fifth comic)
	publisher5 := "Dark Horse Comics"
	writer5 := "Frank Miller"
	artist5 := "Sin City"
	title5 := "Sin City"
	year5 := 1991
	pageNumber5 := 288
	grade5 := 9.6

	// Print information for the fifth comic
	fmt.Println("**Comic Book 5:**")
	fmt.Println("Title:", title5)
	fmt.Println("Published by:", publisher5)
	fmt.Println("Written by:", writer5) // Note: "artist5" should be "artist5:"
	fmt.Println("Drawn by:", artist5)
	fmt.Println("Year:", year5)
	fmt.Println("Pages:", pageNumber5)
	fmt.Printf("Grade: %.2f\n", grade5)
}
