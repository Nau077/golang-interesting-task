package go
import(
	"math"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func MathPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	mathLeftSum := MathPathSum(root.left)
	mathRightSum := MathPathSum(root.right)

	return Max(mathLeftSum, mathRightSum) + root.val
}