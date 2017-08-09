package models

type Image struct {
	Images []ImageInternal `json:"Images"`
}
type ImageInternal struct {
	ColorProfile string  `json:"colorProfile"`
	ColorSpace   string  `json:"colorSpace"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Size         float64 `json:"size"`
	NumLayers    float64 `json:"numLayers"`
}

