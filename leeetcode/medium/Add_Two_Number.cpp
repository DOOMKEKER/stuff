#include <iostream>

struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {}
  explicit ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

class Solution {
 public:
  ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    ListNode* node1 = l1;
    ListNode* node2 = l2;

    ListNode* prev = nullptr;
    ListNode* node;

    int val1;
    int val2;
    int val = 0;

    while (node1 || node2) {
      val1 = 0;
      val2 = 0;

      node = node1 ? node1 : node2;

      if (node1) {
        val1 += node1->val;
        node->next = node1;
      }
      if (node2) {
        val2 += node->val;
        node->next = node2;
      }

      node->val = (val + val1 + val2) % 10;
      val = (val + val1 + val2) / 10;
    }
    if (val) node->next = new ListNode(val);
    node = l1;
    for (; node; node = node->next) {
      std::cout << node->val << std::endl;
    }
    return node;
  }
};

int main() {
  ListNode n1_1(9);
  ListNode n1_2(9, &n1_1);
  ListNode n1_3(9, &n1_2);
  ListNode n1_4(9, &n1_3);
  ListNode n1_5(9, &n1_4);
  ListNode n1_6(9, &n1_5);
  ListNode n1_7(9, &n1_6);
// [9,9,9,9,9,9,9], l2 = [9,9,9,9]

  ListNode n2_1(9);
  ListNode n2_2(9, &n2_1);
  ListNode n2_3(9, &n2_2);
  ListNode n2_4(9, &n2_3);

  Solution sol;

  sol.addTwoNumbers(&n2_4, &n1_7);
  return 1;
}
