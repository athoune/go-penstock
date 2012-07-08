package main

func main() {
	client, err := NewClient("localhost", 4807)
	if err != nil {
		//error
	}
	// some login?
	client.Write([]byte(""), []byte("hello"))
}
