/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type Trie struct {
	end bool
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
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

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	temp := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if temp.children[idx] == nil {
			return false
		}
		
		temp = temp.children[idx]
	}

	return temp.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	temp := this
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if temp.children[idx] == nil {
			return false
		}
		temp = temp.children[idx]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

