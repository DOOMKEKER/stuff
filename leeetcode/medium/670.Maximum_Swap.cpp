#include <iostream>
#include <vector>

#include "../some_funcs.hpp"

using namespace std;

class Solution {
 public:
  int maximumSwap(int num) {
    vector<int> vec;
    int u = 1;
    do {
      vec.push_back(num % 10);
      num = num / 10;
    } while (num);

    int flag = 0;

    int argmax = 0;
    int max = vec[vec.size() - 1];
    for (size_t i = vec.size() - 1; i >= 1; --i) {
      for (int j = i - 1; j >= 0; --j) {
        if (vec[j] > vec[i] && max < vec[j]) {
          max = vec[j];
          argmax = j;
          flag = 1;
        }
      }
      if (flag) {
        vec[argmax] = vec[i];
        vec[i] = max;
        break;
      }
    }

    int ret = 0;
    int pow = 1;
    for (size_t i = 0; i < vec.size(); i++) {
      ret += pow * vec[i];
      pow *= 10;
    }

    return ret;
  }
};

int main() {
  Solution sol;
  std::vector<int> nums = {1, 2, 3};

  int kek = sol.maximumSwap(956);
  cout << kek << endl;
}
