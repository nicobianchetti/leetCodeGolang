package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for _, candie := range candies {
		if candie > max {
			max = candie
		}
	}

	var Bool []bool

	for _, candie := range candies {
		if (candie + extraCandies) >= max {
			Bool = append(Bool, true)
		} else {
			Bool = append(Bool, false)
		}
	}

	return Bool
}
