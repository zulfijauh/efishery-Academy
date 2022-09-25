package main

import "fmt"

// Create function
func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}
// Create struct	
type person struct {
	name string
	age int
}

// Embeded struct
type student struct {
	grade int
	person
}

func main() {
	// Print
	fmt.Print("Hello World \n\n")

	var decimal = 2.62823
	// print bilangan desimal
	fmt.Printf("bilangan desimal: %f\n", decimal)
	// print hanya 3 desimal
	fmt.Printf("bilangan desimal: %.3f\n\n", decimal)

	// Boolean variable
	var exist bool = true
	fmt.Printf("exist? %t \n\n", exist)

	// Contoh variable kalimat / paragraf
	var message = `Nama Saya "Jhon Whick".
salam kenal.
Mari belajar "golang"`
	fmt.Println(message, "\n\n")

	// Print 2 variable
	var firstName string = "John"
	var lastName string = "Wick"
	fmt.Printf("Halo %s %s!\n\n", firstName, lastName)

	// Membuat variable tanpa harus define ulang data type
	var namaDepan string = "Adam"
	namaBelakang := "Levine"
	fmt.Printf("Hallo %s %s! \n\n", namaDepan, namaBelakang)

	// Define variable sekaligus
	var first, second, third = "satu", "dua", "tiga"
	fmt.Print(first, " ", second, " ", third, "\n\n")

	// Gunakan underscore untuk stoce variable yang tidak akan dipakai
	_ = "Belajar Golang"

	// Fixed variable, tidak bisa diganti
	const sureName string = "Hylos"
	fmt.Print("Hello ", sureName, "!\n\n")

	var point = 8
// If function
	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point >= 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("Tidak lulus. Nilai anda %d\n", point)
	}
// switch
	var nilai = 10
	switch nilai {
	case 8:
		fmt.Printf("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Printf("not bad")
	}
// for function
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruits {
		fmt.Printf("elemen %d: %s\n", i, fruit)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}
// swap function
	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	fmt.Print(add(7, 3))

// Use datatype from struct
var s1 person
s1.name = "John Wick"
s1.age = 23

fmt.Println("name: ",  s1.name, "\n")
fmt.Println(("grade: ", s1.age, "\n")

// use datatype from embeded struct
var s2 = student
s2.name = "Johnson" 
s2.age = 17
s2.grade = 8

fmt.Println("name: ",  s2.name, "\n")
fmt.Println(("age: ", s2.age, "\n")
fmt.Println(("age: ", s2.person.age, "\n")
fmt.Println(("grade: ", s2.grade, "\n")
}