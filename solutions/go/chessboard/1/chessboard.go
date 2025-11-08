package chessboard

const lenFile = 8
const minFile = 1

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	fileCount := 0

	if _, exists := cb[file]; !exists {
		return fileCount
	}

	for _, val := range cb[file] {
		if val {
			fileCount += 1
		}
	}

	return fileCount
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	rankCount := 0

	if rank > lenFile || rank < minFile {
		return rankCount
	}

	for _, mapVal := range cb {
		if mapVal[rank-1] {
			rankCount += 1
		}
	}

	return rankCount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	sumLenFiles := 0

	for _, val := range cb {
		sumLenFiles += len(val)
	}

	return sumLenFiles
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	sumOccupy := 0

	for _, val := range cb {
		for _, item := range val {
			if item {
				sumOccupy += 1
			}
		}
	}

	return sumOccupy
}
