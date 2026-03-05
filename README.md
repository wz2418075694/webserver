- project: webserver
-- httpserver
-- router
--bumoudle
--- student



1. 再次整理模块，至少区分出业务和技术模块
2. student完善出来
   3. 设计一个结构存储学生的基础信息，比如 姓名 性别 年龄 电话 城市 
   4. 学生的基础信息需要从api传入，并且支持api获取，获取方式分为批量获取全部和单个
   5. forexample 
      6. --Get http://127.1:8080/students 全量获取
      6. --Get http://127.1:8080/student?name=xxx 单个获取
      7. --Post http://127.1:8080/student  创建(传入)单个或批量的学生信息
   8. 将所有的学生信息都保存在本地文件中，要求单个学生单个文件保存 保存格式为json
9. ID生成器新增定时器模式，每分钟触发定时器自动更新



== project

--webserver 
----httpserver 
----idgener
----router
----serverinfo
----student
----main



= todo @20260227
1. 完善mysql针对student模块的查询
2. 学习golang包，并在项目中全量使用，同时学会使用第三方包
3. 学习gorm包，进一步改进mysql的操作，同时进一步学会使用第三方包
4. 给webserver增加一个配置文件，将写死在代码里的配置抽离到配置文件，webserver.ini 参考：https://github.com/HouGuoFa/goconfiger
5. 学习并使用gin框架 传送门： https://github.com/gin-gonic/gin   中文版：https://gin-gonic.com/zh-cn/docs/
6. 将webserver 翻译成python语言，重新实现 webserver-py  对标gin的框架有 flash


== book
1. 程序员面试宝典 
2. 剑指offer