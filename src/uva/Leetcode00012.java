public class Leetcode00012 {
    public String intToRoman(int num) {
        int nums[] = new int[]{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000};
        String roman[] = new String[]{"I", "IV", "V", "IX","X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"};
        StringBuilder sb = new StringBuilder();
        for(int i = 12; i>=0; ) {
            if(num>=nums[i]) {
                sb.append(roman[i]);
                num -= nums[i];
            } else {
                i--;
            }
        }
        return sb.toString();
    }
}
