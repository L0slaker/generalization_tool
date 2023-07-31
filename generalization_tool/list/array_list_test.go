package list

import (
	"Prove/generalization_tool/internal/errs"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		index   int
		newVal  int
		wantRes []int
		wantErr error
	}{
		{
			name:    "add num to index left",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   0,
			newVal:  0,
			wantRes: []int{0, 1, 2, 3},
		},
		{
			name:    "add num to index right",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   3,
			newVal:  4,
			wantRes: []int{1, 2, 3, 4},
		},
		{
			name:    "add num to index mid",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   1,
			newVal:  4,
			wantRes: []int{1, 4, 2, 3},
		},
		{
			name:    "add num to index -1",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   -1,
			newVal:  4,
			wantErr: errs.NewErrIndexOutOfRange(3, -1),
		},
		{
			name:    "add num to index OutOfRange",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   4,
			newVal:  4,
			wantErr: errs.NewErrIndexOutOfRange(3, 4),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Add(tc.index, tc.newVal)
			assert.Equal(t, tc.wantErr, err)
			// 因为返回了 error，所以我们不用继续往下比较了
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.values)
		})
	}
}

func TestArrayList_Cap(t *testing.T) {
	testCases := []struct {
		name    string
		wantCap int
		list    *ArrayList[int]
	}{
		{
			name:    "the same as actual",
			wantCap: 5,
			list: &ArrayList[int]{
				values: make([]int, 5),
			},
		},
		{
			name:    "nil",
			wantCap: 0,
			list: &ArrayList[int]{
				values: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.list.Cap()
			assert.Equal(t, tc.wantCap, res)
		})
	}
}

func BenchmarkArrayList_Cap(b *testing.B) {
	list := &ArrayList[int]{
		values: make([]int, 0),
	}
	b.Run("Cap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Cap()
		}
	})

	b.Run("Runtime Cap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = cap(list.values)
		}
	})
}

func TestArrayList_Append(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		newVal  []int
		wantRes []int
	}{
		{
			name:    "append non-empty values to non-empty list",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			newVal:  []int{4, 5, 6, 7},
			wantRes: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:    "append empty values to non-empty list",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			newVal:  []int{},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "append nil to non-empty list",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "append non-empty values to empty list",
			list:    NewArrayListOf[int]([]int{}),
			newVal:  []int{4, 5, 6, 7},
			wantRes: []int{4, 5, 6, 7},
		},
		{
			name:    "append empty values to empty list",
			list:    NewArrayListOf[int]([]int{}),
			newVal:  []int{},
			wantRes: []int{},
		},
		{
			name:    "append nil values to empty list",
			list:    NewArrayListOf[int]([]int{}),
			newVal:  nil,
			wantRes: []int{},
		},
		{
			name:    "append non-empty values to nil list",
			list:    NewArrayListOf[int](nil),
			newVal:  []int{4, 5, 6, 7},
			wantRes: []int{4, 5, 6, 7},
		},
		{
			name:    "append empty values to nil list",
			list:    NewArrayListOf[int](nil),
			newVal:  []int{},
			wantRes: []int{},
		},
		{
			name:    "append nil to nil list",
			list:    NewArrayListOf[int](nil),
			newVal:  nil,
			wantRes: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Append(tc.newVal...)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.AsSlice())
		})
	}
}

func TestArrayList_Delete(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		index   int
		wantVal int
		wantRes []int
		wantErr error
	}{
		{
			name: "deleted",
			list: &ArrayList[int]{
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			index:   8,
			wantVal: 9,
			wantRes: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "index out of range",
			list: &ArrayList[int]{
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			index:   10,
			wantErr: errs.NewErrIndexOutOfRange(9, 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.list.Delete(tc.index)
			fmt.Println("val:", val)
			fmt.Println("tc.wantVal:", tc.wantVal)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.values)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestArrayList_Delete_Shrink(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		index   int
		wantVal int
		wantRes []int
		wantErr error
	}{
		{
			name: "deleted",
			list: &ArrayList[int]{
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			index:   8,
			wantVal: 9,
			wantRes: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "index out of range",
			list: &ArrayList[int]{
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			index:   10,
			wantErr: errs.NewErrIndexOutOfRange(9, 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.list.Delete(tc.index)
			fmt.Println("val:", val)
			fmt.Println("tc.wantVal:", tc.wantVal)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.values)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestArrayList_Len(t *testing.T) {

}

func TestArrayList_Get(t *testing.T) {

}
func TestArrayList_Range(t *testing.T) {

}

func TestArrayList_AsSlice(t *testing.T) {

}

func TestArrayList_Set(t *testing.T) {

}
