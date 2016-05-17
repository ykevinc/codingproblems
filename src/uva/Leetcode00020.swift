class Stack<T> {
    var items = [T]()
    func push(i : T) {
        items.append(i)
    }
    func pop() -> T {
        return items.removeLast()
    }
    func count() -> Int {
        return items.count
    }
}

class Solution {
    func isValid(s: String) -> Bool {
        let m : [Character:Character] = [
            "(":")",
            "{":"}",
            "[":"]",
        ]
        var stack = Stack<Character>()
        for c in s.characters {
            switch (c) {
                case let value where m.keys.contains(c):
                    stack.push(c)
                case let value where m.values.contains(c):
                    if stack.count() == 0 || m[stack.pop()]! != c {
                        return false
                    }
                default:
                    return false
            }
        }
        return stack.count() == 0
    }
}
