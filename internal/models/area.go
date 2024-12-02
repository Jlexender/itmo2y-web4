package models

type AreaCheck struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Radius   float64     `json:"radius"`
	Result   bool    `json:"result"`
	SendTime string  `json:"send_time"`
	Author   string  `json:"author"`
}
