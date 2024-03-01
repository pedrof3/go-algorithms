package binarySearch

/*
O algoritmo de binary search é uma maneira eficiente de encontrar um item em uma lista ordenada,
nesse exemplo, caso encontre o item na lista, retorna o index do item no array,
se o item não for encontrada retorna -1.
Em comparação com o algoritmo de busca linear, quanto maior for o tamanha da lista,
maior é vantagem de tempo de execução que a busca binária leva.
A notação big O da busca binária é O(log n), enquanto a de busca linear é O(n),
Por exemplo: em uma lista com 1.000 itens, O(n) leva até verificação para procurar o valor,
enquanto O(log n) leva no executa no máximo 10 verificações
*/
func Search(lista []int, n int) int {
	i, f := 0, len(lista)-1
	for i <= f {
		m := (i + f) / 2
		switch {
		case n < lista[m]:
			f = m - 1

		case n > lista[m]:
			i = m + 1

		default:
			return m
		}

	}
	return -1
}
