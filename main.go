package main

import "fmt"
import "reflect"
// comentario
/*
 comentarios
*/
const Curso string = "Curso de Go" // constante; error si cambias le valor
const ( // declarar varias constantes
	Lunes int = 1
	Martes int 2
)
const ( // asignar secuencia
	Miercoles int = iota+3
	Jueves
	Viernes 
	Sabado
	Domingo
)
func main() {
	fmt.Println("hola mundo")
	//-------------
	var nombre string // ""
	var edad int // 0; si declaro una variable y no la uso da error

	nombre = "jesus" // asignar valor
	edad = 30 // "30" error si usas un valor de tipo diferente

	fmt.Println(nombre)
	fmt.Println(edad)
	//-------------
	apellido := "Gil" // puedes usar <var> o <:=> declarar variables
	var altura = 1.75 // float; go tiene inferencia de tipo

	fmt.Println("datos:", nombre, apellido, edad, altura) // imprimir varios argumentos
	//-------------
	// var name, last_name, country string
	// var name, last_name, country = "Jesus", "Gil", "Colombia"
	name, last_name, country := "Jesus", "Gil", "Colombia" // formas de declarar en una sola linea
	age, height := 30 , 1.75 
	fmt.Println(name, last_name, age, height, country)
	//-------------
	fmt.Println(Curso) // las constantes se declaran fuera del main
	//-------------
	var peso = 70.0
	var comparacion = peso > 5 // > < == >= <= !=
	fmt.Println(comparacion) // true; puedes comparar strings tambien
	//-------------
	resultado := 20 == 20 && 30 == 30 // and <&&> or <||> negacion <!>
	fmt.Println(resultado) // true  
	//-------------
	fmt.Println(Lunes, Martes)
	fmt.Println(Miercoles, Jueves, Viernes, Sabado, Domingo)
	//-------------
	titulo := "Triagulo  de las Bermudas"
	longitud := len(titulo) // int
	primerCaracter := titulo[0] // Char; ultimo [len(Curso)-1]
	fmt.Println("logitud:", longitud)
	fmt.Println("primer caracter ascii", primerCaracter)
	fmt.Println("tipo de dato", reflect.TypeOf(primerCaracter))
	fmt.Printf("%c\n", primerCaracter)
	//-------------
	fmt.Println("Ingresa tu altura")
	fmt.Scanf("%f", &altura) // %d para enteros
	fmt.Println("Ingresa tu Peso")
	fmt.Scanf("%f", &peso)
	imc := peso/altura
	fmt.Println("imc: %.2f\n", imc)
}