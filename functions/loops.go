package functions

func InitLoops() {

	loopTillConditionIsMet()

	traditionalForLoop()

	infiniteLoop()

	loopCollectionWithIndex()

	loopCollectionWithRange()

	loopMaps()

}

func loopMaps() {
	configs := map[string]string{
		"Env":  "DEV",
		"Port": "3000",
		"host": "localhost",
	}

	for key, val := range configs {
		println(key, "=", val)
	}

	// Loop Map Keys only
	for key := range configs {
		println(key)
	}

	// Loop Map Values
	for _, val := range configs {
		println(val)
	}

}

func loopCollectionWithRange() {
	names := []string{"Imran", "Aftab", "Rana"}
	for index, name := range names {
		println("Name at", index, " is", name)
	}
}

func loopCollectionWithIndex() {
	names := []string{"Imran", "Aftab", "Rana"}
	for index := 0; index < len(names); index++ {
		println(names[index])
	}
}

func infiniteLoop() {
	breakInfiniteLoop := false
	for { // similar to while(true)
		if breakInfiniteLoop == false {
			println("Infinit loop")
			breakInfiniteLoop = true
		} else {
			println("Ending infinite loop")
			break
		}

	}
}

func traditionalForLoop() {
	var count int
	for ; count < 10; count++ {
		println("count", count)
		if count == 5 {
			println("skipping")
			continue
		} else if count > 8 {
			println("breaking")
			break
		}
	}

	// Implicit initialization
	for i := 0; i < 10; i++ {
		println("Implicit For: count", i)
		if i == 5 {
			println("skipping")
			continue
		} else if i > 8 {
			println("breaking")
			break
		}
	}

}

func loopTillConditionIsMet() {
	var count int
	for count < 10 {
		if count == 5 {
			println("Skipping ", count)
			count++
			continue
		}

		if count > 8 {
			println("Breaking here.")
			break
		}
		println("count", count)
		count++
	}
}
