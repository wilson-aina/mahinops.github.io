---
title: "Leetcode - Sortest Word Distance(Premium)"
date: 2024-04-15
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Two Pointer]
---


# [Leetcode - Dummy](https://leetcode.com/problems/shortest-word-distance/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-shortest-word-distance.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-shortest-word-distance/)


- Given an array of strings words and two different strings that already exist in the array word1 and word2, return the shortest distance between these two words in the list.


Solution

```
import (
)

type Solution struct{}

// shortestDistance finds the shortest distance between two words in an array.
func (s *Solution) shortestDistance(words []string, word1 string, word2 string) int {
	firstPosition, secondPosition := -1, -1
	distance := len(words)

	for ind, word := range words{
		if word==word1{
			firstPosition = ind
		}else if word==word2{
			secondPosition = ind
		}

		if firstPosition!=-1 && secondPosition!=-1{

			distance= currentDistance(firstPosition, secondPosition, distance)
		}
	}

	return distance
}

func currentDistance(first, second, distance int)int{
	cD := second-first
	if cD<0{
		cD = (-1)*cD
	}

	if cD < distance{
		return cD
	}
	return distance
}

```
