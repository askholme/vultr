package vultr
import ("fmt"
"strconv")
func isInt(str string) (bool) {
  _, err := strconv.Atoi(str)
  if err == nil {
    return true
  }
  return false
}

func errorMsg(err error,msg string) (error) {
  if err != nil {
    return fmt.Errorf(msg,err)
  }
  return nil
}