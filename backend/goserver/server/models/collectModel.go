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

type CountAndProportion struct {
	Count      float64
	Proportion string
}
type CollectedInstallData struct {
	Platform struct {
		Os struct {
			Windows CountAndProportion
			Linux   CountAndProportion
			Mac     CountAndProportion
			Other   CountAndProportion
			Unknown CountAndProportion
		}
		Version struct {
			Windows struct {
				V7      CountAndProportion
				V8      CountAndProportion
				V81     CountAndProportion
				V10     CountAndProportion
				Other   CountAndProportion
			}
			Linux struct {
				Ubuntu1404 CountAndProportion

				Ubuntu1410 CountAndProportion
				Ubuntu1504 CountAndProportion
				Ubuntu1510 CountAndProportion
				Ubuntu1604 CountAndProportion
				Ubuntu1610 CountAndProportion
				Ubuntu1704 CountAndProportion
				Other      CountAndProportion
			}
			Mac struct {
				V1012   CountAndProportion
				Other   CountAndProportion
			}
		}
	}
	CPU struct {
		Architecture struct {
			X86_64  CountAndProportion
			X86     CountAndProportion //на самом деле  x86_64 c маленькой буквы
			Other   CountAndProportion
			Unknown CountAndProportion
		}
		Cores struct {
			C1      CountAndProportion
			C2      CountAndProportion
			C3      CountAndProportion
			C4      CountAndProportion
			C6      CountAndProportion
			C8      CountAndProportion
			Other   CountAndProportion
			Unknown CountAndProportion
		}
	}
	Compiler struct {
		Type struct {
			GCC     CountAndProportion
			Clang   CountAndProportion
			MSVC    CountAndProportion
			Other   CountAndProportion
			Unknown CountAndProportion
		}
		// 	Version struct { //подумать что можно сделать
		// 		V5dot4 float64
		// 		Other  float64
		// 	}
	}
	Locale struct {
		Language struct {
			English CountAndProportion
			Russian CountAndProportion
			Other   CountAndProportion
			Unknown CountAndProportion
		}
	}
}

//Less or equal then mb*
type imageDistribution struct {
	Mb1     CountAndProportion
	Mb5     CountAndProportion
	Mb10    CountAndProportion
	Mb25    CountAndProportion
	Mb50    CountAndProportion
	Mb100   CountAndProportion
	Mb200   CountAndProportion
	Mb400   CountAndProportion
	Mb800   CountAndProportion
	More800 CountAndProportion
}
type colorProfileDistribution struct {
	RGBA      CountAndProportion
	CMYK      CountAndProportion
	Grayscale CountAndProportion
	XYZ       CountAndProportion
	YCbCr     CountAndProportion
	Lab       CountAndProportion
}

//Less then L*
type heightDistribution struct {
	L500  CountAndProportion
	L1000 CountAndProportion
	L2000 CountAndProportion
	L4000 CountAndProportion
	L8000 CountAndProportion
	M8000 CountAndProportion
}
type widthDistribution struct {
	L500  CountAndProportion
	L1000 CountAndProportion
	L2000 CountAndProportion
	L4000 CountAndProportion
	L8000 CountAndProportion
	M8000 CountAndProportion
}
type layersDistribution struct {
	L1  CountAndProportion
	L2  CountAndProportion
	L4  CountAndProportion
	L8  CountAndProportion
	L16 CountAndProportion
	L32 CountAndProportion
	L64 CountAndProportion
	M64 CountAndProportion
}
type ImageCollected struct {
	ID  imageDistribution
	CPD colorProfileDistribution
	HD  heightDistribution
	WD  widthDistribution
	LD  layersDistribution
	//TODO COLORSPACE
}
