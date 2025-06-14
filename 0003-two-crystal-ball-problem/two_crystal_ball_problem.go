package two_crystal_ball_problem

import "math"

func TwoCrystalBalls(breaks []bool) int {
	if len(breaks) == 0 {
		return -1
	}

	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jumpAmount

	for i < len(breaks) {
		if breaks[i] {
			break
		}
		i += jumpAmount
	}

	i -= jumpAmount

	for j := 0; j <= jumpAmount && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}

	return -1
}

func Run() any {
	breaks := []bool{false, false, false, false, false, false, false, true, true, true}
	result := TwoCrystalBalls(breaks)

	return map[string]any{
		"breaks_array":   breaks,
		"breaking_floor": result,
		"description":    "First floor where crystal ball breaks",
	}
}
