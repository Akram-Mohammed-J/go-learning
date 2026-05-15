package main
import ("fmt")

func main() {
  var marks [5] int 
  marks[0] =35
  marks[1] =45
  marks[2] =55
  marks[3] =65
  marks[4] =95
 fmt.Println(marks)


 // declaration + initialization
 var days[7]string = [7]string {"Monday"}
 fmt.Println(days)	

 //inferred type declaration and initialization
 students  :=[4] string {"Fouzi","Akram", "Aaisha", "Abrar"}

 fmt.Println(students)

}