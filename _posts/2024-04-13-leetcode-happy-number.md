---
title: "Leetcode - 202. Happy Number"
date: 2024-04-13
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Hash Table, Math, Two Pointer]
---

# [Leetcode - 202. Happy Number](https://leetcode.com/problems/happy-number/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-happy-number.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-happy-number/)

- Write an algorithm to determine if a number n is happy.

  A happy number is a number defined by the following process:

  Starting with any positive integer, replace the number by the sum of the squares of its digits.
  Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
  Those numbers for which this process ends in 1 are happy.
  Return true if n is a happy number, and false if not.

Solution
```go
func isHappy(n int) bool {
    visited := make(map[int]int)

    for !isVisited(n, visited){
        visited[n]=n
        sumOfSq := sumOfSquare(n)
        if sumOfSq ==1{
            return true
        }
        n = sumOfSq
    }
    return false

}

func isVisited(n int, visited map[int]int) bool{
    _, exists := visited[n]
    return exists
}

func sumOfSquare(n int)int{
    sq := 0
    for n>0{
        rem := n%10
        sq = sq+(rem*rem)
        n = n/10
    }
    return sq
}

```
