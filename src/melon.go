package main

import (
	"fmt"
	"net"
	"strconv"

	"config"
	"protocol/packet"
	"util"
)

func main() {
	cfg := config.NewConfig()

	err := util.ReadJSONFile("config.json", &cfg)
	if err != nil {
		panic(err)
	}

	err = util.WriteJSONFile("config.json", cfg)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(cfg.Port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("Listening on port %v...\n", cfg.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		

		defer conn.Close()

		go func() {
			for {
				pk, err := packet.ReadPacket(conn)
				if err != nil {
					return
				}

				if pk != nil {
					switch pk.ID() {
					case packet.HandshakePacketID:
						if pk.(*packet.HandshakePacket).NextState == 1 {
							response := packet.ResponsePacket{}
							response.Response = "{\"version\":{\"name\":\"1.15.2\",\"protocol\":578},\"players\":{\"max\":100,\"online\":0,\"sample\":[]},\"description\":{\"text\":\"Hello Melon!\"}}"
							packet.WritePacket(conn, &response)
						}
					case packet.PingID:
						if pk.(*packet.HandshakePacket)
					}
				}
			}
		}()
	}
}