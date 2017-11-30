package myPackage

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending notification to %s with %s\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func (u user) changeValueEmail(email string) {
	u.email = email
}

func Type_test() (string, string) {

	fmt.Println("***********************************************")
	fmt.Println("Type_test() function...")

	/*
				type user struct{
					name		string
					email		string
					ext		int
					privileged	bool
				}

			type admin struct {
				person user
				level  string
			}

				var Bill user
				lisa := user {
					name : "Lisa",
					email : "lisa@email.com",
					ext : 123,
					privileged : true, // last comma required..
				}
		fmt.Println("[struct value]", lisa)
	*/

	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeValueEmail("bill@newdomain.com")
	bill.notify()
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	(*lisa).changeValueEmail("lisa@newdomain.com")
	(*lisa).notify()
	(*lisa).changeEmail("lisa@newdomain.com")
	(*lisa).notify()

	return "returned from Type_test()..", "second param"
}
