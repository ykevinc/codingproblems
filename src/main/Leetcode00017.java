import java.util.*;

public class Leetcode00017 {
    public List<String> letterCombinations(String digits) {
        if (digits.length() == 0) {
            return new ArrayList<>(0);
        }
        String map[] = new String[]{
            "",
            "",
            "abc",
            "def",
            "ghi",
            "jkl",
            "mno",
            "pqrs",
            "tuv",
            "wxyz",
        };
        List<String> l = new ArrayList<>();
        combine(l, "", digits, 0, map);
        return l;
    }
    
    public void combine(List<String> list, String desc, String src, int i, String[] m) {
        if (src.length() == i) {
            list.add(desc);
        } else {
            int digit = src.charAt(i) - '0';
            for (int j = 0;j < m[digit].length();j++) {
                combine(list, desc + m[digit].charAt(j), src, i+1, m);
            }
        }
    }
}
