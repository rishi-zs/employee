package main

type emp struct {
	name string
	age  int
}

func checker(e emp) (b bool, val emp) { //checking whether the person age is less than 22 or not.
	if e.age < 22 {
		return
	} else {
		return true, e
	}
}

func main() {
	p := emp{name: "Ram", age: 21}
	checker(p)
}
