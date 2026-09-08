func searchMatrix(matrix [][]int, target int) bool {
	a, b, c, d, n := 0, 0, len(matrix)-1, len(matrix[0])-1, len(matrix[0])
	p1, p2 := plane2pt(n, a, b), plane2pt(n, c, d)
	for p1 <= p2 {
		m := (p1 + p2)/2
		m1, m2 := pt2plane(n, m)
		if matrix[m1][m2] == target {
			return true
		}
		if matrix[m1][m2] > target {
			c, d = pt2plane(n, m - 1)
		} else {
			a, b = pt2plane(n, m + 1)
		}
		p1, p2 = plane2pt(n, a, b), plane2pt(n, c, d)
	} 

	return false
}

func plane2pt(m, x, y int) int {
	return m*x + y
}

func pt2plane(m, p int) (x, y int) {
	x = p/m
	y = p%m
	return
}