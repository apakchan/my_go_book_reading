package main

import "fmt"

type item struct {
	name string
}

func (i *item) use() {
	if i == nil {
		fmt.Printf("no item to use!\n")
		return
	}
	fmt.Printf("use %s \n", i.name)
}

type character struct {
	role     string
	holdItem *item
}

func (c *character) pickup(i *item) {
	if c == nil {
		fmt.Printf("there is nobody...\n")
		return
	}
	if i == nil {
		fmt.Printf("look like %s picks nothing....\n", c.role)
	} else {
		fmt.Printf("%s picks %s now!\n", c.role, i.name)
	}
	c.holdItem = i
}

func (c *character) give(to *character) {
	if c == nil {
		fmt.Printf("there is nobody...\n")
		return
	}
	if to == nil {
		fmt.Printf("%s give to nobody...\n", c.role)
	} else {
		if c.holdItem != nil {
			fmt.Printf("%s gives %s to %s\n", c.role, c.holdItem.name, to.role)
		} else {
			fmt.Printf("%s gives nothing to %s\n", c.role, to.role)
		}
	}
	c.holdItem = nil
}

func main() {
	athor := character{
		role: "athor",
	}
	knight := character{
		role: "knight",
	}
	athor.pickup(&item{
		name: "knife",
	})
	athor.give(&knight)
	knight.give(nil)
}
