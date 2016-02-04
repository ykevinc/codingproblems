import java.util.regex.*;
import java.math.BigInteger;

public class Leetcode00008 {
    public int myAtoi(String str) {
        if (str.length() == 0) {
            return 0;
        }
        if (str.equals("+")||str.equals("-")||str.contains("+-")||str.contains("-+")||str.contains("- ")||str.contains("+ ")||str.contains("++")||str.contains("--")) {
            return 0;
        }
        Matcher m = Pattern.compile("\\s*([-|+]?\\d+).*").matcher(str);
        if (m.matches()) {
            String digits = m.group(1);
            BigInteger l = new BigInteger(digits);
            if (l.compareTo(new BigInteger(String.valueOf(Integer.MAX_VALUE))) > 0) {
                return Integer.MAX_VALUE;
            } else if (l.compareTo(new BigInteger(String.valueOf(Integer.MIN_VALUE))) < 0) {
                return Integer.MIN_VALUE;
            }
            return l.intValue();
        }
        return 0;
    }
}
