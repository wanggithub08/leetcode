package main

/*
 * @Description: 69.Sqrt(x)
 * @Date: 2019-09-22 23:38:59
 * @LastEditors: wangtongli
 * @Email: wangli3508@gmail.com
 * @LastEditTime: 2019-09-22 23:41:55
 */
func mySqrt1(x int) int {
	max, min := x, 0
	if x == 1 {
		return 1
	}
	for max-min > 1 {
		mid := (min + max) / 2
		if mid*mid > x {
			max = mid
		} else {
			min = mid
		}
	}
	return min
}
