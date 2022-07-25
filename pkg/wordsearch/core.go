package wordsearch

import (
	"errors"
	"fmt"
	"math/rand"
	"path/filepath"

	"github.com/cheatsnake/shadify/pkg/assists"
)

var wordsDB []string

func init() {
	nouns, err := assists.ReadFileLineByLine(filepath.Join("data", "nouns.txt"))
	if err != nil {
		fmt.Println(err)
	}
	wordsDB = nouns
}

// Generate a new Core of a given width and heigth,
// corresponding to the size of the Core.Grid
func Generate(w, h int) (Core, error) {
	if w < minGridSideSize || h < minGridSideSize {
		return Core{}, errors.New(minGridSideError)
	}

	if w > maxGridSideSize || h > maxGridSideSize {
		return Core{}, errors.New(maxGridSideError)
	}

	if w*h > maxCellsCount {
		return Core{}, errors.New(tooManyCells)
	}

	grid, pWords := generateGrid(w, h)

	// fill unused cells by random letters
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "" {
				grid[i][j] = alphabet[rand.Intn(26)]
			}
		}
	}

	return Core{
		Width:      w,
		Height:     h,
		WordsCount: len(pWords),
		Grid:       grid,
		Words:      pWords,
	}, nil
}
