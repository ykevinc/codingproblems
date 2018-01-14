import java.io.*;
import java.util.Arrays;

class Euler00007
{
    public static void main (String[] args) throws java.lang.Exception
    {
        //System.out.println("Hello Java");
        int n = Integer.MAX_VALUE/(int)Math.pow(2,10);
        boolean isPrime[] = new boolean[n];
        Arrays.fill(isPrime, true);
        isPrime[2] = true;
        for (int i = 2;i < (int)Math.sqrt(n);i++) {
            for (int j = 2;i*j < n;j++) {
                isPrime[i*j] = false;
                //System.out.println("not prime" + i*j);
            }
        }
        for (int i=2, nth=0;i<n;i++) {
            if (isPrime[i]) {
                nth++;
            }
            if (nth == 10001) {
                System.out.printf("%dth is %d", nth, i);
                break;
            }
        }
    }
}

