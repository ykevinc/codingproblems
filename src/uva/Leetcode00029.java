public class Leetcode00029 {
    public int divide(int dividend, int divisor) {
        if (divisor == 0 ) {
            return Integer.MAX_VALUE;
        }
        if (dividend == Integer.MIN_VALUE && divisor == -1) {
            return Integer.MAX_VALUE;
        }
        int result = 0;
        long p = Math.abs((long)dividend);
        long q = Math.abs((long)divisor);
        while (p >= q) {
            int counter = 0;
            while (p >= (q << counter)) {
                counter++;
            }
            result += 1 << (counter - 1);
            p -= q << (counter - 1);
        }
        return ((dividend > 0) == (divisor > 0))?result:-result;
    }
}
