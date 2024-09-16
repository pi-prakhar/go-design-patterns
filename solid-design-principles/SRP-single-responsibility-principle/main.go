package main

var entryCount = 0

type Journal struct {
	entryies []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++

}
