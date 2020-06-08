package functions

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type ReadFiles struct {

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ChargeBotmaps()  {

	fmt.Println("-----------------")

	//
	f, err := os.Open("./botmaps/index.botmap")
	defer f.Close()

	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------")

}
