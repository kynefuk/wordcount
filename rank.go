package main

// Rank is an object that represents a rank of word.
type Rank struct {
	name  string
	count int
}

// Ranking is an array of Rank that represents a ranking of words.
type Ranking []Rank

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
