import java.util.*;

public class Leetcode00001 {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer,Integer> mapIndexByNum = new HashMap<>();
        for (int index2 = 0;index2 < nums.length ;index2++) {
            int n = nums[index2];
            Integer index1 = mapIndexByNum.get(target-n);
            if (index1 != null) {
                return new int[]{index1+1, index2+1};
            }
            mapIndexByNum.put(n, index2);
        }
        return null;
    }
}
