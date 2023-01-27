что покажет этот код, какие проблемы?

package main

```
import ("fmt")

func main(){
    for i:=0; i<10;i++ {
        go func(){
        fmt.Println(i)
        }()
    }
}
```

==============================

```
type Person struct {
    name String
}

func changeName(person2 *Person) {
    // person1 и person2 указывают на одну и ту же память
    // но это разные указатели,которые лежат в разных местах памяти
     *person2 := Person{
        name: "Alice"
     }
}

func main(){
    person1 = &Person{
        name: "Bob"
    }

    changeName(person1)
    fmt.Println(person1) // Bob
}
```

====================

```
func a() {
    x:=[]int{}       // [] 0 len, 0cap, 0x0
    x = append(x, 0) // [0] 1 len, 1 cap 0x4
    x = append(x, 1) // [0, 1] 2 len, 2 cap, 0x8
    x = append(x, 2) // [0, 1, 2] 3 len, 4 cap, 0x8
    y = append(x, 3) // [0, 1, 2, 3->4] 4 len, 4 cap 0x8
    z = append(x, 4) // [0, 1, 2, 4] 4 len, 4 cap, 0x8

    fmt.Println(y,z) //

}

func main() {
    a()
}
```
