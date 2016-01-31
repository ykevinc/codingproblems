import java.io.*;
import java.util.*;

class Euler00009
{
  
    public static void main (String[] args) throws java.lang.Exception
    {
      System.out.println(find());
    }
  
    public static String find() {
      for (int b=1;b<500;b++) {
        int a = 1000*(500-b)/(1000-b);
        int c = 1000 - a - b;
        if (a*a+b*b==c*c) {
          return (a + " " + b + " " + c);
        }
      }
      return "no solution";
    }
}

