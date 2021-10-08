package offer32_3_levelOrder

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

func levelOrder(root *TreeNode) [][] int{
	var result [][]int
	if root == nil{
		return  nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	level := 0

	for len(queue) != 0 {
		tmp := make([]*TreeNode, 0)
		result = append(result,[]int{})
		if level % 2 == 0{
			for _,v := range queue{
				result[level]= append(result[level], v.Val)
			}
			for i := len(queue)-1 ; i>=0; i--{
				if queue[i].Right != nil{
					tmp = append(tmp,queue[i].Right)
				}
				if queue[i].Left != nil{
					tmp = append(tmp,queue[i].Left)
				}
			}
		} else if level %2 == 1{
			for _, v := range queue{
				result[level]= append(result[level], v.Val)
				if v.Left != nil{
					tmp = append(tmp,v.Left)
				}
				if v.Right != nil{
					tmp = append(tmp,v.Right)
				}
			}
		}
		level++
		queue = tmp
	}
	return result
}
