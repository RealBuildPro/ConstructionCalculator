package construction

import (
	"math"
)

type Calculator interface {
	Drywall(drywallArgs *DrywallArgs) (*DrywallResult, error)
}

type calculator struct {
}

func NewCalculator() Calculator {
	return &calculator{}
}

type Size struct {
	Length float64
	Width  float64
}

type DrywallArgs struct {
	Length         float64
	Width          float64
	Height         float64
	IncludeCeiling bool
	PanelSize      Size
	Holes          []Size
}

type DrywallResult struct {
	Sheets                     int
	PoundsScrews               float64
	TapeRolls                  float64
	PoundsReadyMix             float64
	PoundsQuickSet             float64
	PoundsLightWeight          float64
	GallonsReadyMixLightweight float64
	TotalArea                  float64
}

func (c *calculator) Drywall(drywallArgs *DrywallArgs) (*DrywallResult, error) {
	lengthArea := drywallArgs.Length * drywallArgs.Height
	widthArea := drywallArgs.Width * drywallArgs.Height
	ceiling := 0.00
	if drywallArgs.IncludeCeiling {
		ceiling = drywallArgs.Length * drywallArgs.Width
	}
	sub := 0.00
	for _, hole := range drywallArgs.Holes {
		sub = hole.Length * hole.Width
	}
	drywallResult := &DrywallResult{TotalArea: (((lengthArea + widthArea) * 2) + ceiling) - sub}
	sheet := (drywallResult.TotalArea) / (drywallArgs.PanelSize.Length * drywallArgs.PanelSize.Width)
	sheetTotal := int((drywallResult.TotalArea) / (drywallArgs.PanelSize.Length * drywallArgs.PanelSize.Width))
	if math.Mod(sheet, 10) > 0 {
		sheetTotal = sheetTotal + 1
	}
	drywallResult.Sheets = sheetTotal
	drywallResult.PoundsLightWeight = math.Ceil(drywallResult.TotalArea * .053)
	drywallResult.PoundsQuickSet = math.Ceil(drywallResult.TotalArea * .075)
	drywallResult.TapeRolls = math.Ceil(drywallResult.TotalArea / 200)
	drywallResult.PoundsReadyMix = math.Ceil(drywallResult.TotalArea * .141)
	drywallResult.GallonsReadyMixLightweight = math.Ceil(drywallResult.TotalArea * .01)
	drywallResult.PoundsScrews = math.Ceil(drywallResult.TotalArea * .0055)
	return drywallResult, nil
}
