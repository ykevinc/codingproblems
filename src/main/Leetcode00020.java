import java.util.*;

public class Leetcode00020 {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        Map<Character,Character> m = new HashMap<>();
        m.put('(',')');
        m.put('[',']');
        m.put('{','}');        
        for (int i = 0;i < s.length();i++) {
            char next = s.charAt(i);
            System.out.println(next);
            if (m.keySet().contains(next)) {
                stack.push(next);
            } else {
                if (stack.isEmpty()||next != m.get(stack.pop())) {
                    return false;
                }
            }
        }
        return stack.size() == 0;
    }
}
