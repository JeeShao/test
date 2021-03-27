/**
* @author Jee
* @date 2021/3/28 0:21
 */
package student

import "fmt"

type Stu struct {
	Id      string
	Name    string
	ClassId string
}

func New() *Stu {
	return &Stu{}
}

func (stu *Stu) GetInfo() string {
	fmt.Println("version v1.0.2")
	return fmt.Sprintf("%s`s id is %s, class id is %s", stu.Name, stu.Id, stu.ClassId)
}
