package dootio

import (
	"errors"
	"github.com/dreznel/doot/config"
	"github.com/dreznel/doot/skeletons"
	"github.com/spf13/afero"
)

func DootSkeleton(fs afero.Fs, config config.Config, path string, skeleton skeletons.Skeleton) error {

	directoryIsEmpty, err := afero.IsEmpty(fs, path)

	if err != nil {
		return err
	}

	if !directoryIsEmpty {
		return errors.New("cannot doot in a non-empty directory")
	}

	for _, bone := range skeleton.Bones {
		DootBone(fs, config, bone)
	}
	return nil
}

//TODO: deal with duplicates
func DootBone(fs afero.Fs, config config.Config, bone skeletons.Bone) error {
	if bone.Type == skeletons.Directory {
		err := fs.Mkdir(bone.Name, config.FilePermissions)
		if err != nil {
			return err
		}
		for _, subBone := range bone.SubBones {
			err = DootBone(fs, config, subBone)
		}
		if err != nil {
			return err
		}
	} else if bone.Type == skeletons.File {
		_, err := fs.Create(bone.Name)
		if err != nil {
			return err
		}
	} else {
		panic("invalid bone type.")
	}

	return nil
}