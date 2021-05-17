package main
import "fmt"

func main() {
	var numeros[5] int // declaración de arreglos
	fmt.Println(numeros) // [0,0,0,0,0]

	numeros[0] = 100 // asignación
	fmt.Println(numeros) // [100,0,0,0,0]
	fmt.Println(numeros[0]) // 100

	edades := [5] int {10, 20, 22, 30, 35} // declaración/asignación
	fmt.Println(edades)
	
	fechas := [...] int {1990, 1991, 2007} // go calcula el tamaño
	fmt.Println(fechas)
	//----------
	monedas := [...] string {0:"peso", 2:"euro", 1:"dolar"} // definir índice
	fmt.Println(monedas)
	//----------
	dinamico := [] int {1,2,3,4}
	dinamico = append(dinamico, 5) // agregar valores
	subdinamico := dinamico [0:3] // slice
	dinamico[0] = 100 
	fmt.Println(dinamico) // [100,2,3,4,5]
	fmt.Println(subdinamico) // [100,2,3] los slices apuntan al mismo array
	//----------
	meses := [] string {"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre"}
	tamano := len(meses)
	capacidad := cap(meses)
	fmt.Println(tamano, capacidad)
	meses = meses.append(meses, "octubre")
	tamano = len(meses)
	capacidad = cap(meses)
	fmt.Println(tamano, capacidad)
	//----------
	slice := make([] int, 0, 3) // otra forma de crear slice
	slice[0] = 100
	slice[1] = 200
	slice[2] = 300
	slice = append(slice, 400)
	fmt.Println(slice, len(slice), cap(slice))
	//-----------
	dias := make(map[int]string) // mapas llave valor
	dias[0] = "lunes"
	dias[1] = "martes"
	dias[2] = "miercoles"
	dias[3] = "jueves"
	dias[4] = "viernes"
	delete(dias, 4)
	fmt.Println(dias, dias[0])

	estudiante := make(map[string] []int)
	estudiante["jesus"] = []int {10,9,8,10,5}
	fmt.Println(estudiante)

	usuarios := map[int] string {} // otra forma de mapa
	usuarios[1] = "usuarrio 1"
	usuarios[2] = "usuarrio 2"
	usuarios[3] = "usuarrio 3"
	usuarios[4] = "usuarrio 4"
	for llave, valor := range usuarios {
		fmt.Println(llave, valor)
	}

}