package proxy

import (
	"fmt"

	"github.com/Tnze/go-mc/bot"
	"github.com/melondevs/melon/internal/util"
	"github.com/sandertv/gophertunnel/minecraft"
)

// RunProxy opens and runs the Melon proxy server.
func RunProxy(config util.Config) {
	// Open Bedrock listener.
	listener, err := minecraft.Listen("raknet", config.MelonAddress)
	if err != nil {
		panic(err)
	}

	listener.ServerName = config.Name + util.FormatReset

	defer listener.Close()
	for {
		// Accept incoming Bedrock connection.
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		bedrockClient := conn.(*minecraft.Conn)
		fmt.Println(conn.RemoteAddr().String() + " (" + bedrockClient.ClientData().ThirdPartyName + ") connected.")

		// Create Java client.
		javaClient := bot.NewClient()
		javaClient.Name = bedrockClient.ClientData().ThirdPartyName

		err = javaClient.JoinServer(config.HostAddress, config.HostPort)
		if err != nil {
			panic(err)
		}

		// Start game on Bedrock.
		err = bedrockClient.StartGame(ConvertGameData(javaClient))
		if err != nil {
			javaClient.Disconnect()
		}

		// Process connection.
		go func() {
			// Process Java client.
			go func() {
				javaClient.HandleGame()
			}()

			// Process Bedrock client.
			defer conn.Close()
			for {
				packet, err := bedrockClient.ReadPacket()
				if err != nil {
					javaClient.Disconnect()
					return
				}

				fmt.Println(packet.ID()) // Print packet IDs.
			}
		}()
	}
}
