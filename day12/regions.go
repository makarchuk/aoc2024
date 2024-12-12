package day12

import (
	"io"
	"slices"

	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/set"
)

type RegionsFinder struct {
	farm       field.Field
	attributed set.Set[field.Point]
}

func NewRegionsFinder(in io.Reader) (*RegionsFinder, error) {
	farm, err := field.New(in)
	if err != nil {
		return nil, err
	}
	return &RegionsFinder{
		farm:       *farm,
		attributed: set.New[field.Point](),
	}, nil
}

var neightbourVectors = []field.Point{
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 0, Y: -1},
}

type Price struct {
	Price     int
	BulkPrice int
}

type fenceSpecs struct {
	perimeter int
	area      int
	sides     int
}

func (r *RegionsFinder) PlanFences() Price {
	specs := []fenceSpecs{}

	newRegion := true
	var currentRegionCrop byte
	var currentRegionSurface set.Set[field.Point]
	var currentRegionSpecs fenceSpecs
	var currentRegionSegments map[fenceLine][]int

	for point := range r.farm.Iter() {
		if newRegion {
			if r.attributed.Contains(point) {
				continue
			}
			newRegion = false
			currentRegionSpecs = fenceSpecs{
				perimeter: 0,
				sides:     0,
				area:      1,
			}
			currentRegionSurface = set.New[field.Point]()
			currentRegionSurface.Add(point)
			currentRegionSegments = make(map[fenceLine][]int)
			r.attributed.Add(point)
			currentRegionCrop, _ = r.farm.Get(point)
		}
		for {
			newSurface := set.New[field.Point]()
			for _, point := range currentRegionSurface.List() {
				for _, neighbour := range neightbourVectors {
					possibleNeighbour := point.Add(neighbour)

					//external boundary should contribute to our perimeter calculation as well
					val, _ := r.farm.Get(possibleNeighbour)

					switch val == currentRegionCrop {
					case true:
						if newSurface.Contains(possibleNeighbour) {
							continue
						}
						if r.attributed.Contains(possibleNeighbour) {
							continue
						}
						newSurface.Add(possibleNeighbour)
						r.attributed.Add(possibleNeighbour)
						currentRegionSpecs.area++
					case false:
						side := sideIdentity(point, neighbour)
						currentRegionSegments[side.fenceLine] = append(currentRegionSegments[side.fenceLine], side.Coordinate)
						currentRegionSpecs.perimeter++
					}
				}
			}
			if newSurface.Len() == 0 {
				currentRegionSpecs.sides = sides(currentRegionSegments)
				specs = append(specs, currentRegionSpecs)
				newRegion = true
				break
			}
			currentRegionSurface = newSurface
		}
	}
	totalPrice := Price{}

	for _, spec := range specs {
		totalPrice.Price += spec.perimeter * spec.area
		totalPrice.BulkPrice += spec.sides * spec.area
	}
	return totalPrice
}

type fenceLine struct {
	Direction  field.Point
	Coordinate int
}

type segment struct {
	fenceLine  fenceLine
	Coordinate int
}

func sideIdentity(p field.Point, direction field.Point) segment {
	if direction.X == 0 {
		return segment{
			fenceLine: fenceLine{
				Direction:  direction,
				Coordinate: p.Y,
			},
			Coordinate: p.X,
		}
	}
	return segment{
		fenceLine: fenceLine{
			Direction:  direction,
			Coordinate: p.X,
		},
		Coordinate: p.Y,
	}
}

func sides(segments map[fenceLine][]int) int {
	totalSides := 0
	for _, fences := range segments {
		slices.Sort(fences)
		if len(fences) == 0 {
			continue
		}
		prev := fences[0]

		for _, fence := range fences[1:] {
			if fence-prev > 1 {
				totalSides++
			}
			prev = fence
		}
		//last segment is always unaccounted for
		totalSides++
	}
	return totalSides
}
