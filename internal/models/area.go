package models

type AreaCheck struct {
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Radius   int    `json:"radius"`
	Result   bool   `json:"result"`
	SendTime string `json:"send_time"`
	Author   string `json:"author"`
}
