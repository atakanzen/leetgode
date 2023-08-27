import (
"strings" 
"strconv"
)

func encode(s []string) string {
   var sb strings.Builder
   for _, str := range s {
       sb.WriteString(strconv.Itoa(len(str)))
       sb.WriteString("#")
       sb.WriteString(str)
   }

   return sb.String()
}

func decode(s string) []string {
   result := make([]string, 0)
   var sb strings.Builder

   for i := 0; i < len(s); {
       sb.WriteByte(s[i])
       if string(s[i+1]) == "#" {
           numOfChars, myErr := strconv.Atoi(sb.String())

           if myErr != nil {
               panic(myErr)
           }

           sb.Reset()
           result = append(result, s[i+2:i+2+numOfChars])
           i = i + 2 + numOfChars
       }
   }

   return result
}

