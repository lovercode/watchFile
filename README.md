# watchFile

--------------

### 监控文件变化，自动执行对应的脚本


### 配置

```
[files]
    [files."conf/test.txt"]
         #监控删除动作
         REMOVE = ["conf/test.sh","conf/run.sh"]
         #监控重命名
#        RENAME = ["conf/test.sh","conf/run.sh"]
         #监控创建
#        CREATE = ["conf/test.sh","conf/run.sh"]
         #监控写入
#        WRITE = ["conf/test.sh","conf/run.sh"]
         #监控权限变化
#        CHMOD = ["conf/test.sh","conf/run.sh"]

```