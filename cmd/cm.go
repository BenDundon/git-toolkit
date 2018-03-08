// Copyright © 2018 mritd <mritd1234@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
	"regexp"
	"github.com/mritd/gitflow-toolkit/pkg/consts"
	"io/ioutil"
	"github.com/mritd/gitflow-toolkit/pkg/util"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "git commit-msg hook",
	Long: `
该命令接受一个 git commit message 的文件路径，并检查其格式是否符合 Angular 社区规范`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("check commit message style failed")
			os.Exit(1)
		}

		f,err := os.Open(args[1])
		util.CheckAndExit(err)

		b,err:=ioutil.ReadAll(f)
		util.CheckAndExit(err)

		reg := regexp.MustCompile(consts.CommitMessagePattern)
		if !reg.MatchString(string(b)) {
			fmt.Println("check commit message style failed")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cmCmd)
}
