package main

func numJewelsInStones(J string, S string) int {
	Jewells := 0
	for _, valueJ := range J {
		for _, valueS := range S {
			if valueS == valueJ {
				Jewells++
			}
		}
	}
	return Jewells
}
