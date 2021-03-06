package main

import "testing"

func TestSort(t *testing.T) {
	array := []int{3, 8, 4, 1, 5, 2}
	//
	//
	//  1 8 4 3 5 2
	// 1 3 4 8 5 2


	//  迭代
	// 外层循环，x < y
	// x:0,y:9, target:0->3
	// 从后往前找，找第一个比他小的，找到y:5->2，交换
	// {2, 8, 4, 1, 5, 3, 9, 7, 11, 9}
	// 从前往后找，找第一个比他大的，找到x:1->8，交换
	// {2, 3, 4, 1, 5, 8, 9, 7, 11, 9}
	// x:1,y:5, target:0->2
	// 从右往前找，y:3->1
	// {1, 3, 4, 2, 5, 8, 9, 7, 11, 9}
	// 从前往后找 x：1->3
	// {1, 2, 4, 3, 5, 8, 9, 7, 11, 9}
	// x:1,y:3, target:0->1
	res := QSort(&array, 0, len(array) - 1)
	t.Log(res)
}