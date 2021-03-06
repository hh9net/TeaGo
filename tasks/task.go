package tasks

var channel = make(chan func())

func Start(concurrency int) {
	for i := 0; i < concurrency; i ++ {
		go func(i int) {
			for {
				fn := <-channel
				fn()
			}
		}(i)
	}
}

func Add(fn func()) {
	channel <- fn
}
