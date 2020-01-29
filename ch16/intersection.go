package ch16

type point struct {
	x, y float32
}

type line struct {
	start, end point
	slope      float32
	yintercept float32
}

func newLine(start, end point) line {

	// equation of a line => y = mx + b

	slope := (end.y - start.y) / (end.x - start.x) // m = (y2-y1) / (x2-x1)
	yintercept := start.y - slope*start.x          // b= y-mx

	return line{start: start, end: end, yintercept: yintercept, slope: slope}
}

func intersection(start1, end1, start2, end2 point) *point {

	//rearrange the points in an ascending manner for ease of use
	if start1.x > end1.x {
		start1, end1 = end1, start1
	}

	if start2.x > end2.x {
		start2, end2 = end2, start2
	}

	if start1.x > start2.x {
		start1, end1, start2, end2 = start2, end2, start1, end1
	}
	//rearrange ends here

	line1 := newLine(start1, end1)
	line2 := newLine(start2, end2)

	if line1.slope == line2.slope {
		if line1.yintercept == line2.yintercept && isBetween(start1, start2, end1) {
			return &start2
		}
		return nil
	}

	xIntersect := (line2.yintercept - line1.yintercept) / (line1.slope - line2.slope)
	yIntersect := (line1.slope * xIntersect) + line1.yintercept

	intersectPoint := point{x: xIntersect, y: yIntersect}

	if isBetween(start1, intersectPoint, end1) {
		return &intersectPoint
	}

	return nil
}

func isBetween(start, mid, end point) bool {

	return isBetweenAxis(start.x, mid.x, end.x) && isBetweenAxis(start.y, mid.y, end.y)

}

func isBetweenAxis(start, mid, end float32) bool {

	if start > end {
		return mid <= start && end <= mid
	} else {
		return start <= mid && mid <= end
	}
}
