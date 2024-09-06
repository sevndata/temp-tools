package token_tools

import (
	"encoding/hex"
	"fmt"
	"hash"
)

// MerkleTree represents a simple Merkle Tree
type MerkleTree struct {
	Root *MerkleNode
}

// MerkleNode represents a node in the Merkle Tree
type MerkleNode struct {
	Hash  string
	Left  *MerkleNode
	Right *MerkleNode
}

// NewMerkleNode creates a new Merkle Node
func NewMerkleNode(left, right *MerkleNode, hashInstance hash.Hash) *MerkleNode {
	node := &MerkleNode{}
	if left == nil && right == nil {
		// Leaf node
		return node
	}

	leftHash := ""
	rightHash := ""

	if left != nil {
		leftHash = left.Hash
	}
	if right != nil {
		rightHash = right.Hash
	}

	hashInstance.Reset() // Ensure the hash instance is reset before use
	hashInstance.Write([]byte(leftHash + rightHash))
	node.Hash = hex.EncodeToString(hashInstance.Sum(nil))
	node.Left = left
	node.Right = right
	return node
}

// NewMerkleTree creates a new Merkle Tree from leaf nodes
func NewMerkleTree(leaves []*MerkleNode, hashInstance hash.Hash) *MerkleTree {
	if len(leaves) == 0 {
		return nil
	}

	nodes := leaves
	for len(nodes) > 1 {
		var newLevel []*MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			var right *MerkleNode
			if i+1 < len(nodes) {
				right = nodes[i+1]
			}
			newLevel = append(newLevel, NewMerkleNode(left, right, hashInstance))
		}
		nodes = newLevel
	}

	return &MerkleTree{Root: nodes[0]}
}

// CreateLeafNode creates a new leaf node with the given data
func CreateLeafNode(data string, hashInstance hash.Hash) *MerkleNode {
	hashInstance.Reset() // Ensure the hash instance is reset before use
	hashInstance.Write([]byte(data))
	return &MerkleNode{Hash: hex.EncodeToString(hashInstance.Sum(nil))}
}

// PrintTree prints the Merkle Tree for debugging purposes
func (node *MerkleNode) PrintTree(level int) {
	if node == nil {
		return
	}
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Println(node.Hash)
	if node.Left != nil || node.Right != nil {
		node.Left.PrintTree(level + 1)
		node.Right.PrintTree(level + 1)
	}
}

// GetRootHash returns the hash of the root node
func (tree *MerkleTree) GetRootHash() string {
	if tree.Root == nil {
		return ""
	}
	return tree.Root.Hash
}

// GetTreeHeight returns the height of the tree
func (tree *MerkleTree) GetTreeHeight() int {
	return getTreeHeight(tree.Root)
}

// Helper function to calculate tree height
func getTreeHeight(node *MerkleNode) int {
	if node == nil {
		return 0
	}
	leftHeight := getTreeHeight(node.Left)
	rightHeight := getTreeHeight(node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// GetMerkleTree returns a Merkle Tree from the given data and hash instance
func GetMerkleTree(data []string, hashInstance hash.Hash) *MerkleTree {
	var leaves []*MerkleNode
	for _, d := range data {
		leaves = append(leaves, CreateLeafNode(d, hashInstance))
	}
	return NewMerkleTree(leaves, hashInstance)
}
