
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
func newRoom(room Room) string {
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
func newRoom(room Room) string {

	// IT'S COMMENTED, BECAUSE WE USE GLOBAL VARIABLE - "servers"
	// servers:=getFromRedis()

	// temporary server to find maximum cpu usage
	suitServer := Server{}

	// initializing duration per CPU Usage
	cpuUsage := 0
	if room.Duration == 60 {
		cpuUsage = 80
	} else {
		cpuUsage = 120
	}

	// Find available server by CPU Usage
	for _, server := range servers {

		isAvailable := server.CPUusage+cpuUsage <= server.CPUcapacity
		length := len(server.Rooms)

		if server.Status {						//search (running) servers:

			switch {
			case isAvailable && length != 0:			//-according to its cpuUsage and existing room
				suitServer = server

			case isAvailable && length == 0:			//-according to its cpuUsage and no existing room
				suitServer = server

			}
		} else {							//if not found suitable server yet
			if !server.Status && suitServer.Address == "" {		//-according to existing not running server
				server.Status = true
				suitServer = server
			}
		}
	}

	if suitServer.Address == "" && suitServer.CPUcapacity == 0 {		// if not found above cases, we creaate new server
		suitServer.Address = "newserver"
		suitServer.CPUcapacity = 800
		suitServer.Status = true

		servers = append(servers, suitServer)				// and add to server list
	}

	// add Room to found server
	suitServer.Rooms = append(suitServer.Rooms, Room{Id: room.Id, Duration: room.Duration})

	// random cpu usage
	suitServer.CPUusage += cpuUsage

	for i := 0; i < len(servers); i++ {

		// find servers with no rooms and turn off			//turn off left servers with no rooms if exists
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
```


## Запуск тестов

Для запуска тестов выполните следующую команду
* Выполните тесты одиночно, возможность переполнение (глобальной переменной) __servers__,
```bash
go test
```
```bash
go test -v -cover -short ./...
```
## Тесты
обычный запись комнаты:
```go
func TestHasRooms(t *testing.T) {...}
```
заполняет один из серверов, переходит на другой сервер с комнатами:
```go
func TestServer4Full(t *testing.T) {...}
```
записывает в сервер с пустой комнатой, в случае все серверы уже имеющие комнаты заполнены:
```go
func TestNoRooms(t *testing.T) {...}
```
тест на включения имеющего сервера и запись комнаты, в случае все работающие серверы заполнены:
```go
func TestNoAnyRunningServer(t *testing.T) {...}
```
тест на создание нового сервера, в случае все остальные серверы заполнены:
```go
func TestNoAnyServer(t *testing.T) {...}
```
