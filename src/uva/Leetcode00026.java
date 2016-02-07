import java.util.*;

public class Leetcode00026 {
    public int removeDuplicates(int[] nums) {
        if (nums.length <= 1) {
            return nums.length;
        }
        int l = 0, r = 1;
        while (r < nums.length) {
            if (nums[l] != nums[r]) {
                nums[++l] = nums[r];
            }
            r++;
        }
        return l+1;
    }
}
