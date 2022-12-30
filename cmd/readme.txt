
go run main.go -i 192.168.30.129 -p 1-6000 -m port -li 0 -n 50
go run main.go -i 192.168.30.129 -p 1-65535 -m port -li 0 -n 2000 -ng

-i  设置扫描主机IP或者IP段
-p  设置扫描端口范围
-m  设置功能模块
-li 设置流量大小
-n  设置启动的协程数量大小
-ng 设置没有识别服务的活动端口到ports.txt中查询默认服务
-ns 只进行端口探活,不进行指纹服务识别
-t  设置连接端口超时时间 内网默认2ms(需要适当调大) 外网默认1秒(需要适当调大)
-v6 探测ipv6主机地址
-on 扫描外网主机
-debug 打印输出详细日志信息



3.生成main.exe
  cd /myscan/cmd
  go build main.go

  

4.开启指纹识别扫描任务
   cd /myscan/cmd
   main.exe -i 192.168.30.129 -p 1-65535 -m port -li 0 -n 1000 -ng


5.主机探活
   只要启用-m ping 模块即可
   go run main.go -i 192.168.30.0/24  -m ping


6.端口探活
   添加 -ns 选项参数
   go run main.go -i 192.168.30.129/24 -p 1-65535 -m port -li 0 -n 1000 -ng -ns

7.ipv6探测语法
   go run main.go -i fe80::3de0:b557:a1f:d123%18 -p 1-6000 -m port  -li 0 -n 50 -v6
   go run main.go -i fe80::3de0:b557:a1f:d123%18 -p 80 -m port -v6



statik使用方法
  1.下载http://github.com/rakyll/statik到本地
  2.go install 会在GOPATH/bin下生成statik.exe
  3.把statik.exe 拷贝到/myscan/cmd目录下


  1.cd /myscan/cmd/file
    修改nmap.txt
    修改port.txt

  2.cd /myscan/cmd
  statik.exe -f -src=file 

