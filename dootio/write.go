package dootio

import (
	"errors"
	"fmt"
	"github.com/dreznel/doot/config"
	"github.com/dreznel/doot/skeletons"
	"github.com/spf13/afero"
	"path/filepath"
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
		err := DootBone(fs, config, path,  bone)
		if err != nil {
			return err
		}
	}
	return nil
}

//TODO: deal with duplicates
func DootBone(fs afero.Fs, config config.Config, parentDir string, bone skeletons.Bone) error {
	fmt.Println("Dooting " + filepath.Join(parentDir, bone.Name))
	if bone.Type == skeletons.Directory {
		err := fs.Mkdir(filepath.Join(parentDir, bone.Name), config.FilePermissions)
		if err != nil {
			return err
		}
		for _, subBone := range bone.SubBones {
			err = DootBone(fs, config, filepath.Join(parentDir, bone.Name), subBone)
		}
		if err != nil {
			return err
		}
	} else if bone.Type == skeletons.File {
		_, err := fs.Create(filepath.Join(parentDir, bone.Name))
		if err != nil {
			return err
		}
	} else {
		panic("invalid bone type.")
	}

	return nil
}