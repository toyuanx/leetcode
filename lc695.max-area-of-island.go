// @Author shovenyuan
// @Date: 2022/6/22 17:33

// Package main ...
package main

func maxAreaOfIsland(grid [][]int) int {
	var getArea = func([][]int, int, int) int { return 0 }
	getArea = func(grid [][]int, i, j int) int {
		// 遇见岛屿边界、海洋时退出
		if i < 0 || j < 0 || i > len(grid)-1 || j >= len(grid[0])-1 || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + getArea(grid, i-1, j) + getArea(grid, i+1, j) + getArea(grid, i, j-1) + getArea(grid, i, j+1)
	}

	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 发现陆地啦
			if grid[i][j] == 1 {
				area := getArea(grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
