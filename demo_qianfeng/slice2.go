/*
使用已有的数组创建切片：
slice := arr[start:end)
	切片中的数据：[start:end)
	去数据从头到 end：arr[:end]
	从 start 到末位：arr[start:]
	slice 中的元素包含 arr 中 start 位置上的元素，但不包含 end 位置上的元素

从已有的数组上，直接创建切片，该切片的底层数组就是当前的数组。
	长度是从 start 到 end 切割的数据量。
	但是容量从 start 到数组的末尾。
*/
package main

import (
	"fmt"
)

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("---------------1.使用已有的数组创建切片---------------")

	s1 := a[:5]  // 取 1-5
	s2 := a[3:8] // 取 4-8
	s3 := a[5:]  // 取 6-10
	s4 := a[:]   // 取 1-10

	fmt.Println(a)          // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("s1: ", s1) // [1 2 3 4 5]
	fmt.Println("s2: ", s2) // [4 5 6 7 8]
	fmt.Println("s3: ", s3) // [6 7 8 9 10]
	fmt.Println("s4: ", s4) // [1 2 3 4 5 6 7 8 9 10]

	/*
		注意：
			1、获取数组的地址必须使用 &a 而不是 a
			2、fmt.Printf("s1 的地址：%p\n", s1) 这里面虽然使用的占位符是 %p 而不是 %d，但是得到的结果是 s1 的值，它的值就是 0xc00005a0f0，即 s1 指向的数组的地址，而不是 s1 本身在内存中的地址。
			3、s1 也有自己的内存地址，获取方法是：fmt.Printf("s1 的地址：%p\n", &s1)
	*/
	fmt.Printf(" a 的地址：%p\n", &a) // 0xc00005a0f0
	fmt.Printf("s1 的地址：%p\n", s1) // 0xc00005a0f0 与数组 a 的地址相同
	fmt.Printf("s2 的地址：%p\n", s2) // 0xc00005a108
	fmt.Printf("s3 的地址：%p\n", s3) // 0xc00005a118
	fmt.Printf("s4 的地址：%p\n", s4) // 0xc00005a0f0 与数组 a 的地址相同

	fmt.Println("---------------2.长度和容量---------------")

	fmt.Printf("s1 的长度是：%d，容量是：%d\n", len(s1), cap(s1)) // 5,10
	fmt.Printf("s1 的长度是：%d，容量是：%d\n", len(s2), cap(s2)) // 5,7
	fmt.Printf("s1 的长度是：%d，容量是：%d\n", len(s3), cap(s3)) // 5,5
	fmt.Printf("s1 的长度是：%d，容量是：%d\n", len(s4), cap(s4)) // 10,10

	/*
		为什么切片 s1、s2、s3、s4 的容量是 10、7、5、10 呢？看上去无规律可循。
		这是因为切片是引用类型函数，它指向的是某一个数组的地址。
		所以切片的容量计算要从切片的第一个元素的地址开始到数组的末尾。

		例如：
		数组  a 的元素是：1 2 3 4 5 6 7 8 9 10
		切片 s1 的元素是：1 2 3 4 5
		切片 s2 的元素是：      4 5 6 7 8
		切片 s3 的元素是：          6 7 8 9 10
		切片 s4 的元素是：1 2 3 4 5 6 7 8 9 10

		所以：
		s1 容量是从 1 开始一直到 10，即容量为 10
		同理：
		s2 容量是从 4 开始一直到 10，即容量为 7
		s3 容量是从 6 开始一直到 10，即容量为 5
		s4 容量是从 1 开始一直到 10，即容量为 10
	*/

	fmt.Println("---------------3.更改数组的内容---------------")

	a[4] = 100
	fmt.Println(a)  // [1 2 3 4 100 6 7 8 9 10]
	fmt.Println(s1) // [1 2 3 4 100]
	fmt.Println(s2) // [4 100 6 7 8]
	fmt.Println(s3) // [6 7 8 9 10]
	fmt.Println(s4) // [1 2 3 4 100 6 7 8 9 10]
	/*
		由于切片是指向数组的地址，所以当数组的某个元素发生改变时，引用了该元素的切片也会发生变化。
	*/

	fmt.Println("---------------4.更改切片的内容-修改元素---------------")
	s2[2] = 200
	fmt.Println(a)  // [1 2 3 4 100 200 7 8 9 10]
	fmt.Println(s1) // [1 2 3 4 100]
	fmt.Println(s2) // [4 100 200 7 8]
	fmt.Println(s3) // [200 7 8 9 10]
	fmt.Println(s4) // [1 2 3 4 100 200 7 8 9 10]
	/*
		由于切片是指向数组的地址，所以当某个切片的元素发生改变时，数组的元素一定发生相应的改变。
		其他指向这个数组并引用了该变化元素的切片，也会发生变化。
	*/

	fmt.Println("---------------4.更改切片的内容-追加元素-长度超过容量---------------")

	s1 = append(s1, 1, 1, 1, 1) // 追加后的长度未超出底层数组的容量
	fmt.Println(a)              // [1 2 3 4 100 1 1 1 1 10]
	fmt.Println(s1)             // [1 2 3 4 100 1 1 1 1]
	fmt.Println(s2)             // [4 100 1 1 1]
	fmt.Println(s3)             // [1 1 1 1 10]
	fmt.Println(s4)             // [1 2 3 4 100 1 1 1 1 10]
	/*
		由于切片是指向数组的地址，所以当某个切片追加 n 个元素时（追加后的长度未超出底层数组的容量），数组的长度和包含元素首先发生改变。
		但由于数组是定长的，所以数组的相应变化不可能是追加 n 个元素时，而是将数组中对应该切片的位置之后的 n 个元素替换成在切片中追加的 n 个元素。
		其他指向这个数组并引用了这些变化元素的切片，也会发生变化。
	*/

	fmt.Println("---------------4.更改切片的内容-追加元素-长度未超过容量---------------")
	fmt.Printf("s1 的长度是：%d，容量是：%d\n", len(s1), cap(s1)) // s1 的长度是：9，容量是：10
	s1 = append(s1, 2, 2, 2, 2)                         // 追加后的长度超出底层数组的容量
	fmt.Println(a)                                      // [1 2 3 4 100 1 1 1 1 10]
	fmt.Println(s1)                                     // [1 2 3 4 100 1 1 1 1 2 2 2 2]
	fmt.Println(s2)                                     // [4 100 1 1 1]
	fmt.Println(s3)                                     // [1 1 1 1 10]
	fmt.Println(s4)                                     // [1 2 3 4 100 1 1 1 1 10]

	fmt.Printf("a 的地址：%p\n", &a)  // 0xc00005a0f0，地址不变
	fmt.Printf("s1 的地址：%p\n", s1) // 0xc00008c0a0，地址改变，指向新的数组
	fmt.Printf("s2 的地址：%p\n", s2) // 0xc00005a108，地址不变
	fmt.Printf("s3 的地址：%p\n", s3) // 0xc00005a118，地址不变
	fmt.Printf("s4 的地址：%p\n", s4) // 0xc00005a0f0，地址不变
	/*
		向切片中追加 n 个元素后，由于切片的长度超过了它原来指向的底层数组 a 的容量，
		所以系统创建了一个新的数组，s1 指向新数组。
		也由于这个原因，原来的数组 a 没有发生变化，从而指向这个数组的其他的切片也没有发生。
		从地址的变化可以验证上述结论的正确性。
	*/
}
