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
public:
    int getHeight(TreeNode *root) {
        if (root == nullptr){
            return 0;
        }
        
        return std::max(getHeight(root->left), getHeight(root->right)) + 1;
    }
        
    bool isBalanced(TreeNode* root) {
        if (root == nullptr) {
            return true;
        }
        
        int diff = std::abs(getHeight(root->left) - getHeight(root->right));
        if (diff > 1) {
            return false;
        } else {
            return isBalanced(root->left) && isBalanced(root->right);
        }
        
    }
};