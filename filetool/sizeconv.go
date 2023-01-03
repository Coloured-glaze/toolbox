package filetool

import (
	"fmt"
	"math"
)

// 字节的单位转换 保留两位小数
func FormatSize(fileSize int64) (size string) {
   if fileSize < 1024 {
      //return strconv.FormatInt(fileSize, 10) + "B"
      return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
   } else if fileSize < int64(math.Pow(1024, 2)) {
      return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
   } else if fileSize < int64(math.Pow(1024, 3)) {
      return fmt.Sprintf("%.2fMB", float64(fileSize)/math.Pow(1024, 2))
   } else if fileSize < int64(math.Pow(1024, 4)) {
      return fmt.Sprintf("%.2fGB", float64(fileSize)/math.Pow(1024, 3))
   } else if fileSize < int64(math.Pow(1024, 5)) {
      return fmt.Sprintf("%.2fTB", float64(fileSize)/math.Pow(1024, 4))
   } else if fileSize < int64(math.Pow(1024, 6)) {
      return fmt.Sprintf("%.2fPB", float64(fileSize)/math.Pow(1024, 5))
   } else { // if fileSize < float64(math.Pow(1024, 7)) {
      return fmt.Sprintf("%.2fEB", float64(fileSize)/math.Pow(1024, 6))
   }
}