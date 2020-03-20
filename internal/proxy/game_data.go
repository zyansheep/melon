package proxy

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ConvertGameData converts the data from a Java client into Bedrock edition GameData.
func ConvertGameData(javaClient *bot.Client) minecraft.GameData {
	gameData := minecraft.GameData{}

	gameData.Time = 0 // We need to find this!

	gameData.EntityRuntimeID = uint64(javaClient.EntityID)
	gameData.EntityUniqueID = int64(javaClient.EntityID)

	gameData.Dimension = int32(javaClient.Dimension)
	gameData.PlayerGameMode = int32(javaClient.Gamemode)
	gameData.WorldGameMode = int32(javaClient.Gamemode)

	gameData.PlayerPosition = mgl32.Vec3{float32(javaClient.X), float32(javaClient.Y), float32(javaClient.Z)}
	gameData.WorldSpawn = protocol.BlockPos{int32(javaClient.X), int32(javaClient.Y), int32(javaClient.Z)} // We need to find this!
	gameData.Pitch = javaClient.Pitch
	gameData.Yaw = javaClient.Yaw

	gameData.ServerAuthoritativeMovement = false // For now.

	return gameData
}
