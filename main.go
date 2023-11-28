package main

import (
	"bufio"
	"fmt"
	// "go/printer"
	"os"
	"strings"
)

type Room struct {
	Name string
	AvailableRooms []string
	welcomeString string
	describeString string
	TasksIsActive bool
	Places []Place
}

func (r *Room) GetWelcomeString() (answer string) {
	answer = r.welcomeString + "можно пройти - " + strings.Join(r.AvailableRooms, ", ")
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
	Tasks map[string]bool 
	TasksV2 []string
}

func (p *Player) GetLookAroundString() (answer string) {
	answer += p.room.describeString
	for _, place := range p.room.Places {
		answer += place.Name + " "
		for _, item := range place.Items {
			answer += item.name + ", "
		}
	}
	if p.room.TasksIsActive {
		answer += "надо " + strings.Join(p.TasksV2, " и ") + "."
	}
	answer += " можно пройти - " + strings.Join(p.room.AvailableRooms, ", ")
	return
}

type Backpack struct {
	IsActive bool
	Items []Item
}

var player Player = Player{}
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

	fmt.Println("---------------------Новая игра---------------------")

	player.Tasks = map[string]bool{
		"собрать рюкзак" : false,
		"идти в универ" : false,
	}
	player.TasksV2 = []string{"собрать рюкзак", "идти в универ"}
	

	rooms = map[string]*Room{
		"кухня": {
			Name: "кухня", 
			welcomeString: "кухня, ничего интересного.", 
			describeString: "ты находишься на кухне, ",
			AvailableRooms: []string{"коридор"},
			Places: []Place{{Name: "на столе", Items: []*Item{{"чай", "еда", "ты выпил чай"}}}},
			TasksIsActive: true,
		},
		"комната": {
			Name: "комната", 
			welcomeString: "ты в своей комнате.",
			Places: []Place{
				{
					Name: "на столе",
					Items: []*Item{
						{"ключи", "вещь", "предмет добавлен в инвентарь: ключи"},
						{"конспекты", "вещь", "предмет добавлен в инвентарь: конспекты"},
					},
				},
				{
					Name: "на стуле",
					Items: []*Item{
						{"рюкзак", "сумка", "вы надели: рюкзак"},
					},
				},
			},
			AvailableRooms: []string{"коридор"},
		},
		"коридор": {
			Name: "коридор", 
			welcomeString: "ничего интересного.",
			AvailableRooms: []string{"кухня", "улица", "комната"},
		},
		"улица": {
			Name: "улица", 
			welcomeString: "на улице весна. можно пройти - домой",
			AvailableRooms: []string{"коридор"},
		},
	}
	player.room = rooms["кухня"]
}

func handleCommand(command string) string {
	commandSlice := strings.Split(command, " ")
	switch commandSlice[0] {
	
	case "о":
		return lookAround(&player)
	case "идти":
		return move(&player, commandSlice[1], rooms)
	case "надеть":
		return putOn(&player, commandSlice[1])
	case "взять":
		fmt.Println("")
	
	case "применить":
		fmt.Println("")
	}
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	return "not implemented"
}

func move(player *Player, targetRoom string, rooms map[string]*Room) (answer string) {
	for _, roomName := range player.room.AvailableRooms {
		if roomName == targetRoom {
			player.room = rooms[roomName]
			answer = player.room.GetWelcomeString()
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

func putOnOLD(player *Player, itemName string) (answer string) {

	newPlaces := []Place{}
	
	for _, place := range player.room.Places {
		var itemIndexToDelete*int = nil	
		for idx, item := range place.Items {
			if item.kind == "сумка" && item.name == itemName {
				player.Backpack.IsActive = true
				itemIndexToDelete = &idx
				answer = item.message
			}
		if itemIndexToDelete != nil {
			place.Items = append(place.Items[:*itemIndexToDelete], place.Items[*itemIndexToDelete+1:]...)			
			return answer
			}
		}
	}
	answer = "нет такого"
	return
}

func putOn(player *Player, itemName string) (answer string) {
	
}