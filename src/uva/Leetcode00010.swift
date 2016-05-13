class Solution {
    func isMatch(s: String, _ p: String) -> Bool {
        do {
            let regex = try NSRegularExpression(pattern:"^\(p)$", options: [])
            return regex.matchesInString(s, options: [], range:NSMakeRange(0, s.characters.count)).count > 0
        } catch let error as NSError {
            return false
        }
    }
}
