package main

/*
	Example 1:
	Input: x = 123
	Output: 321
	How it works:
	1. Initialize a variable y to 1 (to handle negative numbers) and z to 0 (to store the reversed number).
	2. If x is negative, set y to -1 and make x positive.
	3. While x is not 0:
	   a. Update z by multiplying it by 10 and adding the last digit of x (x % 10).
	   b. Remove the last digit from x by performing integer division by 10.
	4. After the loop, check if z is within the 32-bit signed integer range. If not, return 0.
	5. Return the reversed number multiplied by y to restore the original sign.
*/

func reverse(x int) int {
	y := 1
	z := 0
	if x < 0 {
		y = -1
		x = -x
	}
	for x != 0 {
		z = x*10 + x%10
		x /= 10
	}

	if z > 1<<31-1 || z < -1<<31 {
		return 0
	}

	return y * z
}
