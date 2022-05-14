# fcache

基于文件的持久化缓存

本程序定位为简易的纯go本地持久化缓存 不需要配置第三方服务

适用于测试环节,将耗时较久的程序如远程请求,复杂计算等缓存在本地

线上生成环境请酌情使用更专业的redis等

例子:

    type User struct {
        Name string
        Age  int
    }

    var result User
	Remember(&result, func() interface{} {
		fmt.Println(`重新获取`)
		return User{
			Name: `xiaofei`,
			Age:  13,
		}
	})
	fmt.Println(`result: `, result)


多线程:

	// 防止多线程调用造成的重复请求 简单判断文件是否存在 若不存在则加锁
	if _, err := os.Stat(path.Join(fc.DirPath, fileName+fc.Suffix)); err != nil {
		mu.Lock()
		defer mu.Unlock()
	}

待改进:

    看是否需要将 bolt 集成进来

    Remember 方法中: 目前情况是不管是不是有缓存都经过一次编码是解码 感觉性能上有所消耗
        但尚未实现如何让一个 interface 赋值给另一个interface类型 即:out=fun() 

    coder/GobCode.go 中EnCode 方法返回的编码后的内容有问题 尚未实现coder

    作为缓存 还应该加上一些过期时间的设置
