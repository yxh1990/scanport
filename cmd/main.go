package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/pterm/pterm"
	_ "yxh/scan/pkg/log"
	"yxh/scan/pkg/common"
	"yxh/scan/run"
	_"yxh/scan/utils"
	"time"
	_ "strings"
	"yxh/scan/finger"
	_ "net/http"
	_ "net/http/pprof"
)

func init() {
	newHeader := pterm.HeaderPrinter{
		TextStyle: pterm.NewStyle(pterm.BgCyan),
		BackgroundStyle: pterm.NewStyle(pterm.BgCyan),
		Margin: 20,
	}

	newHeader.Println("汤尼哥扫描器 v16.8")
	pterm.FgRed.Println("欢迎老板使用发财端口扫描器,祝老板越扫越发财!为了能精准的识别服务,请设置合理的限流参数!!!")
	pterm.FgRed.Println("That's how we do things!")
	pterm.Success.Println(fmt.Sprintf("指纹加载成功,共计[%d]个探针,[%d]条正则,[%d]条TCP端口指纹",len(mynmap.Gnmaps.Probes), mynmap.Gnmaps.Count(), len(mynmap.Gnmaps.Services)))
	flag.Parse()

	if common.NumThreads < 1 || common.NumThreads > 2000 {
		fmt.Println("设置的线程数必须在2000之间")
		os.Exit(-1)
	}

	if common.ShowScanType {
		showmode()
	}

	 //统计规则行数为全局变量
	 /*
	apath, _:=os.Getwd()
    var fingerfile string
    if common.FingerFile == "" {
    	//默认的正则文件路径
    	if strings.Index(apath,"\\") > 0{
    		//windows
    	    fingerfile = utils.GetParentDirectory(apath) +"\\finger\\hc-service-probes"
    	} else {
    		//linux os
    		fingerfile = utils.GetParentDirectory(apath) +"/finger/hc-service-probes"
    	}
    } else {
    	//命令行通过 -ff选项参数传递文件路径
        fingerfile = common.FingerFile
    }
     common.FingerCount,_ = common.GetServiceLines(fingerfile)
     common.GetServiceList(fingerfile)
     //pterm.Success.Println("总共发现服务规则数量为: ",common.FingerCount)
   */
}

func main() {
	/*
      普通context通常这样调用:ctx,cancel := context.WithCancel(context.Background())
	*/
	start := time.Now() //获取当前时间
    ctx, cancel := context.WithCancel(context.Background())
    run.Start(ctx)

    cancel() //达到要求之后触发ctx对象的cancel函数
    pterm.Success.Println("扫描任务结束......")
    elapsed := time.Since(start)
    fmt.Println("该函数执行完成耗时：", elapsed)
    //http.ListenAndServe("0.0.0.0:8000", nil)
}

func showmode(){
	fmt.Println("-m 指定的任务模块不存在.....")
	fmt.Println("-m")
    os.Exit(0)
}
