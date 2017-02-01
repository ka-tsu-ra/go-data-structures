package main

import (
  "fmt"
)

type Person struct {
  name string
  age  int
}

func main() {
  // ! NOTE ! CAN'T DECLARE ANOTHER FUNC INSTIDE 'MAIN' FUNC??

  numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  for _, value := range numbers {
    if value%2 == 0 {
      fmt.Println(value)
    }
  }

  Agnes := Person{"Agnes", 23}
  Maria := Person{"Maria", 35}

  people := []Person{Agnes, Maria}

  for _, person := range people {
    fmt.Println(person.name + ", you were born in " + fmt.Sprint(calculateYear(person.age)))
  }

  // ANOTHER VERSION - TRYING TO REFACTOR STUFF INTO SEPARATE FUNCTIONS
  for _, person := range people {
    statePersonBirthYear(person)
  }

  // ANOTHER VERSION - LESS COUPLED TO PERSON TYPE
  for _, person := range people {
    statePersonBirthYearDecoupled(person.name, calculateYear(person.age))
  }

}

func calculateYear(age int) int {
  return 2016 - age
}

// IF YOU HAVE A RETURN STATEMENT THEN YOU NEED TO SPECIFY THE TYPE OF THING BEING RETURNED.
// BUT IF YOU END THE FUNCTION WITH A FUNCTION CALL LIKE fmt.Println.... THEN YOU DON'T SPECIFY ANY TYPE.

// THIS FUNC MIGHT NOT BE GOOD BECAUSE IT IS TIGHTLY COUPLED TO THE PERSON STRUCT TYPE - I.E. SOMETHING THAT YOU KNOW
// HAS A NAME AND AN AGE PROPERTY.

func statePersonBirthYear(person Person) {
  fmt.Println(person.name + ", you were born in " + fmt.Sprint(calculateYear(person.age)))
}

// ANOTHER VERSION, NOT SO TIGHTLY COUPLED - IE CAN TAKE ANY STRING FOR NAME AND INT FOR AGE AND YOU DECIDE HOW TO FEED
// THOSE TO IT WHEN YOU CALL IT:

func statePersonBirthYearDecoupled(name string, age int) {
  fmt.Println(name + ", you were born in " + fmt.Sprint(age))
}

