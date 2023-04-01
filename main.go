package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var Activities=[]string{
	"logging in",
	"logging out",
	"create record",
	"update record",
	"delete record",
}
type User struct{
	id int
	email string
	logs []Logitem
}

type Logitem struct {
	Activity string
	timestamp time.Time
}


func main(){

	rand.Seed(time.Now().Unix())
	wg := &sync.WaitGroup{}
	t:=time.Now()
	users := GenerateUser(1000)
	for _, user := range users{
		wg.Add(1)
		go SaveUserInfo(user,wg)
	}

	wg.Wait()

	fmt.Println("TIME ELAPSED:",time.Since(t).String())

}

func SaveUserInfo(user User,wg *sync.WaitGroup) error{
	fmt.Printf("Saving file for user id: %d\n",user.id)
	filename:= fmt.Sprintf("logs/uid_%d.txt",user.id)
	file,err:=os.OpenFile(filename,os.O_RDWR | os.O_CREATE, 0644)
	if err!=nil{
		return err
	}
	_,err=file.WriteString(user.GetActivityInfo())
	if err !=nil{
		return err
	}

	wg.Done()

	return nil

}
func (u User) GetActivityInfo() string {
	out:=fmt.Sprintf("ID:%d | email:%s\n Activity Log:\n", u.id, u.email)
	for i,item:=range u.logs{
		out+=fmt.Sprintf("%d. [%s] at %s\n",i+1, item.Activity, item.timestamp)
	}
	return out
}
func GenerateUser(count int) []User{
	users := make([]User,count)
	for  i:=0;i<count;i++{
		users[i]=User{
			id: i+1,
			email: fmt.Sprintf("user%d@mail.ru",i),
			logs: generateLogs(rand.Intn(1000)),
		}
	}
	return users
}



func generateLogs(count int)[]Logitem{
	logs:=make([]Logitem,count)
	for i := 0 ; i<count; i++{
		logs[i]=Logitem{
			timestamp: time.Now(),
			Activity: Activities[rand.Intn(len(Activities))],
		}
	}
	return logs
}

