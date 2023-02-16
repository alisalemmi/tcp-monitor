package dto

type RetransmissionRatio struct {
	Sent         int     `json:"sent"`
	Retransmited int     `json:"retransmited"`
	Ratio        float64 `json:"ratio"`
}
