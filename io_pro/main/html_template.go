package main

import (
	"fmt"
	Thtml "html/template"
	"os"
	Ttext "text/template" //必须起别名
)


//在你阅读本包之前，请务必先去阅读text.template和text.template.parse包，同时非常有必要先去 go中文网 阅读html.template包的说明文档。


//template包（html/template）实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出。
//本包提供了和text/template包相同的接口，无论何时当输出是HTML的时候都应使用本包。
//
//此处的文档关注本包的安全特性。至于如何使用模板，请参照text/template包。
//本包是对text/template包的包装，两个包提供的模板API几无差别，可以安全的随意替换两包。


//本报重要的是看我写的注释！因为全部的api都跟text.template的完全相同，只是多了自动转义这个步骤而已！所以下面我列出来的
// 结构体都是非常重要的！当然你也可以自己去看源码文档！

//// Template is a specialized Template from "text/template" that produces a safe
//// HTML document fragment.
////Template是"text/template"中的专用模板，可生成安全的HTML文档片段。
//type Template struct {
//	// Sticky error if escaping fails, or escapeOK if succeeded.
//	//如果转义失败，则返回下面字段错误，如果成功，则发生转义。
//	escapeErr error
//	// We could embed the text/template field, but it's safer not to because
//	// we need to keep our version of the name space and the underlying
//	// template's in sync.
//	//我们可以嵌入text / template字段，但是这样做比较安全，因为我们需要保持名称空间的版本与底层模板的同步。
//	text *template.Template
//	// The underlying template's parse tree, updated to be HTML-safe.
//	//基础模板的解析树，已更新为HTML安全的。
//	Tree       *parse.Tree
//	*nameSpace // common to all associated templates//所有关联模板通用
//}

		//// escapeOK is a sentinel value used to indicate valid escaping.
		//// escapeOK是用于指示有效转义的标记值。
		//var escapeOK = fmt.Errorf("template escaped correctly")
		//
		//// nameSpace is the data structure shared by all templates in an association.
		//// nameSpace是关联中所有模板共享的数据结构。
		//type nameSpace struct {
		//	mu      sync.Mutex
		//	set     map[string]*Template
		//	escaped bool
		//	esc     escaper
		//}


		//// funcMap maps command names to functions that render their inputs safe.
		//// funcMap将命令名称映射到使输入安全的函数。
		//var funcMap = template.FuncMap{
		//	"_html_template_attrescaper":     attrEscaper,
		//	"_html_template_commentescaper":  commentEscaper,
		//	"_html_template_cssescaper":      cssEscaper,
		//	"_html_template_cssvaluefilter":  cssValueFilter,
		//	"_html_template_htmlnamefilter":  htmlNameFilter,
		//	"_html_template_htmlescaper":     htmlEscaper,
		//	"_html_template_jsregexpescaper": jsRegexpEscaper,
		//	"_html_template_jsstrescaper":    jsStrEscaper,
		//	"_html_template_jsvalescaper":    jsValEscaper,
		//	"_html_template_nospaceescaper":  htmlNospaceEscaper,
		//	"_html_template_rcdataescaper":   rcdataEscaper,
		//	"_html_template_srcsetescaper":   srcsetFilterAndEscaper,
		//	"_html_template_urlescaper":      urlEscaper,
		//	"_html_template_urlfilter":       urlFilter,
		//	"_html_template_urlnormalizer":   urlNormalizer,
		//	"_eval_args_":                    evalArgs,
		//}
		//
		//// escaper collects type inferences about templates and changes needed to make
		//// templates injection safe.
		//// escaper收集有关模板的类型推断以及使模板注入安全所需的更改。
		//type escaper struct {
		//	// ns is the nameSpace that this escaper is associated with.
		//	// ns是与此转义符关联的nameSpace。
		//	ns *nameSpace
		//	// output[templateName] is the output context for a templateName that
		//	// has been mangled to include its input context.
		//	// output [templateName]是templateName的输出上下文，该模板名称已被修改为包括其输入上下文。
		//	output map[string]context
		//	// derived[c.mangle(name)] maps to a template derived from the template
		//	// named name templateName for the start context c.
		//	// //derived[c.mangle（name）]映射到从名为start template c的名为name templateName的模板派生的模板。（derived派生的意思）
		//	//大概表示的就是派生出来的模板对象列表，应该是多个map引用同一堆模板对象
		//	derived map[string]*template.Template
		//	// called[templateName] is a set of called mangled template names.
		//	// named [templateName]是一组被称为损坏的模板名称。
		//	called map[string]bool
		//	// xxxNodeEdits are the accumulated edits to apply during commit.
		//	// Such edits are not applied immediately in case a template set
		//	// executes a given template in different escaping contexts.
		//	// xxxNodeEdits是在提交期间应用的累积编辑。
		//	//如果模板集在不同的转义上下文中执行给定的模板，则不会立即应用此类编辑。
		//	actionNodeEdits   map[*parse.ActionNode][]string//这些Node类型全部都是text.template.parse包下的对象
		//	templateNodeEdits map[*parse.TemplateNode]string
		//	textNodeEdits     map[*parse.TextNode][]byte
		//}


		//// context describes the state an HTML parser must be in when it reaches the
		//// portion of HTML produced by evaluating a particular template node.
		////
		//// The zero value of type context is the start context for a template that
		//// produces an HTML fragment as defined at
		//// https://www.w3.org/TR/html5/syntax.html#the-end
		//// where the context element is null.
		////上下文描述了HTML解析器到达通过评估特定模板节点生成的HTML部分时必须处于的状态。
		//// context类型的零值是模板的起始上下文，该模板会生成一个HTML片段，
		//// 如https://www.w3.org/TR/html5/syntax.html#the-end所定义，其中context元素为null 。
		//type context struct {
		//	state   state
		//	delim   delim
		//	urlPart urlPart
		//	jsCtx   jsCtx
		//	attr    attr
		//	element element
		//	err     *Error
		//}



		//// state describes a high-level HTML parser state.
		////
		//// It bounds the top of the element stack, and by extension the HTML insertion
		//// mode, but also contains state that does not correspond to anything in the
		//// HTML5 parsing algorithm because a single token production in the HTML
		//// grammar may contain embedded actions in a template. For instance, the quoted
		//// HTML attribute produced by
		////     <div title="Hello {{.World}}">
		//// is a single token in HTML's grammar but in a template spans several nodes.
		////状态描述了高级HTML解析器状态。
		////它限制了元素堆栈的顶部，并扩展了HTML插入模式，但还包含了与HTML5解析算法中的任何内容都不对应的状态，因为HTML语法中的单个令牌产生可能包含模板中的嵌入式动作 。
		//// 例如，由<div title =“ Hello {{.World}}”>产生的带引号的HTML属性在HTML语法中是单个标记，但在模板中跨越多个节点。
		//type state uint8
		//
		////go:generate stringer -type state
		//
		//const (
		//	// stateText is parsed character data. An HTML parser is in
		//	// this state when its parse position is outside an HTML tag,
		//	// directive, comment, and special element body.
		//	// stateText是已解析的字符数据。 当HTML解析器的解析位置在HTML标签，指令，注释和特殊元素主体之外时，便处于这种状态。
		//	stateText state = iota
		//	// stateTag occurs before an HTML attribute or the end of a tag.
		//	// stateTag出现在HTML属性或标记结尾之前。
		//	stateTag
		//	// stateAttrName occurs inside an attribute name.
		//	// It occurs between the ^'s in ` ^name^ = value`.
		//	// stateAttrName出现在属性名称内。
		//	// 它发生在`^ name ^ = value`中的^之间。
		//	stateAttrName
		//	// stateAfterName occurs after an attr name has ended but before any
		//	// equals sign. It occurs between the ^'s in ` name^ ^= value`.
		//	// stateAfterName在attr名称结束之后但在任何等号之前发生。 它发生在“ name ^ ^ = value”中的^之间。
		//	stateAfterName
		//	// stateBeforeValue occurs after the equals sign but before the value.
		//	// It occurs between the ^'s in ` name =^ ^value`.
		//	// stateBeforeValue出现在等号之后但在值之前。
		//	// 它发生在`name = ^ ^ value`中的^之间。
		//	stateBeforeValue
		//	// stateHTMLCmt occurs inside an <!-- HTML comment -->.
		//	// stateHTMLCmt出现在<!-- HTML comment -->内部。
		//	stateHTMLCmt
		//	// stateRCDATA occurs inside an RCDATA element (<textarea> or <title>)
		//	// as described at https://www.w3.org/TR/html5/syntax.html#elements-0
		//	// stateRCDATA发生在RCDATA元素（<textarea>或<title>）内部，如https://www.w3.org/TR/html5/syntax.html#elements-0所述
		//	stateRCDATA
		//	// stateAttr occurs inside an HTML attribute whose content is text.
		//	// stateAttr发生在内容为文本的HTML属性中。
		//	stateAttr
		//	// stateURL occurs inside an HTML attribute whose content is a URL.
		//	// stateURL发生在内容为URL的HTML属性内。
		//	stateURL
		//	// stateSrcset occurs inside an HTML srcset attribute.
		//	// stateSrcset出现在HTML srcset属性内。
		//	stateSrcset
		//	// stateJS occurs inside an event handler or script element.
		//	// stateJS发生在事件处理程序或脚本元素内。
		//	stateJS
		//	// stateJSDqStr occurs inside a JavaScript double quoted string.
		//	// stateJSDqStr出现在JavaScript双引号字符串内。
		//	stateJSDqStr
		//	// stateJSSqStr occurs inside a JavaScript single quoted string.
		//	// stateJSSqStr出现在JavaScript单引号字符串内。
		//	stateJSSqStr
		//	// stateJSRegexp occurs inside a JavaScript regexp literal.
		//	// stateJSRegexp发生在JavaScript regexp文字内部。
		//	stateJSRegexp
		//	// stateJSBlockCmt occurs inside a JavaScript /* block comment */.
		//	// stateJSBlockCmt发生在JavaScript /* 块注释 */中。
		//	stateJSBlockCmt
		//	// stateJSLineCmt occurs inside a JavaScript // line comment.
		//	// stateJSLineCmt发生在JavaScript行注释中。
		//	stateJSLineCmt
		//	// stateCSS occurs inside a <style> element or style attribute.
		//	// stateCSS发生在<style>元素或style属性中。
		//	stateCSS
		//	// stateCSSDqStr occurs inside a CSS double quoted string.
		//	// stateCSSDqStr出现在CSS双引号字符串内。
		//	stateCSSDqStr
		//	// stateCSSSqStr occurs inside a CSS single quoted string.
		//	// stateCSSSqStr出现在CSS单引号字符串内。
		//	stateCSSSqStr
		//	// stateCSSDqURL occurs inside a CSS double quoted url("...").
		//	// stateCSSDqURL出现在CSS双引号url（“ ...”）中。
		//	stateCSSDqURL
		//	// stateCSSSqURL occurs inside a CSS single quoted url('...').
		//	// stateCSSSqURL出现在CSS单引号url（'...'）中。
		//	stateCSSSqURL
		//	// stateCSSURL occurs inside a CSS unquoted url(...).
		//	// stateCSSURL发生在CSS未加引号的url（...）中。
		//	stateCSSURL
		//	// stateCSSBlockCmt occurs inside a CSS /* block comment */.
		//	// stateCSSBlockCmt发生在CSS / *块注释* /中。
		//	stateCSSBlockCmt
		//	// stateCSSLineCmt occurs inside a CSS // line comment.
		//	// stateCSSLineCmt发生在CSS行注释中。
		//	stateCSSLineCmt
		//	// stateError is an infectious error state outside any valid
		//	// HTML/CSS/JS construct.
		//	// stateError是任何有效HTML / CSS / JS构造之外的传染性错误状态。
		//	stateError
		//)



		//// delim is the delimiter that will end the current HTML attribute.
		//// delim是将结束当前HTML属性的定界符。
		//type delim uint8
		//
		////go:generate stringer -type delim
		//
		//const (
		//	// delimNone occurs outside any attribute.
		//	// delimNone出现在任何属性之外。
		//	delimNone delim = iota
		//	// delimDoubleQuote occurs when a double quote (") closes the attribute.
		//	//当双引号(")关闭属性时，发生delimDoubleQuote。
		//	delimDoubleQuote
		//	// delimSingleQuote occurs when a single quote (') closes the attribute.
		//	//当单引号(')关闭属性时，发生delimSingleQuote。
		//	delimSingleQuote
		//	// delimSpaceOrTagEnd occurs when a space or right angle bracket (>)
		//	// closes the attribute.
		//	//当空格或直角括号（>）关闭属性时，发生delimSpaceOrTagEnd。
		//	delimSpaceOrTagEnd
		//)




		//// urlPart identifies a part in an RFC 3986 hierarchical URL to allow different
		//// encoding strategies.
		//// urlPart标识RFC 3986分层URL中的一部分，以允许使用不同的编码策略。
		//type urlPart uint8
		//
		////go:generate stringer -type urlPart
		//
		//const (
		//	// urlPartNone occurs when not in a URL, or possibly at the start:
		//	// ^ in "^http://auth/path?k=v#frag".
		//	// urlPartNone不在URL中或可能在开头时出现：^ in "^http://auth/path?k=v#frag".
		//	urlPartNone urlPart = iota
		//	// urlPartPreQuery occurs in the scheme, authority, or path; between the
		//	// ^s in "h^ttp://auth/path^?k=v#frag".
		//	// urlPartPreQuery出现在方案，权限或路径中； 在^s in "h^ttp://auth/path^?k=v#frag"之间。
		//	urlPartPreQuery
		//	// urlPartQueryOrFrag occurs in the query portion between the ^s in
		//	// "http://auth/path?^k=v#frag^".
		//	// urlPartQueryOrFrag出现在"http://auth/path?^k=v#frag^"中^s之间的查询部分中。
		//	urlPartQueryOrFrag
		//	// urlPartUnknown occurs due to joining of contexts both before and
		//	// after the query separator.
		//	// urlPartUnknown发生是由于在查询分隔符之前和之后都加入了上下文。
		//	urlPartUnknown
		//)



		//// jsCtx determines whether a '/' starts a regular expression literal or a
		//// division operator.
		//// jsCtx确定'/'是启动正则表达式文字还是除法运算符。
		//type jsCtx uint8
		//
		////go:generate stringer -type jsCtx
		//
		//const (
		//	// jsCtxRegexp occurs where a '/' would start a regexp literal.
		//	//发生jsCtxRegexp时，“ /”将开始一个regexp文字。
		//	jsCtxRegexp jsCtx = iota
		//	// jsCtxDivOp occurs where a '/' would start a division operator.
		//	// jsCtxDivOp发生在'/'将启动除法运算符的位置。
		//	jsCtxDivOp
		//	// jsCtxUnknown occurs where a '/' is ambiguous due to context joining.
		//	// jsCtxUnknown发生在由于上下文连接而导致“ /”不明确的情况下。
		//	jsCtxUnknown
		//)





		////go:generate stringer -type attr
		//
		//// attr identifies the current HTML attribute when inside the attribute,
		//// that is, starting from stateAttrName until stateTag/stateText (exclusive).
		//// attr标识属性内的当前HTML属性，即从stateAttrName到stateTag / stateText（不包括）。
		//type attr uint8
		//
		//const (
		//	// attrNone corresponds to a normal attribute or no attribute.
		//	// attrNone对应于普通属性或无属性。
		//	attrNone attr = iota
		//	// attrScript corresponds to an event handler attribute.
		//	// attrScript对应于事件处理程序属性。
		//	attrScript
		//	// attrScriptType corresponds to the type attribute in script HTML element
		//	// attrScriptType对应于脚本HTML元素中的type属性
		//	attrScriptType
		//	// attrStyle corresponds to the style attribute whose value is CSS.
		//	// attrStyle对应于样式属性，其值为CSS。
		//	attrStyle
		//	// attrURL corresponds to an attribute whose value is a URL.
		//	// attrURL对应于其值为URL的属性。
		//	attrURL
		//	// attrSrcset corresponds to a srcset attribute.
		//	// attrSrcset对应于srcset属性。
		//	attrSrcset
		//)



		//// element identifies the HTML element when inside a start tag or special body.
		//// Certain HTML element (for example <script> and <style>) have bodies that are
		//// treated differently from stateText so the element type is necessary to
		//// transition into the correct context at the end of a tag and to identify the
		//// end delimiter for the body.
		////元素在开始标签或特殊正文中时标识HTML元素。
		////某些HTML元素（例如<script>和<style>）的正文与stateText的处理方式有所不同，
		//// 因此必须使用元素类型才能在标记的末尾过渡到正确的上下文并标识该末尾的定界符 body。
		//type element uint8
		//
		////go:generate stringer -type element
		//
		//const (
		//	// elementNone occurs outside a special tag or special element body.
		//	// elementNone不在特殊标记tag或特殊元素body之外。
		//	elementNone element = iota
		//	// elementScript corresponds to the raw text <script> element
		//	// with JS MIME type or no type attribute.
		//	// elementScript对应于具有JS MIME类型或没有type属性的原始文本<script>元素。
		//	elementScript
		//	// elementStyle corresponds to the raw text <style> element.
		//	// elementStyle对应于原始文本<style>元素。
		//	elementStyle
		//	// elementTextarea corresponds to the RCDATA <textarea> element.
		//	// elementTextarea对应于RCDATA <textarea>元素。
		//	elementTextarea
		//	// elementTitle corresponds to the RCDATA <title> element.
		//	// elementTitle对应于RCDATA <title>元素。
		//	elementTitle
		//)




		//// Error describes a problem encountered during template Escaping.
		////Error描述转义模板时遇到的问题。
		//type Error struct {
		//	// ErrorCode describes the kind of error.
		//	// ErrorCode描述错误的类型。
		//	ErrorCode ErrorCode
		//	// Node is the node that caused the problem, if known.
		//	// If not nil, it overrides Name and Line.
		//	//Node是导致问题的节点（如果已知）。
		//	//如果不是nil，它将覆盖Name和Line。
		//	Node parse.Node
		//	// Name is the name of the template in which the error was encountered.
		//	// Name是遇到错误的模板的名称。
		//	Name string
		//	// Line is the line number of the error in the template source or 0.
		//	// Line是模板源中错误的行号或0。
		//	Line int
		//	// Description is a human-readable description of the problem.
		//	//Description是人类可读的问题说明。
		//	Description string
		//}
		//
		//// ErrorCode is a code for a kind of error.
		//// ErrorCode是一种错误的错误码，标志标错的类型。
		//type ErrorCode int
		//
		//// We define codes for each error that manifests while escaping templates, but
		//// escaped templates may also fail at runtime.
		////我们为转义模板时出现的每个错误定义代码，但是转义的模板也可能在运行时失败。
		////
		//// Output: "ZgotmplZ"
		//// Example:
		////   <img src="{{.X}}">
		////   where {{.X}} evaluates to `javascript:...`
		//// Discussion:
		////   "ZgotmplZ" is a special value that indicates that unsafe content reached a
		////   CSS or URL context at runtime. The output of the example will be
		////     <img src="#ZgotmplZ">
		////   If the data comes from a trusted source, use content types to exempt it
		////   from filtering: URL(`javascript:...`).
		////讨论：
		////“ ZgotmplZ”是一个特殊值，指示运行时不安全内容到达CSS或URL上下文。 该示例的输出将是
		//// <img src =“＃ZgotmplZ”>
		////如果数据来自受信任的来源，请使用内容类型将其免除过滤：URL（`javascript：...`）。
		//const (
		//	// OK indicates the lack of an error.// OK表示没有错误。
		//	OK ErrorCode = iota
		//
		//	// ErrAmbigContext: "... appears in an ambiguous context within a URL"
		//	// Example:
		//	//   <a href="
		//	//      {{if .C}}
		//	//        /path/
		//	//      {{else}}
		//	//        /search?q=
		//	//      {{end}}
		//	//      {{.X}}
		//	//   ">
		//	// Discussion:
		//	//   {{.X}} is in an ambiguous URL context since, depending on {{.C}},
		//	//  it may be either a URL suffix or a query parameter.
		//	//   Moving {{.X}} into the condition removes the ambiguity:
		//	//   <a href="{{if .C}}/path/{{.X}}{{else}}/search?q={{.X}}">
		//	ErrAmbigContext
		//
		//	// ErrBadHTML: "expected space, attr name, or end of tag, but got ...",
		//	//   "... in unquoted attr", "... in attribute name"
		//	// Example:
		//	//   <a href = /search?q=foo>
		//	//   <href=foo>
		//	//   <form na<e=...>
		//	//   <option selected<
		//	// Discussion:
		//	//   This is often due to a typo in an HTML element, but some runes
		//	//   are banned in tag names, attribute names, and unquoted attribute
		//	//   values because they can tickle parser ambiguities.
		//	//   Quoting all attributes is the best policy.
		//	ErrBadHTML
		//
		//	// ErrBranchEnd: "{{if}} branches end in different contexts"
		//	// Example:
		//	//   {{if .C}}<a href="{{end}}{{.X}}
		//	// Discussion:
		//	//   Package html/template statically examines each path through an
		//	//   {{if}}, {{range}}, or {{with}} to escape any following pipelines.
		//	//   The example is ambiguous since {{.X}} might be an HTML text node,
		//	//   or a URL prefix in an HTML attribute. The context of {{.X}} is
		//	//   used to figure out how to escape it, but that context depends on
		//	//   the run-time value of {{.C}} which is not statically known.
		//	//
		//	//   The problem is usually something like missing quotes or angle
		//	//   brackets, or can be avoided by refactoring to put the two contexts
		//	//   into different branches of an if, range or with. If the problem
		//	//   is in a {{range}} over a collection that should never be empty,
		//	//   adding a dummy {{else}} can help.
		//	ErrBranchEnd
		//
		//	// ErrEndContext: "... ends in a non-text context: ..."
		//	// Examples:
		//	//   <div
		//	//   <div title="no close quote>
		//	//   <script>f()
		//	// Discussion:
		//	//   Executed templates should produce a DocumentFragment of HTML.
		//	//   Templates that end without closing tags will trigger this error.
		//	//   Templates that should not be used in an HTML context or that
		//	//   produce incomplete Fragments should not be executed directly.
		//	//
		//	//   {{define "main"}} <script>{{template "helper"}}</script> {{end}}
		//	//   {{define "helper"}} document.write(' <div title=" ') {{end}}
		//	//
		//	//   "helper" does not produce a valid document fragment, so should
		//	//   not be Executed directly.
		//	ErrEndContext
		//
		//	// ErrNoSuchTemplate: "no such template ..."
		//	// Examples:
		//	//   {{define "main"}}<div {{template "attrs"}}>{{end}}
		//	//   {{define "attrs"}}href="{{.URL}}"{{end}}
		//	// Discussion:
		//	//   Package html/template looks through template calls to compute the
		//	//   context.
		//	//   Here the {{.URL}} in "attrs" must be treated as a URL when called
		//	//   from "main", but you will get this error if "attrs" is not defined
		//	//   when "main" is parsed.
		//	ErrNoSuchTemplate
		//
		//	// ErrOutputContext: "cannot compute output context for template ..."
		//	// Examples:
		//	//   {{define "t"}}{{if .T}}{{template "t" .T}}{{end}}{{.H}}",{{end}}
		//	// Discussion:
		//	//   A recursive template does not end in the same context in which it
		//	//   starts, and a reliable output context cannot be computed.
		//	//   Look for typos in the named template.
		//	//   If the template should not be called in the named start context,
		//	//   look for calls to that template in unexpected contexts.
		//	//   Maybe refactor recursive templates to not be recursive.
		//	ErrOutputContext
		//
		//	// ErrPartialCharset: "unfinished JS regexp charset in ..."
		//	// Example:
		//	//     <script>var pattern = /foo[{{.Chars}}]/</script>
		//	// Discussion:
		//	//   Package html/template does not support interpolation into regular
		//	//   expression literal character sets.
		//	ErrPartialCharset
		//
		//	// ErrPartialEscape: "unfinished escape sequence in ..."
		//	// Example:
		//	//   <script>alert("\{{.X}}")</script>
		//	// Discussion:
		//	//   Package html/template does not support actions following a
		//	//   backslash.
		//	//   This is usually an error and there are better solutions; for
		//	//   example
		//	//     <script>alert("{{.X}}")</script>
		//	//   should work, and if {{.X}} is a partial escape sequence such as
		//	//   "xA0", mark the whole sequence as safe content: JSStr(`\xA0`)
		//	ErrPartialEscape
		//
		//	// ErrRangeLoopReentry: "on range loop re-entry: ..."
		//	// Example:
		//	//   <script>var x = [{{range .}}'{{.}},{{end}}]</script>
		//	// Discussion:
		//	//   If an iteration through a range would cause it to end in a
		//	//   different context than an earlier pass, there is no single context.
		//	//   In the example, there is missing a quote, so it is not clear
		//	//   whether {{.}} is meant to be inside a JS string or in a JS value
		//	//   context. The second iteration would produce something like
		//	//
		//	//     <script>var x = ['firstValue,'secondValue]</script>
		//	ErrRangeLoopReentry
		//
		//	// ErrSlashAmbig: '/' could start a division or regexp.
		//	// Example:
		//	//   <script>
		//	//     {{if .C}}var x = 1{{end}}
		//	//     /-{{.N}}/i.test(x) ? doThis : doThat();
		//	//   </script>
		//	// Discussion:
		//	//   The example above could produce `var x = 1/-2/i.test(s)...`
		//	//   in which the first '/' is a mathematical division operator or it
		//	//   could produce `/-2/i.test(s)` in which the first '/' starts a
		//	//   regexp literal.
		//	//   Look for missing semicolons inside branches, and maybe add
		//	//   parentheses to make it clear which interpretation you intend.
		//	ErrSlashAmbig
		//
		//	// ErrPredefinedEscaper: "predefined escaper ... disallowed in template"
		//	// Example:
		//	//   <div class={{. | html}}>Hello<div>
		//	// Discussion:
		//	//   Package html/template already contextually escapes all pipelines to
		//	//   produce HTML output safe against code injection. Manually escaping
		//	//   pipeline output using the predefined escapers "html" or "urlquery" is
		//	//   unnecessary, and may affect the correctness or safety of the escaped
		//	//   pipeline output in Go 1.8 and earlier.
		//	//
		//	//   In most cases, such as the given example, this error can be resolved by
		//	//   simply removing the predefined escaper from the pipeline and letting the
		//	//   contextual autoescaper handle the escaping of the pipeline. In other
		//	//   instances, where the predefined escaper occurs in the middle of a
		//	//   pipeline where subsequent commands expect escaped input, e.g.
		//	//     {{.X | html | makeALink}}
		//	//   where makeALink does
		//	//     return `<a href="`+input+`">link</a>`
		//	//   consider refactoring the surrounding template to make use of the
		//	//   contextual autoescaper, i.e.
		//	//     <a href="{{.X}}">link</a>
		//	//
		//	//   To ease migration to Go 1.9 and beyond, "html" and "urlquery" will
		//	//   continue to be allowed as the last command in a pipeline. However, if the
		//	//   pipeline occurs in an unquoted attribute value context, "html" is
		//	//   disallowed. Avoid using "html" and "urlquery" entirely in new templates.
		//	ErrPredefinedEscaper
		//	//上面有些没翻译，请自助！
		//	)




//// Strings of content from a trusted source.//来自可信来源的内容字符串。
//type (
//	// CSS encapsulates known safe content that matches any of:
//	//   1. The CSS3 stylesheet production, such as `p { color: purple }`.
//	//   2. The CSS3 rule production, such as `a[href=~"https:"].foo#bar`.
//	//   3. CSS3 declaration productions, such as `color: red; margin: 2px`.
//	//   4. The CSS3 value production, such as `rgba(0, 0, 255, 127)`.
//	// See https://www.w3.org/TR/css3-syntax/#parsing and
//	// https://web.archive.org/web/20090211114933/http://w3.org/TR/css3-syntax#style
//	//
//	// Use of this type presents a security risk:
//	// the encapsulated content should come from a trusted source,
//	// as it will be included verbatim in the template output.
//
//	// CSS封装了符合以下条件的已知安全内容：
//	// 1. CSS3样式表的产生，例如`p { color: purple }`。
//	// 2. CSS3规则产生，例如`a[href=~"https:"].foo#bar`。
//	// 3. CSS3声明产生，例如`color: red; margin: 2px`。
//	// 4. CSS3值的产生，例如`rgba(0, 0, 255, 127)`。
//	//参见https://www.w3.org/TR/css3-syntax/#parsing和https://web.archive.org/web/20090211114933/http://w3.org/TR/css3-syntax# 样式
//	//
//	//使用这种类型会带来安全风险：
//	//封装的内容应来自受信任的来源，因为它将逐字包含在模板输出中。
//
//	CSS string
//
//	// HTML encapsulates a known safe HTML document fragment.
//	// It should not be used for HTML from a third-party, or HTML with
//	// unclosed tags or comments. The outputs of a sound HTML sanitizer
//	// and a template escaped by this package are fine for use with HTML.
//	//
//	// Use of this type presents a security risk:
//	// the encapsulated content should come from a trusted source,
//	// as it will be included verbatim in the template output.
//
//	// HTML封装了一个已知的安全HTML文档片段。
//	//不应将其用于来自第三方的HTML或带有未封闭标签或注释的HTML。 声音HTML清理程序的输出和此程序包转义的模板非常适合与HTML一起使用。
//	//
//	//使用这种类型会带来安全风险：
//	//封装的内容应来自受信任的来源，因为它将逐字包含在模板输出中。
//
//	HTML string
//
//	// HTMLAttr encapsulates an HTML attribute from a trusted source,
//	// for example, ` dir="ltr"`.
//	//
//	// Use of this type presents a security risk:
//	// the encapsulated content should come from a trusted source,
//	// as it will be included verbatim in the template output.
//
//	// HTMLAttr封装了来自受信任来源的HTML属性，例如` dir="ltr"`。
//	//
//	//使用这种类型会带来安全风险：
//	//封装的内容应来自受信任的来源，因为它将逐字包含在模板输出中。
//
//	HTMLAttr string
//
//	// JS encapsulates a known safe EcmaScript5 Expression, for example,
//	// `(x + y * z())`.
//	// Template authors are responsible for ensuring that typed expressions
//	// do not break the intended precedence and that there is no
//	// statement/expression ambiguity as when passing an expression like
//	// "{ foo: bar() }\n['foo']()", which is both a valid Expression and a
//	// valid Program with a very different meaning.
//	//
//	// Use of this type presents a security risk:
//	// the encapsulated content should come from a trusted source,
//	// as it will be included verbatim in the template output.
//	//
//	// Using JS to include valid but untrusted JSON is not safe.
//	// A safe alternative is to parse the JSON with json.Unmarshal and then
//	// pass the resultant object into the template, where it will be
//	// converted to sanitized JSON when presented in a JavaScript context.
//
//	// JS封装了一个已知的安全EcmaScript5表达式，例如`(x + y * z())`。
//	//模板作者负责确保键入的表达式不会破坏预期的优先级，并且在传递诸如"{ foo: bar() }\n['foo']()"之类的表达式时，
//	// 不存在语句/表达式的歧义 ，它既是有效的Expression，又是含义完全不同的有效程序。
//	//
//	//使用这种类型会带来安全风险：
//	//封装的内容应来自受信任的来源，因为它将逐字包含在模板输出中。
//	//
//	//使用JS包含有效但不受信任的JSON是不安全的。
//	//一种安全的选择是使用json.Unmarshal解析JSON，然后将结果对象传递到模板中，在JavaScript上下文中呈现时，该模板将转换为经过净化的JSON。
//
//	JS string
//
//	// JSStr encapsulates a sequence of characters meant to be embedded
//	// between quotes in a JavaScript expression.
//	// The string must match a series of StringCharacters:
//	//   StringCharacter :: SourceCharacter but not `\` or LineTerminator
//	//                    | EscapeSequence
//	// Note that LineContinuations are not allowed.
//	// JSStr("foo\\nbar") is fine, but JSStr("foo\\\nbar") is not.
//	//
//	// Use of this type presents a security risk:
//	// the encapsulated content should come from a trusted source,
//	// as it will be included verbatim in the template output.
//
//	// JSStr封装了一系列要嵌入在JavaScript表达式中的引号之间的字符。
//	//字符串必须与一系列StringCharacters匹配：
//	// StringCharacter :: SourceCharacter，但不是`\`或LineTerminator |  EscapeSequence(转义序列)
//	//请注意，不允许LineContinuations。
//	// JSStr("foo\\nbar") 很好，但是JSStr("foo\\\nbar")不好。
//	//
//	//使用这种类型会带来安全风险：
//	//封装的内容应来自受信任的来源，因为它将逐字包含在模板输出中。
//
//	JSStr string
//
//	// URL encapsulates a known safe URL or URL substring (see RFC 3986).
//	// A URL like `javascript:checkThatFormNotEditedBeforeLeavingPage()`
//	// from a trusted source should go in the page, but by default dynamic
//	// `javascript:` URLs are filtered out since they are a frequently
//	// exploited injection vector.
//	//
//	// Use of this type presents a security risk:
//	// the encapsulated content should come from a trusted source,
//	// as it will be included verbatim in the template output.
//
//	// URL封装了一个已知的安全URL或URL子字符串（请参阅RFC 3986）。
//	//来自可信来源的URL，例如“ javascript：checkThatFormNotEditedBeforeLeavingPage（）”，应该放在页面中，但是默认情况下，动态JavaScriptURL被过滤掉，因为它们是经常被利用的注入向量。
//	//
//	//使用这种类型会带来安全风险：
//	//封装的内容应来自受信任的来源，因为它将逐字包含在模板输出中。
//
//	URL string
//
//	// Srcset encapsulates a known safe srcset attribute
//	// (see https://w3c.github.io/html/semantics-embedded-content.html#element-attrdef-img-srcset).
//	//
//	// Use of this type presents a security risk:
//	// the encapsulated content should come from a trusted source,
//	// as it will be included verbatim in the template output.
//
//	// Srcset封装了一个已知的安全srcset属性（请参阅https://w3c.github.io/html/semantics-embedded-content.html#element-attrdef-img-srcset）。
//	//
//	//使用这种类型会带来安全风险：
//	//封装的内容应来自受信任的来源，因为它将逐字包含在模板输出中。
//
//	Srcset string
//)




func main() {

	//转义与不转义含义：<表达原义，字符串中我怎么写你就怎么存，这就是不转义，但是如果转义的话，则是我怎么写，但是当遇到特殊字符比如<时候会自动将
	//我的输入的<号解释成为原生的"&lt;"来表示，则这就是转义了！
	//		字符		  转义后的字符
	//		`&`, 		"&amp;",
	//		`'`, 		"&#39;",
	//		`<`, 		"&lt;",
	//		`>`, 		"&gt;",
	//		`"`, 		"&#34;",


	fmt.Println("----text.template包与html.template包下生成的字符串区别----")

	fmt.Println("text.template包:")
	t_text, err := Ttext.New("t_text").Parse(`{{define "T1"}}Hello, {{.}}!{{end}}`)
	check_err_html(err)
	fmt.Println(t_text.Tree.Root.Nodes)
	err = t_text.ExecuteTemplate(os.Stdout, "T1", "<script>alert('you have been pwned')</script>")


	fmt.Println()//不可去掉，上面的os.Stdout并没有换行
	fmt.Println("html.template包:")
	t_html, err := Thtml.New("t_html").Parse(`{{define "T2"}}Hello, {{.}}!{{end}}`)
	check_err_html(err)
	fmt.Println(t_html.Tree.Root.Nodes)
	err = t_html.ExecuteTemplate(os.Stdout, "T2", "<script>alert('you have been pwned')</script>")

	//输出：
	//	----text.template包与html.template包下生成的字符串区别----
	//	text.template包:
	//	Hello, <script>alert('you have been pwned')</script>!
	//	html.template包:
	//	Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!
	//从上面可以看得出，text.template包并不会自动对输出的字符串进行转义，除非自己手动转义，而html.template则会自动转义而不用手动转义！


	fmt.Println()
	fmt.Println("--------------html.tempalte和text.template共有的一些方法不会累叙了-------------------")
	//以下方法如果你已经学过了text.template包，那么你应该对这个包下的这些方法会非常了解了!
	//template.New()
	//template.ParseGlob()
	//template.ParseFiles()
	//template.Must()
	//template.IsTrue()
	//template.HTMLEscape()
	//template.HTMLEscapeString()
	//template.JSEscape()
	//template.JSEscaper()
	//template.JSEscapeString()
	//template.URLQueryEscaper()
	////template.URLQueryEscape()//html.template包没有这个方法，这是区别于text.template包的地方


	//下面同样是html.template对象下 的方法，几乎完全和text.template对象的方法一样的！子对象tree也完全相同

	fmt.Println("t_html.Tree.Name:",t_html.Tree.Name)
	fmt.Println("t_html.Tree.Root:",t_html.Tree.Root)
	fmt.Println("t_html.Tree.Root.Nodes:",t_html.Tree.Root.Nodes)//因为我们t_html并不是列表node，所以这里为nil,下面我们故意让她出现NodeList
	fmt.Println("t_html.Tree.Root.String():",t_html.Tree.Root.String())
	fmt.Println("t_html.Tree.ParseName:",t_html.Tree.ParseName)
	fmt.Println("t_html.Tree.Copy():",t_html.Tree.Copy())
	//t_html.Tree.Parse()
	//t_html.Tree.ErrorContext()


	//t_html.ParseFiles()
	//t_html.Parse()
	//t_html.ParseGlob()
	//t_html.New()
	//t_html.Name()
	//t_html.AddParseTree()
	//t_html.DefinedTemplates()
	//t_html.Delims()
	//t_html.Templates()
	//t_html.Clone()
	//t_html.Lookup()
	//t_html.Option()
	//t_html.Execute()
	//t_html.ExecuteTemplate()
	//t_html.Funcs()


	//parse_str:=``

	fmt.Println("--------------")
	//{{template "T4"}}后面不用end注意了
	t_html1, err := Thtml.New("T3").Parse(`{{define "T3"}}{{template "T4" .}}{{range $k, $v:= .ls1}}索引为：{{$k}}，值为：{{println $v}}{{end}}{{end}}`)
	check_err_html(err)

	//事实上此时接不接受返回值都是可以的，因为他们是相同的底层模板索引
	t_html2, err := t_html1.Parse(`{{define "T4"}}{{range $k, $v:= .ls2}}索引为：{{$k}}，值为：{{println $v}}{{end}}{{end}}`)

	check_err_html(err)
	m:=map[string]interface{}{
		"ls1":[]string{"a<>","b\"","c'","d&"},
		//"ls1":[]string{"a<>","b\"","c'","d&"},
		"ls2":[]string{"a1","b1","c1","d1"},
	}
	//上面我们嵌入了一个模板T4才终于有值了！不然的话，Root.Nodes是只会为空！你可以看上面t_html对象
	// 的输出结果就知道了，root是树根节点模板字符串，而nodes是根节点下面的子树模板字符串！如果没有的话则为nil.
	//这个不限于这个包，text.template包和text.template.parse包也是如此！算是对之前的一个知识点的补充吧！
	fmt.Println(t_html1.Tree.Root.Nodes)
	//比text.template.execute()方法多了一个escape转义的过程，很多方法都是这样，不再这里继续累叙了！有疑问的可以自己尝试！
	err = t_html1.Execute(os.Stdout,  m)

	fmt.Println("=====")

	fmt.Println(t_html2.Tree.Root.Nodes)
	err = t_html2.Execute(os.Stdout,  m)
	//我上面之所以要运行t_html2和t_html1就是为了验证他们2个是底层完全相同的text.template对象，而不是不同的.
	//就是因为是底层完全相同的模板，所以，只要有一个模板对象运行了Execute（）方法的话，则底层的模板字符串就会被转义了，只要被转义，
	//后一个被运行的模板也会输出相同的已经被转义的template字符串对象！

	//输出：
	//	--------------html.tempalte和text.template共有的一些方法不会累叙了-------------------
	//	t_html.Tree.Name: t_html
	//	t_html.Tree.Root:
	//	t_html.Tree.Root.Nodes: []
	//	t_html.Tree.Root.String():
	//	t_html.Tree.ParseName: t_html
	//	t_html.Tree.Copy(): &{t_html t_html  {{define "T2"}}Hello, {{.}}!{{end}} [] <nil> [{0 0  0} {0 0  0} {0 0  0}] 0 [] map[]}
	//	--------------
	//	[{{template "T4" .}} {{range $k, $v := .ls1}}索引为：{{$k}}，值为：{{println $v}}{{end}}]
	//	索引为：0，值为：a1
	//	索引为：1，值为：b1
	//	索引为：2，值为：c1
	//	索引为：3，值为：d1
	//	索引为：0，值为：a&lt;&gt;
	//	索引为：1，值为：b&#34;
	//	索引为：2，值为：c&#39;
	//	索引为：3，值为：d&amp;
	//	=====
	//	[{{template "T4" .}} {{range $k, $v := .ls1}}索引为：{{$k | _html_template_htmlescaper}}，值为：{{println $v | _html_template_htmlescaper}}{{end}}]
	//	索引为：0，值为：a1
	//	索引为：1，值为：b1
	//	索引为：2，值为：c1
	//	索引为：3，值为：d1
	//	索引为：0，值为：a&lt;&gt;
	//	索引为：1，值为：b&#34;
	//	索引为：2，值为：c&#39;
	//	索引为：3，值为：d&amp;



	fmt.Println("-------如果要不希望某些字符串被转义---------")


	t_html6, err := Thtml.New("T6").Parse(`{{.str1}}---{{.str2}}`)
	check_err_html(err)

	m6:=map[string]Thtml.HTML{
		"str1":Thtml.HTML("<b>Hello</b>"),
		"str2":"<b>World</b>",
	}//虽然命名类型不可以相互赋值，但是很明显上面的str2已经被string类型已经被go自动转成了Thtml.HTML类型，

	err = t_html6.Execute(os.Stdout, m6)
	check_err_html(err)

	fmt.Println()
	t_html7, err := Thtml.New("T7").Parse(`{{.str1}}---{{.str2}}`)
	m7:=map[string]interface{}{//像这样用这种类型的话就比较好控制到底转义还是不转义了
		"str1":Thtml.HTML("<b>Hello</b>"),
		"str2":"<b>World</b>",
	}//虽然命名类型不可以相互赋值，但是很明显上面的str2已经被string类型已经被go自动转成了Thtml.HTML类型，

	err = t_html7.Execute(os.Stdout, m7)
	check_err_html(err)
	// 输出：
	// -------如果要不希望某些字符串被转义---------
	// <b>Hello</b>---<b>World</b>
	// <b>Hello</b>---&lt;b&gt;World&lt;/b&gt;

	//除了Thtml.HTML这个类之外，还有以下，不再举例逐个说明了！
	//	type CSS string
	//	type HTMLAttr string
	//	type JS string
	//	type JSStr string
	//	type URL string

}




func check_err_html(err error) {
	if err != nil {
		//fmt.Fprintln(os.Stderr,err)
		//上面的这种方式会导致输出顺序不确定，虽然他可以输出红色的字体，但是由于顺序不确定，我们不采用他！
		fmt.Println("出错了，错误信息为：",err)
	}
}
