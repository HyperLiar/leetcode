/*
 * @lc app=leetcode id=720 lang=golang
 *
 * [720] Longest Word in Dictionary
 */

// @lc code=start
func longestWord(words []string) string {
	trie := Trie{end: true}
	maxCount := 0
	maxWord := ""
	for _, word := range words {
		trie.insert(word)
	}

	trie.search(0, "", &maxCount, &maxWord)

	return maxWord
}

type Trie struct {
	end      bool
	children [26]*Trie
}

func (this *Trie) insert(word string) {
	temp := this

	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'

		if temp.children[idx] == nil {
			temp.children[idx] = &Trie{}
		}

		temp = temp.children[idx]
	}

	temp.end = true
}

func (this *Trie) search(count int, word string, maxCount *int, maxWord *string) {
	if !this.end {
		return
	}

	if count > *maxCount || (count == *maxCount && word < *maxWord) {
		*maxCount = count
		*maxWord = word
	}
	for i := 0; i < len(this.children); i++ {
		if this.children[i] == nil {
			continue
		}

		newWord := word + string(byte('a'+i))
		
		this.children[i].search(count+1, newWord, maxCount, maxWord)
	}
}

// @lc code=end

