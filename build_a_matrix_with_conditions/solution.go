package build_a_matrix_with_conditions

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowGraph := make([][]int, k+1)
	colGraph := make([][]int, k+1)
	inDegreeRow := make([]int, k+1)
	inDegreeCol := make([]int, k+1)

	for _, cond := range rowConditions {
		a, b := cond[0], cond[1]
		rowGraph[a] = append(rowGraph[a], b)
		inDegreeRow[b]++
	}

	for _, cond := range colConditions {
		a, b := cond[0], cond[1]
		colGraph[a] = append(colGraph[a], b)
		inDegreeCol[b]++
	}

	rowOrder := topologicalSort(rowGraph, inDegreeRow, k)
	if len(rowOrder) == 0 {
		return [][]int{}
	}

	colOrder := topologicalSort(colGraph, inDegreeCol, k)
	if len(colOrder) == 0 {
		return [][]int{}
	}

	matrix := make([][]int, k)
	for i := range matrix {
		matrix[i] = make([]int, k)
	}
	rowMap := make(map[int]int)
	for i, num := range rowOrder {
		rowMap[num] = i
	}
	for j, num := range colOrder {
		matrix[rowMap[num]][j] = num
	}
	return matrix
}

func topologicalSort(graph [][]int, inDegree []int, n int) []int {
	var queue []int
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	var order []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(order) != n {
		return []int{}
	}
	return order
}
