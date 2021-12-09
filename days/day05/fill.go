package day05

import "fmt"

type skipFilterFunc func(Range) bool

func (d *Solver) fill(skipFn skipFilterFunc) map[string]int {
	fields := map[string]int{}

	for _, data := range d.input {
		if skipFn(data) {
			continue
		}

		xValue, yValue := data.From.X, data.From.Y

		for {
			key := fmt.Sprintf("%.0f,%.0f", xValue, yValue)
			fields[key]++

			// Post check because both ends are included.
			if xValue == data.To.X && yValue == data.To.Y {
				break
			}

			if data.From.X < data.To.X {
				xValue++
			} else if data.From.X > data.To.X {
				xValue--
			}

			if data.From.Y < data.To.Y {
				yValue++
			} else if data.From.Y > data.To.Y {
				yValue--
			}
		}
	}

	return fields
}
