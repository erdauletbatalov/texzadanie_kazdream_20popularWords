package btree

import (
	"kazdream/functions/bytes"

	"os"
)

var stopCount = 0

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Word                []byte
	Count               int
}

// Создает дерево, в котором слова хранятся в алфавитном порядке.
// Сделано это для уменьшения времени поиска уже имеющихся слов в дереве.
func FillBTree(dictionary *TreeNode, text []byte) {
	temp := 0
	for i := 0; i < len(text); i++ {
		// Если видим начало слова
		if bytes.IsAscii(text[i]) {
			start := i
			for j := i; j < len(text); j++ {
				// Итерируемся до тех пор, пока не найдем конец слова
				if !bytes.IsAscii(text[j]) {
					// Захардкодил замую первую ноду, корневую
					if temp == 0 {
						dictionary.Word = bytes.ToLower(text[start:j])
						dictionary.Count++
						temp++
						break
					}
					node := BTreeSearchWord(dictionary, bytes.ToLower(text[start:j]))
					if node == nil {
						_ = BTreeInsertWordByOrder(dictionary, bytes.ToLower(text[start:j]))
					}
					i = j
					break
				}
				// Самый последний элемент в тексте
				if bytes.IsAscii(text[j]) && j == len(text)-1 {
					node := BTreeSearchWord(dictionary, bytes.ToLower(text[start:]))
					if node == nil {
						_ = BTreeInsertWordByOrder(dictionary, bytes.ToLower(text[start:]))
					}
					i = j
					break
				}
			}
		}
	}
}

// Переносим с дерева from в дерево to слова, сортирую по
// полю Count. В данном случае я вытаскиваю элементы с дерева
// from в порядке InOrder.
func SortBTreeByFrequency(from *TreeNode, to *TreeNode) {
	BTreeApplyInorder(from, to, BTreeInsertWordByFrequency)
}

// Выводит в стандартный вывод первые 20 по частоте слова.
//
func Print20ElementsInOrderBTree(btree *TreeNode) {
	BTreeApplyInorderFrequency(btree, printFirst20)
}

func BTreeInsertWordByOrder(bt *TreeNode, word []byte) *TreeNode {
	if bt == nil {
		return &TreeNode{Word: word, Count: 1}
	}

	if bytes.Equal(word, bt.Word) == -1 {
		bt.Left = BTreeInsertWordByOrder(bt.Left, word)
		bt.Left.Parent = bt
	} else if bytes.Equal(word, bt.Word) == 1 {
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

	if bytes.Equal(root.Word, word) == 0 {
		root.Count++
		return root
	}

	if bytes.Equal(word, root.Word) == -1 {
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

func BTreeApplyInorderFrequency(root *TreeNode, f func(stopCount *int, count int, word []byte) bool) {
	if root != nil {
		BTreeApplyInorderFrequency(root.Left, f)
		if f(&stopCount, root.Count, root.Word) {
			return
		}
		BTreeApplyInorderFrequency(root.Right, f)
	}
}

func printFirst20(stopCount *int, count int, word []byte) bool {
	if *stopCount == 20 {
		return true
	}
	os.Stdout.Write(bytes.IntToBytes(count))
	os.Stdout.Write([]byte{32})
	os.Stdout.Write(word)
	os.Stdout.Write([]byte{10})
	*stopCount++
	return false
}
