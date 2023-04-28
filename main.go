package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ProbableMaximumUsageCPU = 120
	ProbableMinimumUsageCPU = 80
)

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

func newRoom(room Room) string {

	// IT'S COMMENTED, BECAUSE WE USE GLOBAL VARIABLE - "servers"
	// servers:=getFromRedis()

	// temporary server to find suitable server
	suitServer := Server{}

	cpuUsage := 0
	if room.Duration == 60 {
		cpuUsage = 80
	} else {
		cpuUsage = 120
	}

	// Find available server with maximum CPU Usage
	for _, server := range servers {

		isAvailable := server.CPUusage+cpuUsage <= server.CPUcapacity
		length := len(server.Rooms)

		if server.Status {

			switch {
			case isAvailable && length != 0:
				suitServer = server

			case isAvailable && length == 0:
				suitServer = server

			}
		} else {
			if !server.Status && suitServer.Address == "" {
				server.Status = true
				suitServer = server
			}
		}
	}

	if suitServer.Address == "" && suitServer.CPUcapacity == 0 {
		suitServer.Address = "newserver"
		suitServer.CPUcapacity = 800
		suitServer.Status = true

		servers = append(servers, suitServer)
	}

	// add Room to found server
	suitServer.Rooms = append(suitServer.Rooms, Room{Id: room.Id, Duration: room.Duration})

	// random cpu usage
	suitServer.CPUusage += cpuUsage

	for i := 0; i < len(servers); i++ {

		// find servers with no rooms and turn off
		if len(servers[i].Rooms) == 0 {
			servers[i].CPUusage = 0
			servers[i].Status = false
		}

		// update chosen server
		if servers[i].Address == suitServer.Address {
			servers[i].Address = suitServer.Address
			servers[i].CPUcapacity = suitServer.CPUcapacity
			servers[i].CPUusage += cpuUsage
			servers[i].Status = suitServer.Status
			servers[i].Rooms = suitServer.Rooms
		}
	}

	return suitServer.Address
}

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
		Rooms:       []Room{{Id: 1111, Duration: 60}},
	},
	{
		Address:     "server4",
		CPUcapacity: 800,
		CPUusage:    120,
		Status:      true,
		Rooms:       []Room{{Id: 1112, Duration: 180}},
	},
}

func main() {

	// example of room for testing
	room := Room{
		Id:       2222,
		Duration: 180,
	}
	room1 := Room{
		Id:       2223,
		Duration: 60,
	}
	room2 := Room{
		Id:       2224,
		Duration: 180,
	}

	serverAddress := newRoom(room)
	serverAddress1 := newRoom(room1)
	serverAddress2 := newRoom(room2)
	serverAddress3 := newRoom(room1)
	serverAddress4 := newRoom(room2)
	serverAddress5 := newRoom(room2)
	serverAddress6 := newRoom(room)
	fmt.Println(serverAddress)
	fmt.Println(serverAddress1)
	fmt.Println(serverAddress2)
	fmt.Println(serverAddress3)
	fmt.Println(serverAddress4)
	fmt.Println(serverAddress5)
	fmt.Println(serverAddress6)

}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
