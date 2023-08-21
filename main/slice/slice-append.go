package main

import "fmt"

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"} // len: 5, capacity: 5
	dwarfs2 := append(dwarfs1, "Orcus")                                 // capacity: 10, len:6  容量翻倍, 底层操作是将 dwarfs1 的元素复制到一个新的底层数组中
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")            // capacity: 10, len: 9  dwarfs2 3 公用一套底层数组, 和 dwarfs1 隔离
	dwarfs3[0] = "newDwarfs"
	fmt.Println(dwarfs1) // [Ceres Pluto Haumea Makemake Eris]
	fmt.Println(dwarfs2) // [newDwarfs Pluto Haumea Makemake Eris Orcus]
	fmt.Println(dwarfs3) // [newDwarfs Pluto Haumea Makemake Eris Orcus Salacia Quaoar Sedna]

	dwarfs4 := append(dwarfs1[0:2], "newEle")
	fmt.Println(dwarfs1) // [Ceres Pluto newEle Makemake Eris]
	fmt.Println(dwarfs4) // [Ceres Pluto newEle]
}
