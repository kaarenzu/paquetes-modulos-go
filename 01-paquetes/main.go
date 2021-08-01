package main

// el paquete main es el paquete ejecutable con el que hemos venido trabajando hasta el momento.
// Y hemos estado utilizando los paquetes oficiales de Go.
// Un paquete es una carpeta que contiene una colección de archivos y nos proveen una funcionalidad
// Es importante tener un nombre de paquete claro y concisos que indiquen claramente que es lo que hace el paquete porque siempre se utilizarán como prefijo es por eso que en Go nunca van a ver un paquete útil que provea otras funcionalidades, siempre debe ser muy especifica la funcionalidad que va a hacer el paquete
import "fmt"

func main() {
	// nombre del paquete.nombre de la función()
	fmt.Println("gokaaa")
}
