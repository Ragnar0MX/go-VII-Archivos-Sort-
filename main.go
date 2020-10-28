package main

import (
	"fmt"
	"os"
	"sort"
)

type ByCadena []string

func (a ByCadena) Len() int           { return len(a) }
func (a ByCadena) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCadena) Less(i, j int) bool { return a[i] < a[j] }

func imprimir(S []string) {
	for _, m := range S {
		println(m)
	}
}

func ascendente(S []string) {
	file, err := os.Create("ascendente.txt")
	if err != nil {
		println(err)
	}

	for _, v := range S {
		file.WriteString(v)
		file.WriteString("\n")
	}
	defer file.Close()

}

func descendete(S []string) {
	file, err := os.Create("descendete.txt")
	if err != nil {
		println(err)
	}
	for _, v := range S {
		file.WriteString(v)
		file.WriteString("\n")
	}
	defer file.Close()

}
func main() {
	S := make([]string, 0)
	var opc int
	var auxP string

	menu := "    Menu    \n1.-AGREGAR\n2.-Salir"
	for ok := true; ok; ok = (opc != 2) {
		fmt.Println(menu)
		fmt.Scan(&opc)
		if opc == 1 {
			fmt.Println("Introduce la palabra:")
			fmt.Scan(&auxP)
			S = append(S, auxP)
		}
		if opc == 2 {
		}
	}
	sort.Sort(ByCadena(S))
	ascendente(S)
	sort.Sort(sort.Reverse(ByCadena(S)))
	descendete(S)

}
