package main

type Node struct {
	tag      string
	id       string
	class    string
	text     string
	src      string
	alt      string
	children []*Node
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
