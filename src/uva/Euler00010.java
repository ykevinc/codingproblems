import java.io.*;
import java.util.*;

class Euler00010
{
  
    public static void main (String[] args) throws java.lang.Exception
    {
      System.out.println(find());
    }
  
    public static long find() {
      int n = 2000000;
      boolean isPrime[] = new boolean[n+1];
      Arrays.fill(isPrime, true);
      for (int i = 4;i < n;i+=2) {
        isPrime[i] = false;
      }
      for (int i = 3;i < Math.sqrt(n);i+=2) {
        for (int j = 2;j*i < n;j+=1) {
          isPrime[j*i] = false;
        }
      }
      long sum = 0;
      for (int i = 2;i < n;i++) {
        if (isPrime[i]) {
          sum += (long)i;
        }
      }
      return sum;
    }
}

