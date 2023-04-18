//#1
func toMap(ss []string) map[string]struct{} {
    result := map[string]struct{}{}
    
    for _, s := range ss {
        result[s] = struct{}{}
    }
    
    return result
}



//#2
type Data struct{
    Value interface{} // 8
    Dimensions map[string]string // {indicator: performance, object: drilingRig, season: winter}
}

// произвольный большой объем входящих данных
var input []*Data

// есть много памяти
// есть продолжительное время между получением данных и запросом на них

func UpdateData(dataS []*Data) map[string]*Data {
    m := make(map[string]*Data)
        
    _,v := range dataS {
         s := []string{}
        
         key, val := range v.Dimesions {
            s := []string{}
            
            s := fmt.Sprintf(i, v)

            
        }
        

    } 
}

// получение данных по запросу пользователя
func getData(dimensions map[string]string) *Data {
    // нужно вернуть данные по полному соответствию dimensions (все ключи+значения) или nil
    // минимальное время выполнения
}