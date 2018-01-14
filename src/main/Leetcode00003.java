import java.util.*;

public class Solution {
    public int lengthOfLongestSubstring(String s) {
        Integer mapIndexByCharacter[] = new Integer[128];
        int max = 0, j = 0;
        for (int i = 0;i < s.length();i++) {
            Integer index = mapIndexByCharacter[s.charAt(i)];
            if (index != null) {
                j = Math.max(j, index+1);
            } 
            mapIndexByCharacter[s.charAt(i)] = i;
            max = Math.max(max, i - j + 1);
        }
        return max;
    }
}
