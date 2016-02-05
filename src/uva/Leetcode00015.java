import java.util.*;

public class Leetcode00015 {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        Map<Integer,Set<Integer>> m = new TreeMap<>();
        int newNum = 0;
        for (int i = 0;i < nums.length;i++) {
            Set<Integer> l = m.get(nums[i]);
            if (l == null) {
                l = new HashSet<Integer>(3);
                m.put(nums[i], l);
            } else if (l.size() == 2 && nums[i]!=0) {
                continue;
            }
            l.add(i);
            newNum++;
        }
        nums = new int[newNum];
        int n = 0;
        for (int key : m.keySet()) {
            Set<Integer> l = m.get(key);
            for (int i = 0;i < l.size();i++) {
                nums[n++] = key;
            }
        }
        m.clear();
        for (int i = 0;i < nums.length;i++) {
            Set<Integer> l = m.get(nums[i]);
            if (l == null) {
                l = new HashSet<Integer>(3);
                m.put(nums[i], l);
            }
            l.add(i);
        }
        Set<String> unique = new HashSet<>();
        List<List<Integer>> ll = new ArrayList<>();
        for (int i = 0;i < nums.length;i++) {
            if (nums[i] > 0) {
                break;
            }
            for (int j = i + 1;j < nums.length;j++) {
    			if (nums[i] + nums[j] > 0 && nums[j] > 0) {
				    break;
    			}
                Set<Integer> l = m.get(-nums[i]-nums[j]);
                if (l != null) {
                    for (int index : l) {
                        if (index != i && index != j && index > i && index > j) {
                            List<Integer> three = new ArrayList<>(3);
                            three.add(nums[i]);
                            three.add(nums[j]);
                            three.add(nums[index]);
                            if (unique.add(three.toString())) {
                                ll.add(three);
                            }
                            break;
                        }
                    }
                }
            }
        }
        return ll;
    }
}
