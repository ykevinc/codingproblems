import java.io.*;
import java.util.*;

class Euler00006
{
    public static void main (String[] args) throws java.lang.Exception
    {
      System.out.println(find(1, 100));
    }
  
    public static int find(int low, int high) {
      int n = (high-low+1);
      int squareOfSum = (int)Math.pow((low + high)*n/2,2);
      int sumOfSquares = n*(n+1)*(2*n+1)/6;
      return Math.abs(squareOfSum-sumOfSquares);
    }

}

