import java.io.*;
import java.util.Arrays;

class Euler00002
{
    public static void main (String[] args) throws java.lang.Exception
    {
        int n = 13;  
        int sum = 0;
        int next1 = 1;
        int next2 = 2;
        int next3 = 3;
        for (;next2 <= n;) {        
            if (next2%2 == 0) {
                sum += next2;
            }
            next1 = next2;
            next2 = next3;
            next3 = next1 + next2;
        }
        System.out.printf("The sum of Fibonacci are %d\n", sum);
    }
}

