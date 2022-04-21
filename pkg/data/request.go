package data

type Request struct {
	Action string
	Type   string
	Data   string `json:",omitempty"`
}
