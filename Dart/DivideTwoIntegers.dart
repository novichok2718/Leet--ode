class Solution 
{
  int divide(int dividend, int divisor) 
  {
    int quotient = 0;
    int sign = dividend.sign * divisor.sign;
    dividend = dividend.abs();
    divisor = divisor.abs();
    int shift = 0;
    while (dividend >= (divisor << shift)) 
    {
      if (dividend < (divisor << (shift + 1)))
      {
        quotient += 1 << shift;
        dividend -= (divisor << shift);
        shift = -1;
      }
      ++shift;
    }
    quotient *= sign;
    if (quotient > (1 << 31) - 1) return (1 << 31) - 1;
    if (quotient < -(1 << 31)) return -(1 << 31);
    return quotient;
  }
}

