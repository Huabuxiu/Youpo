package networks

import (
	"bufio"
	"math/rand"
	"net"
	"strconv"
	"testing"
)

func TestListener(t *testing.T) {
	Listener(":8080")
}

func TestClient(t *testing.T) {

	conn, err := net.Dial("tcp", ":8080")

	for i := 0; i < 30; i++ {
		val := strconv.Itoa(rand.Int())
		_, err = conn.Write([]byte(val + "\n"))
		if err != nil {
			t.Error(err)
			return
		}
		bufReader := bufio.NewReader(conn)
		line, _, err := bufReader.ReadLine()
		if err != nil {
			t.Error(err)
			return
		}
		println(string(line))
	}

}
