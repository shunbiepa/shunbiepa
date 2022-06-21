package main

func main() {
	testMapReadWriteDiffKey()
}

func testMapReadWriteDiffKey() {
	m := make(map[int]int)
	go func() {
		for {
			m[100] = 100
		}
	}()
	go func() {
		for {
			_ = m[12]
		}
	}()
	select {}
}
