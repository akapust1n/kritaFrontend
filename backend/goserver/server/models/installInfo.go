package models

//Request from krita
type Request struct {
	ApplicationVersion struct {
		Version string `json:"value"`
	} `json:"applicationVersion"`
	Compiler struct {
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"compiler"`
	Locale struct {
		Language string `json:"language"`
	} `json:"locale"`
	Opengl struct {
		GlslVersion string `json:"glslVersion"`
		Renderer    string `json:"renderer"`
		Vendor      string `json:"vendor"`
	} `json:"opengl"`
	Platform struct {
		Os      string `json:"os"`
		Version string `json:"version"`
	} `json:"platform"`
	QtVersion struct {
		Version string `json:"value"`
	} `json:"qtVersion"`
	Screens []struct {
		Dpi    float64 `json:"dpi"`
		Height float64 `json:"height"`
		Width  float64 `json:"width"`
	} `json:"screens"`
}
