package maze_with_recursion

import (
	"reflect"
	"testing"
)

func TestSimpleMaze(t *testing.T) {
	maze := []string{
		"xxxxxxx x",
		"x       x",
		"x xxxxxxx",
	}

	start := Point{Row: 2, Col: 1}
	end := Point{Row: 0, Col: 7}

	path, err := SolveMaze(maze, "x", start, end)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(path) == 0 {
		t.Error("Expected non-empty path")
	}

	if path[0].Row != 2 || path[0].Col != 1 {
		t.Errorf("Expected start at (2,1), got (%d,%d)", path[0].Row, path[0].Col)
	}

	lastPoint := path[len(path)-1]
	if lastPoint.Row != 0 || lastPoint.Col != 7 {
		t.Errorf("Expected end at (0,7), got (%d,%d)", lastPoint.Row, lastPoint.Col)
	}
}

func TestComplexMaze(t *testing.T) {
	maze := []string{
		"xxxxxxxx",
		"x   x  x",
		"x x x xx",
		"x x    x",
		"xxxxxxxx",
	}

	start := Point{Row: 1, Col: 1}
	end := Point{Row: 3, Col: 6}

	path, err := SolveMaze(maze, "x", start, end)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(path) == 0 {
		t.Error("Expected non-empty path")
	}

	if path[0].Row != 1 || path[0].Col != 1 {
		t.Errorf("Expected start at (1,1), got (%d,%d)", path[0].Row, path[0].Col)
	}
}

func TestNoSolutionMaze(t *testing.T) {
	maze := []string{
		"xxxxxxx",
		"x  x  x",
		"xxxxxxx",
	}

	start := Point{Row: 1, Col: 1}
	end := Point{Row: 1, Col: 5}

	_, err := SolveMaze(maze, "x", start, end)
	if err == nil {
		t.Error("Expected error for unsolvable maze")
	}
}

func TestSingleCellMaze(t *testing.T) {
	maze := []string{
		"SE",
	}

	start := Point{Row: 0, Col: 0}
	end := Point{Row: 0, Col: 1}

	path, err := SolveMaze(maze, "x", start, end)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []Point{{Row: 0, Col: 0}, {Row: 0, Col: 1}}
	if !reflect.DeepEqual(path, expected) {
		t.Errorf("Expected path %v, got %v", expected, path)
	}
}

func TestLargeMaze(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx",
		"x        x",
		"x xxxxxx x",
		"x        x",
		"x xxxxxx x",
		"x        x",
		"xxxxxxxxxx",
	}

	start := Point{Row: 1, Col: 1}
	end := Point{Row: 5, Col: 8}

	path, err := SolveMaze(maze, "x", start, end)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(path) == 0 {
		t.Error("Expected non-empty path")
	}
}

func TestMazeWithMultiplePaths(t *testing.T) {
	maze := []string{
		"xxxxx",
		"x   x",
		"x   x",
		"xxxxx",
	}

	start := Point{Row: 1, Col: 1}
	end := Point{Row: 1, Col: 3}

	path, err := SolveMaze(maze, "x", start, end)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(path) == 0 {
		t.Error("Expected non-empty path")
	}
}

func TestEmptyMaze(t *testing.T) {
	maze := []string{}
	start := Point{Row: 0, Col: 0}
	end := Point{Row: 0, Col: 0}

	_, err := SolveMaze(maze, "x", start, end)
	if err == nil {
		t.Error("Expected error for empty maze")
	}
}

func TestOutOfBoundsStart(t *testing.T) {
	maze := []string{
		"xxxx x",
		"x    x",
		"xxxxxx",
	}

	start := Point{Row: -1, Col: 0}
	end := Point{Row: 0, Col: 4}

	_, err := SolveMaze(maze, "x", start, end)
	if err == nil {
		t.Error("Expected error for out of bounds start")
	}
}

func TestOutOfBoundsEnd(t *testing.T) {
	maze := []string{
		"xxxxxx",
		"x    x",
		"xxxxxx",
	}

	start := Point{Row: 1, Col: 1}
	end := Point{Row: 10, Col: 10}

	_, err := SolveMaze(maze, "x", start, end)
	if err == nil {
		t.Error("Expected error for out of bounds end")
	}
}

func TestInconsistentMazeRows(t *testing.T) {
	maze := []string{
		"xxxx x",
		"x  x",
		"xxxxxx",
	}

	start := Point{Row: 1, Col: 1}
	end := Point{Row: 0, Col: 4}

	_, err := SolveMaze(maze, "x", start, end)
	if err == nil {
		t.Error("Expected error for inconsistent row lengths")
	}
}

func TestValidateMaze(t *testing.T) {
	maze := []string{
		"xxxxxxx x",
		"x       x",
		"x xxxxxxx",
	}

	err := ValidateMaze(maze)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestValidateMazeMultipleStarts(t *testing.T) {
	maze := []string{
		"xxxxxxx x",
		"x       x",
		"x xxxxxxx",
	}

	err := ValidateMaze(maze)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestValidateMazeMultipleEnds(t *testing.T) {
	maze := []string{
		"xxxxxxx x",
		"x       x",
		"x xxxxxxx",
	}

	err := ValidateMaze(maze)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestValidateMazeNoStart(t *testing.T) {
	maze := []string{
		"xxxxxxx x",
		"x       x",
		"x xxxxxxx",
	}

	err := ValidateMaze(maze)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestValidateMazeNoEnd(t *testing.T) {
	maze := []string{
		"xxxxxxx x",
		"x       x",
		"x xxxxxxx",
	}

	err := ValidateMaze(maze)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestPrintMazeWithPath(t *testing.T) {
	maze := []string{
		"xxx x",
		"x   x",
		"xxxxx",
	}

	path := []Point{
		{Row: 1, Col: 1},
		{Row: 1, Col: 2},
		{Row: 1, Col: 3},
		{Row: 0, Col: 3},
	}

	result := PrintMazeWithPath(maze, path, "x")
	expected := "xxx*x\nx***x\nxxxxx"

	if result != expected {
		t.Errorf("Expected maze with path:\n%s\nGot:\n%s", expected, result)
	}
}

func TestDirectionalMovement(t *testing.T) {
	maze := []string{
		"xxxxx",
		"x   x",
		"xxxxx",
	}

	start := Point{Row: 1, Col: 1}
	end := Point{Row: 1, Col: 3}

	path, err := SolveMaze(maze, "x", start, end)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []Point{
		{Row: 1, Col: 1},
		{Row: 1, Col: 2},
		{Row: 1, Col: 3},
	}

	if !reflect.DeepEqual(path, expected) {
		t.Errorf("Expected path %v, got %v", expected, path)
	}
}

func TestCustomWallCharacter(t *testing.T) {
	maze := []string{
		"xxxxx x",
		"x     x",
		"x xxxxx",
	}

	start := Point{Row: 2, Col: 1}
	end := Point{Row: 0, Col: 5}

	path, err := SolveMaze(maze, "x", start, end)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(path) == 0 {
		t.Error("Expected non-empty path")
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	data, ok := result.(map[string]any)
	if !ok {
		t.Error("Expected map result")
	}

	if !data["success"].(bool) {
		t.Error("Expected successful maze solving")
	}

	pathLength, ok := data["path_length"].(int)
	if !ok || pathLength == 0 {
		t.Error("Expected non-zero path length")
	}
}

func BenchmarkSimpleMaze(b *testing.B) {
	maze := []string{
		"xxxxx x",
		"x     x",
		"x xxxxx",
	}
	start := Point{Row: 2, Col: 1}
	end := Point{Row: 0, Col: 5}

	for i := 0; b.Loop(); i++ {
		SolveMaze(maze, "x", start, end)
	}
}

func BenchmarkComplexMaze(b *testing.B) {
	maze := []string{
		"xxxxxxxxxx",
		"x        x",
		"x xxxxxx x",
		"x        x",
		"x xxxxxx x",
		"x        x",
		"xxxxxxxxxx",
	}
	start := Point{Row: 1, Col: 1}
	end := Point{Row: 5, Col: 8}

	for i := 0; b.Loop(); i++ {
		SolveMaze(maze, "x", start, end)
	}
}

func BenchmarkLargeMaze(b *testing.B) {
	maze := []string{
		"xxxxxxxxxxxxxxxx",
		"x              x",
		"x xxxxxxxxxxxx x",
		"x              x",
		"x xxxxxxxxxxxx x",
		"x              x",
		"x xxxxxxxxxxxx x",
		"x              x",
		"xxxxxxxxxxxxxxxx",
	}
	start := Point{Row: 1, Col: 1}
	end := Point{Row: 7, Col: 14}

	for i := 0; b.Loop(); i++ {
		SolveMaze(maze, "x", start, end)
	}
}
