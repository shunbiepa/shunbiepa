package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/robertkrimen/otto"
)

var configPath string

var mp map[string]string

func main() {

	filePath := "/home/ashun/gosrc/shunbiepa/kabuda/p.js"
	//先读入文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	vm := otto.New()

	_, err = vm.Run(string(bytes))
	if err != nil {
		panic(err)
	}

	data := "eval(function(p,a,c,k,e,d){e=function(c){return(c<a?'':e(parseInt(c/a)))+((c=c%a)>35?String.fromCharCode(c+29):c.toString(36))};if(!''.replace(/^/,String)){while(c--){d[e(c)]=k[c]||e(c)}k=[function(e){return d[e]}];e=function(){return'\\\\w+'};c=1};while(c--){if(k[c]){p=p.replace(new RegExp('\\\\b'+e(c)+'\\\\b','g'),k[c])}}return p}('c[1]=\"b/d/e/g/a/n.f/0\";c[2]=\"b/d/e/g/a/m.f/0\";c[3]=\"b/d/e/g/a/p.f/0\";c[4]=\"b/d/e/g/a/q.f/0\";c[5]=\"b/d/e/g/a/r.f/0\";c[6]=\"b/d/e/g/a/l.f/0\";c[7]=\"b/d/e/g/a/i.f/0\";c[8]=\"b/d/e/g/a/h.f/0\";c[9]=\"b/d/e/g/a/j.f/0\";c[k]=\"b/d/e/g/a/s.f/0\";c[a]=\"b/d/e/g/a/o.f/0\";c[C]=\"b/d/e/g/a/z.f/0\";c[B]=\"b/d/e/g/a/y.f/0\";c[t]=\"b/d/e/g/a/v.f/0\";c[w]=\"b/d/e/g/a/x.f/0\";c[u]=\"b/d/e/g/a/A.f/0\";',39,39,'||||||||||11|images2|photosr|2022|06|jpg|05|56c58ee7c2|565030cbef|5691ef881b|10|56732d0631|560f99d536|56c1d922a2|569c601e2c|56d9537413|56f1a92147|56a86c5759|56efccfb5f|14|16|5625707282|15|5664354bc6|561847812f|561a840bd9|56fa7c5cf9|13|12'.split('|'),0,{}))\n"

	//encodeInp是JS函数的函数名
	value, err := vm.Call("GetPhotoList", nil, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(value.String())

	value2, err := vm.Get("photosr")
	if err != nil {
		panic(err)
	}

	list := strings.Split(value2.String(), ",")

	fmt.Println(list)

}
