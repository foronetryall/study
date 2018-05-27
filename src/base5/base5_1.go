package main

import "fmt"

/**
Go语言-数组类型
     一个数组（Array）就是一个可以容纳若干类型相同的元素的容器。这个容器的大小（即数组的长度）是固定的，且是体现在数组的类型字面量之中的。比如，我们声明了一个数组类型：

type MyNumbers [3]int
注：类型声明语句由关键字type、类型名称和类型字面量组成。

     所谓类型字面量，就是用于表示某个类型的字面表示（或称标记方法）。相对的，用于表示某个类型的值的字面表示可被称为值字面量，或简称为字面量。比如之前提到过的3.7E-2就可被称为浮点数字面量。 类型字面量[3]int由两部分组成。第一部分是由方括号包裹的数组长度，即[3]。这也意味着，一个数组的长度是该数组的类型的组成部分，是固定不变的。该类型字面量的第二个组成部分是int。它代表了该数组可以容纳的元素的类型。说到这里，上面这条类型声明语句实际上是为数组类型[3]int声明了一个别名类型。这使得我们可以把MyNumbers当做数组类型[3]int来使用。


    我们表示这样一个数组类型的值的时候，应该把该类型的类型字面量写在最左边，然后用花括号包裹该值包含的若干元素。各元素之间以（英文半角）逗号分隔，即：

[3]int{1, 2, 3}
    现在，我们把这个数组字面量赋给一个名为numbers的变量：

var numbers = [3]int{1, 2, 3}
注：这是一条变量声明语句。它在声明变量的同时为该变量赋值。

    另一种便捷方法是，在其中的类型字面量中省略代表其长度的数字，像这样：

var numbers = [...]int{1, 2, 3}
    这样就可以免去我们为填入那个数字而数出元素个数的工作了。

    接下来，我们可以很方便地使用索引表达式来访问该变量的值中的任何一个元素，例如：

numbers[0] // 会得到第一个元素
numbers[1] // 会得到第二个元素
numbers[2] // 会得到第三个元素
注：索引表达式由字符串、数组、切片或字典类型的值（或者代表此类值的变量或常量）和由方括号包裹的索引值组成。在这里，索引值的有效范围是[0, 3)。也就是说，对于数组来说，索引值既不能小于0也不能大于或等于数组值的长度。另外要注意，索引值的最小有效值总是0，而不是1。

    相对的，如果我们想修改数组值中的某一个元素值，那么可以使用赋值语句直接达到目的。例如，我们要修改numbers中的第二个元素的话，如此即可：

numbers[1] = 4
    虽然数组的长度已经体现在了它的类型字面量，但是我们在很多时候仍然需要明确的获得它，像这样：

var length = len(numbers)
注：len是Go语言的内建函数的名称。该函数用于获取字符串、数组、切片、字典或通道类型的值的长度。我们可以在Go语言源码文件中直接使用它。

    最后，要注意，如果我们只声明一个数组类型的变量而不为它赋值，那么该变量的值将会是指定长度的、其中各元素均为元素类型的零值（或称默认值）的数组值。例如，若有这样一个变量：

var numbers2 [5]int
  则它的值会是

[5]int{0, 0, 0, 0, 0}
 */

func main() {
	var numbers2 [5]int
	numbers2[0] = 2
	numbers2[3] = numbers2[0] - 3 // -1
	numbers2[1] = numbers2[2] + 5 // 5
	numbers2[4] = len(numbers2) // 5
	sum := (11)
	// “==”用于两个值的相等性判断
	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
}