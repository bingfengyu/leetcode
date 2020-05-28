package main

func main() {

}

type MountainArray struct {
}

func (this *MountainArray) get(index int) int { return index }
func (this *MountainArray) length() int       { return 10 }

//试题地址：https://leetcode-cn.com/problems/find-in-mountain-array/comments/
func findInMountainArray(target int, mountainArr *MountainArray) int {
	result := -1
	lens := mountainArr.length() - 1
	i := 0
	j := lens
	mid := 0
	for i <= j {
		mid = (i + j) / 2
		leftArr := mountainArr.get(mid - 1)
		midArr := mountainArr.get(mid)
		rightArr := mountainArr.get(mid + 1)
		if midArr > leftArr && midArr > rightArr {
			break
		}
		if leftArr < rightArr {
			i = mid
		}
		if leftArr > rightArr {
			j = mid
		}
	}

	i = 0
	j = mid
	for i <= j {
		temp := (i + j) / 2
		tempArr := mountainArr.get(temp)
		if tempArr == target {
			return temp
		}
		if tempArr > target {
			j = temp - 1
		}
		if tempArr < target {
			i = temp + 1
		}
	}

	i = mid
	j = lens
	for i <= j {
		temp := (i + j) / 2
		tempArr := mountainArr.get(temp)
		if tempArr == target {
			return temp
		}
		if tempArr > target {
			i = temp + 1
		}
		if tempArr < target {
			j = temp - 1
		}
	}

	return result
}
