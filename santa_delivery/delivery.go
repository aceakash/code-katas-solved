package santa_delivery

type Position struct {
	X int
	Y int
}

type Direction string

const (
	North   Direction = "North"
	South   Direction = "South"
	East    Direction = "East"
	West    Direction = "West"
	Unknown Direction = ""
)

func CountHousesDeliveredTo(instructions string) int {
	housesDeliveredTo := map[Position]int{}

	directions := parseInstructions(instructions)

	followDirections(directions, housesDeliveredTo)

	return len(housesDeliveredTo)
}

func followDirections(directions []Direction, housesDeliveredTo map[Position]int) {
	santaLocation := Position{X: 0, Y: 0}
	housesDeliveredTo[santaLocation] = 1

	for _, d := range directions {
		switch d {
		case North:
			santaLocation.Y += 1
			break
		case South:
			santaLocation.Y -= 1
			break
		case East:
			santaLocation.X += 1
			break
		case West:
			santaLocation.X -= 1
			break
		default:
			panic("should never happen; maybe change design for better error handling")
		}
		housesDeliveredTo[santaLocation] += 1
	}
}

func parseInstructions(rawInstructions string) []Direction {
	dirs := make([]Direction, len(rawInstructions))
	for i, r := range rawInstructions {
		switch r {
		case '^':
			dirs[i] = North
			break
		case 'v':
			dirs[i] = South
			break

		case '>':
			dirs[i] = East
			break

		case '<':
			dirs[i] = West
			break

		default:
			panic("should never happen; maybe change design for better error handling")
		}
	}
	return dirs
}
