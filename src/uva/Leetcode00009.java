public class Leetcode00009 {
    public boolean isPalindrome(int x) {
        if (x == 0) {
            return true;
        } else if (x < 0) {
            return false;
        } else {
            int reverse = reverse(x);
            if (reverse == 0) {
                return false;
            }
            reverse = Math.abs(reverse);
            x = Math.abs(x);
            while (reverse > 0 || x > 0) {
                if (reverse == 0 && x != 0) {
                    return false;
                }
                if (reverse != 0 && x == 0) {
                    return false;
                }
                if (reverse % 10 != x % 10) {
                    return false;
                }
                reverse /= 10;
                x /= 10;
            }
            return true;
        }
    }
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
