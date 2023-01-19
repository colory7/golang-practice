package trie_tree_demo

import "fmt"

type Trie struct {
	valid      bool                //判断是否是有效节点
	value      interface{}         //节点的值
	childMap   map[interface{}]int //子节点下是否有这个节点的map记录
	childNodes []*Trie             //包含的所有节点
}

func NewTrie() *Trie {
	t := &Trie{}
	return t
}

// 新增数据时间复杂度是O(L),L是新增字符串的长度
func (this *Trie) AddWord(word []byte) {
	if len(word) < 1 {
		return
	}
	length := len(word)
	if this.childNodes == nil {
		//不存在子级，需要重新创建
		this.childNodes = append(this.childNodes, &Trie{
			valid:      length == 1,
			value:      word[0],
			childNodes: nil,
		})
		this.childMap = make(map[interface{}]int)
		this.childMap[word[0]] = 1
		this.childNodes[0].AddWord(word[1:])
	} else {
		index := this.childMap[word[0]]
		if index > 0 {
			//子级里面以及有这个值了
			if len(word) == 1 && this.childNodes[index-1].valid == false {
				//已经存在这样的子级，且此值是有效值，当前节点无效，需要进行修改
				this.childNodes[index-1].valid = true
				return
			} else {
				if length <= 1 {
					return
				}
				//已经存在这样的子级，且此值是有效值，当前节点也有效，则重复，进行下一步
				this.childNodes[index-1].AddWord(word[1:])
			}
		} else {
			//不存在这样的子级
			this.childNodes = append(this.childNodes, &Trie{
				valid:      length == 1,
				value:      word[0],
				childNodes: nil,
			})
			this.childMap[word[0]] = len(this.childMap) + 1
			this.childNodes[len(this.childNodes)-1].AddWord(word[1:])
		}
	}
}

// 删除数据
func (this *Trie) DelWord(word []byte) {
	//不存在则不删除
	if !this.SearchNode(word) {
		return
	}
	length := len(word)
	step := 1
	prefix := this
	index := this.childMap[word[step-1]]
	node := this.childNodes[index-1]
	for step != length {
		step++
		index = node.childMap[word[step-1]]
		prefix = node
		node = node.childNodes[index-1]
	}
	//找到了最后一个word节点
	if len(node.childNodes) != 0 {
		//word节点的后面还有值,只是这个节点不生效
		node.valid = false
	} else {
		//word节点后面没有值了,删除前面节点与后续节点所有关系，依靠golanggc进行回收
		delindex := prefix.childMap[node.value]
		prefix.childMap[node.value] = 0
		prefix.childNodes = append(prefix.childNodes[:delindex], prefix.childNodes[delindex+1:]...)
	}
}

// 按层分布
func PrintTree(tries []*Trie) {
	if len(tries) < 1 {
		return
	}
	//按照树打印，一次打印一个子树
	var inertries []*Trie
	for i := 0; i < len(tries); i++ {
		fmt.Print(string(tries[i].value.(uint8)) + " ")
		inertries = append(inertries, tries[i].childNodes...)
	}
	fmt.Print("\n")
	PrintTree(inertries)
}

// 查询，时间复杂度是O(L),L是查询字符串的长度
func (this *Trie) SearchNode(word []byte) bool {
	if len(word) < 1 {
		return false
	}
	index := this.childMap[word[0]]
	if index == 0 {
		return false
	} else {
		if len(word) == 1 {
			return this.childNodes[index-1].valid
		}
		return this.childNodes[index-1].SearchNode(word[1:])
	}
}
