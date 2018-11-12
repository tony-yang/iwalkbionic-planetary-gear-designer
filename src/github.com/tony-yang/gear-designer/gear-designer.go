package main

import (
  "flag"
  "fmt"
  "math"
)

const bore float64 = 5.0
const facewidth float64 = 8.0
var standardMod = []float64{0.8, 1, 1.25, 1.5, 2, 2.5}
var planet2Variation = []float64{-5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0}

type Gearset struct {
  Sun1, Planet1, Ring1 Gear
  Sun2, Planet2, Ring2 Gear
  Gain float64
  NumberOfPlanets int
}

func(g Gearset) String() string {
  return fmt.Sprintf("Gearset:\n\t%v\n\t%v\n\t%v\n\t%v\n\t%v\n\t%v\n\tGain Ratio: %v\n\tNumber of Planets: %v", g.Sun1, g.Planet1, g.Ring1, g.Sun2, g.Planet2, g.Ring2, g.Gain, g.NumberOfPlanets)
}

type Gear struct {
  Name string
  Gearset int
  Bore float64
  Facewidth float64
  Mod float64
  Teeth int
  ActualTeeth float64
  PitchDiameter float64
}

func (g Gear) String() string {
  return fmt.Sprintf("Gear %v%v: {Bore: %v, Facewidth: %v, Mod: %v, Teeth: %v, ActualTeeth: %v, PitchDiameter: %v}", g.Name, g.Gearset, g.Bore, g.Facewidth, g.Mod, g.Teeth, g.ActualTeeth, g.PitchDiameter)
}

func GetRingGear(sun, planet Gear) Gear {
  ring := Gear{"Ring", sun.Gearset, bore, facewidth, sun.Mod, sun.Teeth + planet.Teeth * 2, sun.ActualTeeth + planet.ActualTeeth * 2, sun.PitchDiameter + planet.PitchDiameter * 2}
  return ring
}

func GetGearDiameter(g Gear) float64 {
  return g.Mod * float64(g.Teeth)
}

func TeethErrorTolerance(sun Gear, planet Gear, ring Gear) bool {
  if math.Abs(float64(sun.Teeth) - sun.ActualTeeth) > 0.11 {
    return false
  } else if math.Abs(float64(planet.Teeth) - planet.ActualTeeth) > 0.11 {
    return false
  } else if math.Abs(float64(ring.Teeth) - ring.ActualTeeth) > 0.11 {
    return false
  }
  return true
}

func Round(n float64, precision int) float64 {
  precisionMultiplier := math.Pow10(precision)
  precisionEpsilon := 0.5 + math.Pow10(-precision - 1)
  return float64(int(n * precisionMultiplier + precisionEpsilon)) / precisionMultiplier
}

func main() {
  TeethSun1Start := flag.Int("sun1-start", 17, "The starting range of number of teeth for gear set #1 sun")
  TeethSun1End := flag.Int("sun1-end", 17, "The ending range of number of teeth for gear set #1 sun")
  TeethPlanet1Start := flag.Int("planet1-start", 20, "The starting range of number of teeth for gear set #1 planet")
  TeethPlanet1End := flag.Int("planet1-end", 20, "The ending range of number of teeth for gear set #1 planet")
  MinimumGain := flag.Int("min-gain", 500, "The minimum gear ratio desired")
  MaximumGain := flag.Int("max-gain", 3000, "The maximum gear ratio desired")
  MaximumRingDiameter := flag.Float64("max-ring-diameter", 80, "The mamimum Ring gear pitch diameter desired")

  flag.Parse()

  var PotentialGearset []Gearset

  for _, mod1 := range standardMod {
    for sun1Teeth := *TeethSun1Start; sun1Teeth <= *TeethSun1End; sun1Teeth++ {
      for planet1Teeth := *TeethPlanet1Start; planet1Teeth <= *TeethPlanet1End; planet1Teeth++ {
        sun1 := Gear{"Sun", 1, bore, facewidth, mod1, sun1Teeth, float64(sun1Teeth), 0.0}
        sun1.PitchDiameter = GetGearDiameter(sun1)
        planet1 := Gear{"Planet", 1, bore, facewidth, mod1, planet1Teeth, float64(planet1Teeth), 0.0}
        planet1.PitchDiameter = GetGearDiameter(planet1)
        ring1 := GetRingGear(sun1, planet1)

        if ring1.PitchDiameter > *MaximumRingDiameter {
          continue
        }

        numberOfPlanets := 4

        for _, mod2 := range standardMod {
          for _, diameterVariation := range planet2Variation {
            planet2PitchDiameter := planet1.PitchDiameter + diameterVariation
            planet2 := Gear{"Planet", 2, bore, facewidth, mod2, int(Round(planet2PitchDiameter / mod2, 0)), planet2PitchDiameter / mod2, 0.0}
            planet2.PitchDiameter = GetGearDiameter(planet2)

            sun2PitchDiameter := sun1.PitchDiameter - diameterVariation
            sun2 := Gear{"Sun", 2, bore, facewidth, mod2, int(Round(sun2PitchDiameter / mod2, 0)), sun2PitchDiameter / mod2, 0.0}
            sun2.PitchDiameter = GetGearDiameter(sun2)

            ring2 := GetRingGear(sun2, planet2)

            if planet2.Teeth < *TeethPlanet1Start || sun2.Teeth < *TeethSun1Start {
              continue
            }

            turnSun1Input := (float64(ring1.Teeth) + float64(sun1.Teeth)) * float64(ring2.Teeth)
            turnRing2Output := float64(sun1.Teeth) * (float64(ring2.Teeth) - float64(ring1.Teeth)/float64(planet1.Teeth)*float64(planet2.Teeth))

            outputGain := turnSun1Input / turnRing2Output

            if sun1.Teeth >= 2 * planet1.Teeth {
              numberOfPlanets = 9
            } else if sun1.Teeth < 2 * planet1.Teeth && sun1.Teeth > planet1.Teeth {
              numberOfPlanets = 6
            } else if sun1.Teeth == planet1.Teeth {
              numberOfPlanets = 5
            }

            sun1Ring1ToPlanet := float64(sun1.Teeth + ring1.Teeth) / float64(numberOfPlanets)
            sun2Ring2ToPlanet := float64(sun2.Teeth + ring2.Teeth) / float64(numberOfPlanets)

            if math.Abs(outputGain) > float64(*MinimumGain) && !math.IsInf(outputGain, 0) && math.Abs(outputGain) < float64(*MaximumGain) && sun1Ring1ToPlanet == float64(int(sun1Ring1ToPlanet)) && sun2Ring2ToPlanet == float64(int(sun2Ring2ToPlanet)) && TeethErrorTolerance(sun2, planet2, ring2) {
              gearset := Gearset{
                sun1, planet1, ring1,
                sun2, planet2, ring2,
                outputGain,
                numberOfPlanets,
              }
              PotentialGearset = append(PotentialGearset, gearset)
            }
          }
        }
      }
    }
  }
  fmt.Println("Potential gearset")
  printPotentialGearset(PotentialGearset)
}

func printPotentialGearset(g []Gearset) {
  for _, v := range g {
    fmt.Println(v)
  }
}
