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
            result := bfs(board, i, j, directions)
            if result > maxResult {
                maxResult = result
            }
        }
    }
    
    return maxResult
}

// State represents a position in the BFS
type State struct {
    row, col int
    number   int
    pathLen  int
    visited  [][]bool
}

func bfs(board [][]int, startRow, startCol int, directions [][]int) int {
    rows := len(board)
    cols := len(board[0])
    maxResult := 0
    
    // Initialize visited matrix for this starting position
    visited := make([][]bool, rows)
    for i := range visited {
        visited[i] = make([]bool, cols)
    }
    
    // Create initial state
    initialVisited := make([][]bool, rows)
    for i := range initialVisited {
        initialVisited[i] = make([]bool, cols)
    }
    initialVisited[startRow][startCol] = true
    
    queue := []State{{
        row:     startRow,
        col:     startCol,
        number:  board[startRow][startCol],
        pathLen: 1,
        visited: initialVisited,
    }}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        // If we've reached path length 4, update max result
        if current.pathLen == 4 {
            if current.number > maxResult {
                maxResult = current.number
            }
            continue
        }
        
        // Try all 4 directions
        for _, dir := range directions {
            newRow := current.row + dir[0]
            newCol := current.col + dir[1]
            
            // Check bounds and if not visited
            if newRow >= 0 && newRow < rows && 
               newCol >= 0 && newCol < cols && 
               !current.visited[newRow][newCol] {
                
                // Create new visited matrix for this path
                newVisited := make([][]bool, rows)
                for i := range newVisited {
                    newVisited[i] = make([]bool, cols)
                    copy(newVisited[i], current.visited[i])
                }
                newVisited[newRow][newCol] = true
                
                // Create new number by appending the digit
                newNumber := current.number*10 + board[newRow][newCol]
                
                // Add to queue
                queue = append(queue, State{
                    row:     newRow,
                    col:     newCol,
                    number:  newNumber,
                    pathLen: current.pathLen + 1,
                    visited: newVisited,
                })
            }
        }
    }
    
    return maxResult
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