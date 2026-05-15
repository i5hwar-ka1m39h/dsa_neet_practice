package sol013
func MaxArea(height []int) int {
   l := 0
   r := len(height) -1
   res := 0

   for l < r{
      area := min(height[l], height[r])*(r-l)

      if area >res{
         res = area
      }
      if height[l] <= height[r]{
         l++
      }else{
         r--
      }
   }
   return  res 
}