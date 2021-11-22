package networks

import (
	"bufio"
	"math/rand"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestListener(t *testing.T) {
	Listener(":8080")
}

func TestClient(t *testing.T) {

	conn, err := net.Dial("tcp", ":8080")

	for i := 0; i < 10; i++ {
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
		time.Sleep(time.Second)
		println(string(line))
	}

}
