#  text.template.parse包说明：

// Package parse builds parse trees for templates as defined by text/template
// and html/template. Clients should use those packages to construct templates
// rather than this one, which provides shared internal data structures not
// intended for general use.

//parse 包为text / template和html / template定义的模板构建解析树。 客户应使用这些包（text / template和html / template）而不是此包parse 来构造模板，后者提供了不用于一般用途的共享内部数据结构。

在开始之前，请你先看text.template包，如果你没看这个包或者你没了解这个包的话，请暂停往下看！不然的话，可能你会不知道我在说什么！

现在请跟我做一遍操作，打开goland编辑器，复制全部的main/text_template_01.go文件内容粘贴到main/compress_zlib.go文件中去，搜索“T_funcs.Parse(templateText)”字符并且在此下断点，右键debug模式执行，你会看到暂停在断点处，接着往下执行一步（快捷键F8），在debug窗口查找到“T_funcs_parsed”对象，单击点开选项，你会看到如下图所示：  

![image-20191201182115354](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201182115354.png)

- 点开common.0.value.Tree.Root.Nodes,
- 继续点开Nodes.0.Text右侧的View，你会看到下图：![image-20191201182557772](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201182557772.png)

这个input是什么呢？现在回到你的断点处代码，滚动轮轴往上翻，你会看到这个：![image-20191201182734459](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201182734459.png)

input就是上面的字符串中的第一行“Input: {{printf "%q" .word}}”中的第一个文本字符串段“input”,那么 {{printf "%q" .word}}在哪里呢？请看下面！

- 继续点开Nodes.1.Pipe.Cmds.0.Args,你会看到下图：![image-20191201183128167](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201183128167.png)
- 分别点开0,1,2，你会看到下图：
![image-20191201183232729](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201183232729.png)

现在到关键时刻了，这些东西和这个parse包有什么联系呢？

由上图你应该注意什么？
![image-20191201183452657](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201183452657.png)

上图标红的地方都是这个包中的对象或者属性，我们在debug窗口中往上追溯：![image-20191201183918097](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201183918097.png)![image-20191201183742563](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201183742563.png)

由上2张图，我们可得到：

Template>Tree>ListNode>[]Node,然后每个Node对象都有以下的字段：

	- NodeType : Node的类型，很多！下面会列举出来
	- Pos : 该Node的字符串（也就是Ident字段或者Text字段中的字符串）在整个解析过程中的整体字符串中的起始位置。
	- tr : 所属树的指针
	- Text : 该节点封装且符合该节点类型的文本字符串
	- Ident : 该节点封装且符合该节点类型的文本字符串切片，跟Text差不多，但是接受多个关键字，这些关键字是运行模板方法execute()中指定的第二个参数对象包含的对象名称，比如在你的当前代码中搜索“m:=map[string]interface{}”，你会看到这个：![image-20191201185235815](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201185235815.png)

而word这个字符串就在第一行，相信你已经注意到了更多的信息，Ident是属于NodeFiled类型的Node才会拥有，而Text是NodeString类型才会拥有，NodeString(16)类型和NodeFiled(8)类型后面的16和8是该类型对象对应的值，每一个Node的类别对象都是对应一个整形数字，如下：

```go

const (
	NodeText       NodeType = iota // Plain text.//文本类型节点
	NodeAction                     // A non-control action such as a field evaluation.
	NodeBool                       // A boolean constant.//boolean类型节点
	NodeChain                      // A sequence of field accesses.//链式调用字段类型节点（如.x.y，常用语map或者slice,array）
    NodeCommand                    // An element of a pipeline.//整个命令行{{开头，}}结尾的节点
    NodeDot                        // The cursor, dot.//{{.}}节点
    nodeElse                       // An else action. Not added to tree.//{{else}}独立节点
    nodeEnd                        // An end action. Not added to tree.//{{end}}独立节点
	NodeField                      // A field or method name.//一个字段或者方法名称节点（相对于结构体）
	NodeIdentifier                 // An identifier; always a function name.//一个内置或者自定义函数名称的节点
	NodeIf                         // An if action.//if语句块节点
	NodeList                       // A list of Nodes.//Nodes列表节点
	NodeNil                        // An untyped nil constant.//未绑定类型的nil常量的节点
	NodeNumber                     // A numerical constant.//数字节点（包含go中的各种类型的数字）
	NodePipe                       // A pipeline of commands.//pipeline管道符连接的节点
	NodeRange                      // A range action.//range语句块保存的节点
	NodeString                     // A string constant.//一个字符串常量节点
    NodeTemplate                   // A template invocation action.//使用关键字template执行模板时候的节点
	NodeVariable                   // A $ variable.//$声明变量后保存变量的类型节点
	NodeWith                       // A with action.//with语句块类型节点
)
```

事实上不必要死记，看上面的几张图可以得到这个关系图：

NodeText

NodeAction>NodePipe>NodeCommand>{NodeIdentifier,NodeString,NodeField}

每个node由NodeText和NodeAction组成，也就是比如“Input: {{printf "%q" .word}}”中，“Input: ”是会解析为NodeText类型的node节点来保存，{{printf "%q" .word}}则会解析为NodeAction类型的Node节点来保存，当然NodeAction还有很多的子节点用来继续拆分：printf被NodeIdentifier类型的节点来保存，"%q"被NodeString类型的几点保存，word则被NodeField类型的节点保存。





- 相对于“Output 20: {{with $x := "output" | printf "%q"}}{{$x}}{{end}}”这条命令的话，请看：		![image-20191201192855541](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201192855541.png)![image-20191201192742193](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/pic/image-20191201192742193.png)
- 更多的信息其实你自己可以在debug窗口点开来查看的！方法已经说了！不再累叙！

下面我们对这个包中的所有的节点对象进行遍历一次：代码来自Node.go，加上我自己的理解和翻译:

```go
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse nodes.

package parse

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var textFormat = "%s" // Changed to "%q" in tests for better error messages.//在测试中更改为“％q”以获取更好的错误消息。

// A Node is an element in the parse tree. The interface is trivial.
// The interface contains an unexported method so that only
// types local to this package can satisfy it.
//节点是解析树中的元素。 该接口是微不足道的。
//接口包含一个未导出的方法，因此只有此程序包本地的类型才能满足它。
//简单理解就是树节点
type Node interface {
	Type() NodeType	//节点的类型
	String() string	//节点的字符串描述
	// Copy does a deep copy of the Node and all its components.
	// To avoid type assertions, some XxxNodes also have specialized
	// CopyXxx methods that return *XxxNode.
    //Copy对节点及其所有组件进行深层复制。
	//为了避免类型断言，某些XxxNode还具有专门的CopyXxx方法，这些方法返回* XxxNode。
	Copy() Node
	Position() Pos // byte position of start of node in full original input string//完整原始输入字符串中节点开始的字节位置，就是从你输入字符串的哪个位置的字节开始是属于这个节点的信息
	// tree returns the containing *Tree.
	// It is unexported so all implementations of Node are in this package.
    // tree返回当前节点被包含的数指针*Tree。每个节点都属于一棵树。
	//它未公开导出，因此Node的所有实现都在此包中。
	tree() *Tree
}

// NodeType identifies the type of a parse tree node.// NodeType标识解析树节点的类型。int类型，不同的int代表不同的节点类型
type NodeType int

// Pos represents a byte position in the original input text from which
// this template was parsed.
// Pos表示解析此模板的原始输入文本中的字节位置。
type Pos int

func (p Pos) Position() Pos {
	return p
}

// Type returns itself and provides an easy default implementation
// for embedding in a Node. Embedded in all non-trivial Nodes.
// Type会返回自身，并为嵌入Node提供简单的默认实现。 嵌入所有非平凡节点中。
// 跟NodeType对象实例t本身的值完全相同
func (t NodeType) Type() NodeType {
	return t
}

const (
	NodeText       NodeType = iota // Plain text.
	NodeAction                     // A non-control action such as a field evaluation.非语句块action，例如字段解析。
	NodeBool                       // A boolean constant.//bool常量
	NodeChain                      // A sequence of field accesses.//链式调用字段，似乎在解析树中并没有用到，可能是在解析时候的某个函数中用到，但是解析后的树节点中我目前没见过
	NodeCommand                    // An element of a pipeline.//一个命令
    NodeDot                        // The cursor, dot.//{{.}}生成的节点
	nodeElse                       // An else action. Not added to tree.//{{else}}生成的节点
	nodeEnd                        // An end action. Not added to tree.//{{end}}生成的节点
	NodeField                      // A field or method name.//字段或者方法名的节点
	NodeIdentifier                 // An identifier; always a function name.//只要是内置或者外置的函数的节点
    NodeIf                         // An if action.//{{if}}节点
    NodeList                       // A list of Nodes.//{{}}节点列表，下面含有多个子节点，比如函数参数有多个时候
	NodeNil                        // An untyped nil constant.//这个东西真没在解析树中见过
	NodeNumber                     // A numerical constant.//go数字类型的节点
	NodePipe                       // A pipeline of commands.//管道符类型节点
    NodeRange                      // A range action.//{{Range}}类型节点
    NodeString                     // A string constant.//go中的字符串类型节点
    NodeTemplate                   // A template invocation action.//{{template}}模板嵌入的节点
    NodeVariable                   // A $ variable.//{{$aa:=...}}的节点
	NodeWith                       // A with action.//这个不知道
)

// Nodes.

// ListNode holds a sequence of nodes.// ListNode包含一系列节点。
type ListNode struct {
	NodeType
	Pos
	tr    *Tree	//几乎所有的节点都含有上面的3个属性
	Nodes []Node // The element nodes in lexical order.//按词法顺序的元素节点。
}

func (t *Tree) newList(pos Pos) *ListNode {
	return &ListNode{tr: t, NodeType: NodeList, Pos: pos}
}

func (l *ListNode) append(n Node) {
	l.Nodes = append(l.Nodes, n)
}

func (l *ListNode) tree() *Tree {
	return l.tr
}


func (l *ListNode) String() string {
	b := new(bytes.Buffer)
	for _, n := range l.Nodes {
		fmt.Fprint(b, n)
	}
	return b.String()
}

func (l *ListNode) CopyList() *ListNode {
	if l == nil {
		return l
	}
	n := l.tr.newList(l.Pos)
	for _, elem := range l.Nodes {
		n.append(elem.Copy())
	}
	return n
}

func (l *ListNode) Copy() Node {
	return l.CopyList()
}
//几乎所有的节点都会含有上面的几个方法或者其中的几个方法，抑或者是加一个Type()方法

// TextNode holds plain text.// TextNode保存纯文本。
type TextNode struct {
	NodeType
	Pos
	tr   *Tree
	Text []byte // The text; may span newlines.// 文本; 可能跨越换行符。
}

func (t *Tree) newText(pos Pos, text string) *TextNode {
	return &TextNode{tr: t, NodeType: NodeText, Pos: pos, Text: []byte(text)}
}

func (t *TextNode) String() string {
	return fmt.Sprintf(textFormat, t.Text)
}

func (t *TextNode) tree() *Tree {
	return t.tr
}

func (t *TextNode) Copy() Node {
	return &TextNode{tr: t.tr, NodeType: NodeText, Pos: t.Pos, Text: append([]byte{}, t.Text...)}
}

// PipeNode holds a pipeline with optional declaration// PipeNode持有带有可选声明的管道
type PipeNode struct {
	NodeType
	Pos
	tr       *Tree
	Line     int             // The line number in the input. Deprecated: Kept for compatibility.//输入中的行号。 不推荐使用：保留兼容性。
    IsAssign bool            // The variables are being assigned, not declared.//是否正在分配未声明的变量，如果没使用{{$Xxx}}的话，则不会为true
	Decl     []*VariableNode // Variables in lexical order.//以词法顺序的变量。跟上面一个字段是相对应的，声明好的变量会在这里存好指针
	Cmds     []*CommandNode  // The commands in lexical order.//命令按词法顺序排列。除了声明变量之外的命令都在这里保存着
}

func (t *Tree) newPipeline(pos Pos, line int, vars []*VariableNode) *PipeNode {
	return &PipeNode{tr: t, NodeType: NodePipe, Pos: pos, Line: line, Decl: vars}
}

func (p *PipeNode) append(command *CommandNode) {
	p.Cmds = append(p.Cmds, command)
}

func (p *PipeNode) String() string {
	s := ""
	if len(p.Decl) > 0 {
		for i, v := range p.Decl {
			if i > 0 {
				s += ", "
			}
			s += v.String()
		}
		s += " := "
	}
	for i, c := range p.Cmds {
		if i > 0 {
			s += " | "
		}
		s += c.String()
	}
	return s
}

func (p *PipeNode) tree() *Tree {
	return p.tr
}

func (p *PipeNode) CopyPipe() *PipeNode {
	if p == nil {
		return p
	}
	var vars []*VariableNode
	for _, d := range p.Decl {
		vars = append(vars, d.Copy().(*VariableNode))
	}
	n := p.tr.newPipeline(p.Pos, p.Line, vars)
	n.IsAssign = p.IsAssign
	for _, c := range p.Cmds {
		n.append(c.Copy().(*CommandNode))
	}
	return n
}

func (p *PipeNode) Copy() Node {
	return p.CopyPipe()
}

// ActionNode holds an action (something bounded by delimiters).
// Control actions have their own nodes; ActionNode represents simple
// ones such as field evaluations and parenthesized pipelines.
// ActionNode持有一个动作（由定界符限制）。
//控制动作有自己的节点； ActionNode代表简单的变量，例如字段解析和带括号的管道。
//比如非{{range}},{{if}},{{with}}的那些命令都是用ActionNode类型包含子节点的！
//但是如果是{{range}},{{if}},{{with}}的那些语句块命令的话，则不会拥有这个节点，而是相应的拥有
//RangeNode,IfNode,WithNode来包住子节点！
type ActionNode struct {
	NodeType
	Pos
	tr   *Tree
	Line int       // The line number in the input. Deprecated: Kept for compatibility.
    //上面的结果属性基本是大节点共有的一些属性字段，怎么理解大节点？自己想想{{range}},{{if}},{{with}}那些就知道了！
	Pipe *PipeNode // The pipeline in the action.
}

func (t *Tree) newAction(pos Pos, line int, pipe *PipeNode) *ActionNode {
	return &ActionNode{tr: t, NodeType: NodeAction, Pos: pos, Line: line, Pipe: pipe}
}

func (a *ActionNode) String() string {
	return fmt.Sprintf("{{%s}}", a.Pipe)

}

func (a *ActionNode) tree() *Tree {
	return a.tr
}

func (a *ActionNode) Copy() Node {
	return a.tr.newAction(a.Pos, a.Line, a.Pipe.CopyPipe())

}

// CommandNode holds a command (a pipeline inside an evaluating action).
// CommandNode拥有一个命令（评估动作内部的管道pipeline）。用来包含PipelineNode对象的，可以有多个管道节点PipelineNode对象
type CommandNode struct {
	NodeType
	Pos
	tr   *Tree
	Args []Node // Arguments in lexical order: Identifier, field, or constant.//以词法顺序的参数：标识符，字段或常量。列表里面可以放IdentifierNode，FieldNode和ConstantNode类型的节点
}

func (t *Tree) newCommand(pos Pos) *CommandNode {
	return &CommandNode{tr: t, NodeType: NodeCommand, Pos: pos}
}

func (c *CommandNode) append(arg Node) {
	c.Args = append(c.Args, arg)
}

func (c *CommandNode) String() string {
	s := ""
	for i, arg := range c.Args {
		if i > 0 {
			s += " "
		}
		if arg, ok := arg.(*PipeNode); ok {
			s += "(" + arg.String() + ")"
			continue
		}
		s += arg.String()
	}
	return s
}

func (c *CommandNode) tree() *Tree {
	return c.tr
}

func (c *CommandNode) Copy() Node {
	if c == nil {
		return c
	}
	n := c.tr.newCommand(c.Pos)
	for _, c := range c.Args {
		n.append(c.Copy())
	}
	return n
}

// IdentifierNode holds an identifier.// IdentifierNode包含一个标识符。
type IdentifierNode struct {
	NodeType
	Pos
	tr    *Tree
	Ident string // The identifier's name.//标识符的名称
}

// NewIdentifier returns a new IdentifierNode with the given identifier name.
// NewIdentifier返回具有给定标识符名称的新IdentifierNode。
func NewIdentifier(ident string) *IdentifierNode {
	return &IdentifierNode{NodeType: NodeIdentifier, Ident: ident}
}

// SetPos sets the position. NewIdentifier is a public method so we can't modify its signature.
// Chained for convenience.
// TODO: fix one day?
// SetPos设置位置。 NewIdentifier是一个公共方法，因此我们无法修改其签名。
func (i *IdentifierNode) SetPos(pos Pos) *IdentifierNode {
	i.Pos = pos
	return i
}

// SetTree sets the parent tree for the node. NewIdentifier is a public method so we can't modify its signature.
// Chained for convenience.
// TODO: fix one day?
// SetTree设置节点的父树。 NewIdentifier是一个公共方法，因此我们无法修改其签名。
func (i *IdentifierNode) SetTree(t *Tree) *IdentifierNode {
	i.tr = t
	return i
}

func (i *IdentifierNode) String() string {
	return i.Ident
}

func (i *IdentifierNode) tree() *Tree {
	return i.tr
}

func (i *IdentifierNode) Copy() Node {
	return NewIdentifier(i.Ident).SetTree(i.tr).SetPos(i.Pos)
}

// AssignNode holds a list of variable names, possibly with chained field
// accesses. The dollar sign is part of the (first) name.
// AssignNode保存变量名称列表，可能具有链接的字段访问权限。 美元符号$是名称的一部分。
type VariableNode struct {
	NodeType
	Pos
	tr    *Tree
	Ident []string // Variable name and fields in lexical order.//按词法顺序排列的变量名称和字段。
}

func (t *Tree) newVariable(pos Pos, ident string) *VariableNode {
	return &VariableNode{tr: t, NodeType: NodeVariable, Pos: pos, Ident: strings.Split(ident, ".")}
}

func (v *VariableNode) String() string {
	s := ""
	for i, id := range v.Ident {
		if i > 0 {
			s += "."
		}
		s += id
	}
	return s
}

func (v *VariableNode) tree() *Tree {
	return v.tr
}

func (v *VariableNode) Copy() Node {
	return &VariableNode{tr: v.tr, NodeType: NodeVariable, Pos: v.Pos, Ident: append([]string{}, v.Ident...)}
}

// DotNode holds the special identifier '.'.// DotNode保留特殊标识符 '.'。
type DotNode struct {
	NodeType
	Pos
	tr *Tree
}

func (t *Tree) newDot(pos Pos) *DotNode {
	return &DotNode{tr: t, NodeType: NodeDot, Pos: pos}
}

func (d *DotNode) Type() NodeType {
	// Override method on embedded NodeType for API compatibility.
	// TODO: Not really a problem; could change API without effect but
	// api tool complains.
    //嵌入式NodeType上的重写方法，以实现API兼容性。
	// TODO：确实不是问题； 可能更改API无效，但api工具抱怨。
	return NodeDot
}

func (d *DotNode) String() string {
	return "."
}

func (d *DotNode) tree() *Tree {
	return d.tr
}

func (d *DotNode) Copy() Node {
	return d.tr.newDot(d.Pos)
}

// NilNode holds the special identifier 'nil' representing an untyped nil constant.
// NilNode保留特殊标识符'nil'，该标识符表示无类型的nil常数。
type NilNode struct {
	NodeType
	Pos
	tr *Tree
}

func (t *Tree) newNil(pos Pos) *NilNode {
	return &NilNode{tr: t, NodeType: NodeNil, Pos: pos}
}

func (n *NilNode) Type() NodeType {
	// Override method on embedded NodeType for API compatibility.
	// TODO: Not really a problem; could change API without effect but
	// api tool complains.
    //嵌入式NodeType上的重写方法，以实现API兼容性。
	// TODO：确实不是问题； 可能更改API无效，但api工具抱怨。
	return NodeNil
}

func (n *NilNode) String() string {
	return "nil"
}

func (n *NilNode) tree() *Tree {
	return n.tr
}

func (n *NilNode) Copy() Node {
	return n.tr.newNil(n.Pos)
}

// FieldNode holds a field (identifier starting with '.').
// The names may be chained ('.x.y').
// The period is dropped from each ident.
// FieldNode保存一个字段（以“。”开头的标识符）。
//名称可以链接在一起（'.x.y'）。
//从每个ident中删除句点‘.’。
type FieldNode struct {
	NodeType
	Pos
	tr    *Tree
	Ident []string // The identifiers in lexical order.//按词法顺序排列的标识符。
}

func (t *Tree) newField(pos Pos, ident string) *FieldNode {
	return &FieldNode{tr: t, NodeType: NodeField, Pos: pos, Ident: strings.Split(ident[1:], ".")} // [1:] to drop leading period// [1：]删除前置句点‘.’
}

func (f *FieldNode) String() string {
	s := ""
	for _, id := range f.Ident {
		s += "." + id
	}
	return s
}

func (f *FieldNode) tree() *Tree {
	return f.tr
}

func (f *FieldNode) Copy() Node {
	return &FieldNode{tr: f.tr, NodeType: NodeField, Pos: f.Pos, Ident: append([]string{}, f.Ident...)}
}

// ChainNode holds a term followed by a chain of field accesses (identifier starting with '.').
// The names may be chained ('.x.y').
// The periods are dropped from each ident.
// ChainNode包含一个术语，后跟一连串的字段访问（以“.”开头的标识符）。
//名称可以链接在一起（'.x.y'）。
//从每个ident中删除句点。
type ChainNode struct {
	NodeType
	Pos
	tr    *Tree
	Node  Node
	Field []string // The identifiers in lexical order.
}

func (t *Tree) newChain(pos Pos, node Node) *ChainNode {
	return &ChainNode{tr: t, NodeType: NodeChain, Pos: pos, Node: node}
}

// Add adds the named field (which should start with a period) to the end of the chain.
// Add将命名字段（应以句点开头）添加到链的末尾。
func (c *ChainNode) Add(field string) {
	if len(field) == 0 || field[0] != '.' {
		panic("no dot in field")
	}
	field = field[1:] // Remove leading dot.
	if field == "" {
		panic("empty field")
	}
	c.Field = append(c.Field, field)
}

func (c *ChainNode) String() string {
	s := c.Node.String()
	if _, ok := c.Node.(*PipeNode); ok {
		s = "(" + s + ")"
	}
	for _, field := range c.Field {
		s += "." + field
	}
	return s
}

func (c *ChainNode) tree() *Tree {
	return c.tr
}

func (c *ChainNode) Copy() Node {
	return &ChainNode{tr: c.tr, NodeType: NodeChain, Pos: c.Pos, Node: c.Node, Field: append([]string{}, c.Field...)}
}

// BoolNode holds a boolean constant.
// BoolNode拥有一个布尔常量。
type BoolNode struct {
	NodeType
	Pos
	tr   *Tree
	True bool // The value of the boolean constant.//布尔常量的值。
}

func (t *Tree) newBool(pos Pos, true bool) *BoolNode {
	return &BoolNode{tr: t, NodeType: NodeBool, Pos: pos, True: true}
}

func (b *BoolNode) String() string {
	if b.True {
		return "true"
	}
	return "false"
}

func (b *BoolNode) tree() *Tree {
	return b.tr
}

func (b *BoolNode) Copy() Node {
	return b.tr.newBool(b.Pos, b.True)
}

// NumberNode holds a number: signed or unsigned integer, float, or complex.
// The value is parsed and stored under all the types that can represent the value.
// This simulates in a small amount of code the behavior of Go's ideal constants.
// NumberNode拥有一个数字：有符号或无符号整数，浮点数或复数( signed or unsigned integer, float, or complex.)。
//解析值并将其存储在可以代表该值的所有类型下。
//这会以少量代码模拟Go理想常数的行为。
type NumberNode struct {
	NodeType
	Pos
	tr         *Tree
	IsInt      bool       // Number has an integral value.//数字具有整数值。
	IsUint     bool       // Number has an unsigned integral value.//数字具有无符号整数值。
	IsFloat    bool       // Number has a floating-point value.// Number有一个浮点值。
	IsComplex  bool       // Number is complex.//数字是复数类型。
    //下面的4个字段对应上面的4个字段的值
	Int64      int64      // The signed integer value.//有符号整数值。
	Uint64     uint64     // The unsigned integer value.//无符号整数值。
	Float64    float64    // The floating-point value.//浮点值。
	Complex128 complex128 // The complex value.//复数值。
	Text       string     // The original textual representation from the input.//输入的原始文本表示形式。
}

func (t *Tree) newNumber(pos Pos, text string, typ itemType) (*NumberNode, error) {
	n := &NumberNode{tr: t, NodeType: NodeNumber, Pos: pos, Text: text}
	switch typ {
	case itemCharConstant:
		rune, _, tail, err := strconv.UnquoteChar(text[1:], text[0])
		if err != nil {
			return nil, err
		}
		if tail != "'" {
			return nil, fmt.Errorf("malformed character constant: %s", text)
		}
		n.Int64 = int64(rune)
		n.IsInt = true
		n.Uint64 = uint64(rune)
		n.IsUint = true
		n.Float64 = float64(rune) // odd but those are the rules.//奇数，但这些是规则。
		n.IsFloat = true
		return n, nil
	case itemComplex:
		// fmt.Sscan can parse the pair, so let it do the work.
        // fmt.Sscan可以解析该对，因此让它完成工作。
		if _, err := fmt.Sscan(text, &n.Complex128); err != nil {
			return nil, err
		}
		n.IsComplex = true
		n.simplifyComplex()
		return n, nil
	}
	// Imaginary constants can only be complex unless they are zero.
    //虚常数只能为复数，除非它们为零。
	if len(text) > 0 && text[len(text)-1] == 'i' {
		f, err := strconv.ParseFloat(text[:len(text)-1], 64)
		if err == nil {
			n.IsComplex = true
			n.Complex128 = complex(0, f)
			n.simplifyComplex()
			return n, nil
		}
	}
	// Do integer test first so we get 0x123 etc.//首先进行整数测试，以便得到0x123等。
	u, err := strconv.ParseUint(text, 0, 64) // will fail for -0; fixed below.//将失败-0; 固定在下面。
	if err == nil {
		n.IsUint = true
		n.Uint64 = u
	}
	i, err := strconv.ParseInt(text, 0, 64)
	if err == nil {
		n.IsInt = true
		n.Int64 = i
		if i == 0 {
			n.IsUint = true // in case of -0.
			n.Uint64 = u
		}
	}
	// If an integer extraction succeeded, promote the float.//如果整数提取成功，则提取浮点数。
	if n.IsInt {
		n.IsFloat = true
		n.Float64 = float64(n.Int64)
	} else if n.IsUint {
		n.IsFloat = true
		n.Float64 = float64(n.Uint64)
	} else {
		f, err := strconv.ParseFloat(text, 64)
		if err == nil {
			// If we parsed it as a float but it looks like an integer,
			// it's a huge number too large to fit in an int. Reject it.
            //如果我们将其解析为浮点型，但看起来像一个整数，则它太大了，无法容纳int。 拒绝它。
			if !strings.ContainsAny(text, ".eEpP") {
				return nil, fmt.Errorf("integer overflow: %q", text)
			}
			n.IsFloat = true
			n.Float64 = f
			// If a floating-point extraction succeeded, extract the int if needed.
            //如果浮点提取成功，则根据需要提取int。
			if !n.IsInt && float64(int64(f)) == f {
				n.IsInt = true
				n.Int64 = int64(f)
			}
			if !n.IsUint && float64(uint64(f)) == f {
				n.IsUint = true
				n.Uint64 = uint64(f)
			}
		}
	}
	if !n.IsInt && !n.IsUint && !n.IsFloat {
		return nil, fmt.Errorf("illegal number syntax: %q", text)
	}
	return n, nil
}

// simplifyComplex pulls out any other types that are represented by the complex number.
// These all require that the imaginary part be zero.
// simpleComplex提取由复数表示的任何其他类型。
//这些都要求虚部为零。
func (n *NumberNode) simplifyComplex() {
	n.IsFloat = imag(n.Complex128) == 0
	if n.IsFloat {
		n.Float64 = real(n.Complex128)
		n.IsInt = float64(int64(n.Float64)) == n.Float64
		if n.IsInt {
			n.Int64 = int64(n.Float64)
		}
		n.IsUint = float64(uint64(n.Float64)) == n.Float64
		if n.IsUint {
			n.Uint64 = uint64(n.Float64)
		}
	}
}

func (n *NumberNode) String() string {
	return n.Text
}

func (n *NumberNode) tree() *Tree {
	return n.tr
}

func (n *NumberNode) Copy() Node {
	nn := new(NumberNode)
	*nn = *n // Easy, fast, correct.//简单，快速，正确。
	return nn
}

// StringNode holds a string constant. The value has been "unquoted".// StringNode保存一个字符串常量。 该值已被“取消引用（unquoted）”。
type StringNode struct {
	NodeType
	Pos
	tr     *Tree
	Quoted string // The original text of the string, with quotes.//字符串的原始文本，带引号。
	Text   string // The string, after quote processing.字符串，经过报价处理。
}

func (t *Tree) newString(pos Pos, orig, text string) *StringNode {
	return &StringNode{tr: t, NodeType: NodeString, Pos: pos, Quoted: orig, Text: text}
}

func (s *StringNode) String() string {
	return s.Quoted
}

func (s *StringNode) tree() *Tree {
	return s.tr
}

func (s *StringNode) Copy() Node {
	return s.tr.newString(s.Pos, s.Quoted, s.Text)
}

// endNode represents an {{end}} action.
// It does not appear in the final parse tree.
// endNode代表一个{{end}}动作。
//它不会出现在最终的分析树中。
type endNode struct {
	NodeType
	Pos
	tr *Tree
}

func (t *Tree) newEnd(pos Pos) *endNode {
	return &endNode{tr: t, NodeType: nodeEnd, Pos: pos}
}

func (e *endNode) String() string {
	return "{{end}}"
}

func (e *endNode) tree() *Tree {
	return e.tr
}

func (e *endNode) Copy() Node {
	return e.tr.newEnd(e.Pos)
}

// elseNode represents an {{else}} action. Does not appear in the final tree.
// elseNode表示一个{{else}}动作。 没有出现在最后的树中。
type elseNode struct {
	NodeType
	Pos
	tr   *Tree
	Line int // The line number in the input. Deprecated: Kept for compatibility.
}

func (t *Tree) newElse(pos Pos, line int) *elseNode {
	return &elseNode{tr: t, NodeType: nodeElse, Pos: pos, Line: line}
}

func (e *elseNode) Type() NodeType {
	return nodeElse
}

func (e *elseNode) String() string {
	return "{{else}}"
}

func (e *elseNode) tree() *Tree {
	return e.tr
}

func (e *elseNode) Copy() Node {
	return e.tr.newElse(e.Pos, e.Line)
}

// BranchNode is the common representation of if, range, and with.
// BranchNode是if，range和with的通用表示形式。
type BranchNode struct {
	NodeType
	Pos
	tr       *Tree
	Line     int       // The line number in the input. Deprecated: Kept for compatibility.
	Pipe     *PipeNode // The pipeline to be evaluated.//要评估解析的管道。
	List     *ListNode // What to execute if the value is non-empty.//如果该值非空，则执行这个节点列表。
	ElseList *ListNode // What to execute if the value is empty (nil if absent).//如果该值为空（不存在则为nil），则执行这个节点列表。
}

func (b *BranchNode) String() string {
	name := ""
	switch b.NodeType {
	case NodeIf:
		name = "if"
	case NodeRange:
		name = "range"
	case NodeWith:
		name = "with"
	default:
		panic("unknown branch type")
	}
	if b.ElseList != nil {
		return fmt.Sprintf("{{%s %s}}%s{{else}}%s{{end}}", name, b.Pipe, b.List, b.ElseList)
	}
	return fmt.Sprintf("{{%s %s}}%s{{end}}", name, b.Pipe, b.List)
}

func (b *BranchNode) tree() *Tree {
	return b.tr
}

func (b *BranchNode) Copy() Node {
	switch b.NodeType {
	case NodeIf:
		return b.tr.newIf(b.Pos, b.Line, b.Pipe, b.List, b.ElseList)
	case NodeRange:
		return b.tr.newRange(b.Pos, b.Line, b.Pipe, b.List, b.ElseList)
	case NodeWith:
		return b.tr.newWith(b.Pos, b.Line, b.Pipe, b.List, b.ElseList)
	default:
		panic("unknown branch type")
	}
}

// IfNode represents an {{if}} action and its commands.// IfNode表示一个{{if}}动作及其命令。
type IfNode struct {
	BranchNode
}

func (t *Tree) newIf(pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) *IfNode {
	return &IfNode{BranchNode{tr: t, NodeType: NodeIf, Pos: pos, Line: line, Pipe: pipe, List: list, ElseList: elseList}}
}

func (i *IfNode) Copy() Node {
	return i.tr.newIf(i.Pos, i.Line, i.Pipe.CopyPipe(), i.List.CopyList(), i.ElseList.CopyList())
}

// RangeNode represents a {{range}} action and its commands.
// RangeNode表示一个{{range}}动作及其命令。
type RangeNode struct {
	BranchNode
}

func (t *Tree) newRange(pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) *RangeNode {
	return &RangeNode{BranchNode{tr: t, NodeType: NodeRange, Pos: pos, Line: line, Pipe: pipe, List: list, ElseList: elseList}}
}

func (r *RangeNode) Copy() Node {
	return r.tr.newRange(r.Pos, r.Line, r.Pipe.CopyPipe(), r.List.CopyList(), r.ElseList.CopyList())
}

// WithNode represents a {{with}} action and its commands.
// WithNode代表一个{{with}}动作及其命令。
type WithNode struct {
	BranchNode
}

func (t *Tree) newWith(pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) *WithNode {
	return &WithNode{BranchNode{tr: t, NodeType: NodeWith, Pos: pos, Line: line, Pipe: pipe, List: list, ElseList: elseList}}
}

func (w *WithNode) Copy() Node {
	return w.tr.newWith(w.Pos, w.Line, w.Pipe.CopyPipe(), w.List.CopyList(), w.ElseList.CopyList())
}

// TemplateNode represents a {{template}} action.// TemplateNode代表一个{{template}}动作。
type TemplateNode struct {
	NodeType
	Pos
	tr   *Tree
	Line int       // The line number in the input. Deprecated: Kept for compatibility.
	Name string    // The name of the template (unquoted).//模板的名称（不带引号）。
	Pipe *PipeNode // The command to evaluate as dot for the template.//评估为dot的模板的命令。
}

func (t *Tree) newTemplate(pos Pos, line int, name string, pipe *PipeNode) *TemplateNode {
	return &TemplateNode{tr: t, NodeType: NodeTemplate, Pos: pos, Line: line, Name: name, Pipe: pipe}
}

func (t *TemplateNode) String() string {
	if t.Pipe == nil {
		return fmt.Sprintf("{{template %q}}", t.Name)
	}
	return fmt.Sprintf("{{template %q %s}}", t.Name, t.Pipe)
}

func (t *TemplateNode) tree() *Tree {
	return t.tr
}

func (t *TemplateNode) Copy() Node {
	return t.tr.newTemplate(t.Pos, t.Line, t.Name, t.Pipe.CopyPipe())
}

```

