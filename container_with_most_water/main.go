package main

import "fmt"

/*
	Explanation:
	   1. l is the first index
	   2. r is the last index
	   3. m is the maximum area
	   4. width = r - l. Example: if l = 1 and r = 4, width = 4 - 1 = 3
	   5. height is the minimum of height[l] and height[r]. Example: if height[l] = 1 and height[r] = 4, height = 1
	   6. area = width * height
	   7. If height[l] < height[r], move l to the right (l++).
	   8. Else, move r to the left (r--).
	   9. Repeat steps 4-8 until l >= r.
	   10. Return m.
*/

func main() {
	height := []int{1, 2}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0

	for l < r {
		h := min(height[r], height[l])

		area := h * (r - l)
		if area > maxArea {
			maxArea = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
