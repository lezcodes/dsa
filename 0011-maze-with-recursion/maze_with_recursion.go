package maze_with_recursion

import (
	"errors"
	"strings"
)

type Point struct {
	Row int
	Col int
}

var directions = []Point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze []string, wall string, current Point, end Point, visited [][]bool, path *[]Point) bool {
	if current.Col < 0 || current.Col >= len(maze[0]) || current.Row < 0 || current.Row >= len(maze) {
		return false
	}

	if string(maze[current.Row][current.Col]) == wall {
		return false
	}

	if current.Row == end.Row && current.Col == end.Col {
		*path = append(*path, current)
		return true
	}

	if visited[current.Row][current.Col] {
		return false
	}

	visited[current.Row][current.Col] = true
	*path = append(*path, current)

	for _, dir := range directions {
		next := Point{Row: current.Row + dir.Row, Col: current.Col + dir.Col}
		if walk(maze, wall, next, end, visited, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func SolveMaze(maze []string, wall string, start Point, end Point) ([]Point, error) {
	if len(maze) == 0 {
		return nil, errors.New("empty maze")
	}

	for i, row := range maze {
		if i > 0 && len(row) != len(maze[0]) {
			return nil, errors.New("inconsistent maze row lengths")
		}
	}

	if start.Row < 0 || start.Row >= len(maze) || start.Col < 0 || start.Col >= len(maze[0]) {
		return nil, errors.New("start position out of bounds")
	}

	if end.Row < 0 || end.Row >= len(maze) || end.Col < 0 || end.Col >= len(maze[0]) {
		return nil, errors.New("end position out of bounds")
	}

	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}

	var path []Point
	if walk(maze, wall, start, end, visited, &path) {
		return path, nil
	}

	return nil, errors.New("no path found from start to end")
}

func ValidateMaze(maze []string) error {
	if len(maze) == 0 {
		return errors.New("empty maze")
	}

	for i, row := range maze {
		if i > 0 && len(row) != len(maze[0]) {
			return errors.New("inconsistent maze row lengths")
		}
	}

	return nil
}

func PrintMazeWithPath(maze []string, path []Point, wall string) string {
	pathSet := make(map[Point]bool)
	for _, p := range path {
		pathSet[p] = true
	}

	var result strings.Builder
	for row := range maze {
		for col := range len(maze[row]) {
			point := Point{Row: row, Col: col}
			char := string(maze[row][col])

			if pathSet[point] && char != wall {
				result.WriteByte('*')
			} else {
				result.WriteByte(maze[row][col])
			}
		}
		if row < len(maze)-1 {
			result.WriteByte('\n')
		}
	}
	return result.String()
}

func Run() any {
	maze := []string{
		"xxxxxxx x",
		"x       x",
		"x xxxxxxx",
	}

	start := Point{Row: 2, Col: 1}
	end := Point{Row: 0, Col: 7}

	err := ValidateMaze(maze)
	if err != nil {
		return map[string]any{
			"error": err.Error(),
			"maze":  maze,
		}
	}

	path, err := SolveMaze(maze, "x", start, end)
	if err != nil {
		return map[string]any{
			"error": err.Error(),
			"maze":  maze,
		}
	}

	mazeWithPath := PrintMazeWithPath(maze, path, "x")

	return map[string]any{
		"original_maze":  maze,
		"start":          start,
		"end":            end,
		"path":           path,
		"path_length":    len(path),
		"maze_with_path": strings.Split(mazeWithPath, "\n"),
		"success":        true,
	}
}
