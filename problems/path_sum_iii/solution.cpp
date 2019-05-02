/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
    int sumUp(TreeNode *node, int pre, int sum) {
        if (node == nullptr) {
            return 0;
        }
        int current = pre + node->val;
        return (current == sum) + sumUp(node->left, current, sum) + sumUp(node->right, current, sum);
    }
public:
    int pathSum(TreeNode* root, int sum) {
        if (root == nullptr) {
            return 0;
        }
        return sumUp(root, 0, sum) + pathSum(root->left, sum) + pathSum(root->right, sum);
    }
};