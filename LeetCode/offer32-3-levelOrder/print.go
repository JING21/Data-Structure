package offer32_3_levelOrder

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}


func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	//利用队列
	queue := []*TreeNode{root}
	var level int = 0
	for len(queue) != 0 {
		//利用临时队列，暂存每个节点的左右子树
		temp := []*TreeNode{}
		//只需考虑在同一层上追加元素
		res = append(res,[]int{})
		//遍历队列，追加队列元素到切片同一层，追加队列元素左右子树到临时队列
		if level %2 ==0{
			//先打印队列元素
			for _,v := range queue {
				res[level] = append(res[level],v.Val)

			}
			//逆序加入队列
			for i:=len(queue)-1; i>=0;i--{
				if queue[i].Right != nil{
					temp = append(temp,queue[i].Right)
				}

				if queue[i].Left != nil {
					temp = append(temp,queue[i].Left)
				}
			}

		}else if level %2 ==1{

			for _,v := range queue {
				res[level] = append(res[level],v.Val)

			}


			for i:=len(queue)-1; i>=0;i--{

				if queue[i].Left != nil{
					temp = append(temp,queue[i].Left)
				}

				if queue[i].Right != nil {
					temp = append(temp,queue[i].Right)
				}

			}

		}

		//层级加1，队列重新复制为队列的左右子树集
		level++
		//队列赋值
		queue = temp
	}
	return res
}