package main

import (
	"Go_Controls_Project/Classes"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var tip int = 1
	if tip == 1 {
		startCrons()
	} else {
	//m := Classes.Feature{PointX:11,PointY:12 }
	//fmt.Println(m.PointY)
	var mode int
	mode = 1
	if mode == 1{
		//channel ornegi buraya
		var myIndexForChannel int = 5
		channelim := make(chan int, myIndexForChannel)
		go ChannelTest(cap(channelim), channelim)
		for res := range(channelim){
			fmt.Println(res)
		}

	} else {
		var studentList []Classes.Student
		var count int
		fmt.Println("Toplam Öğrenci Sayısı :")
		fmt.Scanf("%d", &count)
		for i := 0; i < count; i++ {
			myStudent := Classes.Student{}
			fmt.Println("Öğrencinin Adı :")
			fmt.Scanf("%s", &myStudent.Name)
			fmt.Println("Öğrencinin Soyadı :")
			fmt.Scanf("%s", &myStudent.Surname)

			rand.Seed(time.Now().UnixNano())
			randomExamCount := rand.Intn(20)
			if randomExamCount == 0 {
				randomExamCount += 1
			}

			for examIndex := 0; examIndex < randomExamCount; examIndex++ {
				rand.Seed(time.Now().UnixNano())
				randomExamScore := rand.Intn(100)
				myStudent.CourseGrade = append(myStudent.CourseGrade, randomExamScore)
				time.Sleep(1000)
			}

			studentList = append(studentList, myStudent)
		}

		PrintStudentList(studentList)
	}
	}
}


func ChannelTest(count int, myChannel chan int){

	for i:= 0; i<count; i++ {
		 myChannel <- i*2 // i degerini 2 ile carpip kanal araciligi ile gonderiyorum.
	}

	close(myChannel) // en son kanali kapatiyorum.
}

func PrintStudentList(myStudentList []Classes.Student){
	fmt.Println("*****************************************")
	for i:=0 ; i<len(myStudentList)  ; i++  {
		myStudent := myStudentList[i]
		fmt.Printf("Student Name: %s, Surname: %s \n", myStudent.Name, myStudent.Surname)

		for index := 0; index < len(myStudent.CourseGrade); index++{
			fmt.Printf("Grade [%d]: %d \n",index,myStudent.CourseGrade[index])
		}

		gradeAverage := FindAverage(myStudent.CourseGrade)
		fmt.Printf("Grade Average : %d \n", gradeAverage)

	}
	fmt.Println("*****************************************")
}

func FindAverage(gradeArr []int) int{
	sum := 0
	for i:=0; i< len(gradeArr); i++ {
		sum += gradeArr[i]
		//fmt.Printf("Current Sum: %d \n",sum)
	}
	return  sum/len(gradeArr)
}