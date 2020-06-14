package main

type Rank struct {
	name  string
	count int
}

type Ranking []Rank

func (r Ranking) Len() int { return len(r) }

func (r Ranking) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Ranking) Less(i, j int) bool {
	return r[i].count < r[j].count
}

func (r Ranking) LessByName(i, j int) bool {
	return r[i].name < r[j].name
}
