package layout

import (
	"bytes"
	"strings"
)

type RowMap []int
type SectionMap []RowMap
type SeatMap []SectionMap

func (s SeatMap) String() string {
	var b bytes.Buffer
	for _, sectionMap := range s {
		for _, rowMap := range sectionMap {
			for _, seat := range rowMap {
				if seat == 1 {
					b.WriteString("1")
				} else {
					b.WriteString("0")
				}
			}
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (s SeatMap) FromString(str string) SeatMap {
	lines := strings.Split(str, "\n")

	var sectionMap SectionMap
	for _, line := range lines {
		if line == "" {
			s = append(s, sectionMap)
			sectionMap = SectionMap{}
			continue
		}
		rowMap := make(RowMap, len(line))
		for i, c := range line {
			if c == '1' {
				rowMap[i] = 1
			}
		}
		sectionMap = append(sectionMap, rowMap)
	}

	return s
}

func (s SeatMap) FromHall(h Hall) SeatMap {
	s = make(SeatMap, len(h.Sections()))
	for i, section := range h.Sections() {
		sectionMap := make(SectionMap, len(section.Rows()))
		for _, row := range section.Rows() {
			rowMap := make(RowMap, len(row.Seats()))
			for _, seat := range row.Seats() {
				if !seat.Available() {
					rowMap[seat.Order()] = 1
				}
			}
			sectionMap[row.Order()] = rowMap
		}
		s[i] = sectionMap
	}

	return s
}
