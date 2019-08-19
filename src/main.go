package main

import "poker1/src/poker_service"
func main()  {
	poker_service.NewStartPoker().StartPoker("./source/match.json")
}
