package models

type Macro struct {
	TypeTable     string  `json:"type_name"`
	IdTable       int     `json:"economy_type"`
	Code          string  `json:"economy_code"`
	Name          string  `json:"economy_name"`
	Last          float64 `json:"economy_last"`
	Previous      float64 `json:"economy_previous"`
	Change        float64 `json:"change"`
	ChangeP       float64 `json:"economy_change"`
	ThreeMonthAVG float64 `json:"economy_3month_avg"`
	SixMonthAVG   float64 `json:"economy_6month_avg"`
	OneYearAVG    float64 `json:"economy_1years_avg"`
	ThreeYearAVG  float64 `json:"economy_3years_avg"`
	LastUpdate    string  `json:"economy_lastupdate"`
	Frequency     string  `json:"economy_frequency"`
}

type MacroResponse struct {
	Total int     `json:"total"`
	Data  []Macro `json:"data"`
}
