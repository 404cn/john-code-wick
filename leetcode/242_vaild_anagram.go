package leetcode

// sort
/*
242. Valid Anagram
Easy

1186

129

Add to List

Share
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}
