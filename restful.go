package main

import "fmt"

func main() {

	// string variables
	// var publisher string = "DizzyBooks Publishing Inc."
	var writer string = "Tracey Hatchet"
	var artist string = "Jewel Tampson"
	var title string = "Mr. GoToSLeep"

	// integer variables
	// var year int = 1997
	// var pageNumber int = 14
	var grade float32 = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "grade at ", grade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	// artist = "Phoebe Paperclips"
	// publisher = "DizzyBooks Publishing Inc."
	// year = 2013
	// pageNumber = 160
	grade = 9

	fmt.Println(title, "written by", writer, "drawn by", artist, "grade at ", grade)
}
