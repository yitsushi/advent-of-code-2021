package day05

import "fmt"

type skipFilterFunc func(Range) bool

func (d *Solver) fill(skipFn skipFilterFunc) map[string]int {
	fields := map[string]int{}

	for _, data := range d.input {
		if skipFn(data) {
			continue
		}

		x, y := data.From.X, data.From.Y

		for {
			key := fmt.Sprintf("%.0f,%.0f", x, y)
			fields[key]++

			// Post check because both ends are included.
			if x == data.To.X && y == data.To.Y {
				break
			}

			if data.From.X < data.To.X {
				x++
			} else if data.From.X > data.To.X {
				x--
			}

			if data.From.Y < data.To.Y {
				y++
			} else if data.From.Y > data.To.Y {
				y--
			}

		}
	}

	return fields
}
