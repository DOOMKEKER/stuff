#include <algorithm>
#include <iostream>
#include <map>
#include <vector>

using namespace std;

class Solution {
 public:
  bool isIsomorphic(string s, string t) {
    vector<char> s_abc(300,0);
    vector<char> t_abc(300,0);

    for (size_t i = 0; i < s.size(); i++) {
      if (s_abc[s[i]] != 0) {
        if (s_abc[s[i]] != t[i]) return false;
      }

      if (t_abc[t[i]] != 0) {
        if (t_abc[t[i]] != s[i]) return false;
      }

      s_abc[s[i]] = t[i];
      t_abc[t[i]] = s[i];
    }

    return true;
  }
};

int main(void) {
  Solution sol;
  string kek = "egg";
  string kek2 = "add";
  cout << sol.isIsomorphic(kek, kek2);
  cout << "jopa";
  return 0;
}
