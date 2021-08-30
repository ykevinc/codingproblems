use std::collections::HashMap;

struct Dummy {
    value: i32,
    index: i32
}

impl Solution {
    pub fn two_sum_sort(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut sorted: Vec<Dummy> = Vec::new();
        for (pos, e) in nums.iter().enumerate() {
            sorted.push(Dummy{value: *e, index: pos as i32});
        }
        sorted.sort_by_key(|d| d.value);
        let (mut i, mut j) = (0,  sorted.len()-1);
        loop {
            let s = sorted[i].value + sorted[j].value;
            if (s < target) {
                i = i + 1
            } else if (s > target) {
                j = j - 1
            } else {
                return vec![sorted[i].index, sorted[j].index];
            }
        }
        return nums
    }
    
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map = HashMap::<i32, i32>::new();
        for (pos, e) in nums.iter().enumerate() {
            match map.get(&(target-e)) {
                Some(entry) => return vec![*entry, pos as i32],
                None => { map.insert(*e, pos as i32); }
            }
        }
        return nums
    }
}
