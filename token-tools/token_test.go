package token_tools

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
	"testing"
)

func TestGetMerkleTree(t *testing.T) {

	data := []string{"a", "b", "c", "d", "e", "f", "g"}
	md5Tree := GetMerkleTree(data, md5.New())
	fmt.Println("Merkle Tree with MD5:")
	md5Tree.Root.PrintTree(0)
	fmt.Printf("Root Hash: %s\n", md5Tree.GetRootHash())
	fmt.Printf("Tree Height: %d\n", md5Tree.GetTreeHeight())

	sha512Tree := GetMerkleTree(data, sha512.New())
	fmt.Println("Merkle Tree with SHA-512:")
	sha512Tree.Root.PrintTree(0)
	fmt.Printf("Root Hash: %s\n", sha512Tree.GetRootHash())
	fmt.Printf("Tree Height: %d\n", sha512Tree.GetTreeHeight())

}
