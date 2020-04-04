package graph

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var loaders = map[string]func(*os.File) *Graph{
	".csv": loadCsv,
}

//SupportedEXtensions loadres
func SupportedEXtensions() (extensions []string) {
	extensions = make([]string, 0, len(loaders))
	i := 0
	for k := range loaders {
		extensions[i] = k
		i++
	}
	return
}

//LoadGraph from file
func LoadGraph(filename string) (g *Graph, err error) {
	var f *os.File
	defer func() {
		f.Close()
	}()
	if _, err = os.Stat(filename); err != nil {
		return nil, err
	}
	f, err = os.Open(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println(filepath.Ext(filename))
	g = loaders[filepath.Ext(filename)](f)
	return g, nil
}

//node, edge1,edge2,...
func loadCsv(file *os.File) (g *Graph) {
	g = &Graph{}
	g.SetDirection(true)
	nodes := make(map[string]*Node)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		edges := strings.Split(line, ",")
		for i, v := range edges {
			if _, added := nodes[v]; !added {
				nodes[v] = NewNode(v)
				g.AddNode(nodes[v])
			}
			if i > 0 {
				g.AddEdgeDefaultWeight(nodes[edges[0]], nodes[v])
			}
		}
	}
	return
}
