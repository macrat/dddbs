package dddbs

import (
	"sort"
	"strings"
)

type DataBase struct {
	raw      map[string]string
	inverted map[rune][]string
}

func NewDataBase() (db DataBase) {
	db.raw = make(map[string]string)
	db.inverted = make(map[rune][]string)
	return
}

func (db DataBase) Add(key, value string) {
	v := []rune(value)
	for i := 0; i < len(v); i++ {
		gram := v[i]
		db.inverted[gram] = append(db.inverted[gram], key)
	}
	db.raw[key] = value
}

func (db DataBase) Get(key string) string {
	return db.raw[key]
}

type Result map[string]int

func (result Result) And(target map[string]int) (and Result) {
	and = make(Result)
	for k, rv := range result {
		if tv, ok := target[k]; ok {
			and[k] = rv + tv
		}
	}
	return
}

func (db DataBase) SearchSingleQuery(query string) (found Result) {
	found = make(map[string]int)

	queryRunes := []rune(query)

	for i := 0; i < len(queryRunes); i++ {
		for _, x := range db.inverted[queryRunes[i]] {
			if _, ok := found[x]; ok {
				found[x]++
			} else {
				found[x] = 1
			}
		}
	}

	for k, v := range found {
		if v < len(queryRunes)-1 || !strings.Contains(db.raw[k], query) {
			delete(found, k)
		} else {
			found[k] = strings.Count(db.raw[k], query)
		}
	}

	return
}

func (db DataBase) Search(query string) (found Result) {
	qs := strings.Split(query, " ")
	found = db.SearchSingleQuery(qs[0])

	for _, q := range qs[1:] {
		found = db.SearchSingleQuery(q).And(found)
	}

	return
}

type ResultEntry struct {
	Key   string
	Score int
}

type SortedResult []ResultEntry

func (result Result) Sort() SortedResult {
	var sr SortedResult

	for k, v := range result {
		sr = append(sr, ResultEntry{k, v})
	}

	sort.Sort(sort.Reverse(sr))

	return sr
}

func (sr SortedResult) Len() int {
	return len(sr)
}

func (sr SortedResult) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

func (sr SortedResult) Less(i, j int) bool {
	if sr[i].Score != sr[j].Score {
		return sr[i].Score < sr[j].Score
	} else {
		return sr[i].Key > sr[j].Key
	}
}
