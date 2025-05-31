package routine

// Go golang协程封装
func Go(fn func()) {
	go Safe(fn)
}
