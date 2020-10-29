package gaussiana

import "fmt"

var matriz [][]float64

//PideDatos pide el tamaño de la matriz
func pideDatos() int {
	var n int
	fmt.Println("introduce el tamaño de la matriz")
	fmt.Scanln(&n)
	return n
}

//llenaMatriz llena la matriz con los valores que el usuario quiere que se calculen
func llenarMatriz(n int) [][]float64 {
	var llena float64
	var m = n+1
	matriz := make([][]float64, n)
	for i := 0; i < n; i++ {
		matriz[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			fmt.Println("introduce el valor de ", i, j)
			fmt.Scanln(&llena)
			matriz[i][j]=llena
		}
	}
	return matriz
}

//imprimeMatriz unicamente imprime la matriz
func imprimeMatriz(matriz [][]float64, n int)  {
	for i := 0; i < n; i++ {
		for j := 0; j < (n+1); j++ {
			fmt.Print("\t ", matriz[i][j]," | " )
		}
		fmt.Println("")
	}
}

//calcularPivote toma un valor de la diagonal principal y lo convierte en 1
func calcularPivote(matriz [][]float64, pivote int, n int) [][]float64 {
	var aux float64
	aux = matriz[pivote][pivote]
	for i := 0; i < (n+1); i++ {
		matriz[pivote][i]=matriz[pivote][i]/aux
	}
	return matriz
}

//calcularCeros calcula los ceros despues de haber convertido en 1 el valor pivote en la diagonal
//principal
func calcularCeros(matriz [][]float64, pivote int, n int) [][]float64 {
	var c float64
	for i := 0; i < n; i++ {
		if i!= pivote{
			c = matriz[i][pivote]
			for j := 0; j < (n+1); j++ {
				matriz[i][j]=((-1*c)* matriz[pivote][i])+matriz[i][j]
			}
		}
	}
	return matriz
}

//Unir unicamente une todos los metodos anteriores
func Unir()  {
	pivote :=0
	n:=pideDatos()
	matriz=llenarMatriz(n)
	imprimeMatriz(matriz, n)
	for i := 0; i < n; i++ {
		matriz=calcularPivote(matriz, pivote, n)
		imprimeMatriz(matriz, n)
		matriz=calcularCeros(matriz, pivote, n)
		imprimeMatriz(matriz, n)
		pivote++
	}
	for i := 0; i < n; i++ {
		fmt.Println("el valor de x ", (i+1), " = ", matriz[i][n])
	}
}