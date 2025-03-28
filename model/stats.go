package model

type FileStats struct {
	FileName string  `json:"filename"`
	Average  float64 `json:"average"`
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
}
