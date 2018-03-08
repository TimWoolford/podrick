package status

const upColour = "#4c1"
const downColour = "red"
const warnColour = "#ffc200"
const unknownColour = "#adadba"

type State int

const (
	Up   State = iota + 1
	Fail
	Warn
	Down
)

func (state State) Colour() string {
	switch state {
	case Up:
		return upColour
	case Fail:
		return downColour
	case Warn:
		return warnColour
	default:
		return unknownColour
	}
}
