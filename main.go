package main

// Node represents an HTML element
type Node struct {
	tag      string
	id       string
	class    string
	text     string
	src      string
	alt      string
	children []*Node
}

func findByID(root *Node, id string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.id == id {
			return nextUp
		}
		if len(nextUp.children) > 0 {
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	return nil
}

func main() {
	image := Node{
		tag: "img",
		src: "http://reallyfakesite.com/logo.svg",
		alt: "Example",
	}

	p := Node{
		tag:      "p",
		text:     "Lorem ipsum in a ptag",
		children: []*Node{&image},
	}

	span := Node{
		tag:  "span",
		id:   "copyright",
		text: "some date",
	}

	div := Node{
		tag:      "div",
		class:    "footer",
		text:     "This is a footer",
		children: []*Node{&span},
	}

	h1 := Node{
		tag:  "h1",
		text: "This is a head line",
	}

	body := Node{
		tag:      "body",
		children: []*Node{&h1, &p, &div},
	}

	html := Node{
		tag:      "html",
		children: []*Node{&body},
	}

}
