package collatzconjecture
import "errors"

func CollatzConjecture(n int) (int, error) {
   var steps int
    if n <= 0 {
        return 0, errors.New("number is less than 1")
    }
    for n != 1{
        divideBy2 := n %2 
        if divideBy2 == 0{
            n = n/2
        }else{
            n = (n *3) +1
        }
        steps++
    }

    return steps, nil
}
