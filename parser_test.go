package mark

import (
	"fmt"
	"github.com/k0kubun/pp"
	"testing"
)

func TestParser(t *testing.T) {
	l := lex("1", "foo bar baz\nhello world")
	p := &Tree{lex: l}
	item := p.peek()
	fmt.Println(tokenNames[item.typ], "-->", item.val)
	p.peekCount = 0
	item = p.peek()
	fmt.Println(tokenNames[item.typ], "-->", item.val)
}

func TestParseFn(*testing.T) {
	l := lex("2", "hello\nworld. **ariel**foo  \nenter hahaha  \n~~hello~~ world  \n~foo~  \n_bar_  \n This is my code:`javascript`")
	p := &Tree{lex: l}
	p.parse()

	pp.Printf("[Message]: Tree Node List After Compile\n\n")
	pp.Println(p.Nodes)
	pp.Println("Length of nodes:", len(p.Nodes))
	p.render()
	pp.Printf(p.output + "\n")

	l = lex("3", "```js\nMy Code Block\n```")
	//	l = lex("3", "    MyCodeBlock    \n    again    ")
	p = &Tree{lex: l}
	p.parse()
}
