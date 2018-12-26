# YAML
读取.yaml配置文件(Go)

YAML是一种类似XML、JSON的标记性语言，一种较为广泛的用途是用来编辑配置文件，其扩展名有“.yaml”和“.yml”两种。

该版本的实现依赖于一个已知的结构（yaml.go文件），即使用者应首先定义需要的配置项结构，以下是几点注意事项：

1.执行demo.go文件前，如有必要，请先修改路径信息（如图1）。

![image](https://github.com/mats9693/YAML/blob/master/images/%E5%9B%BE1%EF%BC%9A%E4%BF%AE%E6%94%B9%E8%B7%AF%E5%BE%84%E4%BF%A1%E6%81%AF.png)

图1：修改路径信息

2.人为地将yaml配置项分成了两种，一种是一级标签后紧跟着值的简单配置项（如图2），另一种是拥有低级子标签的复杂配置项（如图3），
相应地，处理办法也有所不同：简单配置项被添加进”Inline”结构体，添加”,inline”标志，当做内置类型（应该是这么叫吧，如图4）；
而复杂配置项则需要声明对应格式的结构体（如图5）。

![image](https://github.com/mats9693/YAML/blob/master/images/%E5%9B%BE2%EF%BC%9A%E7%AE%80%E5%8D%95%E9%85%8D%E7%BD%AE%E9%A1%B9%EF%BC%88%E4%B8%A4%E4%B8%AA%E9%83%BD%E6%98%AF%EF%BC%89.png)

图2：简单配置项（两个都是）

![image](https://github.com/mats9693/YAML/blob/master/images/%E5%9B%BE3%EF%BC%9A%E5%A4%8D%E6%9D%82%E9%85%8D%E7%BD%AE%E9%A1%B9%EF%BC%88Person%EF%BC%89.png)

图3：复杂配置项（Person）

![image](https://github.com/mats9693/YAML/blob/master/images/%E5%9B%BE4%EF%BC%9A%E7%AE%80%E5%8D%95%E9%85%8D%E7%BD%AE%E9%A1%B9%E5%AE%9A%E4%B9%89.png)

图4：简单配置项定义

![image](https://github.com/mats9693/YAML/blob/master/images/%E5%9B%BE5%EF%BC%9A%E5%A4%8D%E6%9D%82%E9%85%8D%E7%BD%AE%E9%A1%B9%E5%AE%9A%E4%B9%89.png)

图5：复杂配置项定义

3.无论哪一种配置项，具体到值时，只要使用了“-”（横线，通常用来表示多个并列元素从属于同一上级）或“[ ]”（list，同上），
即使只有一个"-"或list中只有一个值，也需要修改其类型为相应的切片，并在后面的yaml声明中添加”,flow”标志（如图6）。

![image](https://github.com/mats9693/YAML/blob/master/images/%E5%9B%BE6%EF%BC%9A%E5%8C%85%E5%90%AB%E5%A4%9A%E4%B8%AA%E5%80%BC%E7%9A%84Hobby%EF%BC%8C%E4%BD%BF%E7%94%A8slice%2B%E2%80%9D%2Cflow%E2%80%9D.png)

图6：包含多个值的Hobby，使用slice+”,flow”

4.可以只定义部分结构。如在本例中，删除某一个结构体及其关联地声明/使用也是允许的。
