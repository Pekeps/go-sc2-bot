// Code generated by gen_ids. DO NOT EDIT.
package effect

import "github.com/pekeps/go-sc2ai/api"

func String(e api.EffectID) string {
	return strings[uint32(e)]
}

var strings = map[uint32]string{
	0:  "Invalid",
	1:  "PsiStorm",
	2:  "GuardianShield",
	3:  "TemporalFieldGrowing",
	4:  "TemporalField",
	5:  "ThermalLance",
	6:  "ScannerSweep",
	7:  "NukeDot",
	8:  "LiberatorDefenderZoneSetup",
	9:  "LiberatorDefenderZone",
	10: "BlindingCloud",
	11: "CorrosiveBile",
	12: "LurkerSpines",
}
