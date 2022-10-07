package models

type Status struct {
	Water      int
	WaterLevel string
	Wind       int
	WindLevel  string
}

type SensorData struct {
	Water int `json:"Water"`
	Wind  int `json:"Wind"`
}
