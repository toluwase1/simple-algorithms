package main

import (
	"fmt"
	"sort"
	"strings"
)

func foobar() func() int {
	x := 0
	y := 1
	return func() int {
		x = y
		y = y + x
		 fmt.Println(x)
		return x
	}
}

func search(nums []int, target int) int {
	for i,_:= range nums {
		if nums[i]==target {
			return i
		}
	}
	return -1
}

type MyHashSet struct {
	set[] int
}


func Constructor() MyHashSet {
	return MyHashSet{}
}

//func (this *MyHashSet) Add(key int)  {
//	this.set = append(this.set, key)
//}
//
//
////func (this *MyHashSet) Remove(key int)  {
////		removed:=len(this.set)-2
////		this.set = this.set[:removed]
////}
//
//func (this *MyHashSet) Remove(key int)  {
//	for i, _ := range this.set {
//		if i==key {
//			this.set [i] = this.set[len(this.set)-1]
//			//v, this.set[len(this.set)-1] = this.set[len(this.set)-1], v
//			fmt.Println(this.set)
//		}
//	}
//	fmt.Println(this.set)
//	this.set = this.set[:len(this.set)-1]
//	fmt.Println(this.set)
//}

//
//func (this *MyHashSet) Contains(key int) bool {
//for _,v:= range this.set {
//	if v == key {
//		return true
//	}
//}
//	return false
//}




func isBadVersion (n int) bool {
	return true
}

func valuesLessThanN (n int) int {
	for n < 1 {
		fmt.Println(n)
	}
	return n
}

//MyHashSet myHashSet = new MyHashSet();
//myHashSet.add(1);      // set = [1]
//myHashSet.add(2);      // set = [1, 2]
//myHashSet.contains(1); // return True
//myHashSet.contains(3); // return False, (not found)
//myHashSet.add(2);      // set = [1, 2]
//myHashSet.contains(2); // return True
//myHashSet.remove(2);   // set = [1]
//myHashSet.contains(2); // return False, (already removed)
//func main() {
//	//fmt.Print(2<<10)
//	//f:=foobar()
//	////fmt.Println(f())
//	////fmt.Println("_________")
//	//for i:=0; i<10; i++ {
//	//	if i%2==0 {
//	//		fmt.Println(f(), " ")
//	//	}
//	//}
////	fmt.Println(isBadVersion(4))
////myHashSet:= MyHashSet{}
////myHashSet.Add(10)
////myHashSet.Add(20)
////myHashSet.Add(30)
////myHashSet.Add(40)
////myHashSet.Add(50)
////myHashSet.Add(60)
////myHashSet.Add(70)
////myHashSet.Add(80)
////myHashSet.Remove(20)
////fmt.Println(myHashSet.Contains(60))
////fmt.Println(myHashSet)
////	arr := []int{1,2,3,4,5,6} // 6 elements, last element at index 5
////	fmt.Println(len(arr)) // 6
////	fmt.Println(len(arr)-1) // 5
////	fmt.Println(arr[len(arr)-1]) // 6 <- element at index 5 (last element)
////fmt.Println(isBadVersion(5))
////	fmt.Println(firstBadVersion(8))
//	//n:=2
//	//y:=8
//	////for y >= n {
//	////	fmt.Println(y+50)
//	////}
//	//fmt.Println(valuesLessThanN(10))
//}

func firstBadVersion(n int) int {
	badVersions:= [] int {}
	for n >=1  && isBadVersion(n) {
		badVersions = append(badVersions, n)
		n--
	}
	fmt.Println(badVersions)
	return badVersions[len(badVersions)-1]
	//badVersions:= [] int {}
	//start:=1
	//for start<=4{
	//	if isBadVersion(n) {
	//		badVersions = append(badVersions, start)
	//		start++
	//		fmt.Println(badVersions)
	//	}
	//	fmt.Println(badVersions)
	//}
	//return badVersions[len(badVersions)-1]
	//fmt.Println(n)
	//return n
}

func firstBadVersion3 (n int) int {
	start:=1
	end:= n
	for start < end {
		if isBadVersion(n) {
			return start
		}
	}
	return start
}

func firstBadVersion2(n int) int {
	fmt.Println(n)
	start:=1
	end:=n
	for start < end {
		mid:= (start+end)/2
		if isBadVersion(mid) {
			end = mid
		} else {
			start= mid+1
		}

	}
	return start
}
func CountSheeps(numbers []bool) int {
	numOfSheep := [] bool {}
	for i,j := range numbers {
		if numbers[i] == true {
			numOfSheep = append(numOfSheep, j)
		}
	}
	return len(numOfSheep)// your code here
}



func SortNumbers(numbers []int) []int {
	if (len(numbers)>0) {
		 sort.Ints(numbers)
		return numbers
	}
	return []int{} // your code here
}

func solution(str, ending string) bool {
	//endsWith:= strings.HasPrefix(str, ending)
	if strings.HasSuffix(str, ending) {
		return true
	}
	return false //
}

func main() {
	fmt.Println(firstBadVersion(50))
	fmt.Println(firstBadVersion2(50))
	fmt.Println(solution("abc", "bc"))
}