
#include <algorithm>
#include <chrono>
#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

// Definition for a binary tree node.
struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right)
      : val(x), left(left), right(right) {}
};

TreeNode *kek(int data) {
  struct TreeNode *node = new struct TreeNode;
  node->val = data;
  node->left = NULL;
  node->right = NULL;
  return (node);
}

int main(int argc, char const* argv[]) {
  vector<int> nums{-9, -2, -1, 0, 1, 2, 4, 5};

  int copy;



  class Solution {
   public:
    vector<int> preorderTraversal(TreeNode *root) {
      vector<TreeNode *> nodes;
      vector<int> val;
      nodes.push_back(root);
      if (root == NULL) return val;

      while (nodes.empty() != true) {
        root = nodes.back();
        nodes.pop_back();
        val.push_back(root->val);
        if (root->right) nodes.push_back(root->right);
        if (root->left) nodes.push_back(root->left);
      }
      return val;
    }

    vector<int> inorderTraversal(TreeNode* root) {
      vector<TreeNode *> nodes;
      vector<int> val;
      if (!root) return val;

      do {
        while (root) {
          nodes.push_back(root);
          root = root->left;
        }
        root = nodes.back();
        nodes.pop_back();
        val.push_back(root->val);
        root = root->right;
      } while (!nodes.empty() || root);
      return val;
    }
  };

  class Solution sol;

  struct TreeNode root = TreeNode();

  root.left = kek(3);
  root.right = kek(4);
  root.right->left = kek(5);
  root.left->right = kek(6);
  root.left->left = kek(7);

  auto start = chrono::steady_clock::now();
  vector<int> ha = sol.inorderTraversal(&root);
  auto end = chrono::steady_clock::now();
  cout << endl;
  cout << chrono::duration_cast<chrono::nanoseconds>(end - start).count()
       << " nanoseconds ";
  cout << chrono::duration_cast<chrono::microseconds>(end - start).count()
       << " microseconds" << endl;

  start = chrono::steady_clock::now();
  ha = sol.preorderTraversal(&root);
  end = chrono::steady_clock::now();
  cout << endl;
  cout << chrono::duration_cast<chrono::nanoseconds>(end - start).count()
       << " nanoseconds ";
  cout << chrono::duration_cast<chrono::microseconds>(end - start).count()
       << " microseconds" << endl;

  return 0;
}

