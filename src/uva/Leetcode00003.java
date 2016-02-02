import java.util.*;

public class Leetcode00003 {
    public int lengthOfLongestSubstring(String s) {
        Map<Character,Integer> mapIndexByCharacter = new HashMap<>();
        int max = 0;
        for (int i = 0;i < s.length();i++) {
            Integer index = mapIndexByCharacter.get(s.charAt(i));
            if (index != null) {
                i = index;
                mapIndexByCharacter.clear();
            } else {
                mapIndexByCharacter.put(s.charAt(i), i);
                max = Math.max(max, mapIndexByCharacter.size());
            }
        }
        return max;
    }
}
