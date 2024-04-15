---
title: "Leetcode - 345. Reverse Vowels of a String"
date: 2024-04-14
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Two Pointer, String]
---


# [Leetcode - 345. Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-reverse-vowels-of-a-string.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-reverse-vowels-of-a-string/)


- Given a string s, reverse only all the vowels in the string and return it.

  The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.


Solution

```go
func reverseVowels(s string) string {
    	// as we will replace inplace and strings are immutable in go, hence we converted them to rune
	strToRunes := []rune(s)
	firstInd, lastInd := 0, len(strToRunes)-1

	vowels := "aeiouAEIOU"

	for firstInd < lastInd{
		for firstInd<lastInd && !strings.ContainsRune(vowels, strToRunes[firstInd]){
				firstInd++
			}

		for firstInd<lastInd && !strings.ContainsRune(vowels, strToRunes[lastInd]){
				lastInd--
			}

		strToRunes[firstInd], strToRunes[lastInd] = strToRunes[lastInd], strToRunes[firstInd]
		firstInd++
		lastInd--
	}

	return string(strToRunes) 
}

```
