public class Leetcode000010 {
    
    private boolean isMatch(String s, int sIndex, String p, int pIndex) {
        while (pIndex < p.length()) {
            char sc = sIndex < s.length() ? s.charAt(sIndex) : '\0';
            if (p.charAt(pIndex) != '*') {
                if (pIndex + 1 < p.length()) {
                    if (p.charAt(pIndex+1) != '*') {
                        if (p.charAt(pIndex) != '.' && sc != p.charAt(pIndex)) {
                            return false;
                        }
                        sIndex++;
                        pIndex++;                        
                    } else {
                        if (isMatch(s, sIndex, p, pIndex+2)) {
                            return true;
                        }
                        while (sIndex < s.length() && (p.charAt(pIndex) == '.' || s.charAt(sIndex) == p.charAt(pIndex))) {
                            sIndex++;
                            if (isMatch(s, sIndex, p, pIndex+2)) {
                                return true;
                            }
                        }
                        pIndex += 2;
                    }
                } else if (pIndex == p.length() - 1) {
                    if (p.charAt(pIndex) != '.' && sc != p.charAt(pIndex)) {
                        return false;
                    }
                    sIndex++;
                    pIndex++;
                }
            }
        }
        return sIndex == s.length() && pIndex == p.length();
    }
    
    public boolean isMatch(String s, String p) {
        return isMatch(s, 0, p, 0);
    }
}
