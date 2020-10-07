/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/dreznel/doot/config"
	"github.com/dreznel/doot/dootio"
	"github.com/dreznel/doot/skeletons"
	"github.com/dreznel/doot/templates"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

// RCmd represents the R command
var RCmd = &cobra.Command{
	Use:   "R",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		cfg := config.Config{FilePermissions:os.FileMode(0777)}
		skeleton := createRSkeleton(false, false, "")
		dootio.DootSkeleton(fs, cfg, "./", skeleton )
		fmt.Println("R called")
	},
}

func unpackBones(nodes []templates.TemplateNode) []skeletons.Bone {


	if len(nodes) == 0 {
		return []skeletons.Bone{}
	}

	var bones []skeletons.Bone
	for _, item := range nodes {
		bones = append(bones, skeletons.Bone{
			Name: item.Name,
			Type: convertTypes(item.Type),
			SubBones: unpackBones(item.Contents),
		})
	}



	return bones
}

//TODO: This will be deleted but technically we need to add more types here.
func convertTypes(nodeType string) skeletons.BoneType {
	switch strings.ToLower(nodeType) {
		case "file":
			return skeletons.File
		case "directory":
			return skeletons.Directory
		default:
			return 0
	}
}

func createRSkeleton(usePkgr, useRenv bool, templateUrl string) skeletons.Skeleton {

	//dootTemplate := templates.getDootTemplate

	//home, err := homedir.Dir()
	//if err != nil {
	//	panic("error")
	//}

	//path := filepath.Join("../../..", "doot-templates", "R", "doot.json")
	bytes, err := ioutil.ReadFile("/Users/johncarlos/openspace/doot-templates/R/doot.json")
	if err != nil {
		fmt.Print(err)
		panic("error")
	}

	var rootNode templates.TemplateNode
	err = json.Unmarshal([]byte(bytes), &rootNode)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v\n", rootNode)




	//var boneSlice = []skeletons.Bone
	//for _, item := range rootNode.Contents {
	//	boneSlice = append(boneSlice, skeletons.Bone{
	//		Name: item.Name,
	//		Type: item.Type,
	//		SubBones: item.
	//	})
	//}

	skel := skeletons.Skeleton{
		Name: rootNode.Name,
		Bones: unpackBones(rootNode.Contents),

	}


	fmt.Println("Printing results")
	fmt.Printf("%v\n", skel)
	fmt.Println("Results printed.")
	//skeleton := skeletons.Skeleton{
	//
	//}



	//if usePkgr {
	//
	//}
	//
	//if useRenv {
	//
	//}


	return skel
}


func init() {
	rootCmd.AddCommand(RCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// RCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// RCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
