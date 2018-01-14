import java.util.*;

public class Leetcode00004 {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int merged[] = new int[nums1.length + nums2.length];
        int l = 0;
        int r = 0;
        int i = 0;
        while (l < nums1.length || r < nums2.length ) {
            if (l < nums1.length && r < nums2.length) {
                if (nums1[l] < nums2[r]) {
                    merged[i++] = nums1[l++];
                } else {
                    merged[i++] = nums2[r++];
                }
            } else if (l < nums1.length) {
                merged[i++] = nums1[l++];
            } else {
                merged[i++] = nums2[r++];
            }
        }
        if (merged.length % 2 == 0) {
            return (merged[merged.length/2-1] + merged[merged.length/2])/2.0;
        } else {
            return merged[merged.length/2];
        }
    }

}
