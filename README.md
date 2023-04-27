
# Тестовое задание

Структура сервера:
```go
type Server struct {
	Address     string
	CPUcapacity int    // Max CPU
	CPUusage    int    // CPU usage
	Status      bool   // On or Off
	Rooms       []Room // List of rooms
}
```

Структура комнаты (конференций):
```go
type Room struct {
	Id       int64
	Duration int
}
```

Структура приложения:
```go
// закомментирован для дальнейшей реализации
// func getFromRedis() ([]Server, error) {...}

// * Основная логика
func newRoom(room Room, servers []Server) string {
  ...
}


func main() {
  ...
}

// для получения рандомного числа
func RandomInt(min, max int) int {
  ...
}

```

## Запуск

Клонировать проект

```bash
  git clone https://gitlab.com/ikayrat/testserver.git
```

Запустить

```bash
  go run main.go
```


## Основная часть

```go
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
```


## Запуск тестов

Для запуска тестов выполните следующую команду
```bash
go test
```
```bash
go test -v -cover -short ./...
```
