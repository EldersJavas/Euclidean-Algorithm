local _gcd
_gcd = function(a, b)
  if b == 0 then
    return a
  else
    return _gcd(b, a % b)
  end
end



print(_gcd("1000", 0))


