package status

const upColour = "#4c1"
const downColour = "red"
const defaultColour = "#ffc200"

type State int

const (
	Up   State = iota + 1
	Down
	Warn
)

func (state State) colour() string {
	switch state {
	case Up:
		return upColour
	case Down:
		return downColour
	default:
		return defaultColour
	}
}
