package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Room struct {
	Name string
	AvailableRooms []string
	welcomeString string
	describeString string
	TasksIsActive bool
	PlacesMap map[string]Place
	IsLocked bool
}

func (r *Room) GetWelcomeString() (answer string) {
	answer = r.welcomeString + " можно пройти - " + strings.Join(r.AvailableRooms, ", ")
	return
}

type Item struct{
	name string
	kind string
	message string
}

type Place struct {
	Name string
	Items []*Item
}

type Player struct {
	room *Room
	Backpack
	TasksV2 []string
}

func (p *Player) GetLookAroundString() (answer string) {
	answer += p.room.describeString
	
	placesString := []string{}
	for key, items := range p.room.PlacesMap {
		if len(items.Items) > 0 {
			itemsString := key + ": "
			tempSlice := []string{}
			for _, item := range items.Items {
				tempSlice = append(tempSlice, item.name)
			}
			itemsString += strings.Join(tempSlice, ", ")
			placesString = append(placesString, itemsString)
		}
	}
	answer += strings.Join(placesString, ", ")
	if player.room.Name == "кухня" {
		answer += ", "
	} else {
		answer += ". "
	}
	if p.room.TasksIsActive {
		answer += "надо " + strings.Join(p.TasksV2, " и ") + ". "
	}

	if len(p.Backpack.Items) == 2 && player.room.Name == "комната"{
		answer = "пустая комната. "
	}
	answer += "можно пройти - " + strings.Join(p.room.AvailableRooms, ", ")
	return
}

type Backpack struct {
	IsActive bool
	Items []Item
}

var player Player
var rooms map[string]*Room

func main() {
	initGame()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		answer := handleCommand(scanner.Text())
		fmt.Println(answer)
	}
}

func initGame() {
	player = Player{}
	fmt.Println("---------------------Новая игра---------------------")
	player.TasksV2 = []string{"собрать рюкзак", "идти в универ"}
	

	rooms = map[string]*Room{
		"кухня": {
			Name: "кухня", 
			welcomeString: "кухня, ничего интересного.", 
			describeString: "ты находишься на кухне, ",
			AvailableRooms: []string{"коридор"},
			TasksIsActive: true,
			PlacesMap: map[string]Place{"на столе": {Items: []*Item{{"чай", "еда", "ты выпил чай"}}}},
		},
		"улица": {
			Name: "улица", 
			welcomeString: "на улице весна. можно пройти - домой",
			AvailableRooms: []string{"коридор"},
			IsLocked: true,
		},
		"комната": {
			Name: "комната", 
			welcomeString: "ты в своей комнате.",
			PlacesMap: map[string]Place{
				"на стуле": {
					Items: []*Item{
						{"рюкзак", "сумка", "вы надели: рюкзак"},
					},
				},
				"на столе": {
					Items: []*Item{
						{"ключи", "вещь", "предмет добавлен в инвентарь: ключи"},
						{"конспекты", "вещь", "предмет добавлен в инвентарь: конспекты"},
					},
				},
			},
			AvailableRooms: []string{"коридор"},
		},
		"коридор": {
			Name: "коридор", 
			welcomeString: "ничего интересного.",
			AvailableRooms: []string{"кухня", "комната", "улица"},
		},
	}
	player.room = rooms["кухня"]
}

func handleCommand(command string) string {
	commandSlice := strings.Split(command, " ")
	switch commandSlice[0] {
	
	case "осмотреться":
		return lookAround(&player)
	case "идти":
		return move(&player, commandSlice[1], rooms)
	case "надеть":
		return putOn(&player, commandSlice[1])
	case "взять":
		return take(&player, commandSlice[1])	
	case "применить":
		return use(&player, commandSlice[1], commandSlice[2])
	}
	return "неизвестная команда"
}

func move(player *Player, targetRoom string, rooms map[string]*Room) (answer string) {
	if rooms[targetRoom].IsLocked {
		answer = "дверь закрыта"
		return
	}
	for _, roomName := range player.room.AvailableRooms {
		if roomName == targetRoom {
			player.room = rooms[roomName]
			if player.room.Name == "улица" {
				answer = player.room.welcomeString
			} else {
				answer = player.room.GetWelcomeString()
			}
			return
		}
	}
	answer = "нет пути в " + targetRoom
	return
}

func lookAround(player *Player) (answer string) {
	answer = player.GetLookAroundString()
	return
}

func take(player *Player, itemName string) (answer string) {
	if !player.Backpack.IsActive {
		answer = "некуда класть"
		return answer
	}
	for key, place := range player.room.PlacesMap {
		var itemIndexToDelete*int = nil
		for idx, item := range place.Items {
			if item.kind == "вещь" && item.name == itemName {
				player.Backpack.Items = append(player.Backpack.Items, *item)
				if len(player.Backpack.Items) == 2 {
					player.TasksV2 = deleteElement(player.TasksV2, "собрать рюкзак")
				}
				copyOfIndex := idx
				itemIndexToDelete = &copyOfIndex
				answer = item.message
			}
		}
		if itemIndexToDelete != nil {
			place.Items = append(place.Items[:*itemIndexToDelete], place.Items[*itemIndexToDelete+1:]...)			
			player.room.PlacesMap[key] = place
			
			return answer
		}
	}
	answer = "нет такого"
	return
}

func putOn(player *Player, itemName string) (answer string) {
	for key, place := range player.room.PlacesMap {
		var itemIndexToDelete*int = nil
		for idx, item := range place.Items {
			if item.kind == "сумка" && item.name == itemName {
				player.Backpack.IsActive = true
				itemIndexToDelete = &idx
				answer = item.message
			}
		}
		if itemIndexToDelete != nil {
			place.Items = append(place.Items[:*itemIndexToDelete], place.Items[*itemIndexToDelete+1:]...)			
			player.room.PlacesMap[key] = place
			return answer
			}
	}
	answer = "нет такого"
	return
}

func use(player *Player, itemName string, interactionObjectName string) (answer string) {
	for _, item := range player.Backpack.Items {
		if itemName == item.name {			
			if player.room.Name == "коридор" && interactionObjectName == "дверь"{
				if rooms["улица"].IsLocked {
					rooms["улица"].IsLocked = false
					answer = interactionObjectName + " открыта"
					return
				}
			} else {
				answer = "не к чему применить"
				return
			}
		}
	}
	answer = "нет предмета в инвентаре - " + itemName
	return
}

func deleteElement(slice []string, element string) []string {
    for i := 0; i < len(slice); i++ {
        if slice[i] == element {
            slice = append(slice[:i], slice[i+1:]...)
            return slice
        }
    }
    return slice
}