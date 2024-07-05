package result

import "encoding/json"

type Result struct {
	Status int
	Region string
	Info   string
	Err    error
}

func (r *Result) ToString() string {
	body, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(body)
}
