public class Leetcode00006 {
    public String convert(String s, int numRows) {
        char m[][] = new char[numRows][1000];
        int c = 0;
        int r = 0;
        int dirs[][] = new int[][]{{1,0},{-1,1},{0,1}};
        int dir = 0;
        for (int i = 0;i < s.length();i++) {
            m[r][c] = s.charAt(i);
            if (numRows == 1) {
                dir = 2;
            } else if (r == numRows - 1) {
                dir = 1;
            } else if (r == 0) {
                dir = 0;
            }
            r += dirs[dir][0];
            c += dirs[dir][1];
        }
        StringBuilder sb = new StringBuilder();
        for (r = 0;r < m.length;r++) {
            for (c = 0;c < m[r].length;c++) {
                if (m[r][c] != '\0') {
                    sb.append(m[r][c]);                    
                }
            }
        }
        return sb.toString();
    }
}
