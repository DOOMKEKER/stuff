#include <vector>
#include <iostream>

class Solution {
 public:
  std::vector<std::vector<int>> combinationSum(std::vector<int>& candidates, int target) {
    std::vector<std::vector<int>> ans;
    std::vector<int> buf;
    for (size_t i = 0; i < candidates.size(); ++i) {
      int sum = 0;
      buf.clear();
      for (size_t j = 0; j < i; ++j) {
        while (sum < target) {
          if (sum == target) ans.push_back(buf);
          sum += candidates[j];
        }
      }
    }
  }
};