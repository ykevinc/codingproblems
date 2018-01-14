public class Leetcode00013 {
    public int romanToInt(String s) {
        Map<Character,Integer> nums = new HashMap<>();
        nums.put('I', 1);
        nums.put('V', 5);
        nums.put('X', 10);        
        nums.put('L', 50);
        nums.put('C', 100);
        nums.put('D', 500); 
        nums.put('M', 1000);
        int result=0;
        for(int i = 0;i < s.length();i++) {
           if(i+1<s.length()&&nums.get(s.charAt(i))<nums.get(s.charAt(i+1))) {
               result-=nums.get(s.charAt(i));
           } else {
               result+=nums.get(s.charAt(i));
           }
        }
        return result;
    }
}
