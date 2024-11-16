/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start

import "sort"

func combinationSum(candidates []int, target int) [][]int {
    // Input validation
    if len(candidates) == 0 || target <= 0 {
        return nil
    }

    // Sort to enable early termination
    sort.Ints(candidates)

    result := make([][]int, 0)
    backtrack(&result, make([]int, 0, target/candidates[0]), candidates, 0, target)
    return result
}

func backtrack(result *[][]int, combination []int, candidates []int, start, target int) {
    if target == 0 {
        // Create a new slice to avoid modifying the original
        temp := make([]int, len(combination))
        copy(temp, combination)
        *result = append(*result, temp)
        return
    }
    
    for i := start; i < len(candidates); i++ {
        // Early termination if candidate is too large
        if candidates[i] > target {
            break
        }

        combination = append(combination, candidates[i])
        backtrack(result, combination, candidates, i, target - candidates[i])
        combination = combination[:len(combination)-1]
    }
}

// @lc code=end

