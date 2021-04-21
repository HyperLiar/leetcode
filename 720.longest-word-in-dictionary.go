/*
 * @lc app=leetcode id=720 lang=golang
 *
 * [720] Longest Word in Dictionary
 */

// @lc code=start
func longestWord(words []string) string {
	trie := Trie{}
	maxCount := 0
	maxWord := ""
	for _, word := range words {
		trie.insert(word, &maxCount, &maxWord)
	}

	return maxWord
}

type Trie struct {
	end      bool
	count    int
	children [26]*Trie
}

func (this *Trie) insert(word string, maxCount *int, maxWord *string) {
	temp := this

	count := 0
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'

		if temp.children[idx] == nil {
			temp.children[idx] = &Trie{}
		}
		if temp.end {
			count++
		}

		temp = temp.children[idx]
	}

	temp.end = true
	temp.count = count
	temp.addCount(word, maxCount, maxWord)
}

func (this *Trie) addCount(word string, maxCount *int, maxWord *string) {
	this.count++
	fmt.Println(word, *maxCount, *maxWord)
	if this.count > *maxCount || (this.count == *maxCount && *maxWord > word) {
		*maxWord = word
		*maxCount = this.count
	}
	for i := 0; i < len(this.children); i++ {
		if this.children[i] != nil && this.children[i].end {
			nowWord := word + string(byte('a'+i))
			this.children[i].addCount(nowWord, maxCount, maxWord)
		}
	}
}

// @lc code=end

