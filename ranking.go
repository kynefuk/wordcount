package main

// Word is an object that has word name and word count.
type Word struct {
	name  string
	count int
}

// Ranking is an array of Rank that represents a ranking of words.
type Ranking []Word

func (r Ranking) Len() int { return len(r) }

func (r Ranking) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Ranking) Less(i, j int) bool {
	return r[i].count < r[j].count
}

// LessByName is used for sort alphabetically.
func (r Ranking) LessByName(i, j int) bool {
	return r[i].name < r[j].name
}
