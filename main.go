package main

import (
	"fmt"
    "strings"
    "math/rand"
    "time"
    "strconv"
)

// Global Variable
var size_line int = 75
var max_turns int = 25
var settings map[string]string = map[string]string{
    "difficulty": "EASY",
    "tutorial": "ON",
    "hints": "ON",
}

func main() {
    // Creating random seed for the random functions
    rand.Seed(time.Now().UnixNano())

    game := true
    play := false

    for game {
        for !play {
            title := map[int]string{
                0:  "   _____  ____  _ __      _______ _   _  _____   ",
                1:  "  / ____|/ __ \\| |\\ \\    / /_   _| \\ | |/ ____|  ",
                2:  " | (___ | |  | | | \\ \\  / /  | | |  \\| | |  __           ",
                3:  "  \\___ \\| |  | | |  \\ \\/ /   | | | . ` | | |_ |             ",
                4:  "  ____) | |__| | |___\\  /   _| |_| |\\  | |__| |             ",
                5:  " |_____/_\\____/|______\\/  _|_____|_|_\\_|\\_____|  _ _______  ",
                6:  " |__   __| |  | |  ____| | \\ | |_   _/ ____| |  | |__   __| ",
                7:  "    | |  | |__| | |__    |  \\| | | || |  __| |__| |  | |   ",
                8:  "    | |  |  __  |  __|   | . ` | | || | |_ |  __  |  | |    ",
                9:  "    | |  | |  | | |____  | |\\  |_| || |__| | |  | |  | |    ",
                10: " __ |_|  |_| _|_|______|_|_| \\_|_____\\_____|_|_ |_|_ |_|___ ",
                11: " \\ \\        / /_   _|__   __| |  | | |__   __| |  | |  ____|",
                12: "  \\ \\  /\\  / /  | |    | |  | |__| |    | |  | |__| | |__   ",
                13: "   \\ \\/  \\/ /   | |    | |  |  __  |    | |  |  __  |  __|  ",
                14: "    \\  /\\  /   _| |_   | |  | |  | |    | |  | |  | | |____ ",
                15: "  _  \\/ _\\/___|_____| _|_|_ |_|__|_|    |_|  |_|  |_|______|",
                16: " | \\ | |  ____|  __ \\|  __ \\ / ____|       ",
                17: " |  \\| | |__  | |__) | |  | | (___     ",
                18: " | . ` |  __| |  _  /| |  | |\\___ \\      ",
                19: " | |\\  | |____| | \\ \\| |__| |____) |        ",
                20: " |_| \\_|______|_|  \\_\\_____/|_____/        ",
                21: "",
                22: "",
                23: "[PLAY]",
                24: "[QUIT]",
            }
            
            description := ""
            if (settings["difficulty"] == "EASY") {
                description = "Equipped with a map and your wits, you're ready to save lives."
            } else if (settings["difficulty"] == "MEDIUM") {
                description = "You seem to have lost your map, and the items seem to have changed. You better be prepared."
            } else if (settings["difficulty"] == "HARD") {
                description = "People seem to need these items now..."
            }
            
            settingText := map[int]string{
                0: "Would you like to change any settings?",
                1: "[d] Difficulty - " + settings["difficulty"],
                2: "    Description: " + description,
                3: "[t] Tutorial - " + settings["tutorial"],
                4: "[h] Hints - " + settings["hints"],
                5: "[c] Continue",
            }

            fmt.Println("\033[2J")
            selected := printOption(title)
            if (selected == "PLAY" || selected == "play") {
                play = true
                fmt.Println("\033[2J")
                selectedSetting := printOption(settingText)
                for true {
                    if (selectedSetting == "d") {
                        if (settings["difficulty"] == "EASY") {
                            settings["difficulty"] = "MEDIUM"
                        } else if (settings["difficulty"] == "MEDIUM") {
                            settings["difficulty"] = "HARD"
                        } else if (settings["difficulty"] == "HARD") {
                            settings["difficulty"] = "EASY"
                        }
                    } else if (selectedSetting == "t") {
                        if (settings["tutorial"] == "ON") {
                            settings["tutorial"] = "OFF"
                        } else if (settings["tutorial"] == "OFF") {
                            settings["tutorial"] = "ON"
                        }
                    } else if (selectedSetting == "h") {
                        if (settings["hints"] == "OFF") {
                            settings["hints"] = "ON"
                        } else if (settings["hints"] == "ON") {
                            settings["hints"] = "OFF"
                        }
                    } else if (selectedSetting == "c") {
                        break
                    }

                    if (settings["difficulty"] == "EASY") {
                        description = "Equipped with a map and your wits, you're ready to save lives."
                    } else if (settings["difficulty"] == "MEDIUM") {
                        description = "You seem to have lost your map, and the items seem to have changed. You better be prepared."
                    } else if (settings["difficulty"] == "HARD") {
                        description = "People seem to need these items now..."
                    }

                    settingText = map[int]string{
                        0: "Would you like to change any settings?",
                        1: "[d] Difficulty - " + settings["difficulty"],
                        2: "    Description: " + description,
                        3: "[t] Tutorial - " + settings["tutorial"],
                        4: "[h] Hints - " + settings["hints"],
                        5: "[c] Continue",
                    }

                    fmt.Println("\033[2J")
                    selectedSetting = printOption(settingText)
                }
            } else if (selected == "QUIT" || selected == "quit") {
                return
            }
        }
                                
        // #############################################
        // #            ALL NARRATION TEXT             #
        // #############################################
        introText := map[int]string{
            0: "Dear Pendergast,", 
            1: "I am about to attend a party, and... well... I am worried.",
            2: "I have been tipped off that there is a murder among a group of computer science nerds, and it is up to you to figure out who the murderer is.",
            3: "There are three ways that I believe the killer could attempt to kill someone at this party: poisoning, stabbing, or shooting them.",
            4: "If I were you I would look for specific objects which may be used. For example a gun would probably be used to shoot someone, rat poison may be used to poison someone, and a knife may be used to stab someone.",
            5: "You will have " + strconv.Itoa(max_turns) + " turns to figure out who the murderer is and report them before the murder happens.",
            6: "If you aren't able to stop them... well let's just not let that happen shall we.",
            7: "Within each turn you will be able to: move into any connecting room next to you, search for key items, search someone within a specific room for the items they have, or accuse someone.",
            8: "Don't fail me.",
            9: "Best of luck, Anonymous",
        }
        tutorialText := map[int]string{
            0: "TUTORIAL:",
            1: "Every turn you have 4 actions: acusing someone of being the murderer, searching someone who is in the same room as you for items, inspecting the room for items, and moving to a different room.",
            2: "[a] accuse: accusing someone lists everyone in the game to accuse. After being accused you will find out who the murderer was. You can only accuse ONCE per game.",
            3: "[i] inspect: inspecting the room will tell you if there are any items within the room. If there were no items to begin with 'No items will be found.' If an item was taken from the room, that item will be mentioned and you will notice that it is no longer there.",
            4: "[s] search: searching will list everyone in the room with you. Once you select who you want to search, you will learn what possessions they had on them.",
            5: "[m] move: moving will list all of the rooms adjacent to your room. Before every turn you will see a map at the top with all of the rooms, and it will state what room you are currently in.",
        }
        neverAccusedText := map[int]string{
            0: "After " + strconv.Itoa(max_turns) + " turns you never accused anyone.",
            1: "As a result an innocent person's life is on your hands.",
            2: "And the murderer gets to walk away free. No one will know who they were...",
        }

        // #############################################
        // #       CREATING INITIAL VARIABLES          #
        // #############################################
        var people map[string]map[string]string
        var rooms map[string]map[string]string
        currentRoom := randomRoom()
        var peopleInRoom map[int]string

        listPeople := [14]string{"ELI", "RACHEL", "AARON", "MAK", "MATTHEW", "ANJALI", "ANIKA", "ALEC", "CHRISTOPHER", "TANISHA", "GIDEON", "DANIEL", "EDDIE", "KEVIN"}
        listPeople = randomizeList(listPeople)
        listRooms := [6]string{"LIVING ROOM", "CELLAR", "KITCHEN", "DINING ROOM", "LIBRARY", "COURTYARD"}
        listRooms = randomizeRooms(listRooms)

        game := rand.Intn(3)
        turn := 1
        accused := false

        people = createPeople(game, listPeople)
        rooms = createMap(listRooms)

        // Creating list rooms into order in which items are held
        temp, _ := strconv.Atoi(rooms["KITCHEN"]["id"])
        listRooms[temp] = "KITCHEN"
        temp, _ = strconv.Atoi(rooms["LIVING ROOM"]["id"])
        listRooms[temp] = "LIVING ROOM"
        temp, _ = strconv.Atoi(rooms["DINING ROOM"]["id"])
        listRooms[temp] = "DINING ROOM"
        temp, _ = strconv.Atoi(rooms["LIBRARY"]["id"])
        listRooms[temp] = "LIBRARY"
        temp, _ = strconv.Atoi(rooms["COURTYARD"]["id"])
        listRooms[temp] = "COURTYARD"
        temp, _ = strconv.Atoi(rooms["CELLAR"]["id"])
        listRooms[temp] = "CELLAR"
        fmt.Println(listRooms)

        if (settings["tutorial"] == "ON") {
            fmt.Println("\033[2J")
            printTextBox(tutorialText)
            fmt.Println("\033[2J")
            printTextBox(introText)
        }

        for turn <= max_turns && accused == false {
            fmt.Println("\033[2J")
            people = peopleMovement(people, rooms, listRooms, listPeople, turn, game)
            peopleInRoom = inRoom(people, listPeople, currentRoom)
            currentRoom = takeTurn(currentRoom, people, listPeople, peopleInRoom, rooms, listRooms, turn, game)

            if (currentRoom != "LIBRARY" && currentRoom != "KITCHEN" && currentRoom != "LIVING ROOM" && currentRoom != "DINING ROOM" && currentRoom != "COURTYARD" && currentRoom != "CELLAR") {
                accused = true
            }

            turn++
        }

        if (turn >= max_turns && accused == false) {
            printTextBox(neverAccusedText)
            play = false
        }
        if (accused == true && currentRoom == "true") {
            solvedText := map[int]string{
                0: "You accused " + listPeople[0] + ".",
                1: "After the accusation " + listPeople[0] + " lunged at you. Before they reached you they were tackled to the floor.",
                2: listPeople[0] + " went to court and pleaded guilty for attempted assault and attempted murder.",
                3: "You saved a mans life today. You should be proud.",
            }
            printTextBox(solvedText)
            play = false
        } else if (accused == true) {
            killedText := ""
            if (game == 0) {
                killedText = "stabbed"
            } else if (game == 1) {
                killedText = "poisoned"
            } else {
                killedText = "shot"
            }

            solvedText := map[int]string{
                0: "You accused " + currentRoom + ".",
                1: "After the accusation, " + currentRoom + " bewildered, " + currentRoom + " was arrested.",
                2: "Later that night " + listPeople[0] + " " + killedText + " " + listPeople[13] + ". " + listPeople[13] + " was sent to the hospital, but died that night.",
            }
            printTextBox(solvedText)
            play = false
        }
    }
}

// #############################################
// #         PLAYER TURNS AND OPTIONS          #
// #############################################
func takeTurn(current_room string, people map[string]map[string]string, listPeople [14]string, peopleInRoom map[int]string, rooms map[string]map[string]string, listRooms [6]string, current_turn int, game int) string {
    startSentence := "You are currently in the " + current_room + "."
    for i := 0; i < len(peopleInRoom); i++ {
        if (len(peopleInRoom) == 1) {
            startSentence = startSentence + " In the room with you is " + peopleInRoom[i] + "."
        } else if (len(peopleInRoom) == 2) {
            startSentence = startSentence + " In the room with you is " + peopleInRoom[i] + " and " + peopleInRoom[i + 1] + "."
            break
        } else {
            if (i == 0) {
                startSentence = startSentence + " In the room with you is " + peopleInRoom[i] + ", "
            } else if (i != len(peopleInRoom) - 1) {
                startSentence = startSentence + peopleInRoom[i] + ", "
            } else {
                startSentence = startSentence + "and " + peopleInRoom[i] + "."
            }
        }
    }
    if (len(peopleInRoom) == 0) {
        startSentence = startSentence + " No one else is in the room with you."
    }

    if (current_turn == max_turns) {
        startSentence = startSentence + " There is " + strconv.Itoa(max_turns + 1 - current_turn) + " turn left."
    } else {
        startSentence = startSentence + " There are " + strconv.Itoa(max_turns + 1 - current_turn) + " turns left."
    }

    sequence_rooms_1 := ""
    sequence_rooms_2 := ""

    for i := 0; i < len(listRooms); i++ {
        if (i <= 2) {
            if (listRooms[i] == "LIBRARY") {
                sequence_rooms_1 = sequence_rooms_1 + "[ LI ]"
            } else if (listRooms[i] == "LIVING ROOM") {
                sequence_rooms_1 = sequence_rooms_1 + "[ LR ]"
            } else if (listRooms[i] == "COURTYARD") {
                sequence_rooms_1 = sequence_rooms_1 + "[ CY ]"
            } else if (listRooms[i] == "DINING ROOM") {
                sequence_rooms_1 = sequence_rooms_1 + "[ DR ]"
            } else if (listRooms[i] == "CELLAR") {
                sequence_rooms_1 = sequence_rooms_1 + "[ CE ]"
            } else if (listRooms[i] == "KITCHEN") {
                sequence_rooms_1 = sequence_rooms_1 + "[ KI ]"
            }
        } else {
            if (listRooms[i] == "LIBRARY") {
                sequence_rooms_2 = sequence_rooms_2 + "[ LI ]"
            } else if (listRooms[i] == "LIVING ROOM") {
                sequence_rooms_2 = sequence_rooms_2 + "[ LR ]"
            } else if (listRooms[i] == "COURTYARD") {
                sequence_rooms_2 = sequence_rooms_2 + "[ CY ]"
            } else if (listRooms[i] == "DINING ROOM") {
                sequence_rooms_2 = sequence_rooms_2 + "[ DR ]"
            } else if (listRooms[i] == "CELLAR") {
                sequence_rooms_2 = sequence_rooms_2 + "[ CE ]"
            } else if (listRooms[i] == "KITCHEN") {
                sequence_rooms_2 = sequence_rooms_2 + "[ KI ]"
            }
        }
    }

    var turn map[int]string
    if (settings["difficulty"] == "EASY") {
        turn = map[int]string{
            0: "Map:",
            1: sequence_rooms_1,
            2: sequence_rooms_2,
            3: "",
            4: startSentence,
            5: "What would you like to do?",
            6: "[a] Accuse someone",
            7: "[i] Inspect the room",
            8: "[s] Search someone in the room",
            9: "[m] Move to a different room",
        }
    } else {
        turn = map[int]string{
            0: startSentence,
            1: "What would you like to do?",
            2: "[a] Accuse someone",
            3: "[i] Inspect the room",
            4: "[s] Search someone in the room",
            5: "[m] Move to a different room",
        }
    }
    option := printOption(turn)
    for true {
        if (option == "a") {
            didAccuse := accuse(listPeople)
            if (didAccuse != "c") {
                if didAccuse == listPeople[0] {
                    return "true"
                } else {
                    for i := 0; i < len(listPeople); i++ {
                        if (listPeople[i] == didAccuse) {
                            return didAccuse
                        }
                    }
                }
            }
        } else if (option == "i") {
            inspect(rooms, listRooms, current_room, current_turn, game)
            break
        } else if (option == "s") {
            if (search(people, peopleInRoom) != "c") {
                break
            }
        } else if (option == "m") {
            moveReturn := move(rooms, current_room)
            if(moveReturn != "c") {
                current_room = moveReturn
                break
            }
        }

        option = printOption(turn)
    }

    return current_room
}

func inRoom(people map[string]map[string]string, listPeople [14]string, current_room string) map[int]string {
    var peopleInRoom = make(map[int]string)
    randomizedList := randomizeList(listPeople)
    num := 0

    for i := 0; i < len(people); i++ {
        if (people[randomizedList[i]]["current_room"] == current_room) {
            peopleInRoom[num] = randomizedList[i]
            num++
        }
    }

    return peopleInRoom
}

// #############################################
// #              PLAYER ACTIONS               #
// #############################################
func accuse(listPeople [14]string) string {
    inputText := make(map[int]string)
    textCounter := 0
    inputText[textCounter] = "Who would you like to accuse of being the murderer?"
    textCounter++

    newOrder := listPeople
    for i := 0; i < 7; i++ {
        newOrder = randomizeList(newOrder)
    }

    for i := 0; i < len(newOrder); i++ {
        inputText[textCounter] = "[" + newOrder[i] + "]"
        textCounter++
    }
    inputText[textCounter] = "[c] Cancel"

    option := printOption(inputText)
    for true {
        if (option == "c") {
            return "c"
        }

        for i := 0; i < len(listPeople); i++ {
            if (option == listPeople[i]) {
                return listPeople[i]
            }
        }

        option = printOption(inputText)
    }
    return "c"
}

func move(rooms map[string]map[string]string, currentRoom string) string {
    textCount := 0
    text := make(map[int]string)

    text[textCount] = "Where would you like to move?"
    textCount++
    if (rooms[currentRoom]["up"] != "") {
        text[textCount] = "[u] Move up to the " + rooms[currentRoom]["up"]
        textCount++
    }
    if (rooms[currentRoom]["down"] != "") {
        text[textCount] = "[d] Move down to the " + rooms[currentRoom]["down"]
        textCount++
    }
    if (rooms[currentRoom]["left"] != "") {
        text[textCount] = "[l] Move left to the " + rooms[currentRoom]["left"]
        textCount++
    }
    if (rooms[currentRoom]["right"] != "") {
        text[textCount] = "[r] Move right to the " + rooms[currentRoom]["right"]
        textCount++
    }
    text[textCount] = "[c] Cancel"

    option := printOption(text)
    for true {
        if (option == "u" && rooms[currentRoom]["up"] != "") {
            return rooms[currentRoom]["up"]
        } else if (option == "d" && rooms[currentRoom]["down"] != "") {
            return rooms[currentRoom]["down"]
        } else if (option == "l" && rooms[currentRoom]["left"] != "") {
            return rooms[currentRoom]["left"]
        } else if (option == "r" && rooms[currentRoom]["right"] != "") {
            return rooms[currentRoom]["right"]
        } else if (option == "c") {
            return "c"
        }
        option = printOption(text)
    }
    return "c"
}

func inspect(rooms map[string]map[string]string, listRooms [6]string, currentRoom string, currentTurn int, game int) {
    inspectText := make(map[int]string)
    if (rooms[currentRoom]["item"] == "" && rooms[currentRoom]["taken"] == "") {
        inspectText[0] = "You search the " + currentRoom + ", but nothing was found."
        printTextBox(inspectText)
    } else if (currentRoom == "LIBRARY" && rooms[currentRoom]["taken"] == "") {
        inspectText[0] = "You search the " + currentRoom + ", and under the bookshelf, you open a drawer. Inside, you found a " + rooms[currentRoom]["item"] + "."
        printTextBox(inspectText)
    } else if (currentRoom == "CELLAR" && rooms[currentRoom]["taken"] == "") {
        inspectText[0] = "You search the " + currentRoom + ", and in the corner of the room, you found a " + rooms[currentRoom]["item"] + "."
        printTextBox(inspectText)
    } else if (currentRoom == "KITCHEN" && rooms[currentRoom]["taken"] == "") {
        inspectText[0] = "You search the " + currentRoom + ", and on the counter you found a " + rooms[currentRoom]["item"] + "."
        printTextBox(inspectText)
    } else if (currentRoom == "LIVING ROOM" && rooms[currentRoom]["taken"] == "") {
        inspectText[0] = "You search the " + currentRoom + ", and in the middle of the room, you look upon a table and found a " + rooms[currentRoom]["item"] + "."
        printTextBox(inspectText)
    } else if (currentRoom == "DINING ROOM" && rooms[currentRoom]["taken"] == "") {
        inspectText[0] = "You search the " + currentRoom + ", and on the table you found a " + rooms[currentRoom]["item"] + "."
        printTextBox(inspectText)
    } else if (currentRoom == "COURTYARD" && rooms[currentRoom]["taken"] == "") {
        inspectText[0] = "You search the " + currentRoom + ", and on the ground, by the steps you found a " + rooms[currentRoom]["item"] + "."
        printTextBox(inspectText)
    } else {
        inspectText[0] = "You inspect the " + currentRoom + ", but find nothing. But it seems like a " + rooms[currentRoom]["taken"] + " was taken from the room before you came."
        if (settings["hints"] == "ON") {
            if (game == 0 && currentTurn >= max_turns - 2) {
                if (listRooms[5] == "COURTYARD") {
                    inspectText[1] = "HINT: You notice traces of mud on the ground, soon your blood will be spread around."
                } else if (listRooms[5] == "KITCHEN") {
                    inspectText[1] = "HINT: You've upset the cook, now to the knives we shall look."
                } else if (listRooms[5] == "LIVING ROOM") {
                    inspectText[1] = "HINT: The knife exposes this room of lies, as you will soon meet your demise."
                } else if (listRooms[5] == "DINING ROOM") {
                    inspectText[1] = "HINT: In this room it may be used for decor, but don't be fooled I will still slice you to your core."
                } else if (listRooms[5] == "LIBRARY") {
                    inspectText[1] = "HINT: Knowing the ins and outs of every trade. Don't get stabbed where knowledge is laid."
                } else if (listRooms[5] == "CELLAR") {
                    inspectText[1] = "HINT: Sometimes there is no way to see, but with the knife you will bleed. I'll meet you down here in the ____ and I'll become your fortune teller"
                }
            } else if (game == 1 && currentTurn > max_turns / 3 + 4) {
                if (listRooms[1] == "KITCHEN") {
                    inspectText[1] = "HINT: You think to yourself... if I were to poison someone where would I poison their food?"
                } else if (listRooms[1] == "LIBRARY") {
                    inspectText[1] = "HINT: Where would they have read about poison?"
                } else if (listRooms[1] == "DINING ROOM") {
                    inspectText[1] = "HINT: The cutlery is fine, go to the area where you may..."
                } else if (listRooms[1] == "COURTYARD") {
                    inspectText[1] = "HINT: A glass of wine in fresh air seems like a good place to poison a drink."
                } else if (listRooms[1] == "CELLAR") {
                    inspectText[1] = "HINT: Where is the wine usually stored?"
                } else if (listRooms[1] == "LIVING ROOM") {
                    inspectText[1] = "HINT: Take a sip, and you will say goodbye. Enter this room and you won't be alive."
                }
            } else if (game == 2  && currentTurn > max_turns / 5 + 7 && currentTurn < max_turns * 3 / 7 + 4) {
                inspectText[1] = "HINT: A spare bullet seems to be lying around somewhere..."
            } else if (game == 2  && currentTurn > max_turns * 18 / 25 && currentTurn < max_turns * 21 / 25) {
                if (listRooms[2] == "CELLAR") {
                    inspectText[1] = "HINT: You read a note which states: 'You are lavish with your critique, as you are quite the storyteller. You will meet your demise soon, down in the...'"
                } else if (listRooms[2] == "LIVING ROOM") {
                    inspectText[1] = "HINT: You read a note which states: 'You are lavish with your critique, but soon you won't be living and you'll go boom. You will meet your demise soon, down in the...'"
                } else if (listRooms[2] == "LIBRARY") {
                    inspectText[1] = "HINT: You read a note which states: 'You are quite well read, but I don't think I'm going to need to use my words. You'll meet your end in the...'"
                } else if (listRooms[2] == "DINING ROOM") {
                    inspectText[1] = "HINT: You read a note which states: 'I'm sure people will love the view of your head when they eat. We'll meet in...'"
                }  else if (listRooms[2] == "KITCHEN") {
                    inspectText[1] = "HINT: You read a note which states: 'You'll be prepared medium rare when you meat your end in the...'"
                }  else if (listRooms[2] == "COURTYARD") {
                    inspectText[1] = "HINT: You read a note which states: 'Ahh a duel, working in the mud. You'll meet your end where you belo...'"
                }
                inspectText[2] = "HINT: Part of the note seems to be torn off."
            } 
        }
        printTextBox(inspectText)
    }
}

func search(people map[string]map[string]string, peopleInRoom map[int]string) string {
    textCount := 0
    text := make(map[int]string)
    itemsText := make(map[int]string)

    text[textCount] = "Who would you like to search?"
    textCount++
    for i := 0; i < len(peopleInRoom); i++ {
        text[textCount] = "[" + peopleInRoom[i] + "]"
        textCount++
    }
    text[textCount] = "[c] Cancel"

    option := printOption(text)
    for true {
        for i := 0; i < len(peopleInRoom); i++ {
            if (option == peopleInRoom[i]) {
                if (people[peopleInRoom[i]]["item"] == "") {
                    itemsText[0] = "You searched " + peopleInRoom[i] + ". You found nothing, but " +  peopleInRoom[i] + " seemed annoyed."
                } else if (people[peopleInRoom[i]]["role"] == "decoy_poison") {
                    itemsText[0] = "You searched " + peopleInRoom[i] + ". You found a " + people[peopleInRoom[i]]["item"] + "."
                    itemsText[1] = peopleInRoom[i] + " says, 'I'm really sorry for the confusion, but I was just trying to help out the kitchen. We've had a bad rat problem recently, and they told us to apply it every day.'"
                } else if (people[peopleInRoom[i]]["role"] == "decoy_pew_1" || people[peopleInRoom[i]]["role"] == "decoy_pew_2") {
                    itemsText[0] = "You searched " + peopleInRoom[i] + ". You found a " + people[peopleInRoom[i]]["item"] + "."
                    itemsText[1] = peopleInRoom[i] + " says, 'Well it seems you caught me... I've always envied this rifle, and well... I was hoping I could use it on my next hunting trip in Canada. I know how this looks.'"
                } else if (people[peopleInRoom[i]]["role"] == "decoy_stab") {
                    itemsText[0] = "You searched " + peopleInRoom[i] + ". You found a " + people[peopleInRoom[i]]["item"] + "."
                    itemsText[1] = peopleInRoom[i] + " says, 'I'm almost done preparing the food! We will start dinner shortly.'"
                } else {
                    itemsText[0] = "You searched " + peopleInRoom[i] + ". You found a " + people[peopleInRoom[i]]["item"] + "."
                }
                printTextBox(itemsText)
                return ""
            }
        }
        if (option == "c") {
            return "c"
        }
        option = printOption(text)
    }

    return "c"
}

// #############################################
// #        PEOPLE AND PLAYER MOVEMENT         #
// #############################################
func peopleMovement(people map[string]map[string]string, rooms map[string]map[string]string, listRooms [6]string, listPeople [14]string, turn int, game int) map[string]map[string]string {
    for i := 0; i < len(people); i++ {
        if (game == 0 && people[listPeople[i]]["role"] == "killer") {
            if (turn >= max_turns / 10 && turn <= max_turns / 5) {
                people[listPeople[i]]["current_room"] = listRooms[5]
                continue
            } else if (turn == max_turns / 2) {
                people[listPeople[i]]["current_room"] = listRooms[1]
                people[listPeople[i]]["item"] = rooms[listRooms[1]]["item"]
                rooms[listRooms[1]]["taken"] = rooms[listRooms[1]]["item"]
                rooms[listRooms[1]]["item"] = ""
                continue
            } else if (turn > max_turns / 2 && turn <= max_turns / 2 + 4) {
                people[listPeople[i]]["current_room"] = listRooms[1]
                continue
            } else if (turn >= max_turns - 2 && turn <= max_turns) {
                people[listPeople[i]]["current_room"] = listRooms[5]
                continue
            }
        } else if (game == 1 && people[listPeople[i]]["role"] == "killer") {
            if (turn == max_turns / 3) {
                people[listPeople[i]]["current_room"] = listRooms[2]
                people[listPeople[i]]["item"] = rooms[listRooms[2]]["item"]
                rooms[listRooms[2]]["taken"] = rooms[listRooms[2]]["item"]
                rooms[listRooms[2]]["item"] = ""
                continue
            } else if (turn > max_turns / 3 && turn <= max_turns / 3 + 4) {
                people[listPeople[i]]["current_room"] = listRooms[2]
                continue
            } else if (turn >= max_turns * 2 / 3 && turn <= max_turns * 2 / 3 + 3) {
                people[listPeople[i]]["current_room"] = listRooms[1]
                continue
            } else if (turn >= max_turns - 4 && turn <= max_turns) {
                people[listPeople[i]]["current_room"] = listRooms[4]
                continue
            }
        } else if (game == 2 && people[listPeople[i]]["role"] == "killer") {
            if (turn == max_turns / 5) {
                people[listPeople[i]]["current_room"] = listRooms[0]
                people[listPeople[i]]["item"] = rooms[listRooms[0]]["item"]
                rooms[listRooms[0]]["taken"] = rooms[listRooms[0]]["item"]
                rooms[listRooms[0]]["item"] = ""
                continue
            } else if (turn > max_turns / 5 && turn <= max_turns / 5 + 7) {
                people[listPeople[i]]["current_room"] = listRooms[0]
                continue
            } else if (turn == max_turns * 3 / 7) {
                people[listPeople[i]]["current_room"] = listRooms[3]
                people[listPeople[i]]["item"] = rooms[listRooms[3]]["item"]
                rooms[listRooms[3]]["taken"] = rooms[listRooms[3]]["item"]
                rooms[listRooms[3]]["item"] = ""
                continue
            } else if (turn > max_turns * 3 / 7 && turn <= max_turns * 3 / 7 + 4) {
                people[listPeople[i]]["current_room"] = listRooms[3]
                continue
            } else if (turn >= max_turns * 18 / 25 && turn <= max_turns * 21 / 25) {
                people[listPeople[i]]["current_room"] = listRooms[2]
                continue
            }
        } 
        if (settings["difficulty"] == "HARD" && people[listPeople[i]]["role"] == "decoy_poison") {
            if (turn < max_turns / 5) {
                people[listPeople[i]]["current_room"] = listRooms[2]
                continue
            } else if (turn == max_turns / 5 + 1) {
                people[listPeople[i]]["current_room"] = listRooms[2]
                people[listPeople[i]]["item"] = rooms[listRooms[2]]["item"]
                rooms[listRooms[2]]["taken"] = rooms[listRooms[2]]["item"]
                rooms[listRooms[2]]["item"] = ""
                continue
            } else if (turn > max_turns / 5 && turn <= max_turns / 5 + 7) {
                people[listPeople[i]]["current_room"] = listRooms[1]
                continue
            } else if (turn > max_turns / 2) {
                people[listPeople[i]]["current_room"] = listRooms[2]
                people[listPeople[i]]["item"] = ""
                rooms[listRooms[2]]["item"] = rooms[listRooms[2]]["taken"]
                rooms[listRooms[2]]["taken"] = ""
                continue
            }
        } else if (settings["difficulty"] == "HARD" && people[listPeople[i]]["role"] == "decoy_pew_1") {
            if (turn < max_turns / 10) {
                people[listPeople[i]]["current_room"] = listRooms[0]
                continue
            } else if (turn == max_turns / 10) {
                people[listPeople[i]]["current_room"] = listRooms[0]
                people[listPeople[i]]["item"] = rooms[listRooms[0]]["item"]
                rooms[listRooms[0]]["taken"] = rooms[listRooms[0]]["item"]
                rooms[listRooms[0]]["item"] = ""
                continue
            }
        } else if (settings["difficulty"] == "HARD" && people[listPeople[i]]["role"] == "decoy_pew_2") {
            if (turn > max_turns - 8 && turn < max_turns - 4) {
                people[listPeople[i]]["current_room"] = listRooms[0]
                continue
            } else if (turn == max_turns - 4) {
                people[listPeople[i]]["current_room"] = listRooms[0]
                people[listPeople[i]]["item"] = rooms[listRooms[0]]["item"]
                rooms[listRooms[0]]["taken"] = rooms[listRooms[0]]["item"]
                rooms[listRooms[0]]["item"] = ""
                continue
            } else if (turn > max_turns - 2) {
                people[listPeople[i]]["current_room"] = listRooms[5]
                continue
            }
        } else if (settings["difficulty"] == "HARD" && people[listPeople[i]]["role"] == "decoy_stab") {
            if (turn > max_turns / 2 && turn < max_turns / 2 + 1) {
                people[listPeople[i]]["current_room"] = listRooms[1]
                continue
            } else if (turn == max_turns / 2 + 1) {
                people[listPeople[i]]["current_room"] = listRooms[1]
                people[listPeople[i]]["item"] = rooms[listRooms[1]]["item"]
                rooms[listRooms[1]]["taken"] = rooms[listRooms[1]]["item"]
                rooms[listRooms[1]]["item"] = ""
                continue
            } else if (turn > max_turns / 2 + 1 && turn <= max_turns / 2 + 4) {
                people[listPeople[i]]["current_room"] = listRooms[1]
                continue
            } else if (turn > max_turns / 2 + 4) {
                people[listPeople[i]]["current_room"] = listRooms[1]
                people[listPeople[i]]["item"] = ""
                rooms[listRooms[1]]["item"] = rooms[listRooms[1]]["taken"]
                rooms[listRooms[1]]["taken"] = ""
                continue
            }
        }

        if (people[listPeople[i]]["leave_room"] != "") {
            leave_room, _ := strconv.Atoi(people[listPeople[i]]["leave_room"])
            if (leave_room > 0) {
                leave_room--
                people[listPeople[i]]["leave_room"] = strconv.Itoa(leave_room)
            } else {
                people[listPeople[i]]["current_room"] = randomRoom()
                leave_room = rand.Intn(4) + 2
                people[listPeople[i]]["leave_room"] = strconv.Itoa(leave_room)
            }
        } else {
            people[listPeople[i]]["current_room"] = randomRoom()
            leave_room := rand.Intn(4) + 2
            people[listPeople[i]]["leave_room"] = strconv.Itoa(leave_room)
        }
    }

    return people
}
func randomRoom() string {
    roomNum := rand.Intn(6)

    if roomNum == 0 {
        return "KITCHEN"
    } else if roomNum == 1 {
        return "DINING ROOM"
    } else if roomNum == 2 {
        return "COURTYARD"
    } else if roomNum == 3 {
        return "LIBRARY"
    } else if roomNum == 4 {
        return "LIVING ROOM"
    } else {
        return "CELLAR"
    }
}

// #############################################
// #             CREATING THE MAP              #
// #############################################
func createMap(listRooms [6]string) map[string]map[string]string {
    rooms := map[string]map[string]string{
        "LIBRARY": {
            "id": "0",
            "item": "PISTOL",
            "taken": "",
            "up": "",
            "down": "",
            "left": "",
            "right": "",
        },
        "KITCHEN": {
            "id": "1",
            "item": "KNIFE",
            "taken": "",
            "up": "",
            "down": "",
            "left": "",
            "right": "",
        },
        "COURTYARD": {
            "id": "5",
            "item": "",
            "taken": "",
            "up": "",
            "down": "",
            "left": "",
            "right": "",
        },
        "CELLAR": {
            "id": "2",
            "item": "BOTTLE OF RAT POISON",
            "taken": "",
            "up": "",
            "down": "",
            "left": "",
            "right": "",
        },
        "DINING ROOM": {
            "id": "4",
            "item": "",
            "taken": "",
            "up": "",
            "down": "",
            "left": "",
            "right": "",
        },
        "LIVING ROOM": {
            "id": "3",
            "item": "SPARE BULLET",
            "taken": "",
            "up": "",
            "down": "",
            "left": "",
            "right": "",
        },
    }

    if (settings["difficulty"] != "EASY") {
        rooms[listRooms[0]]["item"] = "PISTOL"
        rooms[listRooms[0]]["id"] = "0"
        rooms[listRooms[1]]["item"] = "KNIFE"
        rooms[listRooms[1]]["id"] = "1"
        rooms[listRooms[2]]["item"] = "BOTTLE OF RAT POISON"
        rooms[listRooms[2]]["id"] = "2"
        rooms[listRooms[3]]["item"] = "SPARE BULLET"
        rooms[listRooms[3]]["id"] = "3"
        rooms[listRooms[4]]["item"] = ""
        rooms[listRooms[4]]["id"] = "4"
        rooms[listRooms[5]]["item"] = ""
        rooms[listRooms[5]]["id"] = "5"
    }

    for i := 0; i < len(listRooms); i++ {
        listRooms = randomizeRooms(listRooms)
    }

    // Setting directions for all of the rooms
    // Upper left room
    rooms[listRooms[0]]["right"] = listRooms[1]
    rooms[listRooms[0]]["down"] = listRooms[3]

    // Upper middle room
    rooms[listRooms[1]]["right"] = listRooms[2]
    rooms[listRooms[1]]["left"] = listRooms[0]
    rooms[listRooms[1]]["down"] = listRooms[4]

    // Upper right room
    rooms[listRooms[2]]["left"] = listRooms[1]
    rooms[listRooms[2]]["down"] = listRooms[5]

    // Bottom left room
    rooms[listRooms[3]]["right"] = listRooms[4]
    rooms[listRooms[3]]["up"] = listRooms[0]

    // Bottom middle room
    rooms[listRooms[4]]["right"] = listRooms[5]
    rooms[listRooms[4]]["left"] = listRooms[3]
    rooms[listRooms[4]]["up"] = listRooms[1]

    // Bottom right room
    rooms[listRooms[5]]["left"] = listRooms[4]
    rooms[listRooms[5]]["up"] = listRooms[2]

    return rooms
}

func randomizeRooms(listRooms [6]string) [6]string{
    for i := 0; i < len(listRooms); i++ {
        randomIndex := rand.Intn(len(listRooms))
        temp := listRooms[i]
        listRooms[i] = listRooms[randomIndex]
        listRooms[randomIndex] = temp
    }
    // fmt.Println(listRooms)
    return listRooms
}

// #############################################
// #      CREATING THE ROLES AND PEOPLE        #
// #############################################
func createPeople(game int, listPeople [14]string) map[string]map[string]string {
    people := map[string]map[string]string{
        "ELI": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "RACHEL": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "AARON": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "MAK": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "MATTHEW": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "ANJALI": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "ANIKA": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "ALEC": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "CHRISTOPHER": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "TANISHA": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "GIDEON": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "DANIEL": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "EDDIE": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
        "KEVIN": {
            "role": "",
            "item": "",
            "current_room": "",
            "leave_room": "",
        },
    }

    people[listPeople[0]]["role"] = "killer"
    if (game == 0) {
        people[listPeople[1]]["role"] = "decoy_poison"
        if (rand.Intn(2) == 0) {
            people[listPeople[2]]["role"] = "decoy_pew_1"
        } else {
            people[listPeople[2]]["role"] = "decoy_pew_2"
        }
    } else if (game == 1) {
        if (rand.Intn(2) == 0) {
            people[listPeople[1]]["role"] = "decoy_pew_1"
        } else {
            people[listPeople[1]]["role"] = "decoy_pew_2"
        }
        people[listPeople[2]]["role"] = "decoy_stab"
    } else {
        people[listPeople[1]]["role"] = "decoy_stab"
        people[listPeople[2]]["role"] = "decoy_poison"
    }

    for i := 0; i < len(listPeople); i++ {
        if (people[listPeople[i]]["role"] == "") {
            people[listPeople[i]]["role"] = "innocent"
        }
    }

    return people
}

func randomizeList(listPeople [14]string) [14]string{
    for i := 0; i < len(listPeople); i++ {
        randomIndex := rand.Intn(len(listPeople))
        temp := listPeople[i]
        listPeople[i] = listPeople[randomIndex]
        listPeople[randomIndex] = temp
    }
    // fmt.Println(listPeople)
    return listPeople
}

// #############################################
// #      CREATING TEXT WITHIN CONSOLE         #
// #############################################
func printTextBox(text map[int]string) {
    var lines [][]string
    var dashes string = ""
    for i := 0; i < size_line + 3; i++ {
        dashes += "-"
    }

    for i := 0; i < len(text); i++ {
        lines = append(lines, createLines(text[i], size_line))
    }
    for i := 0; i < len(lines); i++ {
        fmt.Println(dashes)
        for a := 0; a < len(lines[i]); a++ {
            fmt.Printf("| %-75s|\n", (lines[i][a])) //Adjust when changing line size
        }
        fmt.Println(dashes)
        fmt.Print("[ENTER TO CONTINUE]")
        fmt.Scanln()
    }
}
func printOption(text map[int]string) string {
    var lines [][]string
    var output string
    var dashes string = ""

    for i := 0; i < size_line + 3; i++ {
        dashes += "-"
    }
    for i := 0; i < len(text); i++ {
        lines = append(lines, createLines(text[i], size_line))
    }

    fmt.Println(dashes)
    for i := 0; i < len(lines); i++ {
        for a := 0; a < len(lines[i]); a++ {
            fmt.Printf("| %-75s|\n", (lines[i][a])) //Adjust when changing line size
        }
    }
    fmt.Println(dashes)
    fmt.Print("> ")
    fmt.Scanln(&output)
    return output
}

func createLines(text string, size int) []string {
    var words []string = strings.Split(text, " ")
    var lines []string
    var line string = ""

    for i := 0; i < len(words); i++ {
        if (len(line) + len(words[i]) + 1 <= size && i != len(words) - 1) {
            line += words[i] + " "
        } else if (i == len(words) - 1 && len(line) + len(words[i]) + 1 <= size) {
            line += words[i] + " "
            lines = append(lines, line)
            line = ""
        } else if (i == len(words) - 1 && len(line) + len(words[i]) + 1 >= size) {
            lines = append(lines, line)
            line = ""
            line += words[i] + " "
            lines = append(lines, line)
            line = ""
        } else {
            lines = append(lines, line)
            line = ""
            line += words[i] + " "
        }
    }

    return lines
}
