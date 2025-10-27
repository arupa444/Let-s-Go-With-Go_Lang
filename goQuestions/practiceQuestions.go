package goQuestions

import (
    "fmt"
    "math"
    "strings"
    "golang.org/x/exp/constraints"
    "sort"
)

// Write a program that prints numbers from 1 to 100, but for multiples of 3 print "Fizz",
// for multiples of 5 print "Buzz", and for multiples of both print "FizzBuzz".


func GoFirstQue(){
    for i := 1; i<=100; i++{
        if i%3 == 0 && i%5 == 0{
            fmt.Println("FizzBuzz")
        }else if i%3 == 0{
            fmt.Println("Fizz")
        }else if i%4 == 0{
            fmt.Println("Buzz")
        }else{
            fmt.Println(i)
        }
    }
}

// Write a function that takes a string and returns it reversed.

func GoSecondQue(str string)string{
    storeStr := []rune(str)
//     storeStr1 := []byte(str)
//     fmt.Println(storeStr1)
    for i, j := 0, len(storeStr)-1; i<j; i, j = i+1, j-1{
        storeStr[i], storeStr[j] = storeStr[j], storeStr[i]
    }
    return string(storeStr)
}


// Write a program that checks if a given number is prime or not.

func GoThirdQue(ele int)string{
    ele1 := math.Sqrt(float64(ele))
    for i:=2; i<int(ele1); i++{
        if ele%i == 0{
            return "Not a Prime No."
        }
    }
    return "Prime No."
}


// Write a function that finds the largest element in an array of integers.

// Generic Function

func GoForthQue[T constraints.Ordered](nums []T) T{
    maxVal := nums[0]
    for _, ele := range nums[1:]{
        if ele>maxVal{
            maxVal = ele
        }
    }
    return maxVal
}


// Write a program that counts the number of vowels in a given string.
func GoFifthQue1(str string)int{
    counts := 0
    searchVowel := "aeiouAEIOU"
    for _,ele := range str{
        if strings.ContainsRune(searchVowel, ele){
            counts++;
        }
    }
    return counts
}

func GoFifthQue2(str string)int{
    counts := 0
    for _, ele := range str{
        switch ele{
            case 'a','e','i','o','u','A','E','I','O','U':
                counts++;
        }
    }
    return counts;
}

// Write a function that removes duplicate elements from a slice of integers/Floats/Strings.

func GoSixthQue[T constraints.Ordered](nums []T) []T{
    seen := make(map[T]bool)
    returnSlice := make([]T,0,len(nums))
    for _, ele := range nums{
        if !seen[ele]{
            seen[ele] = true
            returnSlice =append(returnSlice, ele)
        }
    }
    return returnSlice
}

// Write a program that merges two sorted slices into one sorted slice.

func GoSevenQue[T constraints.Ordered](nums, nums1 []T) []T{
    returnSlice := make([]T,0,len(nums)+len(nums1))
    i , j := 0, 0

    for i<len(nums) && j<len(nums1){
        if nums[i] <= nums1[j]{
            returnSlice = append(returnSlice, nums[i])
            i++
        }else{
            returnSlice = append(returnSlice, nums1[j])
            j++
        }
    }
    returnSlice = append(returnSlice, nums[i:]...)
    returnSlice = append(returnSlice, nums1[j:]...)
    return returnSlice
}

// Write a program that merges two slices into one sorted slice.

func GoSevenQue1[T constraints.Ordered](nums, nums1 []T) []T{
    sort.Slice(nums, func(a,b int)bool{return nums[a]<nums[b]})
    sort.Slice(nums1, func(a,b int)bool{return nums1[a]<nums1[b]})

    returnSlice := make([]T,0,len(nums)+len(nums1))
    a, b := 0, 0

    for a < len(nums) && b < len(nums1){
        if nums[a]<=nums1[b]{
            returnSlice = append(returnSlice, nums[a])
            a++;
        }else{
            returnSlice = append(returnSlice, nums1[b])
            b++;
        }
    }
    returnSlice = append(returnSlice, nums[a:]...)
    returnSlice = append(returnSlice, nums1[b:]...)
    return returnSlice
}

// Write a function that rotates a slice to the right by k positions

func GoEightQue[T constraints.Ordered](nums []T, k int) []T{
    if len(nums) == 0 || k == 0{
        return nums
    }
    k %= len(nums)
    returnSlice := make([]T, 0, len(nums))
    returnSlice = append(returnSlice, nums[k:]...)
    returnSlice = append(returnSlice, nums[:k]...)
    return returnSlice
}

// Write a program that finds the second Smallest number in a slice.

func GoNinthQue[T constraints.Ordered](nums []T) T{
    sort.Slice(nums, func(a, b int) bool{ return nums[a] < nums[b]})
    returnVal := nums[0]
    for _, ele := range nums[1:]{
        if ele != returnVal{
            return ele
        }
    }
    return returnVal
}

// Write a program that finds the second largest number in a slice.

func GoNinthQue1[T constraints.Ordered](nums []T) T{
    if len(nums) < 2{
        panic("Slice must have 2 or more element")
    }
    largest := nums[0]
    secondLargest := nums[0]

    for _, ele := range nums{
        if largest < ele{
            secondLargest = largest
            largest = ele
        }else if ele > secondLargest && ele < largest{
            secondLargest = ele
        }
    }
    if largest == secondLargest{
        panic("No second elements exist")
    }
    return secondLargest
}

// Write a function that checks if two slices are equal (contain the same elements in the same order)

func GoTenthQue[T constraints.Ordered](slice1 []T, slice2 []T) bool{
    if len(slice1) != len(slice2){
        return false
    }
    for index, ele := range slice1{
        if ele != slice2[index]{
            return false
        }
    }
    return true
}

// Write a program that counts the frequency of each word in a given sentence using a map.

func GoElevenQue(sentence string) map[string]int{
    storeWordCount := make(map[string]int)

    for _, ele := range strings.Split(sentence, " "){
        if _, exist := storeWordCount[ele]; !exist{
            storeWordCount[ele] = 1
        }else{
            storeWordCount[ele]++
        }
    }
    return storeWordCount
}

// Create a struct for a Book with fields (Title, Author, Pages) and
// write a function that prints book details.

type Book struct{
    Title string
    Authors []string
    Pages int
}

func CallBook(jhaaaat Book){
    fmt.Printf("Title of the book : %#v\n",jhaaaat.Title)
    fmt.Printf("Authors of the book : %v\n",jhaaaat.Authors)
    fmt.Printf("Pages of the book : %v\n",jhaaaat.Pages)
}


func GoTwelveQue(title string, authors []string, pages int) Book{
    var book1 Book
    book1.Title = title
    book1.Authors = authors
    book1.Pages = pages

    return book1
}

// Write a program that finds the most frequent character in a string using a map.

type MostFreq struct{
    ele string
    count int
}

func GoThirteenQue(str string)MostFreq{
    freqChar := make(map[string]int)
    var maxFreq MostFreq
    maxFreq.ele = string(str[0])
    maxFreq.count = 1
    for _, ele := range strings.Split(str, ""){
        freqChar[ele]++
        if maxFreq.count < freqChar[ele]{
            maxFreq.ele = ele
            maxFreq.count = freqChar[ele]
        }
    }
    return maxFreq
}
// Create a struct for Student with Name and Grades (slice of integers),
// and write a method to calculate average grade. 15.

type Student struct{
    Name string
    Grades []int
    AvgGrade float64
}

func GoFourteenQue(name string) Student{
    var grade int
    var student1 Student
    fmt.Print("Accepts all the grades until and unless you give -1 in console and enter: ")
    student1.Name = name
    for true{
        fmt.Scan(&grade)
        if grade == -1{
            break
        }
        student1.AvgGrade += float64(grade)
        student1.Grades = append(student1.Grades, grade)
    }
    student1.AvgGrade /= float64(len(student1.Grades))
    return student1
}

// Write a program that groups a slice of strings by their length using a map.


func GoFifteenQue(listStr []string) map[int][]string{
    groupByLength := make(map[int][]string)
    for _, ele := range listStr{
        groupByLength[len(ele)] = append(groupByLength[len(ele)], ele)
    }
    return groupByLength
}

// Write a recursive function to calculate the nth Fibonacci number.

func GoSixteenQue(nth int) int{
    if nth == 0{
        return 0
    }else if nth == 1{
        return 1
    }
    return GoSixteenQue(nth-1) + GoSixteenQue(nth-2)
}
// linear search in an slice and find the index
func GoLinearSearch(nums []int, ele int)int{
    for index, num := range nums{
        if num == ele{
            return index
        }
    }
    return -1
}

// binary search in an slice and find the index
func GoBinarySearch(nums []int, ele int)int{
    startPivot := 0
    endPivot := len(nums)
    for startPivot < endPivot{
        if nums[(startPivot+endPivot)/2] == ele{
            return (startPivot+endPivot)/2
        }else if nums[(startPivot+endPivot)/2] > ele{
            endPivot = ((startPivot+endPivot)/2)-1
        }else{
            startPivot = ((startPivot+endPivot)/2)+1
        }
    }
    return -1
}

func GoPointers(valuesForP int)*int{
    var p *int
    p = &valuesForP
    return p
}
// 18. Write a program that swaps two numbers using pointers.

func GoPointers(valuesForP int)*int{
    var p *int
    p = &valuesForP
    return p
}
