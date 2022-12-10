/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var sigCh = make(chan os.Signal, 2)
var shutdownCh = make(chan struct{})

// 下载
func download(url string) {
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	cmdArgs := []string{"-i", url, "-vcodec", "copy", "-acodec", "copy", "-absf", "aac_adtstoasc", "output.mp4"}
	cmd := exec.Command("ffmpeg", cmdArgs...)
	go func(c *exec.Cmd) {
		log.Println("开始下载")
		err := c.Run()
		if nil != err {
			log.Printf("录制错误, pid: %d, error: %v", c.Process.Pid, err)
			shutdownCh <- struct{}{}
			return
		}
	}(cmd)
	select {
	case <-sigCh:
		log.Printf("正常退出, pid: %d", cmd.Process.Pid)
	case <-shutdownCh:
		log.Printf("异常退出, pid: %d", cmd.Process.Pid)
	}
}

// 上传
func upload(filename string) {
	f, err := os.Open(filename)
	if nil != err {
		log.Printf("读取文件%s出错", filename)
		return
	}
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(f)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}
	log.Printf("added %s", cid)
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "下载文件",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if nil != err {
			log.Printf("error: %s", err.Error())
			return
		}
		download(url)
	},
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传文件",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fileName, err := cmd.Flags().GetString("file")
		if nil != err {
			log.Printf("error: %s", err.Error())
			return
		}
		upload(fileName)
	},
}

func init() {
	downloadCmd.Flags().StringP("url", "u", "", "直播间地址")
	uploadCmd.Flags().StringP("file", "f", "", "文件")
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(uploadCmd)
}
