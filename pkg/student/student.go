/**
* @author Jee
* @date 2021/3/28 0:21
 */
package student

import (
	"dependence"
	"fmt"
	"github.com/pkg/errors"
)

type Stu struct {
	Id      string
	Name    string
	ClassId string
}

func New() *Stu {
	return &Stu{}
}

func (stu *Stu) GetInfo() string {
	fmt.Println("test version pkg/student/v1.0-fst-r1")
	fmt.Println(dependence.GetVersion())
	println(fmt.Sprintf("%v", errors.New("22222222")))
	return fmt.Sprintf("%s`s id is %s, class id is %s", stu.Name, stu.Id, stu.ClassId)
}
