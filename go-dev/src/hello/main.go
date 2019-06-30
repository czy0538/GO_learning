/*package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}*/

package main

import (
	"github.com/astaxie/beego"
	_ "hello/routers"
)

func main() {
	beego.Run()
}
