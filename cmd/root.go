package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "bj",
	Short: "BaekJoon-CLI",
	Long: `백준 문제풀이 및 파일 관리를 도와줍니다

https://github.com/Changemin/boj-cli 를 참고하세요 👨‍🏫`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}
