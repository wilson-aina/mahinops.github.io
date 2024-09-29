---
title: "Leetcode - 104. Maximum Depth of Binary Tree"
date: 2024-05-01
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Tree, Depth-First Search, DFS, Breadth-First Search, BFS, Binary Tree]
---


# [Leetcode - 104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-maximum-depth-of-binary-tree.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-maximum-depth-of-binary-tree/)


- Given the root of a binary tree, return its maximum depth.

  A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.


Solution

```go
func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }

    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)

    if leftDepth>rightDepth{
        return leftDepth+1
    }
    return rightDepth+1
    
}
```
