---
title: "Leetcode - 226. Invert Binary Tree"
date: 2024-05-02
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Tree, Depth-First Search, DFS, Breadth-First Search, BFS, Binary Tree]
---


# [Leetcode - 226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-invert-binary-tree.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-invert-binary-tree/)


- Given the root of a binary tree, invert the tree, and return its root.


Solution

```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil{
        return root
    }

    tmp := root.Left
    root.Left = root.Right
    root.Right = tmp

    invertTree(root.Left)
    invertTree(root.Right)
    return root
}
```

Another approach...

```go
func invertTree(root *TreeNode) *TreeNode {
    traverse(root)
    return root
}

func traverse(root *TreeNode){
    if root == nil{
        return
    }

    left := root.Left
    right := root.Right

    root.Left = right
    root.Right = left

    traverse(root.Left)
    traverse(root.Right)
}

```
