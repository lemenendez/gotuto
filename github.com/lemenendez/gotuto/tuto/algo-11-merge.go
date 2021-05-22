/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

The number of elements initialized in nums1 and nums2 are m and n respectively. You may assume that nums1 has a size equal to m + n such that it has enough space to hold additional elements from nums2.
*/

package main
import (
	"fmt"
)

func shift(a[] int, pos int) {
	length := len(a)
	if pos>=length || length<2 {
		return
	}
	for I:=len(a)-1; I>=pos && I>0; I-- {
		a[I]=a[I-1]
	}
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    // loop for the input (nums2)
	pos := 0
	new_len := m
	for i:=0; i<n; i++ {
		// loop for the dest arr to find place for num1 element
		fmt.Println("processing:", i, nums2[i])
		// find a place inside nums1 for nums[i]
        for j:=pos; j<len(nums1); j++ {
            if nums2[i] <= nums1[j] || j>=new_len {
				// fmt.Print("inserting")
				shift(nums1, j)
                nums1[j] =nums2[i]
				pos = j+1
				new_len += 1
				break 
            }
        }
		// fmt.Println("new_len", new_len)
		fmt.Println(nums1)
    }
}

func main() {
	// nums1 := []int{1,2,3,0,0,0}
	// nums2 := []int{2,5,6}
	nums1 := []int{2, 0}
	nums2 := []int{1}
	fmt.Println("new input", nums2)
	fmt.Println("dest",nums1)	
	shift(nums1, 0)
	merge(nums1, 1, nums2, 1 )
	//fmt.Println(nums1)
}