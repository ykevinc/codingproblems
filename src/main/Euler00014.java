import java.io.*;
import java.util.*;

class Euler00014
{  
  public static void main (String[] args) throws java.lang.Exception
  {
    System.out.println(find(13, 1000000));
  }

  public static int find(int from, int to) {
    int max = 0;
    int records[] = new int[to+1];
    int c = 0;
    for (int i = from;i < to;i++) {
      max = Math.max(max, collatz(i, records));
    }
    return max;
  }
  
  public static int collatz(long n, int records[]) {
    int c = 1;
    if (n < records.length && records[(int)n] > 0) {
      return records[(int)n];
    } else if (n%2 == 0) {
      c = collatz(n/2, records) + 1;
      if (n < records.length) {
        records[(int)n] = c;
      }
      return c;
    } else if (n == 1) {
      records[1] = 1;
      return 1;
    } else {
      c = collatz(3*n + 1, records) + 1;
      if (n < records.length) {
        records[(int)n] = c;
      }
      return c;
    }
  }
}

