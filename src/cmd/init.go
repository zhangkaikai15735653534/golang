package cmd

import (
	"GoMvc/src/cmd/api"
	"GoMvc/src/router"
	logger "GoMvc/src/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

/**
* @Author: kiki.znag
* @Date: 2021/12/12 3:54 PM
 */

var rootCmd = &cobra.Command{
	Use:               "kiki",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `kiki`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("init error")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `欢迎使用 80K`
		logger.Infof("%s\n", usageStr)
	},
}

func init() {
	api.Usage() //项目启动
	api.Setup() //加载配置项

	//todo: 先放弃这种方式  实在是不会写  学会了再来
	//rootCmd.AddCommand(api.StartCmd)
	//注册系统路由
	r := router.InitRouter()
	//
	//监听端口号
	r.Run(":9999")

}

/**
项目初始化时执行
*/
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
