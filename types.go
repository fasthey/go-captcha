/*
 * @Author: coller
 * @Date: 2021-11-22 13:54:25
 * @LastEditTime: 2023-02-08 12:24:09
 * @Desc:
 */
package captcha

// Point 随机生成的抠图位置
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// CutoutRet 抠图出来的结果
type CutoutRet struct {
	Point        *Point
	BackgroudImg string
	BlockImg     string
}
