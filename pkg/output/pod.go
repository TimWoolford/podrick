package output

type StatusContent map[string]interface{}

type Pod struct {
	name       string
	Error      error         `json:"error"`
	StatusPage StatusContent `json:"statusPage"`
}

func NewPod(name string, content StatusContent, err error) *Pod {
	return &Pod{
		name:       name,
		StatusPage: content,
		Error:      err,
	}
}

func FailedPod(name string, err error) *Pod {
	return &Pod{
		name:  name,
		Error: err,
	}
}
func (p Pod) Status() OutcomeStatus {
	return OutcomeStatusFrom(p.StatusPage["overallStatus"].(string))
}
