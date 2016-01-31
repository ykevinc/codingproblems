import java.io.*;
import java.util.Arrays;

class Euler00001
{
    public static void main (String[] args) throws java.lang.Exception
    {
        int n = 10;  
        int sum = 0;
        for (int i = 1;i < n;i++) {
            if (i%3 == 0||i%5 == 0) {
                sum += i;
            }
        }
        System.out.printf("The sum of multiples are %d\n", sum);
    }
}

