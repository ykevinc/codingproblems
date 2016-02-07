public class Leetcode00028 {
    public int strStr(String haystack, String needle) {
        int i = 0, j = 0;
        if (needle.length() == 0) {
            return 0;
        }
        int[] b = preProcessPattern(needle);
        while (i < haystack.length()) {
            while (j >= 0 && haystack.charAt(i) != needle.charAt(j)) {
                j = b[j];
            }
            i++;
            j++;
            if (j == needle.length()) {
                return (i - needle.length());
            }
        }
        return -1;
    }
    
    public int[] preProcessPattern(String ptrn) {
        int b[] = new int[ptrn.length()+1];
        int j = -1;
        b[0] = j;
        for (int i = 0;i<ptrn.length();) {
            while(j >= 0 && ptrn.charAt(i) != ptrn.charAt(j)) {
                j = b[j];
            }
            b[++i] = ++j;
        }
        return b;
    }
}
