package skeletons

type Skeleton struct {
	Name string
	Bones []Bone
}


type Bone struct {
	Name     string
	Type     BoneType
	SubBones []Bone
}

type BoneType int

const (
	File BoneType = iota
	Directory BoneType = iota
)