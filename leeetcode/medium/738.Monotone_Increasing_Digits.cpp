#include <iostream>
#include <vector>

#include "../some_funcs.hpp"

using namespace std;

class Solution {
 public:
  int monotoneIncreasingDigits(int n) {
    string str = to_string(n);
    auto size = str.size();
    int last = size;

    for (size_t i = size - 1; i > 0; --i) {
      if (str[i - 1] > str[i]) {
        str[i - 1] = str[i - 1] - 1;
        last = i;
      }
    }

    for (size_t i = last; i < size; i++) str[i] = '9';

    return stoi(str);
  }
};

int main() {
  Solution sol;

  cout << sol.monotoneIncreasingDigits(1234) << endl;
  cout << sol.monotoneIncreasingDigits(332) << endl;
  cout << sol.monotoneIncreasingDigits(234567823) << endl;

  return 2;
}
