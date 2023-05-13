package main

/*
	var bar string
	var num int
	var boo bool
	var byt byte

	bar = "Hello"
	num = 12635
	boo = true
	byt = 0x36

	name := "Agustin"

	println(bar)
	println(num)
	println(boo)
	println(byt)
	println(name)

*/

// Esta función es la función main desde donde se ejecuta el programa.
func main() {
	// Yo quiero que, en base a edad
	// se imprima un texto tal que
	// si la edad > 18 => "Mayor de Edad" sino "Menor de Edad"

	age := 1

	// Muchos else-if es mala practica.
	if age > 18 {
		println("Mayor de Edad")
	} else if age < 3 {
		println("Lactante")
	} else if age < 10 {
		println("Infante")
	} else {
		println("Joven")
	}

	// Buena practica.
	switch {
	case age > 18:
		println("Mayor de Edad")
	case age < 3:
		println("Lactante")
	case age < 10:
		println("Infante")
	default:
		println("Joven")
	}
}
