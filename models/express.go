package models

type Delivery struct {
	Message   string    `json:"message"`
	Nu        string    `json:"nu"`
	IsCheck   string    `json:"ischeck"`
	Condition string    `json:"condition"`
	Com       string    `json:"com"`
	Status    string    `json:"status"`
	State     string    `json:"state"`
	Data      []Express `json:"data"`
}

type Express struct {
	Time     string `json:"time"`
	FTime    string `json:"ftime"`
	Context  string `json:"context"`
	Location string `json:"location"`
}
