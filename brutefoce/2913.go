package main

import "fmt"
	

func sumCounts(nums []int) int{
	sum := 0

	for i :=0; i < len(nums); i++{
		arr := make(map[int]struct{})

		for j:=i;j<len(nums); j++{
			arr[nums[j]] = struct{}{}
			sum+=(len(arr) * len(arr))
		}
	}
	return sum
}

func main(){
	var t int

	fmt.Print("Enter the test of the array: ")
	fmt.Scan(&t)

	for ; t > 0;t--{

		var n int

		fmt.Print("Enter the num of element: ")
		fmt.Scan(&n)

		nums := make([]int, n)

		for i := 0 ; i < n; i++{
			fmt.Scan(&nums[i])
		}

		fmt.Println(sumCounts(nums))
	}

}
