package main

import "strings"

type TrieNode struct {
	flag     bool
	children map[rune]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

func (root *TrieNode) Insert(word string) {
	cur := root
	for _, char := range word {
		if _, ok := cur.children[char]; !ok {
			cur.children[char] = NewTrieNode()
		}
		cur = cur.children[char]
	}
	cur.flag = true
}

func (root *TrieNode) Replace(word string) string {
	cur := root
	var sb strings.Builder
	for _, char := range word {
		if child, ok := cur.children[char]; ok {
			sb.WriteRune(char)
			if child.flag {
				return sb.String()
			}
			cur = child
		} else {
			break
		}
	}
	return word
}

func replaceWords(dictionary []string, sentence string) string {
	root := NewTrieNode()
	for _, word := range dictionary {
		root.Insert(word)
	}

	words := strings.Split(sentence, " ")

	var res []string
	for _, word := range words {
		res = append(res, root.Replace(word))
	}

	return strings.Join(res, " ")
}
