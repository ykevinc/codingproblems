import java.util.*;

public class Leetcode00018 {
 
    public List<List<Integer>> fourSum(int[] nums, int target) {
        Arrays.sort(nums);
        
        List<List<Integer>> ll = new ArrayList<>();
        Set<String> unique = new HashSet<>();
        
        for (int i = 0;i < nums.length;i++) {
            for (int j = i+1;j < nums.length;j++) {
                int k = j+1;
                int l = nums.length - 1;
                while (k < l) {
                    int sum = nums[i] + nums[j] + nums[k] + nums[l];
                    if (sum > target) {
                        l--;
                    } else if (sum < target) {
                        k++;
                    } else if (sum == target) {
                        List<Integer> v = Arrays.asList(nums[i], nums[j], nums[k], nums[l]);
                        if (unique.add(v.toString())) {
                            ll.add(v);
                        }
                        l--;
                        k++;
                    }
                }
            }
        }
        return ll;
    }
}
