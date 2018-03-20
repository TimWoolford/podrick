package output

type StatusContent map[string]interface{}

type Pod struct {
	name       string
	Error      string        `json:"error"`
	StatusPage StatusContent `json:"statusPage"`
}

func NewPod(name string, content StatusContent, err error) *Pod {
	var errorMessage string
	if err != nil { errorMessage = err.Error()}
	return &Pod{
		name:       name,
		StatusPage: content,
		Error:      errorMessage,
	}
}

func FailedPod(name string, err error) *Pod {
	return &Pod{
		name:  name,
		Error: err.Error(),
	}
}

func (p Pod) Status() OutcomeStatus {
	status, present := p.StatusPage["overallStatus"]
	if present {
		return OutcomeStatusFrom(status.(string))
	}
	return FAIL
}
