package main

type choice struct {
	items map[int]int
	price int
}

func shoppingOffers(prices []int, specials [][]int, needs []int) int {
	choices := make([]choice, 0, len(prices)+len(specials))
	for i, price := range prices {
		choices = append(choices, choice{map[int]int{i: 1}, price})
	}
	for _, special := range specials {
		price := special[len(special)-1]
		m := make(map[int]int, len(needs))
		for i, v := range special[:len(special)-1] {
			m[i] = v
		}
		choices = append(choices, choice{m, price})
	}
	remain := 0
	for _, n := range needs {
		remain += n
	}
	min := -1
	uses := make([]int, len(needs))
	shoppingOffersSub(choices, 0, needs, uses, remain, 0, &min)
	return min
}

func shoppingOffersSub(choices []choice, k int, needs []int, uses []int, remain int, price int, min *int) {
	if remain == 0 && (*min == -1 || *min > price) {
		*min = price
		return
	}
	if remain <= 0 || (*min != -1 && price > *min) {
		return
	}
	for i := k; i < len(choices); i++ {
		choice := choices[i]
		if isBuyable(choice, needs, uses) {
			for k, p := range choice.items {
				needs[k] -= p
				uses[k] += p
				remain -= p
			}
			shoppingOffersSub(choices, i, needs, uses, remain, price+choice.price, min)
			for k, p := range choice.items {
			t 	needs[k] += p
				uses[k] -= p
				remain += p
			}
		}

	}
}

func isBuyable(c choice, needs []int, uses []int) bool {
	for k, v := range c.items {
		if needs[k]-v < 0 {
			return false
		}
		if uses[k]+v > 6 {
			return false
		}
	}
	return true
}
