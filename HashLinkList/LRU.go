package HashLinkList

var head, end *Node
var limit int



type Node struct {
	Key int
	Value int
	PreNode *Node
	NextNode *Node
}


type LRUCache struct {
	Limit  int
	HashMap map[int]*Node
}


func GetLRUCache(limit int) *LRUCache{
	lruCache := LRUCache{Limit: limit}
	lruCache.HashMap = make(map[int]*Node, limit)

	return &lruCache
}


func (l *LRUCache)get(key int) int{
	if v, ok := l.HashMap[key]; ok{
		l.refreshNode(v)
		return v.Value
	}else {
		return -1
	}
}


func (l *LRUCache) refreshNode(node *Node){
	if node == end {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}


func (l *LRUCache) addNode(node *Node){
	if end != nil{
		end.NextNode = node
		node.PreNode = end
		node.NextNode = nil
	}
	end = node
	if head == nil{
		head = node
	}
}


func (l *LRUCache) removeNode(node *Node) int{
	if node == end{
		end = end.PreNode
	}else if node == head{
		head = head.NextNode
	}else {
		node.PreNode.NextNode = node.NextNode
		node.NextNode.PreNode = node.PreNode
	}
	return node.Key
}