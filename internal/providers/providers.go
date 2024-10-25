package providers

func InitProviders(quit chan bool) {
	go initMongo(quit)
}
