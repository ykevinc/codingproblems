public class Leetcode00005 {
    public String longestPalindrome(String s) {
        int max = 1;
        int maxIndex = 0;
        for (int i = 0;i < s.length();i++) {
            int match = 1;
            for (int l=i-1,r=i+1;l >= 0 && l < s.length() && r >= 0 && r < s.length();l--,r++) {
                if (s.charAt(l) != s.charAt(r) ) {
                    break;
                }
                match += 2;
                if (match > max) {
                    max = match;
                    maxIndex = i;
                }
               
            }
            match = 0;
            for (int l=i,r=i+1;l >= 0 && l < s.length() && r >= 0 && r < s.length();l--,r++) {
                if (s.charAt(l) != s.charAt(r) ) {
                    break;
                }
                match += 2;
                if (match > max) {
                    max = match;
                    maxIndex = i;
                }
            }
        }
        return max%2==0?s.substring(maxIndex+1-max/2, maxIndex+1+max/2):s.substring(maxIndex-max/2, maxIndex+max/2+1);
    }
}
