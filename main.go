package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ProbableMaximumUsageCPU = 120
const ProbableMinimumUsageCPU = 80

type Server struct {
	Address     string
	CPUcapacity int    // Max CPU
	CPUusage    int    // CPU usage
	Status      bool   // On or Off
	Rooms       []Room // List of rooms
}

type Room struct {
	Id       int64
	Duration int
}

// Get servers from Redis
// func getFromRedis() ([]Server, error) {

// 	var servers []Server

// 	// Connect to Redis
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	// Get all servers from Redis
// 	result, err := client.ZRange("servers", 0, -1).Result()
// 	if err != nil {
// 		return servers, err
// 	}

// 	// Unmarshal JSON into Server structs
// 	for _, s := range result {
// 		var server Server
// 		err := json.Unmarshal([]byte(s), &server)
// 		if err != nil {
// 			return servers, err
// 		}
// 		servers = append(servers, server)
// 	}

// 	return servers, nil
// }

func newRoom(room Room, servers []Server) string {

	// servers:=getFromRedis()

	// temporary server to find maximum cpu usage
	maxServer := Server{}

	// Find available server with maximum CPU Usage
	for _, server := range servers {

		// if has rooms
		if len(server.Rooms) != 0 {
			// if usage fit in server capacity
			hasMinUsage := server.CPUusage+ProbableMaximumUsageCPU <= server.CPUcapacity
			if hasMinUsage && server.CPUusage > maxServer.CPUusage && server.Status {
				maxServer = server

			}
			// if has no rooms
		} else if server.Status {
			maxServer = server
		}
	}

	// add Room to found server
	maxServer.Rooms = append(maxServer.Rooms, Room{Id: room.Id, Duration: room.Duration})

	// random cpu usage
	cpuUsage := RandomInt(ProbableMinimumUsageCPU, ProbableMaximumUsageCPU)
	maxServer.CPUusage += cpuUsage

	for i := 0; i < len(servers); i++ {

		// find servers with no rooms and turn off
		if len(servers[i].Rooms) == 0 {
			servers[i].Status = false
		}

		// update choosed server
		if servers[i].Address == maxServer.Address {
			servers[i].Address = maxServer.Address
			servers[i].CPUcapacity = maxServer.CPUcapacity
			servers[i].CPUusage += cpuUsage
			servers[i].Status = maxServer.Status
			servers[i].Rooms = maxServer.Rooms
		}
	}

	return maxServer.Address
}

func main() {

	// example of server with parametres
	var servers = []Server{
		{
			Address:     "server1",
			CPUcapacity: 800,
			CPUusage:    0,
			Status:      true,
			Rooms:       []Room{},
		},
		{
			Address:     "server2",
			CPUcapacity: 800,
			CPUusage:    0,
			Status:      false,
			Rooms:       []Room{},
		},
		{
			Address:     "server3",
			CPUcapacity: 800,
			CPUusage:    80,
			Status:      true,
			Rooms:       []Room{{Id: int64(RandomInt(1000, 1200)), Duration: 60}},
		},
		{
			Address:     "server4",
			CPUcapacity: 800,
			CPUusage:    120,
			Status:      true,
			Rooms: []Room{
				{
					Id:       int64(RandomInt(1000, 1200)),
					Duration: 180,
				},
			},
		},
	}

	// example of room for testing
	room := Room{
		Id:       int64(RandomInt(1000, 1200)),
		Duration: 60,
	}

	serverAddress := newRoom(room, servers)
	fmt.Println(serverAddress)

}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
