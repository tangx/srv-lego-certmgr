package gen

// 包环境变量设置 retry 设置
var (
	// retry 通道， 为了支持多 provider 的情况
	retryChannel = make(map[string]chan string)
)

// // retryApply 错误重试
// func retryApply(prov string) {
// 	ch := make(chan string, 20)
// 	retryChannel[prov] = ch

// 	go func() {
// 		logrus.Infof("启动 %s 重试队列", prov)
// 		for {
// 			domains := <-retryChannel[prov]

// 			for i := 1; i < 4; i++ {
// 				logrus.Infof("%s -> 第 %d 次重试:  %s\n", prov, i, domains)
// 				err := applyCertificate(prov, domains)
// 				if err == nil {
// 					break
// 				}

// 				time.Sleep(30 * time.Second)
// 			}
// 		}
// 	}()
// }
