package output

type OutcomeStatus int

const (
	OK   OutcomeStatus = iota + 1
	WARN
	FAIL
)

func OutcomeStatusFrom(str string) OutcomeStatus {
	switch str {
	case "OK":
		return OK
	case "WARN":
		return WARN
	case "FAIL":
		return FAIL
	default:
		panic("Unknown Status " + str)
	}
}

func (os OutcomeStatus) Prioritise(newOs OutcomeStatus) OutcomeStatus {
	if os > newOs {
		return os
	} else {
		return newOs
	}
}

func (os OutcomeStatus) String() string {
	switch os {
	case OK:
		return "OK"
	case WARN:
		return "WARN"
	case FAIL:
		return "FAIL"
	default:
		panic("Unknown Status " + os.String())
	}
}
