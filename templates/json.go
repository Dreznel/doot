package templates

type TemplateNode struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Contents []TemplateNode `json:"contents"`
	CopyFrom string `json:"copy_from"`
}

type NodeType int

const nodeType {
	directory = iota
	file
	root

}