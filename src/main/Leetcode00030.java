public class Leetcode00030 {
    public List<Integer> findSubstring(String s, String[] words) {
        Map<String,Integer> mapCountByWord = new HashMap<>();
        for (String word : words) {
            Integer count = mapCountByWord.get(word);
            if (count == null) {
                count = 0;
            }
            mapCountByWord.put(word, ++count);
        }
        int wordLen = words[0].length();
        int sentenseLen = wordLen*words.length;
        List<Integer> result = new ArrayList<>();
        for (int i = 0;i <= s.length() - sentenseLen;i++) {
            if (isConcat(s.substring(i, i+sentenseLen), mapCountByWord, wordLen)) {
                result.add(i);
            }
        }
        return result;
    }
    
    public boolean isConcat(String s, Map<String,Integer> mapCountByWord, int wordLen) {
        Map<String,Integer> mapCountByWordMatch = new HashMap<>();
        for (int i = 0;i <= s.length() - wordLen;i+=wordLen) {
            String next = s.substring(i, i+wordLen);
            Integer count = mapCountByWord.get(next);
            Integer countMatch = mapCountByWordMatch.get(next);
            if (count == null) {
                return false;
            } else if (countMatch == null) {
                countMatch = 0;
            } else if (countMatch >= count) {
                return false;
            }
            mapCountByWordMatch.put(next, countMatch+1);
        }
        return true;
    }
}
