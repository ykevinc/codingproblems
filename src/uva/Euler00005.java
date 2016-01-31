import java.io.*;
import java.util.*;

class Euler00005
{
    public static void main (String[] args) throws java.lang.Exception
    {
      System.out.println(find(1, 20));
    }
  
    public static int find(int low, int high) {
      int factors[][] = new int[high+1][high+1];
      for (int i =low;i<=high;i++) {
        findFactors(i, factors);
      }
      int product = 1;
      // each factor
      for (int i=2;i<=high;i++) {
        int max = 0;
        // each num
        for (int j=2;j<=high;j++) {
          max = Math.max(factors[j][i], max);
        }
        product *= Math.pow(i, max);
      }
      return product;
    }
    public static int[][] findFactors(int num, int factors[][]) {
      int k = num;
      for (int i=2;i<=k;i+=1) {
        if (k%i == 0) {
          factors[num][i]++;
          k /= i;
          i=1;
        }
      }
      return factors;
    }
}

