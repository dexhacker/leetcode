package maximum_total_importance_of_roads

import "sort"

type City struct {
	Name  int
	Count int64
	Value int
}

func maximumImportance(n int, roads [][]int) int64 {
	mp := make(map[int]*City)
	for _, road := range roads {
		a := road[0]
		b := road[1]

		if val, ok := mp[a]; ok {
			val.Count += 1
		} else {
			mp[a] = &City{Name: a, Count: 1, Value: 0}
		}

		if val, ok := mp[b]; ok {
			val.Count += 1
		} else {
			mp[b] = &City{Name: b, Count: 1, Value: 0}
		}
	}

	cities := make([]*City, 0)
	for _, v := range mp {
		cities = append(cities, v)
	}

	sort.Slice(cities, func(i, j int) bool {
		return cities[i].Count > cities[j].Count
	})

	for _, city := range cities {
		city.Value = n
		n--
	}

	var answer int64 = 0
	for _, road := range roads {
		a := road[0]
		b := road[1]

		cityA := mp[a]
		cityB := mp[b]
		answer += int64(cityA.Value) + int64(cityB.Value)
	}

	return answer
}
