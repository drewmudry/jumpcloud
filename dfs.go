package main

import (
    "fmt"
)

func Solution(board [][]int) int {
    if len(board) == 0 || len(board[0]) == 0 {
        return 0
    }
    
    rows := len(board)
    cols := len(board[0])
    maxResult := 0
    
    // Directions: up, down, left, right
    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    
    // Try starting from every position
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            // OPTIMIZATION: Skip DFS if starting digit cannot possibly exceed current maxResult
            // If maxResult is 9999 and we're starting with digit 1, 
            // the maximum possible 4-digit number starting with 1 is 1999, which is < 9999
            startingDigit := board[i][j]
            if maxResult > 0 {
                // Calculate the maximum possible 4-digit number starting with this digit
                maxPossibleStartingWithDigit := startingDigit
                for k := 1; k < 4; k++ {
                    maxPossibleStartingWithDigit = maxPossibleStartingWithDigit*10 + 9
                }
                
                // If even the best possible outcome can't exceed current maxResult, skip this position
                if maxPossibleStartingWithDigit <= maxResult {
                    continue
                }
            }
            
            visited := make([][]bool, rows)
            for k := range visited {
                visited[k] = make([]bool, cols)
            }
            
            // Start DFS from position (i, j)
            result := dfs(board, i, j, 1, board[i][j], visited, directions, maxResult)
            if result > maxResult {
                maxResult = result
            }
        }
    }
    
    return maxResult
}

func dfs(board [][]int, row, col, pathLength, currentNumber int, visited [][]bool, directions [][]int, maxResult int) int {
    // If we've reached path length 4, return the current number
    if pathLength == 4 {
        return currentNumber
    }
    
    visited[row][col] = true
    localMax := maxResult
    
    // Try all 4 directions
    for _, dir := range directions {
        newRow := row + dir[0]
        newCol := col + dir[1]
        
        // Check bounds and if not visited
        if newRow >= 0 && newRow < len(board) && 
           newCol >= 0 && newCol < len(board[0]) && 
           !visited[newRow][newCol] {
            
            // Create new number by appending the digit
            newNumber := currentNumber*10 + board[newRow][newCol]
            
            // OPTIMIZATION: Early termination check (commented out)
            // This optimization checks if the remaining path could potentially exceed maxResult
            // before exploring it, which could save significant computation time
            /*
            remainingDigits := 4 - pathLength
            if remainingDigits > 0 {
                // Calculate the maximum possible value for the remaining path
                // Assuming the maximum digit in the board is 9
                maxPossibleValue := newNumber
                for i := 0; i < remainingDigits; i++ {
                    maxPossibleValue = maxPossibleValue*10 + 9
                }
                
                // If even the maximum possible value can't exceed current maxResult, skip this path
                if maxPossibleValue <= localMax {
                    continue
                }
            }
            */
            
            // Continue DFS
            result := dfs(board, newRow, newCol, pathLength+1, newNumber, visited, directions, localMax)
            if result > localMax {
                localMax = result
            }
        }
    }
    
    visited[row][col] = false // Backtrack=
    return localMax
}

// Example usage and test
func main() {
    // Test case 1
    board1 := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Printf("Test 1 - Result: %d\n", Solution(board1))
    
    // Test case 2
    board2 := [][]int{
        {9, 8, 7},
        {6, 5, 4},
        {3, 2, 1},
    }
    fmt.Printf("Test 2 - Result: %d\n", Solution(board2))
    
    // Test case 3 - single row
    board3 := [][]int{
        {9, 8, 7, 6, 5},
    }
    fmt.Printf("Test 3 - Result: %d\n", Solution(board3))
}