package firstdep

import (
	"fmt"
	"github.com/bmizerany/perks/quantile"
)


func ccc() {
	ch := make(chan float64)

	// Compute the 50th, 90th, and 99th percentile.
	q := quantile.NewTargeted(0.50, 0.90, 0.99)
	for v := range ch {
		q.Insert(v)
	}

	fmt.Println("perc50:", q.Query(0.50))
	fmt.Println("perc90:", q.Query(0.90))
	fmt.Println("perc99:", q.Query(0.99))
	fmt.Println("count:", q.Count())
}