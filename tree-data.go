package base

import "fmt"

// TreeNode 树状组织项目接口
type TreeNode interface {
	GetID() string
	GetParentID() string
	InsertChild(item interface{}) error
}

// TreeGenerator 树状组织生成
type TreeGenerator struct {
	root     TreeNode
	itemMap  map[string]TreeNode
	itemList []TreeNode
}

// NewTreeGenerator 创建树状结构生成者
func NewTreeGenerator(root TreeNode) *TreeGenerator {
	return &TreeGenerator{
		root:    root,
		itemMap: map[string]TreeNode{},
	}
}

// AddItem 添加项目
func (g *TreeGenerator) AddItem(item TreeNode) error {
	_, got := g.itemMap[item.GetID()]
	if got {
		return fmt.Errorf("ID [%s] is duplicate", item.GetID())
	}
	g.itemMap[item.GetID()] = item
	g.itemList = append(g.itemList, item)
	return nil
}

// FindNode 在树中查找节点
func (g *TreeGenerator) FindNode(id string) TreeNode {
	item, exists := g.itemMap[id]
	if !exists {
		return nil
	}
	return item
}

// Generate 生成树
func (g *TreeGenerator) Generate() error {
	for _, item := range g.itemList {
		// 定位父节点
		var parent TreeNode
		var got bool
		parent, got = g.itemMap[item.GetParentID()]
		if !got {
			// 使用 根节点 作为父节点
			parent = g.root
		}
		// 插入到父节点的 children 中
		if err := parent.InsertChild(item); err != nil {
			return err
		}
	}
	return nil
}
