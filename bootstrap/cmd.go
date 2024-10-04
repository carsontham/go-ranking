package bootstrap

func Run() {
	s := NewServer(":3000")
	GetDB()
	s.RunServer()
}
