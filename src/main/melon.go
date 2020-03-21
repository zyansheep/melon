package main

import (
	"fmt"
	"net"
	"strconv"

	"../config"
	"../protocol"
	"../packet"
	"../util"
	
)

func main() {
	cfg := config.NewConfig()

	err := util.ReadJSONFile("config.json", &cfg)
	if err != nil { panic(err) }

	err = util.WriteJSONFile("config.json", cfg)
	if err != nil { panic(err) }

	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(cfg.Port))
	if err != nil { panic(err) }
	defer listener.Close()

	fmt.Printf("Listening on port %v...\n", cfg.Port)

	for {
		//When connection established
		conn, err := listener.Accept()
		if err != nil { panic(err) }
		

		defer conn.Close()
		//Create thread
		go func() {
			for {
				pk, err := packet.ReadPacket(conn)
				if err != nil {
					return
				}

				if pk != nil {
					switch pk.ID() {
					case protocol.HandshakePacketID:
						var resp protocol.Response;
						resp.JSON = `{"version\":{"name":"1.15.2","protocol":578},"players":{"max":100,"online":0,"sample":[{"Zyansheep", "e8c3399a-c5c5-4133-9936-2aba905d6770"}]},"description":{"text":"Hello Melon 2!"}}`
						packet.WritePacket(conn, resp);
					case protocol.PingID:
						var resp protocol.PongPacket;
						packet.WritePacket(conn, resp);
					}
				}
			}
		}()
	}
}
