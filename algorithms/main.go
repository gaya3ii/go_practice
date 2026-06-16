// Command algorithms runs a small demo of each solved problem in this module.
package main

import (
	"fmt"

	// "github.com/gayathriad/go_practice/algorithms/hashing"
	// "github.com/gayathriad/go_practice/algorithms/prefix_sum"
	// "github.com/gayathriad/go_practice/algorithms/sliding_window"
	// "github.com/gayathriad/go_practice/algorithms/sorting"
	// "github.com/gayathriad/go_practice/algorithms/stack"
	"github.com/gayathriad/go_practice/algorithms/two_pointers"
)

func main() {
	// fmt.Println("-- hashing --")
	// fmt.Println("GroupAnagrams:", hashing.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// fmt.Println("NonRepeat:", hashing.NonRepeat("loveleetcode"))
	// fmt.Println("WordCount:", hashing.WordCount("golang"))

	// fmt.Println("\n-- prefix_sum --")
	// fmt.Println("ProductExceptSelf:", prefix_sum.ProductExceptSelf([]int{1, 2, 3, 4}))

	// fmt.Println("\n-- sliding_window --")
	// fmt.Println("LengthOfLongestSubstring:", sliding_window.LengthOfLongestSubstring("abcabcbb"))
	// fmt.Println("LongestSubarray:", sliding_window.LongestSubarray([]int{1, 1, 0, 1}))
	// fmt.Println("LongestOnes:", sliding_window.LongestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	// fmt.Println("FindMaxAverage:", sliding_window.FindMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))

	// fmt.Println("\n-- sorting --")
	// fmt.Println("LongestConsecutive:", sorting.LongestConsecutive([]int{100, 4, 200, 1, 3, 2}))

	// fmt.Println("\n-- stack --")
	// fmt.Println("SuperReducedStrings:", stack.SuperReducedStrings("aabccba"))
	// fmt.Println("IsValid:", stack.IsValid("([{}])"))

	fmt.Println("\n-- two_pointers --")
	cycleNode := &two_pointers.ListNode{Value: 2}
	cycleNode.Next = &two_pointers.ListNode{Value: 3, Next: cycleNode}
	cycleHead := &two_pointers.ListNode{Value: 1, Next: cycleNode}
	fmt.Println("HasCycle:", two_pointers.HasCycle(cycleHead))

	list1 := &two_pointers.ListNode{Value: 1, Next: &two_pointers.ListNode{Value: 2, Next: &two_pointers.ListNode{Value: 4}}}
	list2 := &two_pointers.ListNode{Value: 1, Next: &two_pointers.ListNode{Value: 3, Next: &two_pointers.ListNode{Value: 4}}}
	merged := two_pointers.MergeSortedArray(list1, list2)
	fmt.Print("MergeSortedArray: ")
	for n := merged; n != nil; n = n.Next {
		fmt.Print(n.Value, " ")
	}
	fmt.Println()

	fmt.Println("IsPalindrome:", two_pointers.IsPalindrome("A man, a plan, a canal: Panama"))
}
