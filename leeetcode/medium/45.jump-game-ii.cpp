#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int jump(vector<int>& nums) {
    int res = 0;
    if (nums.size() == 1) return 0;
    for (size_t i = 0; i < nums.size();) {
      int max_index = 0;
      int max = 0;
      res++;
      for (size_t j = 1; j <= nums[i]; j++) {
        if (i+j == nums.size()-1) return res;
        if (nums[i + j] + j >= max) {
          max = nums[i + j] + j;
          max_index = i + j;
        }
      }
      i = max_index;
    }
    return res++;
  }
};

int main() {
  Solution sol;
  vector<int> nums = {4,1,1,3,1,1,1};
  cout << sol.jump(nums);

  return 0;
}