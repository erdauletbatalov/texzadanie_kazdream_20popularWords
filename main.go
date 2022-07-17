package main

import (
	"kazdream/functions/btree"
	"kazdream/functions/validation"

	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	text, err := ioutil.ReadFile(args[0])
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
	}

	if validation.CheckValid(text) {
		return
	}

	// Объявляем наше Бинарное дерево,
	// в котором будем хранить все найденные слова
	// и их количество
	dictionary := &btree.TreeNode{}

	btree.FillBTree(dictionary, text)

	// На данном этапе у нас имеется dictionary структура,
	// которая по сути является корнем бинарного дерева.
	// Каждый элемент в дереве - Нода, которая хранит слово,
	// количество этого слова в тексте и три указателя: Родитель, левый и правый элемент.
	// Вставка элементов в дереве dictionary производится по принципу лексикографической проверки слов.

	// Объявляем наше Бинарное дерево,
	// в котором будем хранить слова
	// в порядке возрастания их количиства.
	// Возрастание по дереву -> InOrder.
	// https://kalkicode.com/inorder-tree-traversal картинка, как это выглядит.
	frequencyTree := &btree.TreeNode{}

	btree.SortBTreeByFrequency(dictionary, frequencyTree)

	btree.Print20ElementsInOrderBTree(frequencyTree)
}
