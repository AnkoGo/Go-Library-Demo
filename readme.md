#  《Go标准库例子大全	Go-Library-Demo》

#  项目特点：

- 涵盖广，细节深，全网较全较深的go库学习笔录，目前没有之一！
- 注释多，解释深刻到位！通俗易懂，小白易上手！
- 摒弃上线代码的风格（不易理解），力求通过最小最易懂的例子来讲解api用法，精准定位api同时不拘泥于一种固定代码的写法！
- 缺点：
  - 目前目录结构较为混乱，代码注释不大整齐！

# 涵盖度：
目前大部分库基本写了，除了net，syscall，context包外！目前不打算写runtime，debug包以及加密包！

# 下载和执行须知：  
在这里需要特别指出的是，鉴于学识有限，此项目会存在以下问题：

- 所有代码均在main目录下，其他目录均是辅助目录和文件
- 有部分的注释和代码是有误或者不妥的！待以后跟新！
- 不是所有代码都是我自己写的，有大概百分之10的代码量是参考go源码来写的！但是本人保证绝大部分的代码还是属于自己的！  
- 本代码中的所有文件一定要分开来执行，最好复制main中的文件内容然后放到main3下的任意一个.go文件中执行，我就是这样的！这是因为go不允许多个文件拥有main()函数，因此，项目的几乎全部的.go文件的main方法名都已经改成了main+任意的多个数字！所以在你单独执行时候需要去除这些数字，我还是特别建议复制要执行的.go文件内容然后放到main3目录下去执行，只要不同的目录就可以了！但是有些是必须在main目录下执行的，因为main目录下还有其他的目录和目录下的文件或者子目录，在我的代码执行时候可能会依赖这些目录和文件！如果你把代码放到单独的目录下的单独文件中去执行时候无法正常执行，则说明要么是该执行脚本的目录没有对应的其他必须目录和文件，要么就是被goland等等其他的编辑器清除了某些必须包的导入，此时你可以放回main中去执行（前提是必须保证main目录下只有一个main()函数，这点非常难受），所有遇到的这些问题，我都会在后续持续跟新！  
- 关于代码中使用到的变量命名难看问题，因为api实在太多，多到吓人，我能想到的变量命名都已经用光了！体谅下！  
- 关于本项目中的api用法以及全不全问题，我能在goland中智能提示的go 库的类对象，接口，函数基本都完完整整的写了一遍，并且做了必要的注释！但是我不保证这个注释一定正确无误！但hub主保证尽自己最大的能力了！  
- 关于test目录：这个目录是我学习时候测试的一些东西，不是main目录下的代码的测试！注意了！  
- 其他的一些文件和目录都是为了测试api的用法或者功能而创建的！  
- 本人英语不大好，有些是参考翻译工具的！还有些是参考go中文网！在此非常感谢[go中文网]( https://studygolang.com/ )！以及谷歌翻译和网上的平时不大留意的著作人，因为学习时候截取过很多人的代码，而现在我无法找到原作者！如果您在我的代码中看到了您的代码并且希望标注出处的话，请告诉我，我会及时处理！  
- 关于一个文件代码过长问题，我当初写的时候没想到过开源的！但是已经这样写了，没办法，后期会把大文件分成较小的文件  
- 关于侵权问题，本项目开源，不获取任何利益！如果有其他的任何侵权问题，请私信我2930546527@qq.com  
- 本项目主要是对自己3个月以来的学习的总结！希望交到更多同道的朋友！  

# 目录导航  

1. archive/tar： 
	-  **[archive/tar](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/archive_tar.go)** 
	-  **[archive/tar](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/archive_tar_01.go)** 
	-  **[archive/tar](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/archive_tar_02.go)** 
2. archive/zip：
	-  **[archive/zip](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/archive_zip_01.go)**
	-  **[archive/zip](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/archive_zip_02.go)**
	-  **[archive/zip](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/archive_zip_03.go)**
3. bufio：
	-  **[bufio_Reader](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/bufio_Reader.go)** 
	-  **[bufio_Writer](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/bufio_Writer.go)** 
	-  **[bufio_Scanner](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/bufio_Scanner_01.go)** 
	-  **[bufio_Scanner](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/bufio_Scanner_02.go)** 
	-  **[bufio_Scanner](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/bufio_Scanner_03.go)** 
4. bytes+Strings：
	-  **[bytes+Strings](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/bytes%2Bstrings_01.go)**
	-  **[bytes_Buffer](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/bytes_Buffer_01.go)**
	-  **[bytes_Reader](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/bytes_Reader_01.go)**
5. compress/bzip2：
	-  **[compress/bzip2](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/compress_bzip2.go)** 
6. compress/flate：
	-  **[compress/flate](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/compress_flate.go)** 
7. compress/gzip：
	-  **[compress/gzip](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/compress_gzip.go)** 
	-  **[compress/gzip](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/compress_gzip_01.go)** 
8. compress/lzw：
	-  **[compress/lzw](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/compress_lzw.go)** 
9. compress/zlib：
	-  **[compress/zlib](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/compress_zlib.go)** 
10. database/sql：
	-  **[database/sql](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/sql.go)**  
	-  **[database/sql](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/sql_01.go)** 
	-  **[database/sql](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/sql_regester().go)** 
11. database/sql/driver：
	-  **[database/sql/driver](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/sql_driver.go)**  
12. encoding：
	-  **[encoding](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding.md)**  
13. encoding/ascii85：
	-  **[encoding/ascii85](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_ascii85.go)**  
14. encoding/asn1：
	-  **[encoding/asn1](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_asn1.go)**  
	-  **[encoding/asn1](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_asn1_01.go)**
15. encoding/base32：
	-  **[encoding/base32](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_base32.go)**  
16. encoding/base64：
	-  **[encoding/base64](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_base64.go)**  
17. encoding/binary：
	-  **[encoding/binary](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_binary.go)**  
18. encoding/csv：
	-  **[encoding/csv](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_csv.go)**  
19. encoding/gob：
	-  **[encoding/gob](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_gob.go)** 
	-  **[encoding/gob](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_gob_01.go)** 
20. encoding/hex：
	-  **[encoding/hex](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_hex.go)**  
21. encoding/json：
	-  **[encoding/json](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_json_01.go)** 
	-  **[encoding/json](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_json_02.go)**  
22. encoding/pem：
	-  **[encoding/pem](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_pem.go)** 
23. encoding/xml：
	-  **[encoding/xml](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_xml.go)** 
	-  **[encoding/xml](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_xml_01.go)** 
	-  **[encoding/xml](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/encoding_xml_02.go)**
24. flag：
	-  **[flag](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/flag_01.go)** 
	-  **[flag](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/flag_02.go)**
	-  **[flag](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/flag_03.go)** 
25. fmt：
	-  **[fmt](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/fmt.go)**  
26. io：
	-  **[io](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/io_01.go)**  
27. io/ioutil：
	-  **[io/ioutil](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/io_ioutil_01.go)** 
	-  **[io/ioutil](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/io_ioutil_02.go)**  
28. log：
	-  **[log](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/log.go)**  
29. log/syslog：
	-  **[log/syslog](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/log_syslog.md)**  
30. math：
	-  **[math](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/math_01.go)**  
31. math/big：
	-  **[math/big](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/math_big.go)**  
32. math/cmplx：
	-  **[math/cmplx](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/math_cmplx.md)**  
33. math/rand：
	-  **[math/rand](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/math_rand.go)**   
34. os：
	-  **[os](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_02.go)** 
	-  **[os](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_03.go)** 
	-  **[os](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_04.go)** 
	-  **[os](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_05.go)** 
35. os/exec：
	-  **[os/exec](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_exec_01.go)** 
	-  **[os/exec](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_exec_02.go)** 
	-  **[os/exec](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_exec_03.go)** 
	-  **[os/exec](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_exec_04.go)** 
36. os/signal：
	-  **[os/signal](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_signal_01.go)** 
	-  **[os/signal](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_signal_02.go)**
	-  **[os/signal](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_signal_03.go)**  
37. os/user：
	-  **[os/user](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/os_user_01.go)**  
38. path：
	-  **[path](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/path_01.go)**  
39. path/filepath：
	-  **[path/filepath](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/path_filepath.go)**  
40. reflect：
	-  **[reflect](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/reflect_01.go)** 
	-  **[reflect](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/reflect_02.go)**
	-  **[reflect](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/reflect_03.go)**  
41. regexp：
	-  **[regexp](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/regexp.go)**  
42. regexp/syntax：
	-  **[regexp/syntax](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/regexp_syntax.go)**  
43. sort：
	-  **[sort](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/sort_01.go)** 
44. strconv：
	-  **[strconv](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/strconv_01.go)**   
45. text/scanner：
	-  **[text/scanner](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/text_scanner_01.go)** 
	-  **[text/scanner](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/text_scanner_02.go)** 
46. text/tabwriter：
	-  **[text/tabwriter](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/text_tabwriter_01.go)** 
	-  **[text/tabwriter](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/text_tabwriter_02.go)** 
47. text/template：
	-  **[text/template](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/text_template_01.go)**  
48. text/template/parse：
	-  **[text/template/parse](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/text_template_parse.md)**  
49. time：
	-  **[time](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/time_01.go)** 
	-  **[time](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/time_02.go)** 
50. unicode：
	-  **[unicode](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/unicode_01.go)**  
51. unicode/utf16：
	-  **[unicode/utf16](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/unicode_utf16.md)**  
52. unicode/utf8：
	-  **[unicode/utf8](https://github.com/AnkoGo/Go-Library-Demo/blob/master/io_pro/main/unicode_utf8_01.go)**  


# 后续更新：
本项目会不断的修复旧的go库笔记，同时更新未写的go库笔记！请多多关注！

如果遇到问题请发问！喜欢hub主的话可以给个小星星！感谢！  

#  贡献者：

- [@AnkoGo](https://github.com/AnkoGo)

# 联系与交流： 
- anko：2930546527@qq.com
- Q群：暂无

# 许可：

- 所有文章采用[知识共享署名-非商业性使用-相同方式共享 3.0 中国大陆许可协议](https://creativecommons.org/licenses/by-nc-sa/3.0/cn/)进行许可 