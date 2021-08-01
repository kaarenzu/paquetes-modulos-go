// indica el nombre del paquete al que pertenece en este caso greet que debe corresponder al nombre de la carpeta
package greet

// Hasta el momento hemos visto que hemos declarado variables en la función y las utilizabamos ahi mismo eso se conoce como el scope de la función, pero también podemos declarar variables a nivel de paquete es decir si la declaramos fuera de la función esas variables que declaremos alli las podemos utilizar en cualquier lugar del paquete

// No las puedo utilizar con al asignador de var corta := lo debo hacer con la palabra reservada var y puedo compartir entre paquetes todo lo que yo necesite es decir puedo utilizar variables, constantes,estructuras,funciones que se pueden compartir entre diferentes archivos que pertenezcan al mismo paquete.

// Debo tener en cuenta que cuando yo estoy construyendo un paquete quiero que algunos identificadores sean exportados y otros No por ej si no quiero que la variable emoji la puedan exportar desde otros paquetes pero si la funciones, entonces :
// ¿Cómo le digo a Go qué exportar y que NO Exportar?  es tan sencillo como identificar en la primera letra del identificador(nombre) en máyuscula si quiero exportar y en minuscula si no la exporto
var emoji = "<3"

// GreetEnglish retorna el saludo en ingles
func GreetEnglish() string {
	return "Hi" + emoji
}

// Italiano retorna el saludo en italiano

func Italian() string {
	return "Ciao" + emoji
}
