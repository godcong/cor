# cor

write the server code as:
  corlog.LogToFile("server")
	serv := cor.NewServer()
	//serv.Head().SetRW(true)
	serv.WriteCallback(func(h cor.Header) (interface{}, error) {
		T := cor.Test{
			"congcong",
			1,
			2,
			3,
		}

		return T, nil

	})
	serv.ReadCallback(func(h cor.Header, v interface{}) error {
		log.Println("reader", h, v)
		return nil
	})
	serv.Start()



write the client code as:
  cli := cor.NewClient()
	//cli.Head().SetRW(true)

	cli.WriteCallback(func(h cor.Header) (interface{}, error) {
		return t{"cong", 1}, nil
	})

	cli.ReadCallback(func(h cor.Header, v interface{}) error {
		log.Println(v)
		return nil
	})
	cli.Start()

