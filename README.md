# scanport


选项参数说明
 -i  设置扫描主机IP或者IP段

 -p  设置扫描端口范围

 -m  设置功能模块

 -li 设置流量大小

 -n  设置启动的协程数量大小

 -ng 设置没有识别服务的活动端口到ports.txt中查询默认服务

 -ns 只进行端口探活,不进行指纹服务识别

 -t  设置连接端口超时时间
 
 -v6 探测ipv6主机地址



1.cd /myscan/cmd/file
    修改nmap.txt
    修改port.txt



2.开启指纹识别扫描任务
   main.exe -i 192.168.30.129 -p 1-65535 -m port -li 0 -n 1000 -ng
   go run main.go -i 192.168.30.129 -p 1-6000 -m port -li 0 -n 50
   go run main.go -i 192.168.30.129 -p 1-65535 -m port -li 0 -n 2000 -ng
   
  

3.主机探活
   只要启用-m ping 模块即可
   go run main.go -i 192.168.30.0/24  -m ping


4.端口探活
   添加 -ns 选项参数
   go run main.go -i 192.168.30.129/24 -p 1-65535 -m port -li 0 -n 1000 -ng -ns
   
 ![image](https://user-images.githubusercontent.com/11001852/210042467-dd2bfd74-b334-472f-95de-466c3d4453be.png)





   
