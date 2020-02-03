package main

//5.1 Go语言项目实战：简单爬虫
//原文链接：https://blog.csdn.net/u010986776/article/details/82318137
import (
	"net/http"
	"fmt"
	"io/ioutil"
	"regexp"
)

var (
	//邮箱
	reEmail   = `\w+@\w+\.\w+(\.\w+)?`

	//超链接
	//<a href="http://news.baidu.com/ns?cl=2&rn=20&tn=news&word=%C1%F4%CF%C2%D3%CA%CF%E4%20%B5%BA%B9%FA"
	reLink    = `href="(https?://[\s\S]+?)"`

	//手机号
	//13x xxxx xxxx
	rePhone = `1[345789]\d\s?\d{4}\s?\d{4}`

	//身份证号
	//123456 1990 0817 123X
	reIdcard = `[123456]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dX]`

	//图片链接
	//"http://img2.imgtn.bdimg.com/it/u=2403021088,4222830812&fm=26&gp=0.jpg"
	reImg = `"(https?://[^"]+?(\.((jpg)|(jpeg)|(png)|(gif)|(bmp)|(svg)|(swf)|(ico))))"`
)



func main() {
	var url = "https://github.com/"
	var pageStr = GetPageStr(url)
	fmt.Println("page str:", pageStr)
	fmt.Println("-----------------------------------------------------------")
	SpiderLink(pageStr)
	fmt.Println("-----------------------------------------------------------")
	SpiderImg(pageStr)
}

func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	pageStr = string(pageBytes)
	return pageStr
}

func SpiderLink(pageStr string) {
	re := regexp.MustCompile(reLink)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条Link:\n", len(results))
	for i, result := range results {
		//result[0] -> ref="https://github.blog"
		//result[1] -> https://github.blog
		fmt.Printf("Link[%d]: %s\n", i, result[1])
	}
}

func SpiderImg(pageStr string) {
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条Img:\n", len(results))
	for i, result := range results {
		//result[0] -> https://github.com/fluidicon.png"
		//result[1] -> https://github.com/fluidicon.png
		fmt.Printf("Img[%d]: %s\n", i, result[1])
	}

}

func HandleError(err error, why string) {
	if err != nil {
		fmt.Print(why, err)
	}
}