package main

func main() {


	println(GCD(12,18))
}


// _gcd = function(a, b)
//   if b == 0 then
//     return a
//   else
//     return _gcd(b, a % b)
//   end
// end

func GCD(a,b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
	