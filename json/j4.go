package main

import (
	"fmt"
	"reflect"
)

type T struct {
	f string `one:"1" two:"2"blank:""`
}

// Lookup function returns two values — value associated with key (or blank if not set) 
// and bool indicating if key has been found at all
func main() {
	t := reflect.TypeOf(T{})
	f, _ := t.FieldByName("f")
	fmt.Println(f.Tag) // one:"1" two:"2"blank:""
	v, ok := f.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // 1, true
	v, ok = f.Tag.Lookup("blank")
	fmt.Printf("%s, %t\n", v, ok) // , true
	v, ok = f.Tag.Lookup("five")
	fmt.Printf("%s, %t\n", v, ok)  // , false
	fmt.Println(f.Tag.Get("five")) //
	fmt.Println(f.Tag.Get("two"))  // 2
}
