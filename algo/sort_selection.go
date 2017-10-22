package algo

/*Function to swap two integers from array*/
func swap(array[] int , pos1 int , pos2 int) {
    temp := array[pos1]
    array[pos1] = array[pos2]
    array[pos2] = temp
}

/*Selection sort function*/
func SelectionSort(array []int) {
    for i:=0; i< len(array)-1; i++ {
        min := i
        for j:= i+1; j<len(array);j++ {
            if array[j] < array[min] {
                min = j
            }
        }
        swap(array,min,i)
    }
}

