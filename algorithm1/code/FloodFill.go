package algorithm1

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if newColor == color {
		return image
	}
	changeColor(image, sr, sc, newColor)
	return image
}

func changeColor(image [][]int, sr int, sc int, newColor int) {
	if image[sr][sc] == newColor {
		return
	}

	oldColor := image[sr][sc]
	image[sr][sc] = newColor

	for i := 0; i < 4; i++ {
		if (sr+dir[i][0] >= 0 && sr+dir[i][0] < len(image)) && (sc+dir[i][1] >= 0 && sc+dir[i][1] < len(image[0])) && image[sr+dir[i][0]][sc+dir[i][1]] == oldColor {
			changeColor(image, sr+dir[i][0], sc+dir[i][1], newColor)
		}
	}
}
