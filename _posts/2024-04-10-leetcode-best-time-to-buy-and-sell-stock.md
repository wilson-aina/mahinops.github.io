---
title: "Leetcode - 121. Best Time to Buy and Sell Stock"
date: 2024-04-08
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Dynamic Programming]
---

# [Leetcode - 121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-best-time-to-buy-and-sell-stock.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-best-time-to-buy-and-sell-stock/)

- You are given an array prices where prices[i] is the price of a given stock on the ith day.

    You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

    Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Solution
```
func maxProfit(prices []int) int {
    minPrice := prices[0]
    maxProfit := 0

    for _, price := range prices{
        if price<minPrice{
            minPrice = price
        }else if price-minPrice>=maxProfit{
            maxProfit = price-minPrice
        }
    }
    return maxProfit
}
```
