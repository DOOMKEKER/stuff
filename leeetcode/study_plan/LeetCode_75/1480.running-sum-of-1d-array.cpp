#include <iostream>
#include <vector>

#include "../../some_funcs.hpp"

using namespace std;

class Solution {
 public:
  vector<int> runningSum(vector<int>& nums) {
    for (auto i = 1; i < nums.size(); i++) {
      nums[i] = nums[i] + nums[i - 1];
    }

    return nums;
  }
};

int main() {
  Solution sol;
  vector<int> nums = {1, 2, 3};

  auto ans = sol.runningSum(nums);
  cout_vec(ans);

  return 0;
}
