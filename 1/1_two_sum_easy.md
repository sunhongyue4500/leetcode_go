# 题目



给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

==你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。==

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

# 解题思路

* 第一种解法：暴力求解，O（N^2），空间复杂度O（1）
* 第二种解法：使用哈希表，key-value对应的是（target-nums[index]， index），时间复杂度O（N），空间复杂度为O（N）



# 实现

## Go

### 暴力

```
执行用时 :80 ms, 在所有 golang 提交中击败了5.09%的用户
内存消耗 :2.9 MB, 在所有 golang 提交中击败了79.17%的用户
```

```go
// 暴力
func twoSum(nums []int, target int) []int {
    if len(nums) < 2 {
        return []int{}
    }
    res := []int{}
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                res = append(res, i, j)
            }
        }
    }
    return res
}
```



### 第二种解法

执行用时 :4 ms, 在所有 Go 提交中击败了98.31%的用户

内存消耗 :3.7 MB, 在所有 Go 提交中击败了46.69%的用户

```go
func twoSum(nums []int, target int) []int {
	res := []int{}
	if len(nums) < 1 {
		return res
	}
	mapRes := make(map[int]int, 0)
	for idx, val := range nums {
		if _, ok := mapRes[val]; ok {
			res = append(res, mapRes[val], idx)
			break
		}
		mapRes[target-val] = idx
	}
	return res
}
```

简洁代码

```go
func twoSum(nums []int, target int) []int {
    index := make(map[int]int, len(nums))
    for i, a := range nums {
        if j, ok := index[target-a]; ok {
            return []int{j, i}
        }
        index[a] = i
    }
    return nil
}
```

