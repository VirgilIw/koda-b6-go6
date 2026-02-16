package internal

type Data struct {
	Name string
	Wait int
}

func Queue(ch chan Data) {
	data := []Data{
		{
			Name: "virgil",
			Wait: 5},
		{
			Name: "test",
			Wait: 2},
		{
			Name: "piuu",
			Wait: 3},
		{
			Name: "check",
			Wait: 8},
	}

	// key, index
	for _, i := range data {
		ch <- i
	}

	close(ch)
	// close(channel) di Go itu menandakan bahwa tidak akan ada data lagi yang dikirim ke channel tersebut.
}
