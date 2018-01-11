package parseTCP

import (
  "testing"
)

var packetData = []byte{
  69, 0, 0, 63, 64, 237, 64, 0, 64, 6, 74, 221, 192,
  168, 1, 90, 52, 85, 184, 151, 141, 230, 0, 80,
  174, 147, 86, 192, 18, 107, 243, 149, 128, 24,
  0, 229, 92, 65, 0, 0, 1, 1, 8, 10, 22, 138, 85, 109,
  48, 16, 32, 253, 49, 50, 51, 52, 53, 54, 55, 56,
  57, 48, 10,
}

func TestParse(t *testing.T) {

  packet, err := ParseTCPPacket(packetData)
  if err != nil {
    t.Fatal(err)
  }
  packet.Print()

}
