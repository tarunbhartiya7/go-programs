// package main

// import "fmt"

// const (
// 	Online      = 0
// 	Offline     = 1
// 	Maintenance = 2
// 	Retired     = 3
// )

// func printServerStatus(statusMap map[string]int) {
// 	for server, status := range statusMap {
// 		var statusStr string
// 		switch status {
// 		case Online:
// 			statusStr = "Online"
// 		case Offline:
// 			statusStr = "Offline"
// 		case Maintenance:
// 			statusStr = "Maintenance"
// 		case Retired:
// 			statusStr = "Retired"
// 		default:
// 			panic("Invalid status")
// 		}
// 		fmt.Println(server, " - ", statusStr)
// 	}
// }

// func main() {
// 	servers := []string{"server1", "server2", "server3", "server4", "server5"}

// 	statusMap := make(map[string]int)
// 	for _, server := range servers {
// 		statusMap[server] = Online
// 	}
// 	printServerStatus(statusMap)

// 	statusMap["server1"] = Offline
// 	statusMap["server2"] = Maintenance

// 	printServerStatus(statusMap)

// }
