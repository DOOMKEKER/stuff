#include <algorithm>
#include <iostream>
#include <map>
#include <vector>

using namespace std;

class Solution {
 public:
  bool isIsomorphic(string s, string t) {
    auto str = max(s, t);
    if (s.size() != t.size()) {
      return false;
    };
    map<char, char> s_abc;
    map<char, char> t_abc;
    for (size_t i = 0; i < str.size(); i++) {
      if (s_abc[s[i]] != 0) {
        if (s_abc[s[i]] != t[i]) return false;
      }

      if (t_abc[t[i]] != 0) {
        if (t_abc[s[i]] != s[i]) return false;
      }

      s_abc[s[i]] = t[i];
      t_abc[t[i]] = s[i];
    }
    return true;
  }
};

int main(void) {
  Solution sol;
  string kek = "jopa";
  string kek2 = "jopo";
  cout << sol.isIsomorphic(kek, kek2);
  return 0;
}
