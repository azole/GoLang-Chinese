// 補充說明 (官方 Tour 中沒有)
// 來源：https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.4.md
//
// struct 中可以有匿名欄位，這可以是另外一個 struct，也可以是內建類型與自定義類型
// 很像是物件的組合 (composition)
//
// 當欄位名稱發生衝突時，以外層的為優先，如果要訪問內層的，就要透過多層存取
// 例如: student.Human.name
// 如果是同層級的有共同名稱的欄位，則沒有規定誰會優先存取，
// 必須要透過多層存取，如果沒有，則會有編譯錯誤 (compiler error)

package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，struct
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
}

func main() {
	// 初始化学生Jane
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}
