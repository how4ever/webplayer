# an internal webplayer  
an internal web player  
一个局域网播放器  

1.拷贝release目录下的所有文件到要播放视频的目录  
  copy all files in [release] path to the file path which video existed  
  
2.win+r后输入cmd打开命令行窗口并cd到第1步的视频目录  
  press win+r then input [cmd] command open the command window, then cd to the 1st step file path  

3.运行ipconfig查看本地局域网ip, 运行player  
  run [ipconfig] command view the local ip, then run player  
  
4.打开浏览器输入ip回车选择视频观看  
  open web browser like chrome, then input ip, select a video to watch  
  
5.只支持mp4格式的视频  
  only support mp4 format videos  
  
6.默认为80端口,直接输入ip回车即可, 也可使用自定义端口, 如在第3步运行命令play -12345使用12345端口, 浏览器打开时用ip:12345浏览  
  default use port 80, you can use custom port, for exemple run command [play -12345] use port 12345 in step 3, then use ip:12345 address in web browser  
