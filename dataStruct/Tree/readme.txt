树是均衡了增删改查的所有功能,综合效率最高
避免了array和链表的性能极限的问题


二叉树
线段树
avl平衡树
红黑树
字典树
B+树
区块链默克尔树

      4
  2       6
1   3   5   7


      4
   2       6
1   3   5   7
中序遍历:  123 4 567  
前序遍历:  4 213  657
后序遍历:  132 576 4 

可以看到
中序遍历就是 根节点或者或者父节点在中间
后序遍历就是根节点或者父节点在最后,以此类推
而每一次遍历都是从左节点到右节点