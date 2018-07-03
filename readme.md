├── app #项目主体
│ ├── controller #控制器文件夹
│ │ └── testController.go #控制器文件 小驼峰命名
│ ├── models #模型文件夹
│ │ └── nodeModel.go #模型文件
│ ├── routes.go #路由注册文件
│ ├── conf #配置文件夹
│ │ └── dao.ini #配置文件
│ ├── public #静态文件夹 以后可以再添css html js 等子文件夹
│ ├── utils #工具类文件夹，超过10个以上文件建议再开子文件夹归类
│ │ ├── conf.go #加载配置文件，并暴露Cfg操作配置文件
│ │ └── dao.go #数据层(mysql,redis)加载，并暴露引擎实例
│ │ └── funcs.go #抽离控制器的方法放这
│ │ └── log.go #日志引擎加载和暴露
│ │ └── session.go #session中间件资源初始化
│ │ └── init.go #初始化方法有顺序问题的话就都放这里
│ │ └── websocket.go #websocket的处理，这个项目无
│ ├── vendor #第三方包文件夹 用govendor管理，不然有版本问题
│ └── main.go #入口文件
