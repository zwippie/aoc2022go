package day15

import (
	"aoc2022/utils"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
)

type Pos struct {
	x, y int
}
type Sensor struct {
	pos    Pos
	beacon Pos
	dist   int
}
type SensorGrid []Sensor
type Segment struct {
	x1, y1, x2, y2, slope, yint int
}

// 5525847
func PartA(input []byte) any {
	grid := parseInput(input)
	onRow := onRowY(grid, 2000000)
	return len(onRow)
}

func onRowY(grid SensorGrid, row int) map[Pos]bool {
	result := make(map[Pos]bool)
	for _, sensor := range grid {
		if sensor.pos.Dist(Pos{sensor.pos.x, row}) > sensor.dist {
			continue
		}
		for x := sensor.pos.x - sensor.dist; x <= sensor.pos.x+sensor.dist; x++ {
			p := Pos{x, row}
			if sensor.pos.Dist(p) <= sensor.dist && !sensor.pos.Eq(p) && !sensor.beacon.Eq(p) {
				result[p] = true
			}
		}
	}
	return result
}

// 13340867187704
func PartB(input []byte) any {
	grid := parseInput(input)
	ups, downs := sensorRangeBorders(grid, 1)
	intersections := getIntersections(ups, downs, 4000000)
	fmt.Printf("intersections: %v\n", len(intersections))

	if point, ok := findUncoveredPoint(grid, intersections); ok {
		return point.x*4000000 + point.y
	}
	return ("Point not found")
}

// Missing beacon should be just outside of scanners range,
// so assume all other points in the grid are already covered.
// Create the diagonal line segments that surround the sensor diamond range; 2 / and 2 \
// Every point on the line can be intersected by a line from another scanner,
// that intersection is a possible uncovered spot, so let's check if the spot
// is covered by any scanner on the grid. If it's not, bingo.

func sensorRangeBorders(grid SensorGrid, extra int) ([]Segment, []Segment) {
	ups, downs := []Segment{}, []Segment{} // segments going / and \
	for _, sensor := range grid {
		seg := CreateSegment(sensor.pos.x-sensor.dist-extra, sensor.pos.y, sensor.pos.x, sensor.pos.y-sensor.dist-extra, 1)
		ups = append(ups, seg) // left to top

		seg = CreateSegment(sensor.pos.x, sensor.pos.y+sensor.dist+extra, sensor.pos.x+sensor.dist+extra, sensor.pos.y, 1)
		ups = append(ups, seg) // bottom to right

		seg = CreateSegment(sensor.pos.x-sensor.dist-extra, sensor.pos.y, sensor.pos.x, sensor.pos.y+sensor.dist+extra, -1)
		downs = append(downs, seg) // left to bottom

		seg = CreateSegment(sensor.pos.x, sensor.pos.y-sensor.dist-extra, sensor.pos.x+sensor.dist+extra, sensor.pos.y, -1)
		downs = append(downs, seg) // top to right
	}
	return ups, downs
}

func getIntersections(ups, downs []Segment, maxXY int) []Pos {
	result := make(map[Pos]bool)
	for _, up := range ups {
		for _, down := range downs {
			if pos, ok := Intersection(up, down); ok {
				if pos.x >= 0 && pos.x <= maxXY && pos.y >= 0 && pos.y <= maxXY {
					result[pos] = true
				}
			}
		}
	}
	return maps.Keys(result)
}

func findUncoveredPoint(grid SensorGrid, points []Pos) (Pos, bool) {
	for _, pos := range points {
		covered := false
		for _, scanner := range grid {
			if scanner.Covers(pos) {
				covered = true
			}
		}
		if !covered {
			return pos, true
		}
	}
	return Pos{}, false
}

func CreateSegment(x1, y1, x2, y2, slope int) Segment {
	yint := y1 - slope*x1 // x at y = 0
	return Segment{x1, y1, x2, y2, slope, yint}
}

func Intersection(s1, s2 Segment) (Pos, bool) {
	if s1.slope == s2.slope {
		return Pos{}, false
	}

	x := (s2.yint - s1.yint) / (s1.slope - s2.slope)
	if (s2.yint-s1.yint)%2 != 0 {
		return Pos{}, false // only integer points
	}
	y := s1.slope*x + s1.yint
	return Pos{x, y}, true
}

func (s Sensor) Covers(p Pos) bool {
	return s.pos.Dist(p) <= s.dist
}

func (a Pos) Dist(b Pos) int {
	return utils.Abs(a.x-b.x) + utils.Abs(a.y-b.y)
}

func (a Pos) Eq(b Pos) bool {
	return a.x == b.x && a.y == b.y
}

func parseInput(input []byte) SensorGrid {
	result := SensorGrid{}
	regex := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	for _, line := range strings.Split(string(input), "\n") {
		matches := regex.FindStringSubmatch(line)[1:]
		s := Pos{utils.ParseInt(matches[0]), utils.ParseInt(matches[1])}
		b := Pos{utils.ParseInt(matches[2]), utils.ParseInt(matches[3])}
		result = append(result, Sensor{pos: s, beacon: b, dist: s.Dist(b)})
	}
	return result
}
