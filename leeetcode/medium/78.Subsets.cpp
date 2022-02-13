#include <iostream>
#include <vector>
#include "../some_funcs.hpp"

class Solution {
 public:
  std::vector<std::vector<int>> subsets(std::vector<int>& nums) {
    std::vector<std::vector<int>> ans;
    std::vector<int> buff;
    ans.push_back(buff);
    for (auto& num : nums) {
      size_t size = ans.size();
      for (size_t i = 0; i < size; ++i) {
        buff = ans[i];
        buff.push_back(num);
        ans.push_back(buff);
      }
    }
    return ans;
  }
};

int main() {
  Solution sol;
  std::vector<int> nums = {1, 2, 3};

  std::vector<std::vector<int>> kek = sol.subsets(nums);
  cout_vec(kek);
}
