public class Leetcode00008 {
    public int reverse(int x) {
        if (x == 0) {
            return 0;
        }
        StringBuilder sb = new StringBuilder();
        if (x < 0) {
            if (x == Integer.MIN_VALUE) {
                return 0;
            }
            sb.append("-");
            x = -x;
        }
        while (x > 0) {
            sb.append(x % 10);
            x /= 10;
        }
        Long val = Long.parseLong(sb.toString());
        if (val > Integer.MAX_VALUE || val < Integer.MIN_VALUE) {
            return 0;
        }
        return val.intValue();
    }
}
