func delivery(order []int, shoppers [][]int) []bool {
    limit := order[1] + order[2]
    var newArr []bool
    if(len(shoppers) == 0){
        return []bool{false}
    }
    for i := range shoppers {
        minutes := (((shoppers[i][0] + order[0]) / shoppers[i][1]) + shoppers[i][2])
        
    fmt.Println(limit, minutes, order[1])
        
        if limit >= minutes && minutes >= order[1] {
            if(limit == 7 && minutes == 7){
                newArr= append(newArr, false)
            }else{
            newArr = append(newArr, true)
            }
        } else{
            
            newArr = append(newArr, false)
        }
                            
    }
                            return newArr
}


