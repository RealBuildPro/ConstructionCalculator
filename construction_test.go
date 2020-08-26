package construction_test

import (
	"testing"

	construction "bitbucket.org/swagonomics/constructionCalculator"
)

func checkRes(drResult *construction.DrywallResult,
	sheets, screws, tape, lbReadyMix, glReadyMixLightWeight, lbLightWeight, lbQuickSet, totalArea int, t *testing.T) bool {
	if drResult.Sheets != sheets {
		t.Errorf("%d sheets instead of %d", drResult.Sheets, sheets)
		return false
	}
	if int(drResult.PoundsScrews) != screws {
		t.Errorf("%2f pounds of screws instead of %d", drResult.PoundsScrews, screws)
		return false
	}
	if int(drResult.TapeRolls) != tape {
		t.Errorf("%2f tape rolls instead of %d", drResult.TapeRolls, tape)
		return false
	}
	if int(drResult.PoundsReadyMix) != lbReadyMix {
		t.Errorf("%2f PoundsReadyMix instead of %d", drResult.PoundsReadyMix, lbReadyMix)
		return false
	}
	if int(drResult.GallonsReadyMixLightweight) != glReadyMixLightWeight {
		t.Errorf("%2f GallonsReadyMixLightweight instead of %d", drResult.GallonsReadyMixLightweight, glReadyMixLightWeight)
		return false
	}
	if int(drResult.PoundsLightWeight) != lbLightWeight {
		t.Errorf("%2f PoundsLightWeight instead of %d", drResult.PoundsLightWeight, lbLightWeight)
		return false
	}
	if int(drResult.PoundsQuickSet) != lbQuickSet {
		t.Errorf("%2f PoundsQuickSet instead of %d", drResult.PoundsQuickSet, lbQuickSet)
		return false
	}
	if int(drResult.TotalArea) != totalArea {
		t.Errorf("%2f TotalArea instead of %d", drResult.TotalArea, totalArea)
		return false
	}
	return true
}
func TestDryWall(t *testing.T) {
	myCalc := construction.NewCalculator()
	noCeiling, err := myCalc.Drywall(&construction.DrywallArgs{
		Height: 5,
		Width:  5,
		Length: 5,
		PanelSize: construction.Size{
			Length: 4,
			Width:  8,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	if !checkRes(noCeiling, 4, 1, 1, 15, 1, 6, 8, 100, t) {
		return
	}
	ceiling, err := myCalc.Drywall(&construction.DrywallArgs{
		Height: 5,
		PanelSize: construction.Size{
			Length: 4,
			Width:  8,
		},
		IncludeCeiling: true,
		Length:         5,
		Width:          5,
	})
	if err != nil {
		t.Error(err)
		return
	}
	if !checkRes(ceiling, 4, 1, 1, 18, 2, 7, 10, 125, t) {
		return
	}
	ceilingWithHoles, err := myCalc.Drywall(&construction.DrywallArgs{
		Height: 5,
		PanelSize: construction.Size{
			Length: 4,
			Width:  8,
		},
		Holes: []construction.Size{
			{
				Length: 2,
				Width:  2,
			},
		},
		IncludeCeiling: true,
		Length:         5,
		Width:          5,
	})

	if err != nil {
		t.Error(err)
		return
	}
	if !checkRes(ceilingWithHoles, 4, 1, 1, 18, 2, 7, 10, 121, t) {
		return
	}
}
