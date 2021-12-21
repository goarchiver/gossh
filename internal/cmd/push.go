/*
Copyright © 2021 windvalley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/windvalley/gossh/internal/pkg/sshtask"
	"github.com/windvalley/gossh/pkg/util"
)

var (
	files          []string
	fileDstPath    string
	allowOverwrite bool
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push local files to remote hosts",
	Long: `
Push local files to remote hosts`,
	Example: `
  # Push a local file to host1:/tmp/.
  $ gossh push host1 -f ./foo.txt

  # Specify dest dir by '-d' flag.
  $ gossh push host1 -f ./foo.txt -d /home/user

  # Push local files to remote hosts. 
  $ gossh push host1 -f ./foo.txt -f ./bar.txt 
    or
  $ gossh push host1 -f ./foo.txt,./bar.txt`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if errs := config.Validate(); len(errs) != 0 {
			util.CheckErr(errs)
		}

		if len(files) != 0 {
			for _, file := range files {
				if !util.FileExists(file) {
					util.CheckErr(fmt.Sprintf("file '%s' not found", file))
				}
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		task := sshtask.NewTask(sshtask.PushTask, config)

		task.SetHosts(args)
		task.SetCopyfiles(files)
		task.SetFileOptions(fileDstPath, allowOverwrite)

		task.Start()
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	pushCmd.Flags().StringSliceVarP(&files, "files", "f", nil,
		"files to be copied to remote hosts",
	)
	if err := pushCmd.MarkFlagRequired("files"); err != nil {
		util.CheckErr(err)
	}

	pushCmd.Flags().StringVarP(&fileDstPath, "dest-path", "d", "/tmp",
		"path of remote hosts where files will be copied to",
	)

	pushCmd.Flags().BoolVarP(
		&allowOverwrite,
		"force",
		"F",
		false,
		"allow overwrite files if they already exist on remote hosts",
	)
}
