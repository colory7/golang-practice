package trie_tree_demo

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.AddWord([]byte("abc"))
	trie.AddWord([]byte("abcd"))
	trie.AddWord([]byte("bcd"))
	trie.AddWord([]byte("efg"))
	trie.AddWord([]byte("hi"))
	trie.AddWord([]byte("b"))
	PrintTree(trie.childNodes)

	fmt.Println(trie.SearchNode([]byte("abcd")))
	fmt.Println(trie.SearchNode([]byte("abc")))
	fmt.Println(trie.SearchNode([]byte("b")))
	fmt.Println(trie.SearchNode([]byte("hi")))
	fmt.Println(trie.SearchNode([]byte("efg")))
	fmt.Println(trie.SearchNode([]byte("china")))
	fmt.Println(trie.SearchNode([]byte("ab")))
	fmt.Println(trie.SearchNode([]byte("")))
	trie.DelWord([]byte("ab")) //应该没有影响
	fmt.Println("==================")
	fmt.Println(trie.SearchNode([]byte("abcd")))
	fmt.Println(trie.SearchNode([]byte("abc")))
	fmt.Println(trie.SearchNode([]byte("abcd")))
	fmt.Println(trie.SearchNode([]byte("ab")))
	trie.DelWord([]byte("abc")) //导致abc会找不到
	fmt.Println("===================")
	fmt.Println(trie.SearchNode([]byte("abcd")))
	fmt.Println(trie.SearchNode([]byte("abc")))
	fmt.Println(trie.SearchNode([]byte("abcd")))
	fmt.Println(trie.SearchNode([]byte("ab")))
}

func TestTrie2(t *testing.T) {
	trie := NewTrie()
	trie.AddWord([]byte("9"))
	trie.AddWord([]byte("0"))
	trie.AddWord([]byte(" "))
	trie.AddWord([]byte("("))
	trie.AddWord([]byte("EEEE"))
	trie.AddWord([]byte("E"))
	PrintTree(trie.childNodes)

	fmt.Println(trie.SearchNode([]byte("abcd")))
	fmt.Println(trie.SearchNode([]byte("999")))
}

func TestTrie3(t *testing.T) {
	trie := NewTrie()
	trie.AddWord([]byte("EEEE"))
	trie.AddWord([]byte("EE"))
	PrintTree(trie.childNodes)

	fmt.Println(trie.SearchNode([]byte("EEEE")))
}

func TestTrie4(t *testing.T) {
	trie := NewTrie()
	trie.AddWord([]byte("EEEE"))
	trie.AddWord([]byte("EE"))

	fmt.Println(trie.SearchNode([]byte("E")))
	fmt.Println(trie.SearchNode([]byte("EE")))
}
