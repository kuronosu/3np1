package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kuronosu/3np1/tnpo"
)

func next(n int) int {
	if n == 1 {
		return -1
	}
	var next_n int
	if n%2 == 0 {
		next_n = (n / 2)
	} else {
		next_n = 3*n + 1
	}
	return next_n
}

func generate(tree *tnpo.Tree, n int) *tnpo.Node {
	if n == -1 {
		return nil
	}
	if tree.Contains(n) {
		return tree.GetNode(n)
	}
	parent := generate(tree, next(n))
	return tree.CreateNode(n, parent)
}

func generateTree(n int) *tnpo.Tree {
	defer timeTrack(time.Now(), "Generate Tree")
	tree := tnpo.NewTree()
	for i := 1; i <= n; i++ {
		generate(tree, i)
	}

	return tree
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func GetFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func save(tree tnpo.Tree, n int, r int) {
	defer timeTrack(time.Now(), "Save")
	s := n / r
	ch := make(chan int, r)
	for i := 0; i < r; i++ {
		go func(ii int, ss int, cch chan int) {
			file, err := GetFile(fmt.Sprint("results-", ii, ".txt"))
			if err != nil {
				log.Fatal(err)
			} else {
				defer file.Close()
				m := int(ss * ii)
				for j := m + 1; j < m+ss+1; j++ {
					if _, err := file.Write([]byte(fmt.Sprint(tree.GetNode(j).DataTraceUp()) + "\n")); err != nil {
						log.Fatal(err)
					}
				}
			}
			cch <- 1
		}(i, s, ch)
	}
	results := 0
	for {
		results += <-ch
		if results == r {
			break
		}
	}
}

func deletePrevFiles() {
	defer timeTrack(time.Now(), "Delete Prev Files")
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasPrefix(f.Name(), "results-") && strings.HasSuffix(f.Name(), ".txt") {
			os.Remove(f.Name())
		}
	}

}

func main() {
	n := 10_000_000
	fc := 20
	fmt.Println("Start", n, fc)
	deletePrevFiles()
	save(*generateTree(n), n, fc)
}
