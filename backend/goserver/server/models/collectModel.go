package models

type ActionCollected struct {
	CountUse float64
	Name     string
}
type ToolsCollected struct {
	CountUse float64
	Name     string
}

type ToolsCollectedData struct {
	ToolsUse      []ToolsCollected
	ToolsActivate []ToolsCollected
}

type CollectedInstallData struct {
	Platform struct {
		Os struct {
			Windows float64
			Linux   float64
			Mac     float64
			Other   float64
		}
		Version struct {
			Windows struct {
				V7    float64
				V8    float64
				V81   float64
				V10   float64
				Other float64
			}
			Linux struct {
				Ubuntu1404 float64
				Ubuntu1410 float64
				Ubuntu1504 float64
				Ubuntu1510 float64
				Ubuntu1604 float64
				Ubuntu1610 float64
				Ubuntu1704 float64
				Other      float64
			}
			Mac struct {
				V1012 float64
				Other float64
			}
		}
	}
	CPU struct {
		Architecture struct {
			X86_64 float64
			X86    float64 //на самом деле  x86_64 c маленькой буквы
			Other  float64
		}
		Cores struct {
			C1    float64
			C2    float64
			C3    float64
			C4    float64
			C6    float64
			C8    float64
			Other float64
		}
	}
	Compiler struct {
		Type struct {
			GCC   float64
			Clang float64
			MSVC  float64
			Other float64
		}
		// 	Version struct { //подумать что можно сделать
		// 		V5dot4 float64
		// 		Other  float64
		// 	}
	}
	Locale struct {
		Language struct {
			English float64
			Russian float64
			Other   float64
		}
	}
}

//Less or equal then mb*
type imageDistribution struct {
	Mb1     float64
	Mb5     float64
	Mb10    float64
	Mb25    float64
	Mb50    float64
	Mb100   float64
	Mb200   float64
	Mb400   float64
	Mb800   float64
	More800 float64
}
type colorProfileDistribution struct {
	RGBA      float64
	CMYK      float64
	Grayscale float64
	XYZ       float64
	YCbCr     float64
	Lab       float64
}

//Less then L*
type heightDistribution struct {
	L500  float64
	L1000 float64
	L2000 float64
	L4000 float64
	L8000 float64
	M8000 float64
}
type widthDistribution struct {
	L500  float64
	L1000 float64
	L2000 float64
	L4000 float64
	L8000 float64
	M8000 float64
}
type layersDistribution struct {
	L1  float64
	L2  float64
	L4  float64
	L8  float64
	L16 float64
	L32 float64
	L64 float64
	M64 float64
}
type ImageCollected struct {
	ID  imageDistribution
	CPD colorProfileDistribution
	HD  heightDistribution
	WD  widthDistribution
	LD  layersDistribution
	//TODO COLORSPACE
}
