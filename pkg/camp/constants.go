package camp

const (
	emptyValue int = iota
	treeValue
	tentValue
)

const (
	percentageOfTrees float64 = 0.20 // The maximum number of trees as a percentage of the total number of cells in the field.
	minFieldSize      int     = 5
	maxFieldSize      int     = 15
)

// Error constant
const (
	tooLargeFieldSize    string = "the field size is too large"
	tooSmallFieldSize    string = "the field size is too small"
	incorrectRowTents    string = "incorrect data for the number of tents in the row"
	incorrectColumnTents string = "incorrect data for the number of tents in the column"
	wrongRowTents        string = "the number of tents in each row must match the assignment"
	wrongColumnTents     string = "the number of tents in each column must match the assignment"
	wrongTentPosition    string = "there should be no other tents around the tent"
)
