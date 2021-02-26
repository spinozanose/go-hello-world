package control

import "fmt"

// SimpleLoops is . . .
func SimpleLoops() {

	fmt.Println("Simple Loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("With continue:")
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue // There is also a continue keyword
		}
		fmt.Println(i)
	}

	// for what this is worth:
	fmt.Println("Multiple counters:")
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	fmt.Println("The 'while' form:")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("The infinite loop form:")
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	fmt.Println("Nested Loops:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}

	fmt.Println("Breaking out to a label:")
	OuterLoop:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(i, j, i*j)
			if i * j > 3 {
				break OuterLoop
			}
		}
	}

}

// Collections is . . .
func Collections() {
	fmt.Println("Collection Form:")
	c := []int{2, 4, 6, 7} 
	for k, v := range c {
		fmt.Println(k,v)
	}

	statePopulations := map[string]int {
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}

	fmt.Println("Keys and values:")
	for k, v := range statePopulations {
		fmt.Println(k,v)
	}

	fmt.Println("Just Keys:")
	for k := range statePopulations {
		fmt.Println(k)
	}

	fmt.Println("Just Values:")
	for _, v := range statePopulations {
		fmt.Println(v)
	}
}