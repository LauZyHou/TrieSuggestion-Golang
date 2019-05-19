package logic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//学校结构体(大写字母开头表示是公有的,这样外面才能访问到)
type School struct {
	SchoolId   int    `json:"school_id"`   //学校id,后面表示在JSON传输时自动将该字段名改成school_id
	Province   string `json:"province"`    //省
	City       string `json:"city"`        //市
	SchoolType int    `json:"school_type"` //学校类型
	SchoolName string `json:"school_name"` //学校名字
}

var SchoolList []*School //定义一个切片,存解析后的学校数据的结构体指针

//logic包的初始化工作:解析文件中的学校数据,传入文件名字,返回出错时的error
func InitLogic(filename string) (err error) {
	//打开文件,返回文件的句柄和打开时错误
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("打开文件%s失败,错误%v\n", filename, err)
		return
	}
	//关掉打开的文件
	defer file.Close()

	var id int //用于自增生成id

	//接下来要读取文件内容,这里传入实现了Reader接口的文件对象,返回一个Reader
	reader := bufio.NewReader(file)
	for {
		//传入'\n'作为分隔符,则每次读取的就是一行的数据,这里用errRead是防止和err重名把err覆盖掉
		line, errRead := reader.ReadString('\n')
		//读到文件末尾时会返回EOF作为error
		if errRead == io.EOF {
			break
		}
		//如果错误既不是EOF又不为nil,说明读取文件出错了,这时也要返回
		if errRead != nil {
			fmt.Printf("读取文件%s失败,错误%v\n", filename, err)
			err = errRead //赋值给要返回的err
			return
		}

		//至此,读取这一行是成功的,将其分割出这行的具体要元素,得到的strSplit是一个切片
		strSplit := strings.Split(line, "\t")
		//如果分割完了没有四个元素(省份,市,学校类型,学校名字),那这是一个异常的数据
		if len(strSplit) != 4 {
			fmt.Printf("异常的数据行:%s,仅有%d个元素\n", line, len(strSplit))
			continue //跳过这个异常数据
		}

		//定义一个学校的结构体变量,将前面分割出的strSplit切片中的各元素去掉两端空白符后写入其中
		var school School
		school.Province = strings.TrimSpace(strSplit[0])
		school.City = strings.TrimSpace(strSplit[1])
		school.SchoolName = strings.TrimSpace(strSplit[2])
		//学校类型是int类型,要将读取来的string转换成int,转换过程中可能出错
		schoolType, errAtoi := strconv.Atoi(strings.TrimSpace(strSplit[3]))
		if errAtoi != nil {
			fmt.Printf("字符串%s无法转换为整数\n", strSplit[3])
			continue
		}
		school.SchoolType = schoolType
		id++
		school.SchoolId = id

		//加入到切片中
		SchoolList = append(SchoolList, &school)
		fmt.Printf("school:%+v\n", school) //%+v会将结构体中具体的字段名也打印出来
	}
	return
}
