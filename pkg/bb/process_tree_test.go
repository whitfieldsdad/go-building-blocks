package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestGetProcessTree(t *testing.T) {
	_, err := GetProcessTree()
	assert.Nil(t, err, "Failed to get process tree")
}

func TestGetAncestorPidsOneBranch(t *testing.T) {

	// 1 -> 2 -> 3 -> 4
	tree := NewProcessTree()
	tree.AddProcess(1, 2)
	tree.AddProcess(2, 3)
	tree.AddProcess(3, 4)

	expected := []int{1, 2}
	result := tree.GetAncestorPids(3)
	slices.Sort(result)
	assert.Equal(t, expected, result, "Failed to identify ancestors")
}

func TestGetAncestorPidsMultipleBranches(t *testing.T) {

	// 1 -> 2 -> 3 -> 4
	// 1 -> 5 -> 6 -> 7
	//      5 -> 8 -> 9
	tree := NewProcessTree()
	tree.AddProcess(1, 2)
	tree.AddProcess(2, 3)
	tree.AddProcess(3, 4)
	tree.AddProcess(1, 5)
	tree.AddProcess(5, 6)
	tree.AddProcess(6, 7)
	tree.AddProcess(5, 8)
	tree.AddProcess(8, 9)

	expected := []int{1, 5, 8}
	result := tree.GetAncestorPids(9)
	slices.Sort(result)
	assert.Equal(t, expected, result, "Failed to identify ancestors")
}

func TestGetDescendantPids(t *testing.T) {
	// 1 -> 2 -> 3 -> 4
	// 1 -> 5 -> 6 -> 7
	//      5 -> 8 -> 9
	tree := NewProcessTree()
	tree.AddProcess(1, 2)
	tree.AddProcess(2, 3)
	tree.AddProcess(3, 4)
	tree.AddProcess(1, 5)
	tree.AddProcess(5, 6)
	tree.AddProcess(6, 7)
	tree.AddProcess(5, 8)
	tree.AddProcess(8, 9)

	expected := []int{6, 7, 8, 9}
	result := tree.GetDescendantPids(5)
	slices.Sort(result)
	assert.Equal(t, expected, result, "GetDescendantPids() should return the correct descendant pids")
}

func TestGetSiblingPids(t *testing.T) {
	// 1 -> 2 -> 3 -> 4
	// 1 -> 5 -> 6 -> 7
	//      5 -> 8 -> 9
	tree := NewProcessTree()
	tree.AddProcess(1, 2)
	tree.AddProcess(2, 3)
	tree.AddProcess(3, 4)
	tree.AddProcess(1, 5)
	tree.AddProcess(5, 6)
	tree.AddProcess(6, 7)
	tree.AddProcess(5, 8)
	tree.AddProcess(8, 9)

	expected := []int{8}
	result := tree.GetSiblingPids(6)
	slices.Sort(result)
	assert.Equal(t, expected, result, "GetSiblingPids() should return the correct sibling pids")
}

func TestGetParentPid(t *testing.T) {
	// 1 -> 2 -> 3
	tree := NewProcessTree()
	tree.AddProcess(1, 2)
	tree.AddProcess(2, 3)

	expected := 2
	result, ok := tree.GetParentPid(3)
	assert.True(t, ok, "GetParentPid() should return true if the pid exists")
	assert.Equal(t, expected, result, "GetParentPid() should return the correct parent pid")
}

func TestGetChildPids(t *testing.T) {
	// 1 -> 2 -> 3 -> 4
	//      2 -> 5 -> 6
	tree := NewProcessTree()
	tree.AddProcess(1, 2)
	tree.AddProcess(2, 3)
	tree.AddProcess(3, 4)
	tree.AddProcess(2, 5)
	tree.AddProcess(5, 6)

	expected := []int{3, 5}
	result := tree.GetChildPids(2)
	slices.Sort(result)
	assert.Equal(t, expected, result, "GetChildPids() should return the correct child pids")
}
