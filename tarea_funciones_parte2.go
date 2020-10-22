package main

import "fmt"

// Fibonacci _______________________________________________________________
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Variadic Function ... Número más grande__________________________________
func bigger_num(arg ...int) int {
	bigger := 0

	for _, v := range arg {
		if v > bigger {
			bigger = v
		}
	}
	return bigger
}

// Generador de números impares_____________________________________________
func generadorImpares() func() uint {
	i := uint(1) // i permanecerá en el clousure de la función anónima a retornar
	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

// Intercambiar dos valores_________________________________________________
func intercambia(a, b *int) {
	c := *a
	*a = *b
	*b = c
}

func main() {
	var opc int

	for true {
		fmt.Println("Selecciona una opción: ")
		fmt.Println("1. Secuencia de Fibonacci")
		fmt.Println("2. Obtener el número más grande")
		fmt.Println("3. Generador de Impares")
		fmt.Println("4. Intercambiar dos valores")
		fmt.Println("5. Salir")

		fmt.Scan(&opc)

		switch {
		// Fibonacci _______________________________________________________________
		case opc == 1:
			var n int
			fmt.Println("Dame un número ")
			fmt.Scan(&n)
			fmt.Println(fibonacci(n))

		// Variadic Function ... Número más grande__________________________________
		case opc == 2:
			a := []int{5, 7, 20, 17, 3, 21, 2, 18}
			fmt.Println(bigger_num(a...))

		// Generador de números impares_____________________________________________
		case opc == 3:
			nextImpar := generadorImpares()
			fmt.Println(nextImpar())
			fmt.Println(nextImpar())
			fmt.Println(nextImpar())
			fmt.Println(nextImpar())
			fmt.Println(nextImpar())

		// Intercambiar dos valores_________________________________________________
		case opc == 4:
			var a int
			var b int
			fmt.Scan(&a)
			fmt.Scan(&b)

			intercambia(&a, &b)
			fmt.Println("a = ", a)
			fmt.Println("b = ", b)

		// Salir _______________________________________________________________
		case opc == 5:
			return
		}
	}

}
