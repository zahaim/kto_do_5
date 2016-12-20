package main
import "fmt"

type User struct {
  name string
  days [5]bool
}

type UserList []User

func (ul UserList) SetNames(names []string) UserList {
  for _, name := range names {
    fmt.Println("User to append:", name)
    ul = append(ul, User{name,[5]bool{false,false,false,false,false}})
  }
  for _, user := range ul {
    fmt.Println("Appended:", user)
  }
  return ul
}

func (ul UserList) ChangeThick(who, when int, what bool) {
  ul[who].days[when] = what
  fmt.Println("Changed to:", ul[who].days[when])
}

func main() {
  //days := [5]bool{false,false,false,false,false}
  names := []string{"Ala","Jacek"}
  fmt.Println("Names print:", names)

  users := UserList{}
  users = users.SetNames(names)

  fmt.Printf("We've got %d users in UserList\n", len(users))
  fmt.Println("Array of Users returned:", users)
  fmt.Printf("First user name: %s\n", users[0].name)
  fmt.Printf("Array of %s is %b\n", users[1].name, users[1].days)
  users.ChangeThick(1,2,true)
  users.ChangeThick(1,4,true)
  fmt.Printf("Array of %s is %b\n", users[1].name, users[1].days)
}
