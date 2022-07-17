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
	text, err := ioutil.ReadFile(args[0])
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
	}
	// Объявляем наше Бинарное дерево, в котором будем хранить все найденные слова и их количество
	dictionary := &TreeNode{}
	temp := 0
	for i := 0; i < len(text); i++ {
		// Если видим начало слова
		if isAscii(text[i]) {
			start := i
			for j := i; j < len(text); j++ {
				// Итерируемся до тех пор, пока не найдем конец слова
				if !isAscii(text[j]) {
					// Захардкодил замую первую ноду, корневую
					if temp == 0 {
						root.Word = ToLower(text[start:j])
						root.Count++
						temp++
						break
					}
					node := BTreeSearchWord(dictionary, ToLower(text[start:j]))
					if node == nil {
						_ = BTreeInsertWordByOrder(dictionary, ToLower(text[start:j]))
					}
					i = j
					break
				}
				// Самый последний элемент в тексте
				if isAscii(text[j]) && j == len(text)-1 {
					node := BTreeSearchWord(dictionary, ToLower(text[start:]))
					if node == nil {
						_ = BTreeInsertWordByOrder(dictionary, ToLower(text[start:]))
					}
					i = j
					break
				}
			}
		}
	}

	// На данном этапе у нас имеется dictionary структура,
	// которая по сути является корнем бинарного дерева.
	// Каждый элемент в дереве - Нода, которая хранит слово,
	// количество этого слова в тексте и три указателя: Родитель, левый и правый элемент.
	// Вставка элементов в дереве dictionary производится по принципу лексикографической проверки слов.

	frequencyTree := &TreeNode{}
	// BtreePrint(root)
	BTreeApplyInorder(dictionary, frequencyTree, BTreeInsertWordByFrequency)

	// BtreePrint(frequencyTree)
	BTreeApplyInorder2(frequencyTree, printFirst20)
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

func BTreeInsertWordByOrder(bt *TreeNode, word []byte) *TreeNode {
	if bt == nil {
		return &TreeNode{Word: word, Count: 1}
	}

	if Equal(word, bt.Word) == -1 {
		bt.Left = BTreeInsertWordByOrder(bt.Left, word)
		bt.Left.Parent = bt
	} else if Equal(word, bt.Word) == 1 {
		bt.Right = BTreeInsertWordByOrder(bt.Right, word)
		bt.Right.Parent = bt
	}
	return bt
}

func BTreeInsertWordByFrequency(bt *TreeNode, word []byte, count int) *TreeNode {
	if bt == nil {
		return &TreeNode{Word: word, Count: count}
	}
	if count >= bt.Count {
		bt.Left = BTreeInsertWordByFrequency(bt.Left, word, count)
		bt.Left.Parent = bt
	} else if count < bt.Count {
		bt.Right = BTreeInsertWordByFrequency(bt.Right, word, count)
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
