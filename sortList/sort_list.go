package sortList

/*
Algoritmo de ordenação, que usa duas funções diferentes, uma complementa a outra,
getSmallest retorna o indice do menor valor do array, e então,
sortList adiciona a uma nova lista o valor retornado por getSmallest
*/
func getSmallest(lista []int) int {
	smallest := lista[0]
	index_smallest := 0
	for i := 1; i < len(lista); i++ {
		if lista[i] < smallest {
			smallest = lista[i]
			index_smallest = i
		}
	}
	return index_smallest
}

func Sort(lista []int) []int {
	novaLista := []int{}
	for i := 0; i < len(lista)-1; {
		menor := getSmallest(lista)
		novaLista = append(novaLista, lista[menor])
		lista[menor] = lista[len(lista)-1]
		lista[len(lista)-1] = 0
		lista = lista[:len(lista)-1]
	}
	return novaLista
}
