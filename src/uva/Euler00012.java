import java.io.*;
import java.util.*;

class Euler00012.java
{

  public static void main (String[] args) throws java.lang.Exception
  {
    System.out.println(find(500));
  }

  public static int find(int n) {
    int k = 1;
    int sum = 1;
    while (findDivisions(sum) < n) {
      sum += (k++)+1;
    }
    return sum;
  }
  
  public static int findDivisions(int n) {
    int d = 0;
    for (int i = 1;i <= Math.sqrt(n);i++) {
      if (n%i == 0) {
        d += 1;
        if (i*i != n) {
          d += 1;
        }
      }
    }
    return d;
  }
}

