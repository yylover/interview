二叉树
1. 递归遍历
2. 非递归， 使用栈模拟递归的过程

```
void traverse(TreeNode *root) {
    //前序遍历
    traverse(root->left)
    //中序遍历
    traverse(root->right)
    //后序遍历
}
```

1. 快排的逻辑类似于先序遍历
```quick sort
void sort(int nums[], int lo, int hi) {
    //通过交换元素构建分界点
    int p = partition(nums, lo, hi);
    sort(nums, lo, p-1)
    sort(nums, p+1, hi)
}

```

2. 归并排序的框架类似于后序遍历
```merge sort
void sort(int[] nums, int lo, int hi) {
    int mid = (lo + hi) / 2;
    sort(nums, lo, mid)
    sort(nums, mid+1, hi)

    merge(nums, lo, mid, hi);
}
```

### 递归问题 
1。 只要涉及递归，都可以抽象成二叉树的问题。
2. 递归算法的关键是: 明确函数定义是什么

### 
