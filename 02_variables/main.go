package main

import "fmt"

//city := "Delhi"   //ERR:Shorthand syntax not allowed outside of function scope, compilation error

func main() {

	/*

		MAIN TYPES

		string
		bool
		int
		int int8 int16 int32 int64
		uint uint8 uint16 uint32 uint64 unitptr
		byte - alias for int8
		rune - alias for int32
		float32 float64
		complex64 complex128

	*/

	//Using var
	//var name string = "Lokesh"     //ERR: data-type is automatically inferred
	var name = "Lokesh" //FIX: devoid of keyword 'string' in the declaration

	fmt.Println(name)

	/*
		Commonly used formatters (placeholders)

		General
		===========
		%v - prints the value
		%T - prints the Type
		%% - to print percent sign

		Integer
		============
		%d - base10
		%o - base8

		Floating point
		================
		%f - decimal but no exponent, e.g. 123.456
		%F - synonym for %f
	*/

	var age = 29
	var salary int32 = 50000000
	//var hasBlueEyes = false         //ERR: unused variable throws compilation error

	const hasPlaystation = true
	//hasPlaystation = false		  //ERR: cannot assign to constant variable over and over, compilation error

	fmt.Println(name, age)
	fmt.Printf("age: %v        , type: %T \n", age, age)       //returns 29 & int
	fmt.Printf("salary: %v     , type: %T \n", salary, salary) //returns 50000000 & int32 (unless specified the type of int)

	//Shorthand
	city := "Bangalore"

	fmt.Println(city)

	//Another shorthand (multiple variables declared at once)
	majors, expInYears := "CSE", 7

	fmt.Println(majors, expInYears) //variables once declared cannot be left without using it (comment and see)
}
