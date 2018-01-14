import java.io.*;
import java.util.*;

class Euler00003
{
    public static void main (String[] args) throws java.lang.Exception
    {
      long n = 600851475143L;
      long maxPrime = 1;
      long k = n;
      if (n%2 == 0) {
        k = k/2;
        maxPrime = Math.max(2, maxPrime);
      }
      for (long i=3;i<=k;i+=2) {
        if (n%i == 0) {
          k = k/i;
          maxPrime = Math.max(i, maxPrime);
        }
      }
      System.out.printf("Max prime is %d\n", maxPrime);  
    }
}

