package main

/*
 * @Description: 69.Sqrt(x)
 * @Date: 2019-09-22 16:08:09
 * @LastEditors: wangtongli
 * @Email: wangli3508@gmail.com
 * @LastEditTime: 2019-09-22 23:42:13
 */
func mySqrt(x int) int {
	var i, value int
	//初始化，执行一次；条件，必须；其他
	for value < x {
		i++
		value = i * i
	}
	if value > x {
		return i - 1
	}
	return i
}
