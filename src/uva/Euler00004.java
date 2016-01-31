import java.io.*;
import java.util.*;

class Euler00004
{
    public static void main (String[] args) throws java.lang.Exception
    {
      System.out.println(find(3));
    }
  
    public static Long find(int ndigits) {
      long lowLimit = Long.valueOf("1" + getRepeatedString('0', ndigits - 1));
      long highLimit = Long.valueOf(getRepeatedString('9', ndigits));
      boolean found = false;
      Long max = null;
      for (long i=highLimit;i>=lowLimit;i--) {
        for (long j=highLimit;j>=lowLimit;j--) {
          long product = i*j;
          if (isPalindromic(String.valueOf(product))) {
            max = max != null ? Math.max(product, max) : product;
          }
        }
      }
      return max;
    }
  
    public static boolean isPalindromic(String s) {
      for (int lhs = 0,rhs = s.length()-1;lhs<rhs;lhs++,rhs--) {
        if (s.charAt(lhs) != s.charAt(rhs)) {
          return false;
        }
      }
      return true;
    }
    public static String getRepeatedString(char c, int ndigits) {
      return String.valueOf(new char[ndigits]).replace("\0", String.valueOf(c));
    }
}

