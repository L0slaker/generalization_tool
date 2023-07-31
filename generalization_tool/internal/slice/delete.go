package slice

import "Prove/generalization_tool/internal/errs"

// Delete 区别于切片中的Delete方法，多返回了删除的参数
func Delete[T any](src []T, index int) ([]T, T, error) {
	if index < 0 || index >= len(src) {
		var t T
		return nil, t, errs.NewErrIndexOutOfRange(len(src), index)
	}

	j := 0
	res := src[index]
	for i, v := range src {
		if i != index {
			src[j] = v
			j++
		}
	}
	src = src[:j]

	return src, res, nil
}
