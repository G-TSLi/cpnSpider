package scheduler


// 注册资源队列
func AddMatrix(spiderName string) *Matrix {
	matrix := newMatrix(spiderName)
	return matrix
}