package main

import (
  "sort"
)

func QueueTime(customers []int, n int) int {
  queueTime := make([]int, n)
  
  for _, customerTime := range customers {
    queueTime[0] += customerTime
    sort.Ints(queueTime)
  }
  return queueTime[n-1]
}