# WeeklyReportTool
## What can we do?
- E-mail your weekly report to a specific E-mail address automatically.  
- Remind you of updateing your weekly report by sending a warning email.  

## 本程序的主要功能  
- 自动发送你已经编辑完的周报至指定的邮箱。  
- 如果你忘了写周报，发送提示邮件至你指定的邮箱来提醒你。  
*************************************  
编辑email.env文件来进行配置,右键打开方式，用记事本打开。  

------------------------------------------------------------------------
  
腾讯企业邮箱服务器 HOST="smtp.exmail.qq.com"  
腾讯企业邮箱服务器端口 PORT="465"  
(注：如果你的邮箱是腾讯企业邮箱，那么以上两行都不需要更改。如果你是网易邮箱，则HOST="smtp.163.com"，PORT="25",并且网易邮箱需要设置客户端授权密码，在网易邮箱网页版点击“设置”,然后点击左侧的“客户端授权密码”，然后再打开页面点击“重置授权码”，新设置的密码填入下方的“PASSWORD=”中，腾讯企业邮箱是直接输入邮箱密码就可以)  
你的腾讯企业邮箱 USER="xxxxx@xxxxx.com"  
腾讯企业邮箱密码 PASSWORD="xxxxxxxx"  
对方的邮箱 TO="example@xx.com"   
周报目录 ATTACH="D:\xxxx\xxxxx"  
你的名字 NAME="张三"  
周报命名格式 周报-张三-2017-09-22.doc  
  
  
------------------------------------------------------------------------
  
  
配置完之后双击 sendmail.exe 便可以发送邮件。
  
  
配置完以后还不能自动地定时地发送邮件，你可以继续配置Windows系统自带的任务计划程序，参考http://jingyan.baidu.com/article/6181c3e0435026152ef153d0.html
