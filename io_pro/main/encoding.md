encoding包定义了供其它包使用的可以将数据在字节水平和文本表示之间转换的接口。encoding/gob、encoding/json、encoding/xml三个包都会检查使用这些接口。因此，只要实现了这些接口一次，就可以在多个包里使用。标准包内建类型time.Time和net.IP都实现了这些接口。接口是成对的，分别产生和还原编码后的数据。

主要有下面的结果方法：

[type BinaryMarshaler]

[type BinaryUnmarshaler]

[type TextMarshaler]

[type TextUnmarshaler]

####  这个包一般不被用到，他会被其他包继承！所以不写！事实上在其他encoding包中已经被写到了，请去了解其他包中的实现类！

