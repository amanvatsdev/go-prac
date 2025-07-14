/*
["Spring", "is", "my", "favorite", "season", "because", "of", "the", "beautiful", "changes", "that", "happen", "in", "nature.", "As", "the", "weather", "warms", "up,", "the", "flowers", "bloom", "in", "vibrant", "colors,", "creating", "a", "stunning", "landscape.", "The", "sweet", "scent", "of", "blossoms", "fills", "the", "air,", "inviting", "bees", "and", "butterflies", "to", "join", "the", "party.", "Every", "morning,", "I", "love", "to", "walk", "outside", "and", "hear", "the", "cheerful", "chirping", "of", "birds", "returning", "from", "their", "winter", "homes.", "The", "gentle", "breeze", "is", "refreshing,", "and", "it", "feels", "like", "a", "hug", "from", "nature.", "I", "also", "enjoy", "seeing", "the", "trees", "regain", "their", "green", "leaves,", "which", "makes", "everything", "look", "alive", "again.", "Spring", "is", "a", "]()

*/

package main

import (
	"fmt"
)

func main() {
	a := []string{"the", "is", "my", "favorite", "season", "because", "of", "the", "beautiful", "changes", "that", "happen", "in", "nature.", "the", "the", "weather", "warms", "up,", "the", "flowers", "bloom", "in", "vibrant", "colors,", "creating", "a", "the", "landscape.", "The", "sweet", "scent", "of", "blossoms", "fills", "the", "air,", "inviting", "bees", "and", "the", "to", "join", "the", "the.", "Every", "morning,", "I", "love", "to", "walk", "outside", "and", "the", "the", "cheerful", "chirping", "of", "birds", "returning", "from", "their", "winter", "homes.", "The", "gentle", "breeze", "is", "refreshing,", "and", "it", "feels", "like", "a", "hug", "from", "nature.", "I", "also", "enjoy", "seeing", "the", "trees", "regain", "their", "green", "leaves,", "which", "makes", "everything", "look", "alive", "again.", "Spring", "is", "a"}

	// sort.Strings(a)

	b := make(map[string]int)

	for i,v:=range a{
		b[v]++
		fmt.Printf("%v\t%v\n",i,v)
	}
	fmt.Println("---------------------------------------------------")
	for s, t := range b {
		fmt.Printf("%v\t%v\n", s, t)
	}

}
