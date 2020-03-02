package main

//位向量
//go语言的集合 通常使用 map[T]bool 实现

//集合的操作 多是  求并集 和交集  位向量是个理想的数据结构

//IntSet 是一个 包含非负整数的集合
//零值代表空集合
type IntSet struct {
	words []uint64
}

//是否存在非负数x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//add 添加非负数到集合中
func (s *IntSet) Add(x int) {

}

//对s t做并集 并保存在s中
func (s *IntSet) UnionWith(t *IntSet) {

}
