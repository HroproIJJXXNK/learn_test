package main

import (
		"fmt"
		"io/ioutil"
		"strings"
		"sort"
		)
/*
读入一个英文文件，统计每个单词的词频，然后
										1）输出top 10
										2）输入一个单词，输出词频
思路：
	1）读取文件并输出string
	2）将读到的string按一个个单词输出
	3）单词->字典
	4）将字典排序
	5）输出top 10
	6）输入一个单词，输出词频
*/
var result = ""
func CountTest(fileName string){
	//result := ""

	//读取文件
	contents,err := ioutil.ReadFile(fileName)
	if err == nil{
		//fmt.Printf("Type: %T \n",err)
		//contents存储的内容为bytes[]类型
		//fmt.Println(contents)
		result = strings.Replace(string(contents),"\n","",1)
		//result = string(contents)
		//fmt.Println(result)
	}
	
	m := stringToMap(result)
	fmt.Println(m)
	sortMapToList(m)
}

func stringToMap(s string)(map[string]int){
	wordMap := make(map[string]int)
	
	//处理str将其根据空格等分隔符划分为字符串数组
	strWord := strings.FieldsFunc(s, split)
	
	//遍历字符串数组
	for _, str := range strWord{
		//fmt.Println(_, str)
		
		//检查元素是否存在在map中
		if num, ok := wordMap[str];ok{
			num++;
			wordMap[str] = num;
		}else {
			wordMap[str] = 1;
		}
	}
	return wordMap
}

func split(s rune) bool{
	if s == ',' || s == '.' || s == '?' || s == ' ' || s == '	' || s == '\n' || s == '\r' || s == '!'{
		return true
	}
	return false
}

//定义一个Pair->PairList类型，
//同时写出swap、Len、Less函数->用于满足sort.Sort的调用条件
type Pair struct{
	key string
	value int
}
type PairList []Pair
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].value < p[j].value }

func sortMapToList(m map[string]int)PairList{
	p := make(PairList, len(m))
	i := 0; 
	for k, v := range m{
		p[i] = Pair{k, v}
		i++
	}
	//sort.Sort(p)//递增排序
	sort.Sort(sort.Reverse(p))//递减排序
	fmt.Println(p.Len())
	
	
	//输出top 5
	outLen := len(m)
	if outLen > 5{
		outLen = 5
	}
	for j := 0 ; j < outLen ;j++{
		fmt.Println(p[j])
	}
	return p
}

//输入单词，输出词频
func wordToCount(words string){
	m := stringToMap(result)
	fmt.Println(m[words])
}
func main() {
	CountTest("D:\\fileTest.txt")
	//fmt.Println("我爱我家")
	fmt.Printf("please enter a word:")
	word := ""
	fmt.Scanln(&word)
	fmt.Printf("单词词频为：")
	wordToCount(word)
}

