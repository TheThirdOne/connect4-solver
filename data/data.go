package data

var boards []map[int64]int8
var length int
var sets, gets, dels int
func Init(sections int){
  boards = make([]map[int64]int8,sections)
  length = sections
  for i, _ := range boards {
    boards[i]= make(map[int64]int8)
  }
}

func moves(hash int64) int{
  if hash < 0 {
    hash *= -1
  }
  shifted := int(hash >> 42)
  sum := 0
  for i := 0; i < 7; i++{
    sum += shifted % 8
    shifted = shifted >> 3
  }
  return sum
}
func Get(hash int64) (int8,bool) {
  gets++
  //select current sub-map
  tmp := boards[moves(hash)%length]
  value, result := tmp[hash];
  if result{
    //decrement counter for pruning
    if value < 0 {
      value += 2
    }else{
      value -= 2
    }
    //delete or update
    if value*value < 2 {
      dels++
      delete(tmp,hash)
    }else{
      tmp[hash] = value
    }
  }
  
  return value%2, result
}
func Set(hash int64, value int8){
  sets++
  tmp := boards[moves(hash)%length]
  if value < 0 {
    value -= 12
  }else{
    value += 12
  }
  tmp[hash] = value
}

func GetVals() (int,int,int){
  return sets, gets, dels
}