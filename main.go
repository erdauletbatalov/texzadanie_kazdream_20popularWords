package main

import (
	"io/ioutil"
	"os"
)

var stopCount = 0

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}

	// text := []byte{97, 32, 231, 136, 190, 32, 97, 10, 71, 117, 116, 101, 110, 98, 101, 114, 103, 39, 83, 10, 71, 117, 116, 101, 110, 98, 101, 114, 103, 39, 115, 10, 71, 117, 116, 101, 110, 98, 101, 114, 103, 231, 136, 190, 115, 10, 71, 117, 116, 101, 110, 98, 101, 114, 103, 39, 115, 10, 71, 117, 116, 101, 110, 98, 101, 114, 103, 39, 115, 10, 71, 117, 116, 101, 110, 98, 101, 114, 103, 39, 115}
	text, err := ioutil.ReadFile(args[0])
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
	}
	fileLen := len(text)
	root := &TreeNode{}
	temp := 0
	for i := 0; i < len(text); i++ {
		if isAscii(text[i]) {
			start := i
			for j := i; j < len(text); j++ {
				if !isAscii(text[j]) {
					if temp == 0 {
						root.Word = ToLower(text[start:j])
						root.Count++
						temp++
						break
					}
					node := BTreeSearchWord(root, ToLower(text[start:j]))
					if node == nil {
						_ = BTreeInsertWord(root, ToLower(text[start:j]))
					}
					// root = BTreeInsertAndIncrementWord(root, ToLower(text[start:j]))
					// words = append(words, text[start:j])
					i = j
					break
				}
				if isAscii(text[j]) && j == fileLen-1 {
					node := BTreeSearchWord(root, ToLower(text[start:]))
					if node == nil {
						_ = BTreeInsertWord(root, ToLower(text[start:]))
					}
					// root = BTreeInsertAndIncrementWord(root, ToLower(text[start:]))
					// words = append(words, text[start:])
					i = j
					break
				}
			}
		}
	}

	root2 := &TreeNode{}
	// BtreePrint(root)
	BTreeApplyInorder(root, root2, BTreeInsertWordInfo)

	// BtreePrint(root2)
	BTreeApplyInorder2(root2, printFirst20)
}

func printFirst20(stopCount *int, count int, word []byte) bool {
	if *stopCount == 20 {
		return true
	}
	os.Stdout.Write(intToBytes(count))
	os.Stdout.Write([]byte{32})
	os.Stdout.Write(word)
	os.Stdout.Write([]byte{10})
	*stopCount++
	return false
}

func isAscii(char byte) bool {
	if ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') {
		return true
	}
	return false
}

func ToLower(word []byte) []byte {
	hasUpper := false
	for i := 0; i < len(word); i++ {
		char := word[i]
		hasUpper = ('A' <= char && char <= 'Z')
		if hasUpper {
			break
		}
	}
	if !hasUpper {
		return word
	}
	result := make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		loweredChar := word[i]
		if 'A' <= loweredChar && loweredChar <= 'Z' {
			loweredChar += 'a' - 'A'
		}
		result[i] = loweredChar
	}
	return result
}

func Equal(a, b []byte) int {
	if len(a) < len(b) {
		for i, v := range a {
			if v > b[i] {
				return 1
			} else if v < b[i] {
				return -1
			}
		}
		return -1
	}

	if len(a) > len(b) {
		for i, v := range b {
			if v < a[i] {
				return 1
			} else if v > a[i] {
				return -1
			}
		}
		return 1
	}
	if len(a) == len(b) {
		for i, v := range a {
			if v > b[i] {
				return 1
			} else if v < b[i] {
				return -1
			}
		}
		return 0
	}

	return 0
}

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Word                []byte
	Count               int
}

func BTreeInsertWord(bt *TreeNode, word []byte) *TreeNode {
	if bt == nil {
		return &TreeNode{Word: word, Count: 1}
	}

	if Equal(word, bt.Word) == -1 {
		bt.Left = BTreeInsertWord(bt.Left, word)
		bt.Left.Parent = bt
	} else if Equal(word, bt.Word) == 1 {
		bt.Right = BTreeInsertWord(bt.Right, word)
		bt.Right.Parent = bt
	}
	return bt
}

func BTreeInsertWordInfo(bt *TreeNode, word []byte, count int) *TreeNode {
	if bt == nil {
		return &TreeNode{Word: word, Count: count}
	}
	if count >= bt.Count {
		bt.Left = BTreeInsertWordInfo(bt.Left, word, count)
		bt.Left.Parent = bt
	} else if count < bt.Count {
		bt.Right = BTreeInsertWordInfo(bt.Right, word, count)
		bt.Right.Parent = bt
	}
	return bt
}

func BTreeSearchWord(root *TreeNode, word []byte) *TreeNode {
	if root == nil {
		return nil
	}

	if Equal(root.Word, word) == 0 {
		root.Count++
		return root
	}

	if Equal(word, root.Word) == -1 {
		return BTreeSearchWord(root.Left, word)
	}

	return BTreeSearchWord(root.Right, word)
}

func BTreeApplyInorder(root *TreeNode, bt *TreeNode, f func(bt *TreeNode, word []byte, count int) *TreeNode) {
	if root != nil {
		BTreeApplyInorder(root.Left, bt, f)
		_ = f(bt, root.Word, root.Count)
		BTreeApplyInorder(root.Right, bt, f)
	}
}

func BTreeApplyInorder2(root *TreeNode, f func(stopCount *int, count int, word []byte) bool) {
	if root != nil {
		BTreeApplyInorder2(root.Left, f)
		if f(&stopCount, root.Count, root.Word) {
			return
		}
		BTreeApplyInorder2(root.Right, f)
	}
}

func intToBytes(x int) []byte {
	arr := []byte{}
	for x != 0 {
		arr = append(arr, byte(x%10+48))
		x /= 10
	}
	return reverse(arr)
}

func reverse(a []byte) []byte {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
