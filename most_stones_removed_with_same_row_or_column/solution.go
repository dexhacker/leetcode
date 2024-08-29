package most_stones_removed_with_same_row_or_column

type Point struct {
	X int
	Y int
}

var mapX map[int][]*Point
var mapY map[int][]*Point
var used map[*Point]bool

func dfs(point *Point) int {
	if used[point] {
		return 0
	}

	answer := 1
	used[point] = true

	for _, item := range mapX[point.X] {
		answer += dfs(item)
	}
	for _, item := range mapY[point.Y] {
		answer += dfs(item)
	}
	return answer
}

func removeStones(stones [][]int) int {
	mapX = make(map[int][]*Point)
	mapY = make(map[int][]*Point)
	used = make(map[*Point]bool)
	points := make([]*Point, 0)

	for _, stone := range stones {
		point := &Point{X: stone[0], Y: stone[1]}
		points = append(points, point)
		mapX[point.X] = append(mapX[point.X], point)
		mapY[point.Y] = append(mapY[point.Y], point)
	}

	answer := 0
	for _, point := range points {
		result := dfs(point)
		if result > 1 {
			answer += result - 1
		}
	}
	return answer
}
